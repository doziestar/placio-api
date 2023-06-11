package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/auth0/go-auth0/management"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"placio-app/models"
	"placio-app/utility"
	"placio-pkg/hash"
	"strings"

	"gorm.io/gorm"
)

const (
	auth0TokenCacheKey      = "auth0_mgmt_token"
	auth0TokenEncryptionKey = "eyJhbGciOiJSUzIC"
)

type UserService interface {
	GetUser(authOID string) (*models.User, error)
	CreateBusinessAccount(userID, name, role string) (*models.BusinessAccount, error)
	GetUserBusinessAccounts(userID string) ([]models.BusinessAccount, error)
	CanPerformAction(userID, businessAccountID string, action string) (bool, error)
	RemoveUserFromBusinessAccount(userID, businessAccountID uint) error
	GetUsersForBusinessAccount(businessAccountID string) ([]models.User, error)
	GetBusinessAccountsForUser(userID string) ([]models.BusinessAccount, error)
	AssociateUserWithBusinessAccount(userID, businessAccountID, role string) error
	AcceptInvitation(invitationID uint) error
	InviteUserToBusinessAccount(email string, businessAccountID uint, role string) (*models.Invitation, error)
	RejectInvitation(invitationID uint) error
	TransferBusinessAccountOwnership(currentOwnerID uint, newOwnerID uint, businessAccountID uint) error
	GetUserInvitations(userID uint) ([]*models.Invitation, error)
	//UpdateAuth0UserData(userID string, userData *models.Auth0UserData, appData *models.AppMetadata, userMetaData *models.Metadata) (*models.Auth0UserData, error)
	GetAuth0UserData(userID string) (*management.User, error)
	UpdateAuth0UserMetadata(userID string, userMetaData *models.Metadata) (*management.User, error)
	UpdateAuth0UserInformation(userID string, userData *models.Auth0UserData) (*management.User, error)
	// GetAuth0ManagementToken GetAuth0UserMetaData(userID string, IdToken string) (models.Metadata, error)
	//GetAuth0AppMetaData(userID string, IdToken string) (models.AppMetadata, error)
	//GetAuth0UserRoles(userID string, IdToken string) ([]string, error)
	//GetAuth0UserPermissions(userID string, IdToken string) ([]string, error)
	//GetAuth0UserGroups(userID string, IdToken string) ([]string, error)
	//GetAuth0UserRolesPermissions(userID string, IdToken string) ([]string, error)
	//AuthorizeUser(userID string, IdToken string, roles []string, permissions []string, groups []string) error
	//DeAuthorizeUser(userID string, IdToken string, roles []string, permissions []string, groups []string) error
	GetAuth0ManagementToken(ctx context.Context) (string, error)
	UpdateAuth0AppMetadata(userID string, appData *models.AppMetadata) (*management.User, error)
}

type UserServiceImpl struct {
	db    *gorm.DB
	cache *utility.RedisClient
}

type Auth0TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func NewUserService(db *gorm.DB, cache *utility.RedisClient) *UserServiceImpl {
	return &UserServiceImpl{db: db, cache: cache}
}

func (s *UserServiceImpl) GetUser(auth0ID string) (*models.User, error) {
	log.Println("GetUser", auth0ID)
	var user models.User
	if err := s.db.Preload("Relationships").Where("auth0_id = ?", auth0ID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// The user does not exist in our database, so let's create one
			newUser := models.User{Auth0ID: auth0ID, UserID: models.GenerateID()}
			if err := s.db.Create(&newUser).Error; err != nil {
				return nil, err
			}
			return &newUser, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetUserBusinessAccounts retrieves all the business accounts
// associated with a specific user from the database.
func (s *UserServiceImpl) GetUserBusinessAccounts(userID string) ([]models.BusinessAccount, error) {
	// Define a slice to hold the UserBusinessRelationship instances.
	var relationships []models.UserBusinessRelationship

	// Use the GORM Preload method to automatically load the BusinessAccount
	// instances associated with each UserBusinessRelationship when fetching
	// the UserBusinessRelationship instances from the database.
	if err := s.db.Preload("BusinessAccount").Where("user_id = ?", userID).Find(&relationships).Error; err != nil {
		// If an error occurs during database query, return it.
		return nil, err
	}

	// Define a slice to hold the BusinessAccount instances.
	businessAccounts := make([]models.BusinessAccount, len(relationships))

	// Iterate over the UserBusinessRelationship instances.
	for i, relationship := range relationships {
		// Extract the BusinessAccount from each UserBusinessRelationship
		// and place it in the BusinessAccount slice.
		businessAccounts[i] = relationship.BusinessAccount
	}

	// Return the BusinessAccount slice.
	return businessAccounts, nil
}

func (s *UserServiceImpl) CanPerformAction(userID, businessAccountID string, action string) (bool, error) {
	var relationship models.UserBusinessRelationship
	if err := s.db.Where("user_id = ? AND business_account_id = ?", userID, businessAccountID).First(&relationship).Error; err != nil {
		return false, err
	}

	// Check if the user's role within the business account allows the action
	// This will depend on how you define the capabilities of each role
	if relationship.Role == "admin" && action == "delete_account" {
		return true, nil
	}

	return false, nil
}

// CreateBusinessAccount creates a new Business Account and associates it with a user.
func (s *UserServiceImpl) CreateBusinessAccount(userID, name, role string) (*models.BusinessAccount, error) {
	businessAccount := &models.BusinessAccount{Name: name, ID: models.GenerateID()}
	relationship := &models.UserBusinessRelationship{UserID: userID, BusinessAccount: *businessAccount, Role: "owner", BusinessAccountID: businessAccount.ID, ID: models.GenerateID()}

	tx := s.db.Begin()

	if err := tx.Create(businessAccount).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(relationship).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	// Now we need to fetch the created business account with its relationships
	var createdBusinessAccount models.BusinessAccount
	if err := s.db.Preload("Relationships").Where("id = ?", businessAccount.ID).First(&createdBusinessAccount).Error; err != nil {
		return nil, err
	}

	return &createdBusinessAccount, nil
}

// AssociateUserWithBusinessAccount associates a user with a Business Account.
func (s *UserServiceImpl) AssociateUserWithBusinessAccount(userID, businessAccountID, role string) error {
	relationship := &models.UserBusinessRelationship{UserID: userID, BusinessAccountID: businessAccountID, Role: role}
	return s.db.Create(relationship).Error
}

// GetBusinessAccountsForUser returns all Business Accounts associated with a user.
func (s *UserServiceImpl) GetBusinessAccountsForUser(userID string) ([]models.BusinessAccount, error) {
	var relationships []models.UserBusinessRelationship
	if err := s.db.Preload("BusinessAccount").Where("user_id = ?", userID).Find(&relationships).Error; err != nil {
		return nil, err
	}

	businessAccounts := make([]models.BusinessAccount, len(relationships))
	for i, relationship := range relationships {
		businessAccounts[i] = relationship.BusinessAccount
	}
	return businessAccounts, nil
}

// GetUsersForBusinessAccount returns all Users associated with a Business Account.
func (s *UserServiceImpl) GetUsersForBusinessAccount(businessAccountID string) ([]models.User, error) {
	var relationships []models.UserBusinessRelationship
	if err := s.db.Preload("User").Where("business_account_id = ?", businessAccountID).Find(&relationships).Error; err != nil {
		return nil, err
	}

	users := make([]models.User, len(relationships))
	for i, relationship := range relationships {
		users[i] = relationship.User
	}
	return users, nil
}

func (s *UserServiceImpl) UpdateAuth0UserInformation(userID string, userData *models.Auth0UserData) (*management.User, error) {
	mergedData, err := s.prepareUserData(userID, userData)
	if err != nil {
		return nil, err
	}
	return s.updateAuth0Data(userID, mergedData, "user information")
}

func (s *UserServiceImpl) UpdateAuth0UserMetadata(userID string, userMetaData *models.Metadata) (*management.User, error) {
	mergedData, err := s.prepareUserData(userID, userMetaData)
	if err != nil {
		return nil, err
	}
	return s.updateAuth0Data(userID, mergedData, "user_metadata")
}

func (s *UserServiceImpl) UpdateAuth0AppMetadata(userID string, appData *models.AppMetadata) (*management.User, error) {
	log.Println("Updating app metadata")
	newAppData, err := utility.StructToMap(&appData)
	//mergedData, err := s.prepareUserData(userID, appData)
	if err != nil {
		return nil, err
	}
	return s.updateAuth0Data(userID, newAppData, "app_metadata")
}

func (s *UserServiceImpl) prepareUserData(userID string, data interface{}) (map[string]interface{}, error) {
	currUserData, err := s.GetAuth0UserData(userID)
	if err != nil {
		log.Println("Error getting current user data", err)
		return nil, err
	}

	currUserDataMap, err := utility.StructToMap(&currUserData)
	if err != nil {
		return nil, err
	}
	newDataMap, err := utility.StructToMap(data)
	if err != nil {
		return nil, err
	}

	return utility.MergeMaps(currUserDataMap, newDataMap), nil
}

func (s *UserServiceImpl) updateAuth0Data(userID string, mergedUserData map[string]interface{}, dataType string) (*management.User, error) {
	log.Println("Updating auth0 data", dataType)
	client := &http.Client{}

	if dataType != "user information" {
		mergedUserData = map[string]interface{}{
			dataType: mergedUserData,
		}
	}

	jsonPayload, err := json.Marshal(mergedUserData)
	if err != nil {
		return nil, err
	}
	log.Println("JSON payload", string(jsonPayload))

	req, err := http.NewRequest("PATCH", fmt.Sprintf("https://auth.placio.io/api/v2/users/%s", userID), bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println("Error creating request", err)
		return nil, err
	}

	managementToken, err := s.GetAuth0ManagementToken(context.Background())
	if err != nil {
		log.Println("Error getting management token", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", managementToken))

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request", err)
		return nil, err
	}
	defer resp.Body.Close()

	log.Println("Response status code", resp.StatusCode)
	log.Println("Response status", resp.Status)

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	log.Println("Successfully updated", dataType, "for user", userID)
	// unmarshal the response body into management.User
	var user management.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetAuth0UserData retrieves the current user data from Auth0.
func (s *UserServiceImpl) GetAuth0UserData(userID string) (*management.User, error) {
	log.Println("Getting Auth0 user data", userID)
	// Create an HTTP client
	client := &http.Client{}

	// Create the request
	req, err := http.NewRequest("GET", fmt.Sprintf("https://auth.placio.io/api/v2/users/%s", userID), nil)
	if err != nil {
		log.Println("Error creating request", err)
		return &management.User{}, err
	}

	//Get the token
	managementToken, err := s.GetAuth0ManagementToken(context.Background())
	if err != nil {
		log.Println("Error getting management token", err)
		return &management.User{}, err
	}
	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", managementToken))

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request", err)
		return &management.User{}, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return &management.User{}, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse the response body
	var userData management.User
	err = json.NewDecoder(resp.Body).Decode(&userData)
	if err != nil {
		return &management.User{}, err
	}

	return &userData, nil
}

func (s *UserServiceImpl) GetAuth0ManagementToken(ctx context.Context) (string, error) {

	// Check if token is in cache
	encryptedTokenBytes, err := s.cache.GetCache(ctx, auth0TokenCacheKey)
	if err != nil {
		log.Println("Error retrieving token from cache:", err)
	} else {
		// If token is in cache, decrypt and return
		encryptedToken := string(encryptedTokenBytes)
		encryptedToken = strings.Trim(encryptedToken, "\"") // Remove quotes from the string
		//log.Println("Retrieved encrypted token from cache:", encryptedToken) // Added log line
		token, err := hash.DecryptString(encryptedToken, auth0TokenEncryptionKey)
		//log.Println("Decrypted token:", token) // Added log line
		if err != nil {
			log.Println("Error decrypting token:", err)
		} else {
			return token, nil
		}
	}

	log.Println("Token not in cache, retrieving new token")

	// If token is not in cache or there was an error, retrieve a new one
	token, err := s.retrieveAuth0Token(ctx) // Replace with actual function to retrieve token
	if err != nil {
		return "", err
	}

	// Encrypt and cache the new token
	encryptedToken, err := hash.EncryptString(token, auth0TokenEncryptionKey)
	if err != nil {
		return "", err
	}

	log.Println("Caching encrypted token:", encryptedToken) // Added log line
	err = s.cache.SetCache(ctx, auth0TokenCacheKey, encryptedToken)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserServiceImpl) retrieveAuth0Token(ctx context.Context) (string, error) {

	payload := strings.NewReader(`{
		"client_id": "wORDxmfRFTkBvoSVU06Af4HFQAo25gUI",
		"client_secret": "_tbGyuaBZ7j5zp659MU5AYqkZessCVeNs2bv8Yl1Hp6XUj_hUdQAW9a5zw8hIA3F",
		"audience": "https://dev-qlb0lv3d.us.auth0.com/api/v2/",
		"grant_type": "client_credentials"
	}`)

	log.Println("payload", payload)
	req, _ := http.NewRequest("POST", "https://dev-qlb0lv3d.us.auth0.com/oauth/token", payload)
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error retrieving token from Auth0:", err)
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", errors.New("auth0 request not successful")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return "", err
	}

	var tokenResponse Auth0TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		log.Println("Error unmarshalling response body:", err)
		return "", err
	}

	return tokenResponse.AccessToken, nil
}

// RemoveUserFromBusinessAccount removes a User's association with a Business Account.
func (s *UserServiceImpl) RemoveUserFromBusinessAccount(userID, businessAccountID uint) error {
	return s.db.Where("user_id = ? AND business_account_id = ?", userID, businessAccountID).Delete(&models.UserBusinessRelationship{}).Error
}

func (s *UserServiceImpl) GetUserInvitations(userID uint) ([]*models.Invitation, error) {
	// Implementation goes here
	return nil, nil
}

func (s *UserServiceImpl) TransferBusinessAccountOwnership(currentOwnerID uint, newOwnerID uint, businessAccountID uint) error {
	// Implementation goes here
	return nil
}

func (s *UserServiceImpl) RejectInvitation(invitationID uint) error {
	// Implementation goes here
	return nil
}

func (s *UserServiceImpl) AcceptInvitation(invitationID uint) error {
	// Implementation goes here
	return nil
}

func (s *UserServiceImpl) InviteUserToBusinessAccount(email string, businessAccountID uint, role string) (*models.Invitation, error) {
	// Implementation goes here
	return nil, nil
}

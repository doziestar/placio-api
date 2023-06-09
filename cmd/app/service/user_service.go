package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"placio-app/models"
	"placio-app/utility"

	"gorm.io/gorm"
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
	UpdateAuth0UserData(userID string, IdToken string, userData *models.Auth0UserData, appData *models.AppMetadata, userMetaData *models.Metadata) (*models.Auth0UserData, error)
	GetAuth0UserData(userID string, IdToken string) (models.Auth0UserData, error)
}

type UserServiceImpl struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{db: db}
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

// UpdateAuth0UserData updates the user_metadata, app_metadata, and other fields in Auth0.
func (s *UserServiceImpl) UpdateAuth0UserData(userID string, IdToken string, userData *models.Auth0UserData, appData *models.AppMetadata, userMetaData *models.Metadata) (*models.Auth0UserData, error) {
	log.Println("Updating Auth0 user data", userID, IdToken, userData, appData, userMetaData)
	// Create an HTTP client
	client := &http.Client{}

	// Get the current user data
	currUserData, err := s.GetAuth0UserData(userID, IdToken)
	if err != nil {
		log.Println("Error getting current user data", err)
		return nil, err
	}
	log.Println("Current user data", currUserData)

	// Convert the current data and the new data into maps
	currUserDataMap, err := utility.StructToMap(&currUserData)
	if err != nil {
		return nil, err
	}
	newUserDataMap, err := utility.StructToMap(userData)
	if err != nil {
		return nil, err
	}
	appDataMap, err := utility.StructToMap(appData)
	if err != nil {
		return nil, err
	}
	userMetaDataMap, err := utility.StructToMap(userMetaData)
	if err != nil {
		return nil, err
	}

	// Merge the new data with the current data
	mergedUserData := utility.MergeMaps(currUserDataMap, newUserDataMap)

	// Overwrite the metadata fields with the new data
	mergedUserData["user_metadata"] = userMetaDataMap
	mergedUserData["app_metadata"] = appDataMap

	// Create the JSON payload
	jsonPayload, err := json.Marshal(mergedUserData)
	if err != nil {
		return nil, err
	}

	// Create the request
	req, err := http.NewRequest("PATCH", fmt.Sprintf("https://auth.placio.io/api/v2/users/%s", userID), bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", IdToken))

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Return the updated user data
	updatedUserData, err := s.GetAuth0UserData(userID, IdToken)
	if err != nil {
		return nil, err
	}

	return &updatedUserData, nil
}

// GetAuth0UserData retrieves the current user data from Auth0.
func (s *UserServiceImpl) GetAuth0UserData(userID string, IdToken string) (models.Auth0UserData, error) {
	// Create an HTTP client
	client := &http.Client{}

	// Create the request
	req, err := http.NewRequest("GET", fmt.Sprintf("https://auth.placio.io/api/v2/users/%s", userID), nil)
	if err != nil {
		return models.Auth0UserData{}, err
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", IdToken))

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return models.Auth0UserData{}, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return models.Auth0UserData{}, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse the response body
	var userData models.Auth0UserData
	err = json.NewDecoder(resp.Body).Decode(&userData)
	if err != nil {
		return models.Auth0UserData{}, err
	}

	return userData, nil
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

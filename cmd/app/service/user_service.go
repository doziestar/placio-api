package service

import (
	"errors"
	"gorm.io/gorm"
	"placio-app/models"
)

type UserService interface {
	GetUser(authOID string) (*models.User, error)
	CreateBusinessAccount(name string, userID string, role string) (*models.BusinessAccount, error)
	GetUserBusinessAccounts(userID string) ([]models.BusinessAccount, error)
	CanPerformAction(userID, businessAccountID string, action string) (bool, error)
}

type UserServiceImpl struct {
	db *gorm.DB
}

func (s *UserServiceImpl) GetUser(auth0ID string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("auth0_id = ?", auth0ID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// The user does not exist in our database, so let's create one
			newUser := models.User{Auth0ID: auth0ID}
			if err := s.db.Create(&newUser).Error; err != nil {
				return nil, err
			}
			return &newUser, nil
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserServiceImpl) CreateBusinessAccount(name string, userID uint, role string) (*models.BusinessAccount, error) {
	newBusinessAccount := models.BusinessAccount{Name: name}
	if err := s.db.Create(&newBusinessAccount).Error; err != nil {
		return nil, err
	}

	relationship := models.UserBusinessRelationship{
		UserID:            userID,
		BusinessAccountID: newBusinessAccount.ID,
		Role:              role,
	}

	if err := s.db.Create(&relationship).Error; err != nil {
		return nil, err
	}

	return &newBusinessAccount, nil
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

func (s *UserServiceImpl) CanPerformAction(userID uint, businessAccountID uint, action string) (bool, error) {
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

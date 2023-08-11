package settings

import (
	"placio-app/models"

	"gorm.io/gorm"
)

type GeneralSettingsService interface {
	CreateGeneralSettings(userID string, db *gorm.DB) error
	UpdateGeneralSettings(userID string, updatedSettings *models.GeneralSettings, db *gorm.DB)
	GetGeneralSettings(userID string, db *gorm.DB) (*models.GeneralSettings, error)
	DeleteGeneralSettings(userID string, db *gorm.DB) error
	// add other necessary methods
}

type NotificationsSettingsService interface {
	CreateNotificationsSettings(userID string, db *gorm.DB) error
	GetNotificationsSettings(userID string, db *gorm.DB) (*models.NotificationsSettings, error)
	UpdateNotificationsSettings(userID string, updatedSettings *models.NotificationsSettings, db *gorm.DB) error
	DeleteNotificationsSettings(userID string, db *gorm.DB) error
	// add other necessary methods
}

type AccountSettingsService interface {
	CreateAccountSettings(accountID string, db *gorm.DB) (*models.AccountSettings, error)
	GetAccountSettings(id string, d *gorm.DB) (*models.AccountSettings, error)
	AddBlockedUser(userID string, db *gorm.DB) error
	RemoveBlockedUser(userID string, db *gorm.DB) error
	AddMutedUser(userID string, db *gorm.DB) error
	RemoveMutedUser(userID string, db *gorm.DB) error
	UpdateAccountSettings(accountID string, updatedSettings models.AccountSettings, db *gorm.DB) error
	DeleteAccountSettings(accountID string, db *gorm.DB) error
}

type ConnectedAccountService interface {
	CreateConnectedAccount(userID string, provider string, db *gorm.DB) error
	GetConnectedAccount(accountID string, db *gorm.DB) (*models.ConnectedAccount, error)
	UpdateConnectedAccount(accountID string, updatedAccount *models.ConnectedAccount, db *gorm.DB) error
	DeleteConnectedAccount(accountID string, db *gorm.DB) error
	DeleteContentSettings(userID string, db *gorm.DB) error
}

type ContentSettingsService interface {
	CreateContentSettings(userID string, db *gorm.DB) error
	GetContentSettings(userID string, db *gorm.DB) (*models.ContentSettings, error)
	UpdateContentSettings(userID string, updatedSettings models.ContentSettings, db *gorm.DB) error
	DeleteContentSettings(userID string, db *gorm.DB) error
	// add other necessary methods
}

type SettingsService struct {
	db                    *gorm.DB
	generalSettings       *models.GeneralSettings
	notificationsSettings *models.NotificationsSettings
	accountSettings       *models.AccountSettings
	contentSettings       *models.ContentSettings
}

func NewSettingsService(db *gorm.DB) *SettingsService {
	return &SettingsService{db: db}
}

// For GeneralSettings
func (g *SettingsService) GetGeneralSettings(userID string, db *gorm.DB) (*models.GeneralSettings, error) {
	var settings models.GeneralSettings
	result := db.First(&settings, "user_id = ?", userID)
	return &settings, result.Error
}

// For NotificationsSettings
func (n *SettingsService) GetNotificationsSettings(userID string, db *gorm.DB) (*models.NotificationsSettings, error) {
	var settings models.NotificationsSettings
	result := db.First(&settings, "user_id = ?", userID)
	return &settings, result.Error
}

// For ContentSettings
func (c *SettingsService) GetContentSettings(userID string, db *gorm.DB) (*models.ContentSettings, error) {
	var settings models.ContentSettings
	result := db.First(&settings, "user_id = ?", userID)
	return &settings, result.Error
}

// For GeneralSettings
func (g *SettingsService) UpdateGeneralSettings(userID string, updatedSettings *models.GeneralSettings, db *gorm.DB) error {
	result := db.Model(&models.GeneralSettings{}).Where("user_id = ?", userID).Updates(updatedSettings)
	return result.Error
}

// For NotificationsSettings
func (n *SettingsService) UpdateNotificationsSettings(userID string, updatedSettings *models.NotificationsSettings, db *gorm.DB) error {
	result := db.Model(&n.notificationsSettings).Where("user_id = ?", userID).Updates(updatedSettings)
	return result.Error
}

// For AccountSettings
func (a *SettingsService) UpdateAccountSettings(accountID string, updatedSettings models.AccountSettings, db *gorm.DB) error {
	result := db.Model(&a.accountSettings).Where("account_id = ?", accountID).Updates(updatedSettings)
	return result.Error
}

// For ContentSettings
func (c *SettingsService) UpdateContentSettings(userID string, updatedSettings models.ContentSettings, db *gorm.DB) error {
	result := db.Model(&models.ContentSettings{}).Where("user_id = ?", userID).Updates(updatedSettings)
	return result.Error
}

// For GeneralSettings
func (g *SettingsService) DeleteGeneralSettings(userID string, db *gorm.DB) error {
	result := db.Where("user_id = ?", userID).Delete(&models.GeneralSettings{})
	return result.Error
}

// For NotificationsSettings
func (n *SettingsService) DeleteNotificationsSettings(userID string, db *gorm.DB) error {
	result := db.Where("user_id = ?", userID).Delete(&models.NotificationsSettings{})
	return result.Error
}

// For AccountSettings
func (a *SettingsService) DeleteAccountSettings(accountID string, db *gorm.DB) error {
	result := db.Where("account_id = ?", accountID).Delete(&models.AccountSettings{})
	return result.Error
}

// For ContentSettings
func (c *SettingsService) DeleteContentSettings(userID string, db *gorm.DB) error {
	result := db.Where("user_id = ?", userID).Delete(&models.ContentSettings{})
	return result.Error
}

// AddBlockedUser method adds a new user to the BlockedUsers list
func (a *SettingsService) AddBlockedUser(userID string, db *gorm.DB) error {
	a.accountSettings.BlockedUsers = append(a.accountSettings.BlockedUsers, userID)
	result := db.Model(a).Updates(models.AccountSettings{BlockedUsers: a.accountSettings.BlockedUsers})
	return result.Error
}

// RemoveBlockedUser method removes a user from the BlockedUsers list
func (a *SettingsService) RemoveBlockedUser(userID string, db *gorm.DB) error {
	for i, u := range a.accountSettings.BlockedUsers {
		if u == userID {
			a.accountSettings.BlockedUsers = append(a.accountSettings.BlockedUsers[:i], a.accountSettings.BlockedUsers[i+1:]...)
			break
		}
	}
	result := db.Model(a).Updates(models.AccountSettings{BlockedUsers: a.accountSettings.BlockedUsers})
	return result.Error
}

// AddMutedUser method adds a new user to the MutedUsers list
func (a *SettingsService) AddMutedUser(userID string, db *gorm.DB) error {
	a.accountSettings.MutedUsers = append(a.accountSettings.MutedUsers, userID)
	result := db.Model(a).Updates(models.AccountSettings{MutedUsers: a.accountSettings.MutedUsers})
	return result.Error
}

// RemoveMutedUser method removes a user from the MutedUsers list
func (a *SettingsService) RemoveMutedUser(userID string, db *gorm.DB) error {
	for i, u := range a.accountSettings.MutedUsers {
		if u == userID {
			a.accountSettings.MutedUsers = append(a.accountSettings.MutedUsers[:i], a.accountSettings.MutedUsers[i+1:]...)
			break
		}
	}
	result := db.Model(a).Updates(models.AccountSettings{MutedUsers: a.accountSettings.MutedUsers})
	return result.Error
}

// CreateConnectedAccount method creates a new connected account
func (ca *SettingsService) CreateConnectedAccount(userID string, provider string, db *gorm.DB) error {
	// ca.accountSettings.UserID = userID
	// ca.Provider = provider
	// ca.ID = GenerateID()
	result := db.Create(&ca)
	return result.Error
}

// GetConnectedAccount method retrieves a connected account based on its ID
func (ca *SettingsService) GetConnectedAccount(accountID string, db *gorm.DB) (*models.ConnectedAccount, error) {
	var account models.ConnectedAccount
	result := db.First(&account, "id = ?", accountID)
	return &account, result.Error
}

// UpdateConnectedAccount method updates a connected account
func (ca *SettingsService) UpdateConnectedAccount(accountID string, updatedAccount models.ConnectedAccount, db *gorm.DB) error {
	result := db.Model(&models.ConnectedAccount{}).Where("id = ?", accountID).Updates(updatedAccount)
	return result.Error
}

// DeleteConnectedAccount method deletes a connected account
func (ca *SettingsService) DeleteConnectedAccount(accountID string, db *gorm.DB) error {
	result := db.Where("id = ?", accountID).Delete(&models.ConnectedAccount{})
	return result.Error
}

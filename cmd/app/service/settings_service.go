package service

import (
	"gorm.io/gorm"
	"placio-app/models"
)

type UserSettingsService interface {
	GetGeneralSettings(userID string) (*models.GeneralSettings, error)
	UpdateUserSettings(userID string, settings *models.GeneralSettings) error
}

type NotificationsSettingsService interface {
	GetNotificationsSettings(userID string) (*models.NotificationsSettings, error)
	UpdateNotificationsSettings(userID string, settings *models.NotificationsSettings) error
}

type AccountSettingsService interface {
	//GetAccountSettings(userID string) (*models.AccountSettings, error)
	UpdateAccountSettings(userID string, settings *models.AccountSettings) error
	ConnectAccount(userID string, connectedAccount *models.ConnectedAccount) error
	DisconnectAccount(userID string, provider string) error
}

type ContentSettingsService interface {
	GetContentSettings(userID string) (*models.ContentSettings, error)
	UpdateContentSettings(userID string, settings *models.ContentSettings) error
}

type ISettingsService interface {
	UserSettingsService
	NotificationsSettingsService
	AccountSettingsService
	ContentSettingsService
}

type SettingsService struct {
	db *gorm.DB
}

func NewSettingsService(db *gorm.DB) *SettingsService {
	return &SettingsService{db: db}
}

func (s *SettingsService) GetGeneralSettings(userID string) (*models.GeneralSettings, error) {
	generalSettings := &models.GeneralSettings{UserID: userID}

	if err := s.db.Preload("ContentSettings").Preload("NotificationsSettings").Preload("AccountSettings").Where("user_id = ?", userID).First(generalSettings).Error; err != nil {
		return nil, err
	}

	return generalSettings, nil
}

func (s *SettingsService) UpdateUserSettings(userID string, settings *models.GeneralSettings) error {
	if err := s.db.Model(&models.GeneralSettings{}).Where("user_id = ?", userID).Updates(settings).Error; err != nil {
		return err
	}
	return nil
}

func (s *SettingsService) GetNotificationsSettings(userID string) (*models.NotificationsSettings, error) {
	generalSettings, err := s.GetGeneralSettings(userID)
	if err != nil {
		return nil, err
	}
	return &generalSettings.Notifications, nil
}

func (s *SettingsService) UpdateNotificationsSettings(userID string, settings *models.NotificationsSettings) error {
	if err := s.db.Model(&models.GeneralSettings{}).Where("user_id = ?", userID).Updates(settings).Error; err != nil {
		return err
	}
	return nil
}

//
//func (s *SettingsService) GetAccountSettings(userID string) (*models.AccountSettings, error) {
//	generalSettings, err := s.GetGeneralSettings(userID)
//	if err != nil {
//		return nil, err
//	}
//	return &generalSettings.Account, nil
//}

func (s *SettingsService) UpdateAccountSettings(userID string, settings *models.AccountSettings) error {
	if err := s.db.Model(&models.GeneralSettings{}).Where("user_id = ?", userID).Updates(settings).Error; err != nil {
		return err
	}
	return nil
}

func (s *SettingsService) ConnectAccount(userID string, connectedAccount *models.ConnectedAccount) error {
	if err := s.db.Model(&models.GeneralSettings{}).Where("user_id = ?", userID).Updates(connectedAccount).Error; err != nil {
		return err
	}
	return nil
}

func (s *SettingsService) DisconnectAccount(userID string, provider string) error {
	//if err := s.db.Model(&GeneralSettings{}).Where("user_id = ?", userID).Updates(connectedAccount).Error; err != nil {
	//	return err
	//}
	return nil
}

func (s *SettingsService) GetContentSettings(userID string) (*models.ContentSettings, error) {
	generalSettings, err := s.GetGeneralSettings(userID)
	if err != nil {
		return nil, err
	}
	return &generalSettings.Content, nil
}

func (s *SettingsService) UpdateContentSettings(userID string, settings *models.ContentSettings) error {
	if err := s.db.Model(&models.GeneralSettings{}).Where("user_id = ?", userID).Updates(settings).Error; err != nil {
		return err
	}
	return nil
}

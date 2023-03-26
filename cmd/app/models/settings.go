package models

import (
	"gorm.io/gorm"
	"time"
)

type GeneralSettings struct {
	//gorm.Model
	CreatedAt     time.Time  `gorm:"column:created_at"`
	UpdatedAt     time.Time  `gorm:"column:updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at"`
	ID            string     `gorm:"primaryKey"`
	Language      string
	Theme         string
	UserID        string `gorm:"ForeignKey:ID"`
	Privacy       string `gorm:"default:'public'"`
	Notifications NotificationsSettings
	Account       AccountSettings
	Content       ContentSettings
}

// BeforeCreate OnCreateGeneralSettings /*
func (g *GeneralSettings) BeforeCreate(tx *gorm.DB) error {
	g.ID = GenerateID()
	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()
	return nil
}

type NotificationsSettings struct {
	EmailNotifications         bool
	PushNotifications          bool
	DirectMessageNotifications bool
	LikeNotifications          bool
	CommentNotifications       bool
	MentionNotifications       bool
	FollowNotifications        bool
}

type AccountSettings struct {
	TwoFactorAuthentication bool
	ConnectedAccounts       []ConnectedAccount
	BlockedUsers            []string
	MutedUsers              []string
}

type ConnectedAccount struct {
	Provider string
	UserID   string
}

type ContentSettings struct {
	MediaVisibility       string
	ExplicitContentFilter string
	DefaultPostPrivacy    string
	AutoplayVideos        bool
	DisplaySensitiveMedia bool
}

type UserSettingsService interface {
	GetGeneralSettings(userID string) (*GeneralSettings, error)
	UpdateUserSettings(userID string, settings *GeneralSettings) error
}

type NotificationsSettingsService interface {
	GetNotificationsSettings(userID string) (*NotificationsSettings, error)
	UpdateNotificationsSettings(userID string, settings *NotificationsSettings) error
}

type AccountSettingsService interface {
	GetAccountSettings(userID string) (*AccountSettings, error)
	UpdateAccountSettings(userID string, settings *AccountSettings) error
	ConnectAccount(userID string, connectedAccount *ConnectedAccount) error
	DisconnectAccount(userID string, provider string) error
}

type ContentSettingsService interface {
	GetContentSettings(userID string) (*ContentSettings, error)
	UpdateContentSettings(userID string, settings *ContentSettings) error
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

func (s *SettingsService) GetGeneralSettings(userID string) (*GeneralSettings, error) {
	generalSettings := &GeneralSettings{UserID: userID}

	if err := s.db.Where("user_id = ?", userID).First(generalSettings).Error; err != nil {
		return nil, err
	}

	return generalSettings, nil
}

func (s *SettingsService) UpdateUserSettings(userID string, settings *GeneralSettings) error {
	if err := s.db.Model(&GeneralSettings{}).Where("user_id = ?", userID).Updates(settings).Error; err != nil {
		return err
	}
	return nil
}

func (s *SettingsService) GetNotificationsSettings(userID string) (*NotificationsSettings, error) {
	generalSettings, err := s.GetGeneralSettings(userID)
	if err != nil {
		return nil, err
	}
	return &generalSettings.Notifications, nil
}

func (s *SettingsService) UpdateNotificationsSettings(userID string, settings *NotificationsSettings) error {
	if err := s.db.Model(&GeneralSettings{}).Where("user_id = ?", userID).Updates(settings).Error; err != nil {
		return err
	}
	return nil
}

func (s *SettingsService) GetAccountSettings(userID string) (*AccountSettings, error) {
	generalSettings, err := s.GetGeneralSettings(userID)
	if err != nil {
		return nil, err
	}
	return &generalSettings.Account, nil
}

func (s *SettingsService) UpdateAccountSettings(userID string, settings *AccountSettings) error {
	if err := s.db.Model(&GeneralSettings{}).Where("user_id = ?", userID).Updates(settings).Error; err != nil {
		return err
	}
	return nil
}

func (s *SettingsService) ConnectAccount(userID string, connectedAccount *ConnectedAccount) error {
	if err := s.db.Model(&GeneralSettings{}).Where("user_id = ?", userID).Updates(connectedAccount).Error; err != nil {
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

func (s *SettingsService) GetContentSettings(userID string) (*ContentSettings, error) {
	generalSettings, err := s.GetGeneralSettings(userID)
	if err != nil {
		return nil, err
	}
	return &generalSettings.Content, nil
}

func (s *SettingsService) UpdateContentSettings(userID string, settings *ContentSettings) error {
	if err := s.db.Model(&GeneralSettings{}).Where("user_id = ?", userID).Updates(settings).Error; err != nil {
		return err
	}
	return nil
}

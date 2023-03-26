package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
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
	UserID        string                `gorm:"ForeignKey:ID"`
	Privacy       string                `gorm:"default:'public'"`
	Notifications NotificationsSettings `gorm:"ForeignKey:ID"`
	//NotificationID string
	//AccountID      string
	Account AccountSettings `gorm:"ForeignKey:ID"`
	//ContentID      string
	Content ContentSettings `gorm:"ForeignKey:ID"`
}

// BeforeCreate OnCreateGeneralSettings /*
func (g *GeneralSettings) BeforeCreate(tx *gorm.DB) error {
	g.ID = GenerateID()
	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()
	return nil
}

type NotificationsSettings struct {
	ID                         string `gorm:"primaryKey"`
	UserID                     string `gorm:"ForeignKey:ID"`
	EmailNotifications         bool
	PushNotifications          bool
	DirectMessageNotifications bool
	LikeNotifications          bool
	CommentNotifications       bool
	MentionNotifications       bool
	FollowNotifications        bool
}

type AccountSettings struct {
	ID                      string `gorm:"primaryKey"`
	UserID                  string `gorm:"ForeignKey:ID"`
	TwoFactorAuthentication bool
	//ConnectedAccounts       []ConnectedAccount `gorm:"type:json,ForeignKey:ID"`
	BlockedUsers []string `gorm:"type:json"`
	MutedUsers   []string `gorm:"type:json"`
}

type ConnectedAccount struct {
	ID       string `gorm:"primaryKey"`
	Provider string
	UserID   string
}

type ContentSettings struct {
	ID                    string `gorm:"primaryKey"`
	MediaVisibility       string
	ExplicitContentFilter string
	DefaultPostPrivacy    string
	AutoplayVideos        bool
	DisplaySensitiveMedia bool
	UserID                string `gorm:"ForeignKey:ID"`
}

type StringSlice []string

func (s *StringSlice) Scan(src interface{}) error {
	if src == nil {
		*s = nil
		return nil
	}

	source, ok := src.([]byte)
	if !ok {
		return errors.New("incompatible type for StringSlice")
	}

	var tempSlice []string
	err := json.Unmarshal(source, &tempSlice)
	if err != nil {
		return err
	}

	*s = StringSlice(tempSlice)
	return nil
}

func (s *StringSlice) Value() (driver.Value, error) {
	if s == nil {
		return nil, nil
	}

	return nil, nil
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

	if err := s.db.Preload("ContentSettings").Preload("NotificationsSettings").Preload("AccountSettings").Where("user_id = ?", userID).First(generalSettings).Error; err != nil {
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

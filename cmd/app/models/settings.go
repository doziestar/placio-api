package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"time"
)

type GeneralSettings struct {
	CreatedAt     time.Time  `gorm:"column:created_at"`
	UpdatedAt     time.Time  `gorm:"column:updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at"`
	ID            string     `gorm:"primaryKey"`
	Language      string
	Theme         string
	UserID        string                `gorm:"unique"`
	Privacy       string                `gorm:"default:'public'"`
	Notifications NotificationsSettings `gorm:"foreignKey:UserID;references:ID"`
	Content       ContentSettings       `gorm:"foreignKey:UserID;references:ID"`
}

func (g *GeneralSettings) BeforeCreate(tx *gorm.DB) error {
	//g.ID = GenerateID()
	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()
	return nil
}

type NotificationsSettings struct {
	ID                         string `gorm:"primaryKey"`
	UserID                     string `gorm:"unique"`
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
	AccountID               string `gorm:"unique"`
	TwoFactorAuthentication bool
	BlockedUsers            []string `gorm:"type:json"`
	MutedUsers              []string `gorm:"type:json"`
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
	UserID                string `gorm:"unique"`
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

// CreateGeneralSettings /*
func (g GeneralSettings) CreateGeneralSettings(userID string, db *gorm.DB) error {
	g.UserID = userID
	g.Privacy = "public"
	g.Language = "en"
	g.Theme = "light"
	g.ID = GenerateID()

	result := db.Create(&g)

	return result.Error
}

// CreateNotificationsSettings /*
func (n *NotificationsSettings) createNotificationsSettings(userID string, db *gorm.DB) error {
	n.ID = GenerateID()
	n.UserID = userID
	n.EmailNotifications = true
	n.PushNotifications = true
	n.DirectMessageNotifications = true
	n.LikeNotifications = true
	n.CommentNotifications = true
	n.MentionNotifications = true
	n.FollowNotifications = true
	result := db.Create(&n)
	return result.Error
}

// CreateAccountSettings /*
func (a *AccountSettings) createAccountSettings(accountID string, db *gorm.DB) (*AccountSettings, error) {
	a.ID = GenerateID()
	a.AccountID = accountID
	a.TwoFactorAuthentication = false
	a.BlockedUsers = []string{}
	a.MutedUsers = []string{}
	result := db.Create(&a)
	return a, result.Error
}

func (a *AccountSettings) GetAccountSettings(id string, d *gorm.DB) (*AccountSettings, error) {
	var accountSettings AccountSettings
	result := d.First(&accountSettings, "account_id = ?", id)
	return &accountSettings, result.Error
}

// CreateContentSettings /*
func (c *ContentSettings) createContentSettings(userID string, db *gorm.DB) error {
	c.ID = GenerateID()
	c.UserID = userID
	c.MediaVisibility = "public"
	c.ExplicitContentFilter = "off"
	c.DefaultPostPrivacy = "public"
	c.AutoplayVideos = true
	c.DisplaySensitiveMedia = true
	result := db.Create(&c)
	return result.Error
}

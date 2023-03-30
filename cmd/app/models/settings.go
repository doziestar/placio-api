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

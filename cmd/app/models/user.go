package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

func UnmarshalAuth0UserData(data []byte) (Auth0UserData, error) {
	var r Auth0UserData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Auth0UserData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Auth0UserData struct {
	Blocked           bool     `json:"blocked"`
	EmailVerified     bool     `json:"email_verified"`
	Email             string   `json:"email"`
	PhoneNumber       string   `json:"phone_number"`
	PhoneVerified     bool     `json:"phone_verified"`
	UserMetadata      Metadata `json:"user_metadata"`
	AppMetadata       Metadata `json:"app_metadata"`
	GivenName         string   `json:"given_name"`
	FamilyName        string   `json:"family_name"`
	Name              string   `json:"name"`
	Nickname          string   `json:"nickname"`
	Picture           string   `json:"picture"`
	VerifyEmail       bool     `json:"verify_email"`
	VerifyPhoneNumber bool     `json:"verify_phone_number"`
	Password          string   `json:"password"`
	Connection        string   `json:"connection"`
	ClientID          string   `json:"client_id"`
	Username          string   `json:"username"`
}

func UnmarshalAuth0UserSettings(data []byte) (Auth0UserSettings, error) {
	var r Auth0UserSettings
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Auth0UserSettings) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Auth0UserSettings struct {
	CreatedAt     string        `json:"CreatedAt"`
	UpdatedAt     string        `json:"UpdatedAt"`
	DeletedAt     interface{}   `json:"DeletedAt"`
	ID            string        `json:"ID"`
	Language      string        `json:"Language"`
	Theme         string        `json:"Theme"`
	UserID        string        `json:"UserID"`
	Privacy       string        `json:"Privacy"`
	Notifications Notifications `json:"Notifications"`
	Content       Content       `json:"Content"`
}

type Content struct {
	ID                    string `json:"ID"`
	MediaVisibility       string `json:"MediaVisibility"`
	ExplicitContentFilter string `json:"ExplicitContentFilter"`
	DefaultPostPrivacy    string `json:"DefaultPostPrivacy"`
	AutoplayVideos        bool   `json:"AutoplayVideos"`
	DisplaySensitiveMedia bool   `json:"DisplaySensitiveMedia"`
	UserID                string `json:"UserID"`
}

type Notifications struct {
	ID                         string `json:"ID"`
	UserID                     string `json:"UserID"`
	EmailNotifications         bool   `json:"EmailNotifications"`
	PushNotifications          bool   `json:"PushNotifications"`
	DirectMessageNotifications bool   `json:"DirectMessageNotifications"`
	LikeNotifications          bool   `json:"LikeNotifications"`
	CommentNotifications       bool   `json:"CommentNotifications"`
	MentionNotifications       bool   `json:"MentionNotifications"`
	FollowNotifications        bool   `json:"FollowNotifications"`
}

type Metadata struct {
	AppMetadata *Auth0UserData `json:"app_metadata"`
}

type User struct {
	// gorm.Model
	UserID        string `gorm:"primaryKey"`
	Auth0ID       string
	Relationships []UserBusinessRelationship `gorm:"foreignKey:UserID"`
	// Settings      GeneralSettings            `gorm:"foreignKey:UserID"`
	Posts     []Post         `gorm:"foreignKey:UserID"`
	Active    bool           `gorm:"default:false"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeleteAt  gorm.DeletedAt `gorm:"index"`
}

type BusinessAccount struct {
	// gorm.Model
	ID              string `gorm:"primaryKey"`
	Name            string
	Relationships   []UserBusinessRelationship `gorm:"foreignKey:BusinessAccountID"`
	AccountSettings AccountSettings            `gorm:"foreignKey:BusinessAccountID"`
	Posts           []Post                     `gorm:"foreignKey:BusinessAccountID"`
	Active          bool                       `gorm:"default:false"`
	CreatedAt       time.Time                  `gorm:"autoCreateTime"`
	UpdatedAt       time.Time                  `gorm:"autoUpdateTime"`
	DeleteAt        gorm.DeletedAt             `gorm:"index"`
}

type UserBusinessRelationship struct {
	ID                string `gorm:"primaryKey"`
	UserID            string
	User              User
	BusinessAccountID string
	BusinessAccount   BusinessAccount
	Role              string
	CreatedAt         time.Time      `gorm:"autoCreateTime"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime"`
	DeleteAt          gorm.DeletedAt `gorm:"index"`
}

type Invitation struct {
	gorm.Model
	Email             string
	Role              string
	BusinessAccountID string
	BusinessAccount   BusinessAccount
	CreatedAt         time.Time      `gorm:"autoCreateTime"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime"`
	DeleteAt          gorm.DeletedAt `gorm:"index"`
}

package models

import "time"

type GoogleAccount struct {
	AccessToken  string
	RefreshToken string
	UserID       string `gorm:"column:user_id"`
	Email        string
	DateCreated  time.Time
}

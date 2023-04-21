package models

import "time"

type TwitterAccount struct {
	AccessToken  string
	RefreshToken string
	UserID       string `gorm:"column:user_id"`
	UserName     string `gorm:"column:user_name"`
	CodeVerifier string `gorm:"column:code_verifier"`
	State        string
	Name         string
	DateCreated  time.Time `gorm:"column:date_created"`
	ExpiresIn    time.Time `gorm:"column:expires_in"`
}

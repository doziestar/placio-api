package models

import "time"

type Notification struct {
	ID         uint `gorm:"primary_key"`
	Message    string
	UserID     uint
	BusinessID uint
	CreatedAt  time.Time
	UpdatedAt  time.Time

	User     User
	Business Business
}

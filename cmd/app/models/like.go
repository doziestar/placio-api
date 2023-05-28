package models

import "time"

type Like struct {
	ID        string    `gorm:"primaryKey,unique"`
	UserId    string    `gorm:"index"`
	PostID    string    `gorm:"index"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

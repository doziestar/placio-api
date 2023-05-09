package models

import "time"

type Comment struct {
	ID        string    `gorm:"primaryKey,unique"`
	AccountID string    `gorm:"index;foreignKey:ID"`
	PostID    string    `gorm:"index"`
	Content   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

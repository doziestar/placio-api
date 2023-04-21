package models

import "time"

type Media struct {
	ID        string    `gorm:"primaryKey,unique"`
	PostID    string    `gorm:"index"`
	URL       string    `gorm:"type:text"`
	MediaType string    `gorm:"type:text"` // "image", "gif", or "video"
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

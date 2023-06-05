package models

import "time"

type Post struct {
	ID string `gorm:"primaryKey,unique"`
	// AccountID string    `gorm:"index;foreignKey:ID"`
	UserID    string    `gorm:"index"`
	BusinessAccountID string `gorm:"index"`
	Content   string    `gorm:"type:text"`
	Medias    []Media   `gorm:"foreignKey:PostID"`
	Comments  []Comment `gorm:"foreignKey:PostID"`
	Likes     []Like    `gorm:"foreignKey:PostID"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

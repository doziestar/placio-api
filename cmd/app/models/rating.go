package models

import "time"

type Rating struct {
	ID        string `gorm:"primaryKey,unique"`
	EventID   string
	UserID    string
	Score     int
	Review    string
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

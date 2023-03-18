package models

import "time"

type Notification struct {
	ID         string `gorm:"primary_key"`
	Message    string
	UserID     string
	BusinessID string
	CreatedAt  time.Time
	UpdatedAt  time.Time

	// User     User
	// Business Business
}

func (n *Notification) TableName() string {
	return "notifications"
}

func (n *Notification) BeforeCreate() {
	n.ID = GenerateID()
	n.CreatedAt = time.Now()
	n.UpdatedAt = time.Now()
}

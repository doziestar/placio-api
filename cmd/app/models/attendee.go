package models

import "time"

type Attendee struct {
	ID        string `gorm:"primaryKey,unique"`
	EventID   string
	UserID    string
	TicketID  string
	Attended  bool
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

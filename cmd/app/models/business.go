package models

import (
	"time"

	"gorm.io/gorm"
)

type Business struct {
	*gorm.Model
	ID          string `gorm:"primary_key"`
	Name        string
	Description string `gorm:"type:text"`
	// CreatedAt   time.Time
	// UpdatedAt   time.Time

	// Events        []Event
	// Notifications []Notification
	// Tickets       []Ticket
	// Bookings      []Booking
	// Payments      []Payment
	// Users         []User
}

func (b *Business) TableName() string {
	return "business"
}

func (b *Business) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = GenerateID()
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	return nil
}

func (b *Business) GetID() string {
	return b.ID
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	*gorm.Model
	ID          string `gorm:"primaryKey"`
	Name        string `gorm:"index"`
	Date        time.Time
	Time        time.Time
	Location    string
	Description string    `gorm:"type:text"`
	// BusinessID  string    `gorm:"foreignKey:Business"`
	// CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	// Business    Business  `gorm:"foreignKey:BusinessID"`
	// Bookings    []Booking `gorm:"foreignKey:EventID"`
	// Tickets     []Ticket  `gorm:"foreignKey:EventID"`
}

func (e *Event) TableName() string {
	return "events"
}

func (e *Event) GetID() string {
	return e.ID
}

func BuildEvent(event *Event, businessID string) *Event {
	return &Event{
		ID:          event.ID,
		Name:        event.Name,
		Date:        event.Date,
		Time:        event.Time,
		Location:    event.Location,
		Description: event.Description,
		// BusinessID:  businessID,
	}
}

func (event *Event) BeforeCreate(tx *gorm.DB) (err error) {
	event.ID = GenerateID()
	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()
	return nil
}

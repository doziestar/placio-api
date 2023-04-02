package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	*gorm.Model
	ID             string `gorm:"primaryKey"`
	Name           string `gorm:"index"`
	Date           time.Time
	Time           time.Time
	EndDate        time.Time
	EndTime        time.Time
	Location       string
	Address        string
	City           string
	State          string
	Country        string
	Description    string `gorm:"type:text"`
	Category       string
	Tags           []string `gorm:"type:text[]"`
	ImageURL       string
	Organizer      string
	OrganizerEmail string
	OrganizerPhone string
	Website        string
	TicketURL      string
	PriceRange     string
	Capacity       int
	IsFree         bool
	IsPublic       bool
	IsOnline       bool
	AccountID      string    `gorm:"column:account_id"` // Foreign key to Account
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
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

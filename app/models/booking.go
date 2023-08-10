package models

import (
	"gorm.io/gorm"
	"time"
)

type Booking struct {
	*gorm.Model
	ID        string    `gorm:"primaryKey"`
	UserID    string    `gorm:"foreignKey:User"`
	EventID   string    `gorm:"foreignKey:Event"`
	TicketID  string    `gorm:"foreignKey:Ticket"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	// User      User      `gorm:"foreignKey:UserID"`
	// Event     Event     `gorm:"foreignKey:EventID"`
	// Ticket    Ticket    `gorm:"foreignKey:TicketID"`
	// Payment   Payment   `gorm:"foreignKey:BookingID"`
}

func (b *Booking) TableName() string {
	return "bookings"
}

func (b *Booking) GetID() string {
	return b.ID
}

func BuildBooking(booking *Booking, userID, eventID, ticketID string) *Booking {
	return &Booking{
		ID: booking.ID,
		// UserID:   userID,
		// EventID:  eventID,
		// TicketID: ticketID,
	}
}

func (b *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = GenerateID()
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	return nil
}

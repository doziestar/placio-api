package models

import "time"

type Ticket struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Price     int
	EventID   uint
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Event     Event     `gorm:"foreignKey:EventID"`
	Bookings  []Booking `gorm:"foreignKey:TicketID"`
}

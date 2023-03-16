package models

import "time"

type Ticket struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Price     int
	EventID   string
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	// Event     Event     `gorm:"foreignKey:EventID"`
	// Bookings  []Booking `gorm:"foreignKey:TicketID"`
}

func (t *Ticket) TableName() string {
	return "tickets"
}

func (t *Ticket) GetID() string {
	return t.ID
}

func BuildTicket(ticket *Ticket, eventID string) *Ticket {
	return &Ticket{
		ID:      ticket.ID,
		Name:    ticket.Name,
		Price:   ticket.Price,
		EventID: eventID,
	}
}
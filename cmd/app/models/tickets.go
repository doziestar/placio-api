package models

import "time"

type Ticket struct {
	ID            string `gorm:"primaryKey"`
	Name          string
	Price         int
	EventID       string
	CreatedAt     time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	TicketOptions []TicketOption `gorm:"foreignKey:EventID"`
	Attendees     []Attendee     `gorm:"foreignKey:EventID"`
	Comments      []Comment      `gorm:"foreignKey:EventID"`
	Ratings       []Rating       `gorm:"foreignKey:EventID"`
	// Event     Event     `gorm:"foreignKey:EventID"`
	// Bookings  []Booking `gorm:"foreignKey:TicketID"`
}

type TicketOption struct {
	ID        string `gorm:"primaryKey,unique"`
	EventID   string
	Name      string
	Price     float64
	Quantity  int
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
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

package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID        string `gorm:"primaryKey"`
	UserID    string
	BookingID string
	Amount    int
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	// User      User      `gorm:"foreignKey:UserID"`
	// Booking   *Booking  `gorm:"foreignKey:BookingID"`
}

func (p *Payment) TableName() string {
	return "payments"
}

func (p *Payment) GetID() string {
	return p.ID
}

func BuildPayment(payment *Payment, userID, bookingID string) *Payment {
	return &Payment{
		ID:        payment.ID,
		UserID:    userID,
		BookingID: bookingID,
		Amount:    payment.Amount,
	}
}

func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = GenerateID()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return nil
}

package models

import "time"

type Payment struct {
	ID        string `gorm:"primaryKey"`
	UserID    string
	BookingID string
	Amount    int
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	User      User      `gorm:"foreignKey:UserID"`
	Booking   *Booking  `gorm:"foreignKey:BookingID"`
}

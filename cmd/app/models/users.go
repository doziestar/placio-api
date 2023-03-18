package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"placio-app/Dto"
)

type User struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Profile  []Profile `gorm:"foreignKey:UserID"`
	// MessagesSent       []Message      `gorm:"foreignKey:SenderID"`
	// MessagesReceived   []Message      `gorm:"foreignKey:RecipientID"`
	// Conversations      []Conversation `gorm:"many2many:conversation_participant"`
	// Groups             []Group        `gorm:"many2many:group_membership"`
	// VoiceNotesSent     []VoiceNote    `gorm:"foreignKey:SenderID"`
	// VoiceNotesReceived []VoiceNote    `gorm:"foreignKey:RecipientID"`
	// Notifications      []Notification
	// Bookings           []Booking
	// Payments           []Payment
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return nil
}

func (u *User) ComparePassword(password string) bool {
	return u.Password == password
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) GetID() string {
	return u.ID
}

// GenerateToken generates a new token for the user
func (u *User) GenerateToken(user User) (Dto.Token, error) {
	return Dto.Token{}, nil
}

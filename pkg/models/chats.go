package models

import (
	"gorm.io/gorm"
	"time"
)

type Conversation struct {
	gorm.Model
	LastMessageTimestamp time.Time
	Participants         []*User `gorm:"many2many:conversation_participant;"`
	Messages             []Message
}

type Group struct {
	gorm.Model
	Name        string
	Description string
	Members     []*User `gorm:"many2many:group_membership;"`
	Messages    []Message
}

type Message struct {
	gorm.Model
	SenderID       uint
	RecipientID    uint
	ConversationID uint
	GroupID        uint
	MessageText    string
	MessageType    string
	Timestamp      time.Time
	Delivered      bool
	Read           bool
	Sender         User `gorm:"foreignKey:SenderID"`
	Recipient      User `gorm:"foreignKey:RecipientID"`
	Conversation   Conversation
	Group          Group
}

type VoiceNote struct {
	gorm.Model
	SenderID      uint
	RecipientID   uint
	VoiceNoteFile string
	Timestamp     time.Time
	Delivered     bool
	Read          bool
	Sender        User `gorm:"foreignKey:SenderID"`
	Recipient     User `gorm:"foreignKey:RecipientID"`
}

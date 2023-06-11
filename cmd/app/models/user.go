package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	UserID        string `gorm:"primaryKey"`
	Auth0ID       string
	Relationships []UserBusinessRelationship `gorm:"foreignKey:UserID"`
	Posts         []Post                     `gorm:"foreignKey:UserID"`
	CreatedAt     time.Time                  `gorm:"autoCreateTime"`
	UpdatedAt     time.Time                  `gorm:"autoUpdateTime"`
	DeleteAt      gorm.DeletedAt             `gorm:"index"`
	//management.User
}

type BusinessAccount struct {
	// gorm.Model
	ID              string `gorm:"primaryKey"`
	Name            string
	Relationships   []UserBusinessRelationship `gorm:"foreignKey:BusinessAccountID"`
	AccountSettings AccountSettings            `gorm:"foreignKey:BusinessAccountID"`
	Posts           []Post                     `gorm:"foreignKey:BusinessAccountID"`
	Active          bool                       `gorm:"default:false"`
	CreatedAt       time.Time                  `gorm:"autoCreateTime"`
	UpdatedAt       time.Time                  `gorm:"autoUpdateTime"`
	DeleteAt        gorm.DeletedAt             `gorm:"index"`
}

type UserBusinessRelationship struct {
	ID                string `gorm:"primaryKey"`
	UserID            string
	User              User
	BusinessAccountID string
	BusinessAccount   BusinessAccount
	Role              string
	CreatedAt         time.Time      `gorm:"autoCreateTime"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime"`
	DeleteAt          gorm.DeletedAt `gorm:"index"`
}

type Invitation struct {
	gorm.Model
	Email             string
	Role              string
	BusinessAccountID string
	BusinessAccount   BusinessAccount
	CreatedAt         time.Time      `gorm:"autoCreateTime"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime"`
	DeleteAt          gorm.DeletedAt `gorm:"index"`
}

package models

import "gorm.io/gorm"

type AccountType string

const (
	UserAccount     AccountType = "user"
	BusinessAccount AccountType = "business"
	Admin           AccountType = "admin"
)

// Profile is a struct that represents a user's profile
type Profile struct {
	ID       string `gorm:"primaryKey"`
	Phone    string
	UserID   string `gorm:"foreignKey:User,constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CoverPic string `gorm:"default:''"`
	Avatar   string `gorm:"default:''"`
	// tyoe should be an enum of user, business, admin
	Type AccountType `gorm:"default:'user'"` // user, business, admin
}

// TableName returns the name of the table
func (p *Profile) TableName() string {
	return "profiles"
}

// GetID returns the ID of the profile
func (p *Profile) GetID() string {
	return p.ID
}

// BuildProfile builds a profile
func BuildProfile(profile *Profile, userID string) *Profile {
	return &Profile{
		ID:     profile.ID,
		Phone:  profile.Phone,
		UserID: userID,
	}
}

// BeforeCreate is a hook that is called before creating a profile
func (p *Profile) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = GenerateID()
	return nil
}

package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Token represents the token data in the database
type Token struct {
	ID       string `gorm:"primaryKey"`
	Provider string `gorm:"not null"`
	Jwt      string
	Access   string
	Refresh  string
	UserID   string `gorm:"not null"`
}

// Save creates new or updates an existing token
func (t *Token) Save(provider string, data map[string]string, user string, db *gorm.DB) error {
	if data["access"] != "" {
		// data["access"] = crypto.Encrypt(data["access"])
	}
	if data["refresh"] != "" {
		// data["refresh"] = crypto.Encrypt(data["refresh"])
	}

	// is there already a token for this provider?
	var tokenData Token
	if err := db.Where("provider = ? AND user_id = ?", provider, user).First(&tokenData).Error; err == nil {
		// update existing token
		tokenData.Jwt = data["jwt"]
		tokenData.Access = data["access"]
		tokenData.Refresh = data["refresh"]
		return db.Save(&tokenData).Error
	}

	// create a new token
	newToken := Token{
		ID:       uuid.NewString(),
		Provider: provider,
		Jwt:      data["jwt"],
		Access:   data["access"],
		Refresh:  data["refresh"],
		UserID:   user,
	}
	return db.Create(&newToken).Error
}

// Get returns the token for the given user and provider
func (t *Token) Get(id, provider, user string, skipDecryption bool, db *gorm.DB) ([]Token, error) {
	var data []Token
	query := db.Where("user_id = ?", user)
	if id != "" {
		query = query.Where("id = ?", id)
	}
	if provider != "" {
		query = query.Where("provider = ?", provider)
	}
	if err := query.Find(&data).Error; err != nil {
		return nil, err
	}

	if !skipDecryption {
		for i := range data {
			if data[i].Access != "" {
				// data[i].Access = crypto.Decrypt(data[i].Access)
			}
			if data[i].Refresh != "" {
				// data[i].Refresh = crypto.Decrypt(data[i].Refresh)
			}
		}
	}

	return data, nil
}

// Verify checks if a token is present for the given user and provider
func (t *Token) Verify(provider, user string, db *gorm.DB) bool {
	var count int64
	db.Model(&Token{}).Where("provider = ? AND user_id = ?", provider, user).Count(&count)
	return count > 0
}

// Delete deletes the token with the given ID, provider, and user
func (t *Token) Delete(id, provider, user string, db *gorm.DB) error {
	query := db.Where("user_id = ?", user)
	if id != "" {
		query = query.Where("id = ?", id)
	}
	if provider != "" {
		query = query.Where("provider = ?", provider)
	}
	return query.Delete(&Token{}).Error
}

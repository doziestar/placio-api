package models

import (
	"context"
	"placio-pkg/logger"
	"time"

	"gorm.io/gorm"
)

type UserAndAccount struct {
	User
	Account
}

type Account struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Permission  string
	AccountType string
	AccountID   string
	Onboarded   bool
	// Interests                  []string `gorm:"type:text[]"`
	UserID                     string `gorm:"column:user_id"`
	Plan                       string
	Active                     bool
	StripeCustomerID           string `gorm:"column:stripe_customer_id"`
	StripeSubscriptionID       string `gorm:"column:stripe_subscription_id"`
	StripeSubscriptionStatus   string
	PayStackSubscriptionID     string
	PayStackSubscriptionStatus string
	PayStackCustomerID         string
	Status                     string    `gorm:"column:status"`
	LastActive                 time.Time `gorm:"column:last_active"`
	Disabled                   bool      `gorm:"column:disabled"`
}

// BeforeCreate /*
func (a *Account) BeforeCreate(tx *gorm.DB) error {
	a.ID = GenerateID()
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	return nil
}

// CreateAccount /*
func (a *Account) CreateAccount(userID, permission string, db *gorm.DB) (*Account, error) {
	logger.Info(context.Background(), "Creating account")
	a.Active = true
	a.Permission = permission
	a.UserID = userID
	// a.Interests = []string{}

	result := db.Create(&a)
	return a, result.Error
}

// GetAccount /*
func (a *Account) GetUserAccount(id string) (*UserAndAccount, error) {
	result := db.Where("id = ?", id).First(&a)
	if result.Error != nil {
		return nil, result.Error
	}

	// get owner info
	var userData *User

	result = db.Table("user").
		Select("name, email").
		Where("account_id = ? AND (permission = ? OR permission = ?)",
			id, "owner", "master").
		Scan(&userData)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	if userData != nil {
		userAccount := UserAndAccount{
			User:    *userData,
			Account: *a,
		}
		return &userAccount, nil
	}
	return nil, nil
}

// GetAccount /*
func (a *Account) GetAccount(id string) (*Account, error) {
	logger.Info(context.Background(), "Getting account")
	result := db.Where("id = ?", id).First(&a)
	return a, result.Error
}

// Subscription /*
//func (a *Account) Subscription(id string) (string, error) {
//	result := db.Where("id = ?", id).First(&a)
//	if result.Error != nil {
//		return "", result.Error
//	}
//
//	if a.Plan == "free" {
//		return "active", nil
//	}
//
//	if a.StripeSubscriptionID == "" {
//		return "inactive", nil
//	}
//
//	subscription, err := stripe.Subscription(a.StripeSubID)
//	if err != nil {
//		return "", err
//	}
//
//	status := subscription.Status
//	if status != "active" && status != "trialing" {
//		// update account to inactive
//		db.Model(&a).Update("active", false)
//	}
//
//	return status, nil
//}

// UpdateAccount Update /*
func (a *Account) UpdateAccount(data map[string]interface{}) error {
	result := db.Model(&a).Updates(data)
	return result.Error
}

// DeleteAccount /*
func (a *Account) DeleteAccount(id string) error {

	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// delete all users
	result := tx.Where("account_id = ?", id).Delete(&User{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// delete account
	result = tx.Where("id = ?", id).Delete(&a)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	return tx.Commit().Error
}

func (a *Account) GetSubscription(id string) (*Account, error) {
	result := db.Where("id = ?", id).First(&a)
	return a, result.Error
}

//
//func (a *Account) GetUserAccounts(id string) (interface{}, interface{}) {
//
//}

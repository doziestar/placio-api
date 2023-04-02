package models

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"placio-app/Dto"
	"placio-app/database"
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
	Name        string `gorm:"column:name"`
	// Interests                  []string `gorm:"type:text[]"`
	UserID                     string `gorm:"column:user_id"`
	Plan                       string
	Active                     bool
	StripeCustomerID           string `gorm:"column:stripe_customer_id"`
	StripeSubscriptionID       string `gorm:"column:stripe_subscription_id"`
	StripeSubscriptionStatus   string
	PayStackSubscriptionID     string
	PayStackSubscriptionStatus string
	Selected                   bool
	Default                    bool
	AccountSetting             AccountSettings `gorm:"foreignKey:AccountID"`
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
func (a *Account) CreateAccount(userID, permission, accountType string, db *gorm.DB) (*Account, error) {
	logger.Info(context.Background(), "Creating account")
	id := GenerateID()
	a.Active = true
	a.Permission = permission
	a.UserID = userID
	a.ID = id
	a.AccountID = id
	a.Onboarded = false
	a.AccountType = accountType
	a.AccountSetting = AccountSettings{
		ID:                      GenerateID(),
		AccountID:               id,
		TwoFactorAuthentication: false,
		BlockedUsers:            nil,
		MutedUsers:              nil,
	}
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
func (a *Account) GetAccount(id string, db *gorm.DB) (*Account, error) {
	logger.Info(context.Background(), "Getting account")
	result := db.Preload("AccountSetting").Where("id = ?", id).First(&a)
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

func (a *Account) CreateBusinessAccount(user *User, data Dto.AddAccountDto, d *gorm.DB) (*Account, error) {
	a.Active = true
	a.Permission = "owner"
	a.UserID = user.ID
	a.AccountID = GenerateID()
	a.Onboarded = false
	a.AccountType = "business"
	a.Name = data.AccountName
	a.Status = "active"
	a.LastActive = time.Now()
	a.Disabled = false
	a.AccountSetting = AccountSettings{
		ID:                      GenerateID(),
		AccountID:               a.AccountID,
		TwoFactorAuthentication: false,
		BlockedUsers:            nil,
		MutedUsers:              nil,
	}
	result := d.Create(&a)
	return a, result.Error
}

func (a *Account) GenerateUserAccountResponse(db *gorm.DB) *Dto.UserAccountResponse {
	logger.Info(context.Background(), "Generating user account response")
	// get owner info
	var userData *User
	user, err := userData.GetUserById(a.UserID, database.DB)
	if err != nil {
		sentry.CaptureException(err)
		return nil
	}

	//logger.Info(context.Background(), fmt.Sprintf("user: %+v", user))
	// get the account that matches the user's account id
	var accountData *Account
	for _, account := range user.Accounts {
		// change the active account
		account.Active = false
		db.Save(&account)
		if account.ID == a.ID {
			accountData = &account
		}
	}

	if accountData == nil {
		return nil
	}

	// set the active account
	accountData.Active = true
	db.Save(&accountData)

	//logger.Info(context.Background(), fmt.Sprintf("account is active: %+v", accountData))
	// get the account settings
	//var accountSettingsData *AccountSettings
	//accountSettings, err := accountSettingsData.GetAccountSettings(accountData.ID, db)
	//if err != nil {
	//	sentry.CaptureException(err)
	//	return nil
	//}

	//accountData.AccountSetting = *accountSettings

	logger.Info(context.Background(), fmt.Sprintf("account: %+v", accountData))
	// generate the response
	response := accountData.GenerateUserAccountData(user)

	return response
}

func (a *Account) GenerateUserAccountData(user *User) *Dto.UserAccountResponse {
	return &Dto.UserAccountResponse{
		Name:           user.Name,
		Email:          user.Email,
		Disabled:       user.Disabled,
		SupportEnabled: user.SupportEnabled,
		UserId:         user.ID,
		Account: &Dto.Account{
			ID:          a.ID,
			Permission:  a.Permission,
			AccountType: a.AccountType,
			AccountID:   a.AccountID,
			AccountSetting: Dto.AccountSetting{
				ID:                      a.AccountSetting.ID,
				AccountID:               a.AccountSetting.AccountID,
				TwoFactorAuthentication: a.AccountSetting.TwoFactorAuthentication,
				BlockedUsers:            a.AccountSetting.BlockedUsers,
				MutedUsers:              a.AccountSetting.MutedUsers,
			},
			Onboarded: a.Onboarded,
			//Interests:     a.Interests,
			UserID:   a.UserID,
			Plan:     a.Plan,
			Active:   a.Active,
			Status:   a.Status,
			Disabled: a.Disabled,
		},
	}

}

//
//func (a *Account) GetUserAccounts(id string) (interface{}, interface{}) {
//
//}

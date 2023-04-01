package Dto

type AddAccountDto struct {
	AccountType string `json:"account_type" validate:"required"`
	AccountName string `json:"account_name" validate:"required"`
}

type UserAccountResponse struct {
	Name           string   `json:"name"`
	Email          string   `json:"email"`
	Disabled       bool     `json:"disabled"`
	SupportEnabled bool     `json:"support_enabled"`
	UserId         string   `json:"user_id"`
	Account        *Account `json:"account"`
}

type AccountSetting struct {
	ID                      string `gorm:"primaryKey"`
	AccountID               string `gorm:"unique"`
	TwoFactorAuthentication bool
	BlockedUsers            []string `gorm:"type:json"`
	MutedUsers              []string `gorm:"type:json"`
}

type Account struct {
	ID             string         `json:"ID"`
	Permission     string         `json:"Permission"`
	AccountType    string         `json:"AccountType"`
	AccountID      string         `json:"AccountID"`
	AccountSetting AccountSetting `json:"AccountSetting"`
	Onboarded      bool           `json:"Onboarded"`
	Interests      []string       `json:"Interests"`
	UserID         string         `json:"UserID"`
	Plan           string         `json:"Plan"`
	Active         bool           `json:"Active"`
	Status         string         `json:"Status"`
	Disabled       bool           `json:"Disabled"`
}

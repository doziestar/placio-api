package Dto

import (
	"placio-app/models"
)

type AddAccountDto struct {
	AccountType string `json:"account_type" validate:"required"`
	AccountName string `json:"account_name" validate:"required"`
}

type UserAccountResponse struct {
	Name           string          `json:"name"`
	Email          string          `json:"email"`
	Disabled       bool            `json:"disabled"`
	SupportEnabled bool            `json:"support_enabled"`
	UserId         string          `json:"user_id"`
	Account        *models.Account `json:"account"`
}

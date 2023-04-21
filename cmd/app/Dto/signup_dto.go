package Dto

import (
	"encoding/json"
	"placio-app/errors"
)

type SignUpDto struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8"`
	Username        string `json:"username" validate:"required"`
	Name            string `json:"name" validate:"required"`
	Phone           string `json:"phone" validate:"required"`
	AccountType     string `json:"account_type" validate:"required"`
}

// MarshalJSON is a custom JSON marshaler for SignUpDto.
func (m *SignUpDto) MarshalJSON() ([]byte, error) {
	type Alias SignUpDto
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	})
}

// UnmarshalJSON is a custom JSON unmarshaler for SignUpDto.
func (m *SignUpDto) UnmarshalJSON(data []byte) error {
	type Alias SignUpDto
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

func (m *SignUpDto) IsValid() (SignUpDto, error) {
	if m.Password != m.ConfirmPassword {
		return SignUpDto{}, errors.ErrInvalid
	}
	if m.AccountType != "business" && m.AccountType != "user" {
		return SignUpDto{}, errors.ErrInvalid
	}
	return SignUpDto{
		Email:       m.Email,
		Password:    m.Password,
		Name:        m.Name,
		Phone:       m.Phone,
		AccountType: m.AccountType,
	}, nil

}

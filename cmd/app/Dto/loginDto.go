package Dto

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// MarshalJSON is a custom JSON marshaler for LoginDto.
func (m *LoginDto) MarshalJSON() ([]byte, error) {
	type Alias LoginDto
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	})
}

// UnmarshalJSON is a custom JSON unmarshaler for LoginDto.
func (m *LoginDto) UnmarshalJSON(data []byte) error {
	type Alias LoginDto
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

func (m *LoginDto) IsValid() (interface{}, interface{}) {
	// check data is valid
	if m.Email == "" || m.Password == "" {
		return nil, nil
	}
	return m.Email, m.Password
}

// SigninRequest defines the request body for the signin route
type SigninRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
	Provider string `json:"provider,omitempty"`
	// ProviderID   string `json:"provider_id,omitempty"`
	MagicViewURL string `json:"magic_view_url,omitempty"`
}

// SigninResponse defines the response body for the signin route

func (s *SigninRequest) IsValid() error {
	if s.Email == "" && s.Token == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Email or token is required")
	}
	if s.Email != "" && s.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Password is required")
	}
	return nil
}

// covert SigninRequest to json
func (s *SigninRequest) ToJson() map[string]interface{} {
	mySigninRequestMap := map[string]interface{}{
		"email":    s.Email,
		"password": s.Password,
		"token":    s.Token,
		"provider": s.Provider,
		// "provider_id":    s.ProviderID,
		"magic_view_url": s.MagicViewURL,
	}
	return mySigninRequestMap
}

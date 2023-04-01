package Dto

import (
	"encoding/json"
	"time"
)

type UserResponseDto struct {
	ID                    string        `json:"id"`
	Email                 string        `json:"email"`
	Name                  string        `json:"name"`
	AccessToken           string        `json:"accessToken"`
	AccessTokenExpiresIn  time.Duration `json:"accessTokenExpiresIn"`
	RefreshToken          string        `json:"refreshToken"`
	RefreshTokenExpiresIn time.Duration `json:"refreshTokenExpiresIn"`
}

// MarshalJSON is a custom JSON marshaller for UserResponseDto.
func (m *UserResponseDto) MarshalJSON() ([]byte, error) {
	type Alias UserResponseDto
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	})
}

// UnmarshalJSON is a custom JSON unmarshaler for UserResponseDto.
func (m *UserResponseDto) UnmarshalJSON(data []byte) error {
	type Alias UserResponseDto
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

type UserToken struct {
	UserID           string `json:"UserID"`
	Access           string `json:"Access"`
	AccessExpiresIn  int64  `json:"AccessExpiresIn"`
	Refresh          string `json:"Refresh"`
	RefreshExpiresIn int64  `json:"RefreshExpiresIn"`
}

type User struct {
	ID                   string          `json:"ID"`
	Name                 string          `json:"Name"`
	Email                string          `json:"Email"`
	Disabled             bool            `json:"Disabled"`
	HasPassword          bool            `json:"HasPassword"`
	Onboarded            bool            `json:"Onboarded"`
	Account              []Account       `json:"Account"`
	Permission           string          `json:"Permission"`
	CurrentActiveAccount Account         `json:"CurrentActiveAccount"`
	GeneralSettings      GeneralSettings `json:"GeneralSettings"`
}

type GeneralSettings struct {
	ID       string `json:"ID"`
	Language string `json:"Language"`
	Theme    string `json:"Theme"`
}

type UserResponse struct {
	User  *User      `json:"user"`
	Token *UserToken `json:"token"`
}

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

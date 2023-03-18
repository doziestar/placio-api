package Dto

import (
	"encoding/json"
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

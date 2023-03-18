package Dto

import "encoding/json"

type SignUpDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Role     string `json:"role" validate:"required"`
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

func (m *SignUpDto) IsValid() (interface{}, interface{}) {
	return m.Email, m.Password
}

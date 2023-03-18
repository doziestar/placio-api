package Dto

import "encoding/json"

type UserResponseDto struct {
	ID    uint     `json:"id"`
	Email string   `json:"email"`
	Name  string   `json:"name"`
	Token TokenDto `json:"token"`
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

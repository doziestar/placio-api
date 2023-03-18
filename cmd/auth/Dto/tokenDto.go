package Dto

import "encoding/json"

type TokenDto struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// MarshalJSON is a custom JSON marshaler for TokenDto.
func (m *TokenDto) MarshalJSON() ([]byte, error) {
	type Alias TokenDto
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	})
}

// UnmarshalJSON is a custom JSON unmarshaler for TokenDto.
func (m *TokenDto) UnmarshalJSON(data []byte) error {
	type Alias TokenDto
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

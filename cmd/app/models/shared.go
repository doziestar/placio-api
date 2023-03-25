package models

import "github.com/google/uuid"

func GenerateID() string {
	uid, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	return uid.String()
}

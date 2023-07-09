package Dto

import "placio-app/ent"

type BusinessDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	CoverImage  string `json:"cover_image"`
	Website     string `json:"website"`
	Location    string `json:"location"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
}

type BusinessAccountPlacesAndEvents struct {
	Places []*ent.Place `json:"places"`
	Events []*ent.Event `json:"events"`
}

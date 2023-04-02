package Dto

type EventDto struct {
	Name string `json:"name" validate:"required"`
	Date string `json:"date" validate:"required"`
	Time string `json:"time" validate:"required"`
}

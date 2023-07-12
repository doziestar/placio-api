package Dto

import "time"

type EventDTO struct {
	ID                   string
	Name                 string
	EventType            string
	Status               string
	Location             string
	URL                  string
	Title                string
	TimeZone             string
	StartTime            time.Time
	EndTime              time.Time
	StartDate            string
	EndDate              string
	Frequency            string
	FrequencyInterval    string
	FrequencyDayOfWeek   string
	FrequencyDayOfMonth  string
	FrequencyMonthOfYear string
	VenueType            string
	VenueName            string
	VenueAddress         string
	VenueCity            string
	VenueState           string
	VenueCountry         string
	VenueZIP             string
	VenueLat             string
	VenueLon             string
	VenueURL             string
	VenuePhone           string
	VenueEmail           string
	Tags                 []string
	Description          string
	EventSettings        map[string]interface{}
	CoverImage           string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

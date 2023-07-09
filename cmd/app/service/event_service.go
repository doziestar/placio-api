package service

import (
	"context"
	"errors"
	"fmt"
	"placio-app/Dto"
	"placio-app/ent"
	"placio-app/ent/event"
	"placio-app/models"
	"time"
)

type IEventService interface {
	CreateEvent(ctx context.Context, businessId string, data Dto.EventDTO) (*ent.Event, error)
	UpdateEvent(ctx context.Context, eventId string, businessId string, data Dto.EventDTO) (*ent.Event, error)
	//GetEventByID(eventId string) (*models.Event, error)
	//GetEventByLocation(locationId string) (*[]models.Event, error)
	//GetEventByCategory(categoryId string) (*[]models.Event, error)
	//GetEventByDate(date string) (*[]models.Event, error)
	//DeleteEvent(eventId string) error
	//UpdateEvent(eventId string, data *Dto.EventDto) (*models.Event, error)
	//GetEventParticipants(eventId string) error
	//GetEventsByAccount(accountID string) ([]models.Event, error)
}

type EventService struct {
	client        *ent.Client
	searchService SearchService

	// account *models.Account
}

func NewEventService(client *ent.Client, searchService SearchService) *EventService {
	return &EventService{client: client, searchService: searchService}
}

func (s *EventService) CreateEvent(ctx context.Context, businessId string, data Dto.EventDTO) (*ent.Event, error) {
	// get the user from the context
	user := ctx.Value("user").(string)
	// get user from database
	userEnt, err := s.client.User.Get(ctx, user)
	if err != nil {
		return nil, err
	}

	var businessEnt *ent.Business

	// get business from database
	if businessId != "" {
		businessEnt, err = s.client.Business.Get(ctx, businessId)
		if err != nil {
			return nil, err
		}
	}

	typeEnum, err := parseEventType(data.EventType)
	if err != nil {
		return nil, err
	}
	frequencyEnum, err := parseFrequencyType(data.Frequency)
	if err != nil {
		return nil, err
	}
	venueTypeEnum, err := parseVenueType(data.VenueType)
	if err != nil {
		return nil, err
	}

	event, err := s.client.Event.
		Create().
		SetID(data.ID).
		SetName(data.Name).
		SetEventType(typeEnum).
		SetStatus(data.Status).
		SetLocation(data.Location).
		SetURL(data.URL).
		SetTitle(data.Title).
		SetTimeZone(data.TimeZone).
		SetStartTime(data.StartTime).
		SetEndTime(data.EndTime).
		SetStartDate(data.StartDate).
		SetEndDate(data.EndDate).
		SetFrequency(frequencyEnum).
		SetFrequencyInterval(data.FrequencyInterval).
		SetFrequencyDayOfWeek(data.FrequencyDayOfWeek).
		SetFrequencyDayOfMonth(data.FrequencyDayOfMonth).
		SetFrequencyMonthOfYear(data.FrequencyMonthOfYear).
		SetVenueType(venueTypeEnum).
		SetVenueName(data.VenueName).
		SetVenueAddress(data.VenueAddress).
		SetVenueCity(data.VenueCity).
		SetVenueState(data.VenueState).
		SetVenueCountry(data.VenueCountry).
		//SetVenueZIP(data.VenueZIP).
		SetVenueLat(data.VenueLat).
		SetVenueLon(data.VenueLon).
		SetVenueURL(data.VenueURL).
		SetVenuePhone(data.VenuePhone).
		SetVenueEmail(data.VenueEmail).
		SetTags(data.Tags).
		SetDescription(data.Description).
		SetEventSettings(data.EventSettings).
		SetCoverImage(data.CoverImage).
		SetCreatedAt(data.CreatedAt).
		SetUpdatedAt(data.UpdatedAt).
		SetOwnerUser(userEnt).
		SetOwnerBusiness(businessEnt).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (s *EventService) UpdateEvent(ctx context.Context, eventId string, businessId string, data Dto.EventDTO) (*ent.Event, error) {
	// get the user from the context
	user := ctx.Value("user").(string)

	event, err := s.client.Event.Get(ctx, eventId)
	if err != nil {
		return nil, err
	}

	// check if the business is the owner of the event
	if businessId != "" && event.Edges.OwnerBusiness.ID != businessId {
		return nil, errors.New("unauthorized: You can only update events that you own")
	}

	// check if the user is the owner of the event
	if event.Edges.OwnerUser.ID != user {
		return nil, errors.New("unauthorized: You can only update events that you own")
	}

	typeEnum, err := parseEventType(data.EventType)
	if err != nil {
		return nil, err
	}
	frequencyEnum, err := parseFrequencyType(data.Frequency)
	if err != nil {
		return nil, err
	}
	venueTypeEnum, err := parseVenueType(data.VenueType)
	if err != nil {
		return nil, err
	}

	upd := s.client.Event.UpdateOne(event)

	if data.Name != "" {
		upd.SetName(data.Name)
	}

	upd.SetEventType(typeEnum).
		SetStatus(data.Status).
		SetLocation(data.Location).
		SetURL(data.URL).
		SetTitle(data.Title).
		SetTimeZone(data.TimeZone).
		SetStartTime(data.StartTime).
		SetEndTime(data.EndTime).
		SetStartDate(data.StartDate).
		SetEndDate(data.EndDate).
		SetFrequency(frequencyEnum).
		SetFrequencyInterval(data.FrequencyInterval).
		SetFrequencyDayOfWeek(data.FrequencyDayOfWeek).
		SetFrequencyDayOfMonth(data.FrequencyDayOfMonth).
		SetFrequencyMonthOfYear(data.FrequencyMonthOfYear).
		SetVenueType(venueTypeEnum).
		SetVenueName(data.VenueName).
		SetVenueAddress(data.VenueAddress).
		SetVenueCity(data.VenueCity).
		SetVenueState(data.VenueState).
		SetVenueCountry(data.VenueCountry).
		//SetVenueZIP(data.VenueZIP).
		SetVenueLat(data.VenueLat).
		SetVenueLon(data.VenueLon).
		SetVenueURL(data.VenueURL).
		SetVenuePhone(data.VenuePhone).
		SetVenueEmail(data.VenueEmail).
		SetTags(data.Tags).
		SetDescription(data.Description).
		SetCoverImage(data.CoverImage).
		SetUpdatedAt(time.Now())

	// Merge the existing and new settings.
	newSettings := data.EventSettings
	for k, value := range event.EventSettings {
		if _, exists := newSettings[k]; !exists {
			newSettings[k] = value
		}
	}

	event, err = upd.Save(ctx)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (s *EventService) GetEventByID(eventId string) (*models.Event, error) {
	return nil, nil
}

func (s *EventService) GetEventByLocation(locationId string) (*[]models.Event, error) {
	return nil, nil
}

func (s *EventService) GetEventByCategory(categoryId string) (*[]models.Event, error) {
	return nil, nil
}

func (s *EventService) GetEventByDate(date string) (*[]models.Event, error) {
	return nil, nil
}

func (s *EventService) DeleteEvent(eventId string) error {
	return nil
}

func (s *EventService) GetEventParticipants(eventId string) error {
	return nil
}

//func (s *EventService) GetEventsByAccount(accountID string) ([]models.Event, error) {
//	var events []models.Event
//	err := s.db.Where("account_id = ?", accountID).Find(&events).Error
//	return events, err
//}

func parseEventType(s string) (event.EventType, error) {
	switch s {
	case "event":
		return event.EventTypeEvent, nil
	case "place":
		return event.EventTypePlace, nil
	case "business":
		return event.EventTypeBusiness, nil
	default:
		return "", fmt.Errorf("invalid EventType: %s", s)
	}
}

func parseFrequencyType(s string) (event.Frequency, error) {
	switch s {
	case "once":
		return event.FrequencyOnce, nil
	case "daily":
		return event.FrequencyDaily, nil
	case "weekly":
		return event.FrequencyWeekly, nil
	case "monthly":
		return event.FrequencyMonthly, nil
	case "yearly":
		return event.FrequencyYearly, nil
	default:
		return "", fmt.Errorf("invalid FrequencyType: %s", s)
	}
}

func parseVenueType(s string) (event.VenueType, error) {
	switch s {
	case "online":
		return event.VenueTypeOnline, nil
	case "in_person":
		return event.VenueTypeInPerson, nil
	case "hybrid":
		return event.VenueTypeHybrid, nil
	default:
		return "", fmt.Errorf("invalid VenueType: %s", s)
	}
}

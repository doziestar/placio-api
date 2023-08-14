package events_management

import (
	"context"
	"errors"
	"fmt"
	"log"
	"placio-app/domains/search"
	"placio-app/ent"
	"placio-app/ent/event"
	"strings"
	"time"
)

type EventFilter struct {
	EventType string
	Status    string
	Location  string
	Title     string
	TimeZone  string
	StartDate struct {
		From string
		To   string
	}
	EndDate struct {
		From string
		To   string
	}
	StartTime struct {
		From time.Time
		To   time.Time
	}
	EndTime struct {
		From time.Time
		To   time.Time
	}
}

type IEventService interface {
	CreateEvent(ctx context.Context, businessId string, data EventDTO) (*ent.Event, error)
	UpdateEvent(ctx context.Context, eventId string, businessId string, data EventDTO) (*ent.Event, error)
	GetEventByID(ctx context.Context, id string) (*ent.Event, error)
	DeleteEvent(ctx context.Context, eventId string) error
	GetEvents(ctx context.Context, filter *EventFilter, page int, pageSize int) ([]*ent.Event, error)
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
	searchService search.SearchService

	// account *models.Account
}

func NewEventService(client *ent.Client, searchService search.SearchService) *EventService {
	return &EventService{client: client, searchService: searchService}
}

func (s *EventService) CreateEvent(ctx context.Context, businessId string, data EventDTO) (*ent.Event, error) {
	// get the user from the context
	user := ctx.Value("user").(string)
	// get user from db
	userEnt, err := s.client.User.Get(ctx, user)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}

	var businessEnt *ent.Business

	// get business from db
	if businessId != "" {
		businessEnt, err = s.client.Business.Get(ctx, businessId)
		if err != nil {
			log.Println("error: ", err)
			return nil, err
		}
	}

	log.Println("data.EventType", data.EventType)

	//typeEnum, err := parseEventType(data.EventType)
	//if err != nil {
	//	log.Println("error: ", err)
	//	return nil, err
	//}
	//frequencyEnum, err := parseFrequencyType(data.Frequency)
	//if err != nil {
	//	log.Println("error: ", err)
	//	return nil, err
	//}
	//venueTypeEnum, err := parseVenueType(data.VenueType)
	//if err != nil {
	//	log.Println("error: ", err)
	//	return nil, err
	//}

	event, err := s.client.Event.
		Create().
		SetID(data.ID).
		SetName(data.Name).
		//SetEventType(typeEnum).
		SetStatus(data.Status).
		SetLocation(data.Location).
		SetURL(data.URL).
		SetTitle(data.Title).
		SetTimeZone(data.TimeZone).
		SetStartTime(data.StartTime).
		SetEndTime(data.EndTime).
		SetStartDate(data.StartDate).
		SetEndDate(data.EndDate).
		//SetFrequency(frequencyEnum).
		SetFrequencyInterval(data.FrequencyInterval).
		SetFrequencyDayOfWeek(data.FrequencyDayOfWeek).
		SetFrequencyDayOfMonth(data.FrequencyDayOfMonth).
		SetFrequencyMonthOfYear(data.FrequencyMonthOfYear).
		//SetVenueType(venueTypeEnum).
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
		//SetVenueCapacity(data.VenueCapacity).
		// TODO: SetTags(data.Tags).
		//SetTags(data.Tags).
		SetDescription(data.Description).
		SetEventSettings(data.EventSettings).
		SetCoverImage(data.CoverImage).
		SetCreatedAt(data.CreatedAt).
		SetUpdatedAt(data.UpdatedAt).
		SetOwnerUser(userEnt).
		SetOwnerBusiness(businessEnt).
		Save(ctx)

	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}

	return event, nil
}

func (s *EventService) UpdateEvent(ctx context.Context, eventId string, businessId string, data EventDTO) (*ent.Event, error) {
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
		//SetTags(data.Tags).
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

func (s *EventService) GetEventByID(ctx context.Context, id string) (*ent.Event, error) {
	// Get event by ID
	event, err := s.client.
		Event.
		Query().
		Where(event.IDEQ(id)).
		WithOwnerUser().
		WithOwnerBusiness().
		WithFaqs().
		WithTickets().
		First(ctx)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (s *EventService) DeleteEvent(ctx context.Context, eventId string) error {
	// Try to delete the event from the db.
	err := s.client.Event.
		DeleteOneID(eventId).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (s *EventService) GetEvents(ctx context.Context, filter *EventFilter, page int, pageSize int) ([]*ent.Event, error) {
	query := s.client.Event.
		Query().
		WithOwnerUser().
		WithOwnerBusiness()

	//// Apply filters
	//if filter.EventType != "" {
	//	eventTypeEnum, err := parseEventType(filter.EventType)
	//	if err != nil {
	//		return nil, err
	//	}
	//	query = query.Where(event.HasEventTypeWith(event.EventTypeEqual(eventTypeEnum)))
	//}
	//if filter.Status != "" {
	//	query = query.Where(event.StatusEqual(filter.Status))
	//}
	//if filter.Location != "" {
	//	query = query.Where(event.LocationEqual(filter.Location))
	//}
	//if filter.Title != "" {
	//	query = query.Where(event.TitleEqual(filter.Title))
	//}
	//// ... You can add similar checks for other string fields ...
	//
	//// Apply date ranges
	//if filter.StartDate.From != "" {
	//	from, err := time.Parse("2006-01-02", filter.StartDate.From)
	//	if err == nil {
	//		query = query.Where(event.StartDateGTE(from))
	//	}
	//}
	//if filter.StartDate.To != "" {
	//	to, err := time.Parse("2006-01-02", filter.StartDate.To)
	//	if err == nil {
	//		query = query.Where(event.StartDateLTE(to))
	//	}
	//}
	//
	//if filter.EndDate.From != "" {
	//	from, err := time.Parse("2006-01-02", filter.EndDate.From)
	//	if err == nil {
	//		query = query.Where(event.EndDateGTE(from))
	//	}
	//}
	//if filter.EndDate.To != "" {
	//	to, err := time.Parse("2006-01-02", filter.EndDate.To)
	//	if err == nil {
	//		query = query.Where(event.EndDateLTE(to))
	//	}
	//}

	// Apply time ranges
	if !filter.StartTime.From.IsZero() {
		query = query.Where(event.StartTimeGTE(filter.StartTime.From))
	}
	if !filter.StartTime.To.IsZero() {
		query = query.Where(event.StartTimeLTE(filter.StartTime.To))
	}
	if !filter.EndTime.From.IsZero() {
		query = query.Where(event.EndTimeGTE(filter.EndTime.From))
	}
	if !filter.EndTime.To.IsZero() {
		query = query.Where(event.EndTimeLTE(filter.EndTime.To))
	}

	// Apply pagination
	query = query.Offset((page - 1) * pageSize).Limit(pageSize)

	// Execute query
	events, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func parseEventType(s string) (event.EventType, error) {
	if s == "" {
		return "", nil
	}
	switch strings.ToLower(s) {
	case strings.ToLower(string(event.EventTypeEvent)):
		return event.EventTypeEvent, nil
	case strings.ToLower(string(event.EventTypePlace)):
		return event.EventTypePlace, nil
	case strings.ToLower(string(event.EventTypeBusiness)):
		return event.EventTypeBusiness, nil
	default:
		return "", fmt.Errorf("invalid EventType: %s", s)
	}
}

func parseFrequencyType(s string) (event.Frequency, error) {
	if s == "" {
		return "", nil
	}
	switch strings.ToLower(s) {
	case strings.ToLower(string(event.FrequencyOnce)):
		return event.FrequencyOnce, nil
	case strings.ToLower(string(event.FrequencyDaily)):
		return event.FrequencyDaily, nil
	case strings.ToLower(string(event.FrequencyWeekly)):
		return event.FrequencyWeekly, nil
	case strings.ToLower(string(event.FrequencyMonthly)):
		return event.FrequencyMonthly, nil
	case strings.ToLower(string(event.FrequencyYearly)):
		return event.FrequencyYearly, nil
	default:
		return "", fmt.Errorf("invalid FrequencyType: %s", s)
	}
}

func parseVenueType(s string) (event.VenueType, error) {
	if s == "" {
		return "", nil
	}
	switch strings.ToLower(s) {
	case strings.ToLower(string(event.VenueTypeOnline)):
		return event.VenueTypeOnline, nil
	case strings.ToLower(string(event.VenueTypeInPerson)):
		return event.VenueTypeInPerson, nil
	case strings.ToLower(string(event.VenueTypeHybrid)):
		return event.VenueTypeHybrid, nil
	default:
		return "", fmt.Errorf("invalid VenueType: %s", s)
	}
}

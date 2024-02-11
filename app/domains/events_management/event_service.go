package events_management

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"placio-app/domains/search"
	"placio-app/ent"
	"placio-app/ent/event"
	"placio-app/ent/eventorganizer"
	"strings"
	"time"
)

type IEventService interface {
	CreateEvent(ctx context.Context, businessId string, data *ent.Event) (*ent.Event, error)
	AddOrganizers(ctx context.Context, eventID string, organizers []OrganizerInput) error
	UpdateEvent(ctx context.Context, eventId string, businessId string, data EventDTO) (*ent.Event, error)
	GetEventByID(ctx context.Context, id string) (*ent.Event, error)
	DeleteEvent(ctx context.Context, eventId string) error
	GetEvents(ctx context.Context, filter *EventFilter, page int, pageSize int) ([]*ent.Event, error)
	CheckInAttendee(ctx context.Context, eventId string, attendeeId string, method CheckInMethod) error
	ManageTicketing(ctx context.Context, eventId string, ticketDetails *ent.Ticket) error
	ManageAttendees(ctx context.Context, eventId string, attendeeDetails *ent.User) error
	RunAdsOnEvent(ctx context.Context, eventId string, adDetails AdDTO) error
	UpdateEventMedia(ctx context.Context, eventId string, mediaDetails *ent.Media) error
	GenerateEventAnalytics(ctx context.Context, eventId string) (*EventAnalyticsDTO, error)
	SynchronizeEventWithSocialMedia(ctx context.Context, eventId string, platforms []SocialPlatform) error
	SendEventNotifications(ctx context.Context, eventId string, notification EventNotificationDTO) error
	HandleOnSiteTools(ctx context.Context, eventId string, tools OnSiteToolsDTO) error
	EnforceComplianceRules(ctx context.Context, eventId string, rules ComplianceRulesDTO) error
	// AdvancedTicketManagement Advanced ticket management, including seat reservations and special pricing
	AdvancedTicketManagement(ctx context.Context, eventId string, ticketOptions AdvancedTicketOptions) error

	// PersonalizeAttendeeExperience Personalized attendee experience based on historical data and preferences
	PersonalizeAttendeeExperience(ctx context.Context, eventId string, attendeeId string, preferences PersonalizationPreferences) error

	// InteractiveVenueMap Interactive venue map with real-time updates and attendee tracking
	InteractiveVenueMap(ctx context.Context, eventId string) (*VenueMapDetails, error)

	// GamifyEventExperience Gamification of the event experience, including rewards and leaderboards
	GamifyEventExperience(ctx context.Context, eventId string, gamificationOptions GamificationOptions) error

	// FacilitateAttendeeNetworking Real-time chat and networking facilitation for attendees
	FacilitateAttendeeNetworking(ctx context.Context, eventId string, networkingOptions NetworkingOptions) error

	// AdvancedAnalyticsAndPredictions Advanced analytics with predictive modelling and actionable insights
	AdvancedAnalyticsAndPredictions(ctx context.Context, eventId string) (*AdvancedAnalyticsDTO, error)

	// ManageVendorsAndSponsors Vendor and sponsor management, including booths and sponsored sessions
	ManageVendorsAndSponsors(ctx context.Context, eventId string, vendorDetails VendorDTO) error

	// IncidentReportingAndResponse Real-time incident reporting and response system for event management
	IncidentReportingAndResponse(ctx context.Context, eventId string, incidentDetails IncidentReportDTO) error

	// CustomEventAppCreation Custom event app creation with features specific to each event
	CustomEventAppCreation(ctx context.Context, eventId string, appFeatures CustomAppFeatures) (*CustomAppDetails, error)

	// IntegrateWithExternalServices Integration with external services and APIs for additional functionalities
	IntegrateWithExternalServices(ctx context.Context, eventId string, integrationDetails IntegrationDetailsDTO) error

	// LoyaltyAndRewardsProgram Loyalty and rewards program for frequent attendees
	LoyaltyAndRewardsProgram(ctx context.Context, eventId string, loyaltyOptions LoyaltyOptionsDTO) error

	// ProvideAccessibilityServices Real-time accessibility services for attendees with disabilities
	ProvideAccessibilityServices(ctx context.Context, eventId string, accessibilityOptions AccessibilityOptionsDTO) error

	// MultiLanguageSupport Multi-language support for international attendees
	MultiLanguageSupport(ctx context.Context, eventId string, languageOptions LanguageSupportDTO) error
}

type EventService struct {
	client        *ent.Client
	searchService search.SearchService
}

func (s *EventService) CheckInAttendee(ctx context.Context, eventId string, attendeeId string, method CheckInMethod) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) ManageTicketing(ctx context.Context, eventId string, ticketDetails *ent.Ticket) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) ManageAttendees(ctx context.Context, eventId string, attendeeDetails *ent.User) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) RunAdsOnEvent(ctx context.Context, eventId string, adDetails AdDTO) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) UpdateEventMedia(ctx context.Context, eventId string, mediaDetails *ent.Media) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) GenerateEventAnalytics(ctx context.Context, eventId string) (*EventAnalyticsDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) SynchronizeEventWithSocialMedia(ctx context.Context, eventId string, platforms []SocialPlatform) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) SendEventNotifications(ctx context.Context, eventId string, notification EventNotificationDTO) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) HandleOnSiteTools(ctx context.Context, eventId string, tools OnSiteToolsDTO) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) EnforceComplianceRules(ctx context.Context, eventId string, rules ComplianceRulesDTO) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) AdvancedTicketManagement(ctx context.Context, eventId string, ticketOptions AdvancedTicketOptions) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) PersonalizeAttendeeExperience(ctx context.Context, eventId string, attendeeId string, preferences PersonalizationPreferences) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) InteractiveVenueMap(ctx context.Context, eventId string) (*VenueMapDetails, error) {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) GamifyEventExperience(ctx context.Context, eventId string, gamificationOptions GamificationOptions) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) FacilitateAttendeeNetworking(ctx context.Context, eventId string, networkingOptions NetworkingOptions) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) AdvancedAnalyticsAndPredictions(ctx context.Context, eventId string) (*AdvancedAnalyticsDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) ManageVendorsAndSponsors(ctx context.Context, eventId string, vendorDetails VendorDTO) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) IncidentReportingAndResponse(ctx context.Context, eventId string, incidentDetails IncidentReportDTO) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) CustomEventAppCreation(ctx context.Context, eventId string, appFeatures CustomAppFeatures) (*CustomAppDetails, error) {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) IntegrateWithExternalServices(ctx context.Context, eventId string, integrationDetails IntegrationDetailsDTO) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) LoyaltyAndRewardsProgram(ctx context.Context, eventId string, loyaltyOptions LoyaltyOptionsDTO) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) ProvideAccessibilityServices(ctx context.Context, eventId string, accessibilityOptions AccessibilityOptionsDTO) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventService) MultiLanguageSupport(ctx context.Context, eventId string, languageOptions LanguageSupportDTO) error {
	//TODO implement me
	panic("implement me")
}

func NewEventService(client *ent.Client, searchService search.SearchService) *EventService {
	return &EventService{client: client, searchService: searchService}
}

const OrganizerContextKey ContextKey = "organizers"

func (s *EventService) CreateEvent(ctx context.Context, userID string, data *ent.Event) (*ent.Event, error) {
	// Extract organizers from context
	organizers, ok := ctx.Value(OrganizerContextKey).([]OrganizerInfo)
	if !ok {
		return nil, errors.New("organizer information missing")
	}

	// Start a transaction
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	// Use eventData directly to create the event, assuming SetOwnerUserID or similar linkage
	event, err := tx.Event.Create().
		SetName(data.Name).
		SetID(uuid.NewString()).
		SetNillableName(&data.Name). // SetNillable used for optional fields
		SetNillableEventType(&data.EventType).
		SetNillableStatus(&data.Status).
		SetNillableLocation(&data.Location).
		SetNillableURL(&data.URL).
		SetNillableTitle(&data.Title).
		SetNillableTimeZone(&data.TimeZone).
		SetNillableStartTime(&data.StartTime).
		SetNillableEndTime(&data.EndTime).
		SetNillableStartDate(&data.StartDate).
		SetNillableEndDate(&data.EndDate).
		//SetNillableFrequency(event.Frequency(data.Frequency)).
		SetNillableFrequencyInterval(&data.FrequencyInterval).
		SetNillableFrequencyDayOfWeek(&data.FrequencyDayOfWeek).
		SetNillableFrequencyDayOfMonth(&data.FrequencyDayOfMonth).
		SetNillableFrequencyMonthOfYear(&data.FrequencyMonthOfYear).
		//SetNillableVenueType(event.VenueType(data.VenueType)).
		SetNillableVenueName(&data.VenueName).
		SetNillableVenueAddress(&data.VenueAddress).
		SetNillableVenueCity(&data.VenueCity).
		SetNillableVenueState(&data.VenueState).
		//SetNillableVenueCountry(data.VenueCountry).
		SetNillableVenueZip(&data.VenueZip).
		SetNillableVenueLat(&data.VenueLat).
		SetNillableVenueLon(&data.VenueLon).
		SetNillableVenueURL(&data.VenueURL).
		SetNillableVenuePhone(&data.VenuePhone).
		SetNillableVenueEmail(&data.VenueEmail).
		SetTags(data.Tags).
		SetNillableDescription(&data.Description).
		//SetNillableEventSettings(data.EventSettings).
		SetCoverImage(data.CoverImage).
		SetCreatedAt(data.CreatedAt).
		SetUpdatedAt(data.UpdatedAt).
		//SetNillableMapCoordinates(data.MapCoordinates).
		SetNillableLongitude(&data.Longitude).
		SetNillableLatitude(&data.Latitude).
		SetNillableSearchText(&data.SearchText).
		SetNillableRelevanceScore(&data.RelevanceScore).
		SetFollowerCount(data.FollowerCount).
		SetFollowingCount(data.FollowingCount).
		SetIsPremium(data.IsPremium).
		SetIsPublished(data.IsPublished).
		SetIsOnline(data.IsOnline).
		SetIsFree(data.IsFree).
		SetIsPaid(data.IsPaid).
		SetIsPublic(data.IsPublic).
		SetIsOnlineOnly(data.IsOnlineOnly).
		SetIsInPersonOnly(data.IsInPersonOnly).
		SetIsHybrid(data.IsHybrid).
		SetIsOnlineAndInPerson(data.IsOnlineAndInPerson).
		SetIsOnlineAndInPersonOnly(data.IsOnlineAndInPersonOnly).
		SetIsOnlineAndInPersonOrHybrid(data.IsOnlineAndInPersonOrHybrid).
		SetLikedByCurrentUser(data.LikedByCurrentUser).
		SetFollowedByCurrentUser(data.FollowedByCurrentUser).
		SetNillableRegistrationType(&data.RegistrationType).
		SetNillableRegistrationURL(&data.RegistrationURL).
		SetIsPhysicallyAccessible(data.IsPhysicallyAccessible).
		SetNillableAccessibilityInfo(&data.AccessibilityInfo).
		SetIsVirtuallyAccessible(data.IsVirtuallyAccessible).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Process each organizer, adding them to the event
	for _, org := range organizers {
		// Assume a method to process and add organizers based on their type
		err := addOrganizerToEvent(tx, event, org)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Attempt to commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return event, nil
}

// Helper function to add an organizer to an event, distinguishing between user and business types
func addOrganizerToEvent(tx *ent.Tx, event *ent.Event, org OrganizerInfo) error {
	switch org.Type {
	case "user":
		_, err := tx.EventOrganizer.Create().
			SetEvent(event).
			SetOrganizerID(org.ID).
			SetOrganizerType("user").
			Save(context.Background()) // Consider passing the original context if available
		return err
	case "business":
		_, err := tx.EventOrganizer.Create().
			SetEvent(event).
			SetOrganizerID(org.ID).
			SetOrganizerType("business").
			Save(context.Background()) // Consider passing the original context if available
		return err
	default:
		return errors.New("invalid organizer type")
	}
}

func (s *EventService) AddOrganizers(ctx context.Context, eventID string, organizers []OrganizerInput) error {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return err
	}

	for _, organizer := range organizers {
		if organizer.OrganizerType != "user" && organizer.OrganizerType != "business" {
			tx.Rollback()
			return errors.New("invalid organizer type")
		}

		_, err := tx.EventOrganizer.
			Create().
			SetOrganizerID(organizer.OrganizerID).
			SetOrganizerType(organizer.OrganizerType).
			SetEventID(eventID).
			Save(ctx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (s *EventService) GetOrganizersForEvent(ctx context.Context, eventID string) ([]interface{}, error) {
	organizers, err := s.client.EventOrganizer.
		Query().
		Where(eventorganizer.HasEventWith(event.ID(eventID))).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var result []interface{}
	for _, org := range organizers {
		switch org.OrganizerType {
		case "user":
			user, err := s.client.User.Get(ctx, org.OrganizerID)
			if err != nil {
				return nil, err
			}
			result = append(result, user)
		case "business":
			business, err := s.client.Business.Get(ctx, org.OrganizerID)
			if err != nil {
				return nil, err
			}
			result = append(result, business)
		default:
			// Handle invalid type
		}
	}
	return result, nil
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

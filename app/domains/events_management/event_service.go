package events_management

import (
	"context"
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"mime/multipart"
	"placio-app/domains/media"
	"placio-app/domains/search"
	"placio-app/ent"
	"placio-app/ent/business"
	"placio-app/ent/event"
	"placio-app/ent/eventorganizer"
	"strings"
	"time"
)

type IEventService interface {
	CreateEvent(ctx context.Context, businessId string, data *EventDTO) (*ent.Event, error)
	GetEventByBusinessID(ctx context.Context, businessID string) ([]*ent.Event, error)
	AddOrganizers(ctx context.Context, eventID string, organizers []OrganizerInfo) error
	RemoveOrganizer(ctx context.Context, eventID string, organizerID string) error
	GetOrganizersForEvent(ctx context.Context, eventID string) ([]interface{}, error)
	GetEventsByOrganizerID(ctx context.Context, organizerId string) ([]*ent.Event, error)
	UpdateEvent(ctx context.Context, eventId string, businessId string, data *EventDTO) (*ent.Event, error)
	GetEventByID(ctx context.Context, id string) (*ent.Event, error)
	DeleteEvent(ctx context.Context, eventId string) error
	AddMediaToEvent(ctx context.Context, eventID string, files []*multipart.FileHeader) (*ent.Event, error)
	RemoveMediaFromEvent(ctx context.Context, eventID string, mediaID string) error
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
	mediaService  media.MediaService
	searchService search.SearchService
}

func NewEventService(client *ent.Client, searchService search.SearchService, mediaService media.MediaService) *EventService {
	return &EventService{client: client, searchService: searchService, mediaService: mediaService}
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

func (s *EventService) CreateEvent(ctx context.Context, userID string, data *EventDTO) (*ent.Event, error) {

	organizers := data.OrganizerInfo
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
		//SetNillableEventType(event.EventType(data.EventType)).
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
		//SetNillableRegistrationType(&data.RegistrationType).
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
		err := addOrganizerToEvent(ctx, tx, event, org)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// add event to business
	_, err = tx.Business.UpdateOneID(userID).
		AddEvents(event).
		Save(ctx)

	// Attempt to commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return event, nil
}

func (s *EventService) GetEventByBusinessID(ctx context.Context, businessID string) ([]*ent.Event, error) {
	events, err := s.client.Event.Query().
		Where(event.HasOwnerBusinessWith(business.ID(businessID))).
		WithEventCategories().
		WithEventCategoryAssignments().
		WithOwnerUser().
		WithOwnerBusiness().
		WithUserFollowers().
		WithBusinessFollowers().
		WithEventOrganizers().
		WithMedia().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return events, nil

}

// Helper function to add an organizer to an event, distinguishing between user and business types
func addOrganizerToEvent(ctx context.Context, tx *ent.Tx, event *ent.Event, org OrganizerInfo) error {
	switch org.Type {
	case "user":
		_, err := tx.EventOrganizer.Create().
			SetEvent(event).
			SetID(uuid.New().String()).
			SetOrganizerID(org.ID).
			SetOrganizerType("user").
			Save(ctx)
		return err
	case "business":
		_, err := tx.EventOrganizer.Create().
			SetEvent(event).
			SetID(uuid.New().String()).
			SetOrganizerID(org.ID).
			SetOrganizerType("business").
			Save(ctx)
		return err
	default:
		return errors.New("invalid organizer type")
	}
}

func (s *EventService) GetEventsByOrganizerID(ctx context.Context, organizerId string) ([]*ent.Event, error) {
	events, err := s.client.EventOrganizer.
		Query().
		Where(eventorganizer.OrganizerID(organizerId)).
		WithEvent().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var result []*ent.Event
	for _, event := range events {
		result = append(result, event.Edges.Event)
	}
	return result, nil
}

func (s *EventService) RemoveMediaFromEvent(ctx context.Context, eventID string, mediaID string) error {
	// Start a transaction
	tx, err := s.client.Tx(ctx)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	// get media
	media, err := tx.Media.Get(ctx, mediaID)
	if err != nil {
		tx.Rollback()
		sentry.CaptureException(err)
		return err
	}

	err = tx.Event.UpdateOneID(eventID).
		RemoveMedia(media).
		Exec(ctx)
	if err != nil {
		tx.Rollback()
		sentry.CaptureException(err)
		return err
	}

	if err := tx.Commit(); err != nil {
		sentry.CaptureException(err)
		return err
	}

	return nil
}

func (s *EventService) AddOrganizers(ctx context.Context, eventID string, organizers []OrganizerInfo) error {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return err
	}

	for _, organizer := range organizers {
		if organizer.Type != "user" && organizer.Type != "business" {
			continue
		}

		exists, err := tx.EventOrganizer.
			Query().
			Where(
				eventorganizer.OrganizerID(organizer.ID),
				eventorganizer.HasEventWith(event.ID(eventID)),
				eventorganizer.OrganizerType(organizer.Type),
			).
			Exist(ctx)
		if err != nil {
			tx.Rollback()
			return err
		}

		if !exists {
			_, err := tx.EventOrganizer.
				Create().
				SetID(uuid.New().String()).
				SetOrganizerID(organizer.ID).
				SetOrganizerType(organizer.Type).
				SetEventID(eventID).
				Save(ctx)
			if err != nil {
				tx.Rollback()
				return err
			}
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

func (s *EventService) RemoveOrganizer(ctx context.Context, eventID string, organizerID string) error {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return err
	}

	_, err = tx.EventOrganizer.Delete().
		Where(
			eventorganizer.HasEventWith(event.ID(eventID)),
			eventorganizer.OrganizerID(organizerID),
		).
		Exec(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s *EventService) UpdateEvent(ctx context.Context, eventId string, businessId string, data *EventDTO) (*ent.Event, error) {
	//userID, exist := ctx.Value("userId").(string)
	//if !exist {
	//	return nil, errors.New("user not found")
	//}

	// Load the event with its owner edges to check ownership.
	event, err := s.client.Event.Query().
		Where(event.IDEQ(eventId)).
		WithOwnerUser().
		WithOwnerBusiness().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	// Authorization checks
	if event.Edges.OwnerBusiness != nil && businessId != event.Edges.OwnerBusiness.ID {
		return nil, errors.New("unauthorized: You can only update events that your business owns")
	}
	//if event.Edges.OwnerUser != nil && userID != event.Edges.OwnerUser.ID {
	//	return nil, errors.New("unauthorized: You can only update events that you own")
	//}

	// Begin updating the event fields
	upd := s.client.Event.UpdateOneID(eventId)

	// Directly set fields if they are not nil in the DTO.
	if data.Name != "" {
		upd.SetName(data.Name)
	}
	if data.EventType != "" {
		// Assuming parseEventType returns an ent.EventType enum and handles conversion.
		if eventType, err := parseEventType(data.EventType); err == nil {
			upd.SetEventType(eventType)
		} else {
			return nil, err
		}
	}
	if data.Status != "" {
		upd.SetStatus(data.Status)
	}
	if data.Location != "" {
		upd.SetLocation(data.Location)
	}
	if data.URL != "" {
		upd.SetURL(data.URL)
	}
	if data.Title != "" {
		upd.SetTitle(data.Title)
	}
	if data.TimeZone != "" {
		upd.SetTimeZone(data.TimeZone)
	}
	if data.StartTime != time.Now() {
		upd.SetStartTime(data.StartTime)
	}
	if data.EndTime != time.Now() {
		upd.SetEndTime(data.EndTime)
	}
	if data.StartDate != "" {
		upd.SetStartDate(data.StartDate)
	}
	if data.EndDate != "" {
		upd.SetEndDate(data.EndDate)
	}
	if data.Frequency != "" {
		frequencyEnum, err := parseFrequencyType(string(data.Frequency))
		if err != nil {
			return nil, err
		}
		upd.SetFrequency(frequencyEnum)
	}
	if data.FrequencyInterval != "" {
		upd.SetFrequencyInterval(data.FrequencyInterval)
	}
	if data.FrequencyDayOfWeek != "" {
		upd.SetFrequencyDayOfWeek(data.FrequencyDayOfWeek)
	}
	if data.FrequencyDayOfMonth != "" {
		upd.SetFrequencyDayOfMonth(data.FrequencyDayOfMonth)
	}
	if data.FrequencyMonthOfYear != "" {
		upd.SetFrequencyMonthOfYear(data.FrequencyMonthOfYear)
	}
	if data.VenueType != "" {
		venueTypeEnum, err := parseVenueType(string(data.VenueType)) // Convert to enum
		if err != nil {
			return nil, err
		}
		upd.SetVenueType(venueTypeEnum)
	}
	if data.VenueName != "" {
		upd.SetVenueName(data.VenueName)
	}
	if data.VenueAddress != "" {
		upd.SetVenueAddress(data.VenueAddress)
	}
	if data.VenueCity != "" {
		upd.SetVenueCity(data.VenueCity)
	}
	if data.VenueState != "" {
		upd.SetVenueState(data.VenueState)
	}
	if data.VenueCountry != "" {
		upd.SetVenueCountry(data.VenueCountry)
	}
	if data.VenueZip != "" {
		upd.SetVenueZip(data.VenueZip)
	}
	if data.VenueLat != "" {
		upd.SetVenueLat(data.VenueLat)
	}
	if data.VenueLon != "" {
		upd.SetVenueLon(data.VenueLon)
	}
	if data.VenueURL != "" {
		upd.SetVenueURL(data.VenueURL)
	}
	if data.VenuePhone != "" {
		upd.SetVenuePhone(data.VenuePhone)
	}
	if data.VenueEmail != "" {
		upd.SetVenueEmail(data.VenueEmail)
	}
	if len(data.Tags) > 0 {
		upd.SetTags(data.Tags)
	}
	if data.Description != "" {
		upd.SetDescription(data.Description)
	}
	if data.CoverImage != "" {
		upd.SetCoverImage(data.CoverImage)
	}
	if data.Longitude != "" {
		upd.SetLongitude(data.Longitude)
	}
	if data.Latitude != "" {
		upd.SetLatitude(data.Latitude)
	}
	if data.SearchText != "" {
		upd.SetSearchText(data.SearchText)
	}
	//if data.RelevanceScore != "" {
	//	upd.SetRelevanceScore(data.RelevanceScore)
	//}
	if data.RegistrationURL != "" {
		upd.SetRegistrationURL(data.RegistrationURL)
	}
	upd.SetIsPremium(data.IsPremium).
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
		SetIsPhysicallyAccessible(data.IsPhysicallyAccessible).
		SetIsVirtuallyAccessible(data.IsVirtuallyAccessible)

	// Execute the update
	updatedEvent, err := upd.Save(ctx)
	if err != nil {
		return nil, err
	}

	return updatedEvent, nil
}

func (s *EventService) GetEventByID(ctx context.Context, id string) (*ent.Event, error) {
	// Get event by ID
	event, err := s.client.
		Event.
		Query().
		Where(event.IDEQ(id)).
		WithOwnerUser().
		WithOwnerBusiness().
		WithEventOrganizers().
		WithMedia().
		WithEventComments().
		WithAdditionalOrganizers().
		WithPlace(func(query *ent.PlaceQuery) {

		}).
		WithPerformers().
		WithFaqs().
		WithTickets().
		First(ctx)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (s *EventService) DeleteEvent(ctx context.Context, eventId string) error {
	err := s.client.Event.
		DeleteOneID(eventId).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

// AddMediaToEvent adds media files to an event. It starts a transaction and fetches the event with the given eventID. Then, it uploads and creates media files using the media service
func (s *EventService) AddMediaToEvent(ctx context.Context, eventID string, files []*multipart.FileHeader) (*ent.Event, error) {
	// Start a transaction
	tx, err := s.client.Tx(ctx)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	// Fetch event within transaction
	eventData, err := tx.Event.Get(ctx, eventID)
	if err != nil {
		sentry.CaptureException(err)
		tx.Rollback()
		return nil, err
	}

	uploadedFiles, err := s.mediaService.UploadAndCreateMedia(ctx, files)
	if err != nil {
		sentry.CaptureException(err)
		tx.Rollback()
		return nil, err
	}

	eventData, err = tx.Event.UpdateOne(eventData).
		AddMedia(uploadedFiles...).
		Save(ctx)
	if err != nil {
		sentry.CaptureException(err)
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	event, err := s.client.Event.Query().
		Where(event.IDEQ(eventID)).
		WithOwnerUser().
		WithOwnerBusiness().
		WithEventOrganizers().
		WithMedia().
		WithEventComments().
		WithAdditionalOrganizers().
		WithPlace(func(query *ent.PlaceQuery) {
			query.WithMedias()
		}).First(ctx)

	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	return event, nil
}

func (s *EventService) GetEvents(ctx context.Context, filter *EventFilter, page int, pageSize int) ([]*ent.Event, error) {
	query := s.client.Event.
		Query().
		WithOwnerUser().
		WithOwnerBusiness()

	if filter.EventType != "" {
		eventType, err := parseEventType(filter.EventType)
		if err != nil {
		}
		query = query.Where(event.EventTypeIn(eventType))
	}
	if filter.Status != "" {
		query = query.Where(event.StatusContainsFold(filter.Status))
	}
	if filter.Location != "" {
		query = query.Where(event.LocationContainsFold(filter.Location))
	}
	if filter.Title != "" {
		query = query.Where(event.TitleContainsFold(filter.Title))
	}
	if filter.TimeZone != "" {
		query = query.Where(event.TimeZoneContainsFold(filter.TimeZone))
	}

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

	query = query.Offset((page - 1) * pageSize).Limit(pageSize)

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
	case strings.ToLower(string(event.EventTypeFree)):
		return event.EventTypeFree, nil
	case strings.ToLower(string(event.EventTypePaid)):
		return event.EventTypePaid, nil
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

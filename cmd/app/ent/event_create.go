// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/businessfollowevent"
	"placio-app/ent/category"
	"placio-app/ent/categoryassignment"
	"placio-app/ent/event"
	"placio-app/ent/faq"
	"placio-app/ent/place"
	"placio-app/ent/rating"
	"placio-app/ent/ticket"
	"placio-app/ent/ticketoption"
	"placio-app/ent/user"
	"placio-app/ent/userfollowevent"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EventCreate is the builder for creating a Event entity.
type EventCreate struct {
	config
	mutation *EventMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ec *EventCreate) SetName(s string) *EventCreate {
	ec.mutation.SetName(s)
	return ec
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ec *EventCreate) SetNillableName(s *string) *EventCreate {
	if s != nil {
		ec.SetName(*s)
	}
	return ec
}

// SetEventType sets the "EventType" field.
func (ec *EventCreate) SetEventType(et event.EventType) *EventCreate {
	ec.mutation.SetEventType(et)
	return ec
}

// SetNillableEventType sets the "EventType" field if the given value is not nil.
func (ec *EventCreate) SetNillableEventType(et *event.EventType) *EventCreate {
	if et != nil {
		ec.SetEventType(*et)
	}
	return ec
}

// SetStatus sets the "status" field.
func (ec *EventCreate) SetStatus(s string) *EventCreate {
	ec.mutation.SetStatus(s)
	return ec
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ec *EventCreate) SetNillableStatus(s *string) *EventCreate {
	if s != nil {
		ec.SetStatus(*s)
	}
	return ec
}

// SetLocation sets the "location" field.
func (ec *EventCreate) SetLocation(s string) *EventCreate {
	ec.mutation.SetLocation(s)
	return ec
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (ec *EventCreate) SetNillableLocation(s *string) *EventCreate {
	if s != nil {
		ec.SetLocation(*s)
	}
	return ec
}

// SetURL sets the "url" field.
func (ec *EventCreate) SetURL(s string) *EventCreate {
	ec.mutation.SetURL(s)
	return ec
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (ec *EventCreate) SetNillableURL(s *string) *EventCreate {
	if s != nil {
		ec.SetURL(*s)
	}
	return ec
}

// SetTitle sets the "title" field.
func (ec *EventCreate) SetTitle(s string) *EventCreate {
	ec.mutation.SetTitle(s)
	return ec
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ec *EventCreate) SetNillableTitle(s *string) *EventCreate {
	if s != nil {
		ec.SetTitle(*s)
	}
	return ec
}

// SetTimeZone sets the "time_zone" field.
func (ec *EventCreate) SetTimeZone(s string) *EventCreate {
	ec.mutation.SetTimeZone(s)
	return ec
}

// SetNillableTimeZone sets the "time_zone" field if the given value is not nil.
func (ec *EventCreate) SetNillableTimeZone(s *string) *EventCreate {
	if s != nil {
		ec.SetTimeZone(*s)
	}
	return ec
}

// SetStartTime sets the "start_time" field.
func (ec *EventCreate) SetStartTime(t time.Time) *EventCreate {
	ec.mutation.SetStartTime(t)
	return ec
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (ec *EventCreate) SetNillableStartTime(t *time.Time) *EventCreate {
	if t != nil {
		ec.SetStartTime(*t)
	}
	return ec
}

// SetEndTime sets the "end_time" field.
func (ec *EventCreate) SetEndTime(t time.Time) *EventCreate {
	ec.mutation.SetEndTime(t)
	return ec
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (ec *EventCreate) SetNillableEndTime(t *time.Time) *EventCreate {
	if t != nil {
		ec.SetEndTime(*t)
	}
	return ec
}

// SetStartDate sets the "start_date" field.
func (ec *EventCreate) SetStartDate(s string) *EventCreate {
	ec.mutation.SetStartDate(s)
	return ec
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (ec *EventCreate) SetNillableStartDate(s *string) *EventCreate {
	if s != nil {
		ec.SetStartDate(*s)
	}
	return ec
}

// SetEndDate sets the "end_date" field.
func (ec *EventCreate) SetEndDate(s string) *EventCreate {
	ec.mutation.SetEndDate(s)
	return ec
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (ec *EventCreate) SetNillableEndDate(s *string) *EventCreate {
	if s != nil {
		ec.SetEndDate(*s)
	}
	return ec
}

// SetFrequency sets the "frequency" field.
func (ec *EventCreate) SetFrequency(e event.Frequency) *EventCreate {
	ec.mutation.SetFrequency(e)
	return ec
}

// SetNillableFrequency sets the "frequency" field if the given value is not nil.
func (ec *EventCreate) SetNillableFrequency(e *event.Frequency) *EventCreate {
	if e != nil {
		ec.SetFrequency(*e)
	}
	return ec
}

// SetFrequencyInterval sets the "frequency_interval" field.
func (ec *EventCreate) SetFrequencyInterval(s string) *EventCreate {
	ec.mutation.SetFrequencyInterval(s)
	return ec
}

// SetNillableFrequencyInterval sets the "frequency_interval" field if the given value is not nil.
func (ec *EventCreate) SetNillableFrequencyInterval(s *string) *EventCreate {
	if s != nil {
		ec.SetFrequencyInterval(*s)
	}
	return ec
}

// SetFrequencyDayOfWeek sets the "frequency_day_of_week" field.
func (ec *EventCreate) SetFrequencyDayOfWeek(s string) *EventCreate {
	ec.mutation.SetFrequencyDayOfWeek(s)
	return ec
}

// SetNillableFrequencyDayOfWeek sets the "frequency_day_of_week" field if the given value is not nil.
func (ec *EventCreate) SetNillableFrequencyDayOfWeek(s *string) *EventCreate {
	if s != nil {
		ec.SetFrequencyDayOfWeek(*s)
	}
	return ec
}

// SetFrequencyDayOfMonth sets the "frequency_day_of_month" field.
func (ec *EventCreate) SetFrequencyDayOfMonth(s string) *EventCreate {
	ec.mutation.SetFrequencyDayOfMonth(s)
	return ec
}

// SetNillableFrequencyDayOfMonth sets the "frequency_day_of_month" field if the given value is not nil.
func (ec *EventCreate) SetNillableFrequencyDayOfMonth(s *string) *EventCreate {
	if s != nil {
		ec.SetFrequencyDayOfMonth(*s)
	}
	return ec
}

// SetFrequencyMonthOfYear sets the "frequency_month_of_year" field.
func (ec *EventCreate) SetFrequencyMonthOfYear(s string) *EventCreate {
	ec.mutation.SetFrequencyMonthOfYear(s)
	return ec
}

// SetNillableFrequencyMonthOfYear sets the "frequency_month_of_year" field if the given value is not nil.
func (ec *EventCreate) SetNillableFrequencyMonthOfYear(s *string) *EventCreate {
	if s != nil {
		ec.SetFrequencyMonthOfYear(*s)
	}
	return ec
}

// SetVenueType sets the "venue_type" field.
func (ec *EventCreate) SetVenueType(et event.VenueType) *EventCreate {
	ec.mutation.SetVenueType(et)
	return ec
}

// SetNillableVenueType sets the "venue_type" field if the given value is not nil.
func (ec *EventCreate) SetNillableVenueType(et *event.VenueType) *EventCreate {
	if et != nil {
		ec.SetVenueType(*et)
	}
	return ec
}

// SetVenueName sets the "venue_name" field.
func (ec *EventCreate) SetVenueName(s string) *EventCreate {
	ec.mutation.SetVenueName(s)
	return ec
}

// SetNillableVenueName sets the "venue_name" field if the given value is not nil.
func (ec *EventCreate) SetNillableVenueName(s *string) *EventCreate {
	if s != nil {
		ec.SetVenueName(*s)
	}
	return ec
}

// SetVenueAddress sets the "venue_address" field.
func (ec *EventCreate) SetVenueAddress(s string) *EventCreate {
	ec.mutation.SetVenueAddress(s)
	return ec
}

// SetNillableVenueAddress sets the "venue_address" field if the given value is not nil.
func (ec *EventCreate) SetNillableVenueAddress(s *string) *EventCreate {
	if s != nil {
		ec.SetVenueAddress(*s)
	}
	return ec
}

// SetVenueCity sets the "venue_city" field.
func (ec *EventCreate) SetVenueCity(s string) *EventCreate {
	ec.mutation.SetVenueCity(s)
	return ec
}

// SetNillableVenueCity sets the "venue_city" field if the given value is not nil.
func (ec *EventCreate) SetNillableVenueCity(s *string) *EventCreate {
	if s != nil {
		ec.SetVenueCity(*s)
	}
	return ec
}

// SetVenueState sets the "venue_state" field.
func (ec *EventCreate) SetVenueState(s string) *EventCreate {
	ec.mutation.SetVenueState(s)
	return ec
}

// SetNillableVenueState sets the "venue_state" field if the given value is not nil.
func (ec *EventCreate) SetNillableVenueState(s *string) *EventCreate {
	if s != nil {
		ec.SetVenueState(*s)
	}
	return ec
}

// SetVenueCountry sets the "venue_country" field.
func (ec *EventCreate) SetVenueCountry(s string) *EventCreate {
	ec.mutation.SetVenueCountry(s)
	return ec
}

// SetNillableVenueCountry sets the "venue_country" field if the given value is not nil.
func (ec *EventCreate) SetNillableVenueCountry(s *string) *EventCreate {
	if s != nil {
		ec.SetVenueCountry(*s)
	}
	return ec
}

// SetVenueZip sets the "venue_zip" field.
func (ec *EventCreate) SetVenueZip(s string) *EventCreate {
	ec.mutation.SetVenueZip(s)
	return ec
}

// SetNillableVenueZip sets the "venue_zip" field if the given value is not nil.
func (ec *EventCreate) SetNillableVenueZip(s *string) *EventCreate {
	if s != nil {
		ec.SetVenueZip(*s)
	}
	return ec
}

// SetVenueLat sets the "venue_lat" field.
func (ec *EventCreate) SetVenueLat(s string) *EventCreate {
	ec.mutation.SetVenueLat(s)
	return ec
}

// SetNillableVenueLat sets the "venue_lat" field if the given value is not nil.
func (ec *EventCreate) SetNillableVenueLat(s *string) *EventCreate {
	if s != nil {
		ec.SetVenueLat(*s)
	}
	return ec
}

// SetVenueLon sets the "venue_lon" field.
func (ec *EventCreate) SetVenueLon(s string) *EventCreate {
	ec.mutation.SetVenueLon(s)
	return ec
}

// SetNillableVenueLon sets the "venue_lon" field if the given value is not nil.
func (ec *EventCreate) SetNillableVenueLon(s *string) *EventCreate {
	if s != nil {
		ec.SetVenueLon(*s)
	}
	return ec
}

// SetVenueURL sets the "venue_url" field.
func (ec *EventCreate) SetVenueURL(s string) *EventCreate {
	ec.mutation.SetVenueURL(s)
	return ec
}

// SetNillableVenueURL sets the "venue_url" field if the given value is not nil.
func (ec *EventCreate) SetNillableVenueURL(s *string) *EventCreate {
	if s != nil {
		ec.SetVenueURL(*s)
	}
	return ec
}

// SetVenuePhone sets the "venue_phone" field.
func (ec *EventCreate) SetVenuePhone(s string) *EventCreate {
	ec.mutation.SetVenuePhone(s)
	return ec
}

// SetNillableVenuePhone sets the "venue_phone" field if the given value is not nil.
func (ec *EventCreate) SetNillableVenuePhone(s *string) *EventCreate {
	if s != nil {
		ec.SetVenuePhone(*s)
	}
	return ec
}

// SetVenueEmail sets the "venue_email" field.
func (ec *EventCreate) SetVenueEmail(s string) *EventCreate {
	ec.mutation.SetVenueEmail(s)
	return ec
}

// SetNillableVenueEmail sets the "venue_email" field if the given value is not nil.
func (ec *EventCreate) SetNillableVenueEmail(s *string) *EventCreate {
	if s != nil {
		ec.SetVenueEmail(*s)
	}
	return ec
}

// SetTags sets the "tags" field.
func (ec *EventCreate) SetTags(s string) *EventCreate {
	ec.mutation.SetTags(s)
	return ec
}

// SetNillableTags sets the "tags" field if the given value is not nil.
func (ec *EventCreate) SetNillableTags(s *string) *EventCreate {
	if s != nil {
		ec.SetTags(*s)
	}
	return ec
}

// SetDescription sets the "description" field.
func (ec *EventCreate) SetDescription(s string) *EventCreate {
	ec.mutation.SetDescription(s)
	return ec
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ec *EventCreate) SetNillableDescription(s *string) *EventCreate {
	if s != nil {
		ec.SetDescription(*s)
	}
	return ec
}

// SetEventSettings sets the "event_settings" field.
func (ec *EventCreate) SetEventSettings(m map[string]interface{}) *EventCreate {
	ec.mutation.SetEventSettings(m)
	return ec
}

// SetCoverImage sets the "cover_image" field.
func (ec *EventCreate) SetCoverImage(s string) *EventCreate {
	ec.mutation.SetCoverImage(s)
	return ec
}

// SetNillableCoverImage sets the "cover_image" field if the given value is not nil.
func (ec *EventCreate) SetNillableCoverImage(s *string) *EventCreate {
	if s != nil {
		ec.SetCoverImage(*s)
	}
	return ec
}

// SetCreatedAt sets the "createdAt" field.
func (ec *EventCreate) SetCreatedAt(t time.Time) *EventCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (ec *EventCreate) SetNillableCreatedAt(t *time.Time) *EventCreate {
	if t != nil {
		ec.SetCreatedAt(*t)
	}
	return ec
}

// SetUpdatedAt sets the "updatedAt" field.
func (ec *EventCreate) SetUpdatedAt(t time.Time) *EventCreate {
	ec.mutation.SetUpdatedAt(t)
	return ec
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (ec *EventCreate) SetNillableUpdatedAt(t *time.Time) *EventCreate {
	if t != nil {
		ec.SetUpdatedAt(*t)
	}
	return ec
}

// SetMapCoordinates sets the "map_coordinates" field.
func (ec *EventCreate) SetMapCoordinates(m map[string]interface{}) *EventCreate {
	ec.mutation.SetMapCoordinates(m)
	return ec
}

// SetLongitude sets the "longitude" field.
func (ec *EventCreate) SetLongitude(s string) *EventCreate {
	ec.mutation.SetLongitude(s)
	return ec
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (ec *EventCreate) SetNillableLongitude(s *string) *EventCreate {
	if s != nil {
		ec.SetLongitude(*s)
	}
	return ec
}

// SetLatitude sets the "latitude" field.
func (ec *EventCreate) SetLatitude(s string) *EventCreate {
	ec.mutation.SetLatitude(s)
	return ec
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (ec *EventCreate) SetNillableLatitude(s *string) *EventCreate {
	if s != nil {
		ec.SetLatitude(*s)
	}
	return ec
}

// SetSearchText sets the "search_text" field.
func (ec *EventCreate) SetSearchText(s string) *EventCreate {
	ec.mutation.SetSearchText(s)
	return ec
}

// SetNillableSearchText sets the "search_text" field if the given value is not nil.
func (ec *EventCreate) SetNillableSearchText(s *string) *EventCreate {
	if s != nil {
		ec.SetSearchText(*s)
	}
	return ec
}

// SetRelevanceScore sets the "relevance_score" field.
func (ec *EventCreate) SetRelevanceScore(f float64) *EventCreate {
	ec.mutation.SetRelevanceScore(f)
	return ec
}

// SetNillableRelevanceScore sets the "relevance_score" field if the given value is not nil.
func (ec *EventCreate) SetNillableRelevanceScore(f *float64) *EventCreate {
	if f != nil {
		ec.SetRelevanceScore(*f)
	}
	return ec
}

// SetID sets the "id" field.
func (ec *EventCreate) SetID(s string) *EventCreate {
	ec.mutation.SetID(s)
	return ec
}

// AddTicketIDs adds the "tickets" edge to the Ticket entity by IDs.
func (ec *EventCreate) AddTicketIDs(ids ...string) *EventCreate {
	ec.mutation.AddTicketIDs(ids...)
	return ec
}

// AddTickets adds the "tickets" edges to the Ticket entity.
func (ec *EventCreate) AddTickets(t ...*Ticket) *EventCreate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ec.AddTicketIDs(ids...)
}

// AddTicketOptionIDs adds the "ticket_options" edge to the TicketOption entity by IDs.
func (ec *EventCreate) AddTicketOptionIDs(ids ...string) *EventCreate {
	ec.mutation.AddTicketOptionIDs(ids...)
	return ec
}

// AddTicketOptions adds the "ticket_options" edges to the TicketOption entity.
func (ec *EventCreate) AddTicketOptions(t ...*TicketOption) *EventCreate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ec.AddTicketOptionIDs(ids...)
}

// AddPlaceIDs adds the "place" edge to the Place entity by IDs.
func (ec *EventCreate) AddPlaceIDs(ids ...string) *EventCreate {
	ec.mutation.AddPlaceIDs(ids...)
	return ec
}

// AddPlace adds the "place" edges to the Place entity.
func (ec *EventCreate) AddPlace(p ...*Place) *EventCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ec.AddPlaceIDs(ids...)
}

// AddEventCategoryIDs adds the "event_categories" edge to the Category entity by IDs.
func (ec *EventCreate) AddEventCategoryIDs(ids ...string) *EventCreate {
	ec.mutation.AddEventCategoryIDs(ids...)
	return ec
}

// AddEventCategories adds the "event_categories" edges to the Category entity.
func (ec *EventCreate) AddEventCategories(c ...*Category) *EventCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ec.AddEventCategoryIDs(ids...)
}

// AddEventCategoryAssignmentIDs adds the "event_category_assignments" edge to the CategoryAssignment entity by IDs.
func (ec *EventCreate) AddEventCategoryAssignmentIDs(ids ...string) *EventCreate {
	ec.mutation.AddEventCategoryAssignmentIDs(ids...)
	return ec
}

// AddEventCategoryAssignments adds the "event_category_assignments" edges to the CategoryAssignment entity.
func (ec *EventCreate) AddEventCategoryAssignments(c ...*CategoryAssignment) *EventCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ec.AddEventCategoryAssignmentIDs(ids...)
}

// SetOwnerUserID sets the "ownerUser" edge to the User entity by ID.
func (ec *EventCreate) SetOwnerUserID(id string) *EventCreate {
	ec.mutation.SetOwnerUserID(id)
	return ec
}

// SetNillableOwnerUserID sets the "ownerUser" edge to the User entity by ID if the given value is not nil.
func (ec *EventCreate) SetNillableOwnerUserID(id *string) *EventCreate {
	if id != nil {
		ec = ec.SetOwnerUserID(*id)
	}
	return ec
}

// SetOwnerUser sets the "ownerUser" edge to the User entity.
func (ec *EventCreate) SetOwnerUser(u *User) *EventCreate {
	return ec.SetOwnerUserID(u.ID)
}

// SetOwnerBusinessID sets the "ownerBusiness" edge to the Business entity by ID.
func (ec *EventCreate) SetOwnerBusinessID(id string) *EventCreate {
	ec.mutation.SetOwnerBusinessID(id)
	return ec
}

// SetNillableOwnerBusinessID sets the "ownerBusiness" edge to the Business entity by ID if the given value is not nil.
func (ec *EventCreate) SetNillableOwnerBusinessID(id *string) *EventCreate {
	if id != nil {
		ec = ec.SetOwnerBusinessID(*id)
	}
	return ec
}

// SetOwnerBusiness sets the "ownerBusiness" edge to the Business entity.
func (ec *EventCreate) SetOwnerBusiness(b *Business) *EventCreate {
	return ec.SetOwnerBusinessID(b.ID)
}

// AddUserFollowerIDs adds the "userFollowers" edge to the UserFollowEvent entity by IDs.
func (ec *EventCreate) AddUserFollowerIDs(ids ...string) *EventCreate {
	ec.mutation.AddUserFollowerIDs(ids...)
	return ec
}

// AddUserFollowers adds the "userFollowers" edges to the UserFollowEvent entity.
func (ec *EventCreate) AddUserFollowers(u ...*UserFollowEvent) *EventCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ec.AddUserFollowerIDs(ids...)
}

// AddBusinessFollowerIDs adds the "businessFollowers" edge to the BusinessFollowEvent entity by IDs.
func (ec *EventCreate) AddBusinessFollowerIDs(ids ...string) *EventCreate {
	ec.mutation.AddBusinessFollowerIDs(ids...)
	return ec
}

// AddBusinessFollowers adds the "businessFollowers" edges to the BusinessFollowEvent entity.
func (ec *EventCreate) AddBusinessFollowers(b ...*BusinessFollowEvent) *EventCreate {
	ids := make([]string, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ec.AddBusinessFollowerIDs(ids...)
}

// AddFaqIDs adds the "faqs" edge to the FAQ entity by IDs.
func (ec *EventCreate) AddFaqIDs(ids ...string) *EventCreate {
	ec.mutation.AddFaqIDs(ids...)
	return ec
}

// AddFaqs adds the "faqs" edges to the FAQ entity.
func (ec *EventCreate) AddFaqs(f ...*FAQ) *EventCreate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ec.AddFaqIDs(ids...)
}

// AddRatingIDs adds the "ratings" edge to the Rating entity by IDs.
func (ec *EventCreate) AddRatingIDs(ids ...string) *EventCreate {
	ec.mutation.AddRatingIDs(ids...)
	return ec
}

// AddRatings adds the "ratings" edges to the Rating entity.
func (ec *EventCreate) AddRatings(r ...*Rating) *EventCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ec.AddRatingIDs(ids...)
}

// Mutation returns the EventMutation object of the builder.
func (ec *EventCreate) Mutation() *EventMutation {
	return ec.mutation
}

// Save creates the Event in the database.
func (ec *EventCreate) Save(ctx context.Context) (*Event, error) {
	if err := ec.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EventCreate) SaveX(ctx context.Context) *Event {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EventCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EventCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EventCreate) defaults() error {
	if _, ok := ec.mutation.CoverImage(); !ok {
		v := event.DefaultCoverImage
		ec.mutation.SetCoverImage(v)
	}
	if _, ok := ec.mutation.CreatedAt(); !ok {
		if event.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized event.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := event.DefaultCreatedAt()
		ec.mutation.SetCreatedAt(v)
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		if event.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized event.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := event.DefaultUpdatedAt()
		ec.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ec *EventCreate) check() error {
	if v, ok := ec.mutation.EventType(); ok {
		if err := event.EventTypeValidator(v); err != nil {
			return &ValidationError{Name: "EventType", err: fmt.Errorf(`ent: validator failed for field "Event.EventType": %w`, err)}
		}
	}
	if v, ok := ec.mutation.Frequency(); ok {
		if err := event.FrequencyValidator(v); err != nil {
			return &ValidationError{Name: "frequency", err: fmt.Errorf(`ent: validator failed for field "Event.frequency": %w`, err)}
		}
	}
	if v, ok := ec.mutation.VenueType(); ok {
		if err := event.VenueTypeValidator(v); err != nil {
			return &ValidationError{Name: "venue_type", err: fmt.Errorf(`ent: validator failed for field "Event.venue_type": %w`, err)}
		}
	}
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Event.createdAt"`)}
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "Event.updatedAt"`)}
	}
	return nil
}

func (ec *EventCreate) sqlSave(ctx context.Context) (*Event, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Event.ID type: %T", _spec.ID.Value)
		}
	}
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EventCreate) createSpec() (*Event, *sqlgraph.CreateSpec) {
	var (
		_node = &Event{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(event.Table, sqlgraph.NewFieldSpec(event.FieldID, field.TypeString))
	)
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ec.mutation.Name(); ok {
		_spec.SetField(event.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ec.mutation.EventType(); ok {
		_spec.SetField(event.FieldEventType, field.TypeEnum, value)
		_node.EventType = value
	}
	if value, ok := ec.mutation.Status(); ok {
		_spec.SetField(event.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := ec.mutation.Location(); ok {
		_spec.SetField(event.FieldLocation, field.TypeString, value)
		_node.Location = value
	}
	if value, ok := ec.mutation.URL(); ok {
		_spec.SetField(event.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := ec.mutation.Title(); ok {
		_spec.SetField(event.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := ec.mutation.TimeZone(); ok {
		_spec.SetField(event.FieldTimeZone, field.TypeString, value)
		_node.TimeZone = value
	}
	if value, ok := ec.mutation.StartTime(); ok {
		_spec.SetField(event.FieldStartTime, field.TypeTime, value)
		_node.StartTime = value
	}
	if value, ok := ec.mutation.EndTime(); ok {
		_spec.SetField(event.FieldEndTime, field.TypeTime, value)
		_node.EndTime = value
	}
	if value, ok := ec.mutation.StartDate(); ok {
		_spec.SetField(event.FieldStartDate, field.TypeString, value)
		_node.StartDate = value
	}
	if value, ok := ec.mutation.EndDate(); ok {
		_spec.SetField(event.FieldEndDate, field.TypeString, value)
		_node.EndDate = value
	}
	if value, ok := ec.mutation.Frequency(); ok {
		_spec.SetField(event.FieldFrequency, field.TypeEnum, value)
		_node.Frequency = value
	}
	if value, ok := ec.mutation.FrequencyInterval(); ok {
		_spec.SetField(event.FieldFrequencyInterval, field.TypeString, value)
		_node.FrequencyInterval = value
	}
	if value, ok := ec.mutation.FrequencyDayOfWeek(); ok {
		_spec.SetField(event.FieldFrequencyDayOfWeek, field.TypeString, value)
		_node.FrequencyDayOfWeek = value
	}
	if value, ok := ec.mutation.FrequencyDayOfMonth(); ok {
		_spec.SetField(event.FieldFrequencyDayOfMonth, field.TypeString, value)
		_node.FrequencyDayOfMonth = value
	}
	if value, ok := ec.mutation.FrequencyMonthOfYear(); ok {
		_spec.SetField(event.FieldFrequencyMonthOfYear, field.TypeString, value)
		_node.FrequencyMonthOfYear = value
	}
	if value, ok := ec.mutation.VenueType(); ok {
		_spec.SetField(event.FieldVenueType, field.TypeEnum, value)
		_node.VenueType = value
	}
	if value, ok := ec.mutation.VenueName(); ok {
		_spec.SetField(event.FieldVenueName, field.TypeString, value)
		_node.VenueName = value
	}
	if value, ok := ec.mutation.VenueAddress(); ok {
		_spec.SetField(event.FieldVenueAddress, field.TypeString, value)
		_node.VenueAddress = value
	}
	if value, ok := ec.mutation.VenueCity(); ok {
		_spec.SetField(event.FieldVenueCity, field.TypeString, value)
		_node.VenueCity = value
	}
	if value, ok := ec.mutation.VenueState(); ok {
		_spec.SetField(event.FieldVenueState, field.TypeString, value)
		_node.VenueState = value
	}
	if value, ok := ec.mutation.VenueCountry(); ok {
		_spec.SetField(event.FieldVenueCountry, field.TypeString, value)
		_node.VenueCountry = value
	}
	if value, ok := ec.mutation.VenueZip(); ok {
		_spec.SetField(event.FieldVenueZip, field.TypeString, value)
		_node.VenueZip = value
	}
	if value, ok := ec.mutation.VenueLat(); ok {
		_spec.SetField(event.FieldVenueLat, field.TypeString, value)
		_node.VenueLat = value
	}
	if value, ok := ec.mutation.VenueLon(); ok {
		_spec.SetField(event.FieldVenueLon, field.TypeString, value)
		_node.VenueLon = value
	}
	if value, ok := ec.mutation.VenueURL(); ok {
		_spec.SetField(event.FieldVenueURL, field.TypeString, value)
		_node.VenueURL = value
	}
	if value, ok := ec.mutation.VenuePhone(); ok {
		_spec.SetField(event.FieldVenuePhone, field.TypeString, value)
		_node.VenuePhone = value
	}
	if value, ok := ec.mutation.VenueEmail(); ok {
		_spec.SetField(event.FieldVenueEmail, field.TypeString, value)
		_node.VenueEmail = value
	}
	if value, ok := ec.mutation.Tags(); ok {
		_spec.SetField(event.FieldTags, field.TypeString, value)
		_node.Tags = value
	}
	if value, ok := ec.mutation.Description(); ok {
		_spec.SetField(event.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ec.mutation.EventSettings(); ok {
		_spec.SetField(event.FieldEventSettings, field.TypeJSON, value)
		_node.EventSettings = value
	}
	if value, ok := ec.mutation.CoverImage(); ok {
		_spec.SetField(event.FieldCoverImage, field.TypeString, value)
		_node.CoverImage = value
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.SetField(event.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.UpdatedAt(); ok {
		_spec.SetField(event.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ec.mutation.MapCoordinates(); ok {
		_spec.SetField(event.FieldMapCoordinates, field.TypeJSON, value)
		_node.MapCoordinates = value
	}
	if value, ok := ec.mutation.Longitude(); ok {
		_spec.SetField(event.FieldLongitude, field.TypeString, value)
		_node.Longitude = value
	}
	if value, ok := ec.mutation.Latitude(); ok {
		_spec.SetField(event.FieldLatitude, field.TypeString, value)
		_node.Latitude = value
	}
	if value, ok := ec.mutation.SearchText(); ok {
		_spec.SetField(event.FieldSearchText, field.TypeString, value)
		_node.SearchText = value
	}
	if value, ok := ec.mutation.RelevanceScore(); ok {
		_spec.SetField(event.FieldRelevanceScore, field.TypeFloat64, value)
		_node.RelevanceScore = value
	}
	if nodes := ec.mutation.TicketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.TicketsTable,
			Columns: []string{event.TicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.TicketOptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.TicketOptionsTable,
			Columns: []string{event.TicketOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticketoption.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.PlaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.PlaceTable,
			Columns: []string{event.PlaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.EventCategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.EventCategoriesTable,
			Columns: []string{event.EventCategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.EventCategoryAssignmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.EventCategoryAssignmentsTable,
			Columns: []string{event.EventCategoryAssignmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(categoryassignment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.OwnerUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   event.OwnerUserTable,
			Columns: []string{event.OwnerUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_owned_events = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.OwnerBusinessIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.OwnerBusinessTable,
			Columns: []string{event.OwnerBusinessColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.business_events = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.UserFollowersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   event.UserFollowersTable,
			Columns: []string{event.UserFollowersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userfollowevent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.BusinessFollowersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   event.BusinessFollowersTable,
			Columns: []string{event.BusinessFollowersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessfollowevent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.FaqsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   event.FaqsTable,
			Columns: event.FaqsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(faq.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.RatingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.RatingsTable,
			Columns: []string{event.RatingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rating.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EventCreateBulk is the builder for creating many Event entities in bulk.
type EventCreateBulk struct {
	config
	builders []*EventCreate
}

// Save creates the Event entities in the database.
func (ecb *EventCreateBulk) Save(ctx context.Context) ([]*Event, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Event, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EventCreateBulk) SaveX(ctx context.Context) []*Event {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EventCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EventCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}

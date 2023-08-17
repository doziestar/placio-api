// Code generated by ent, DO NOT EDIT.

package event

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEventType holds the string denoting the eventtype field in the database.
	FieldEventType = "event_type"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldTimeZone holds the string denoting the time_zone field in the database.
	FieldTimeZone = "time_zone"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the end_time field in the database.
	FieldEndTime = "end_time"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// FieldFrequency holds the string denoting the frequency field in the database.
	FieldFrequency = "frequency"
	// FieldFrequencyInterval holds the string denoting the frequency_interval field in the database.
	FieldFrequencyInterval = "frequency_interval"
	// FieldFrequencyDayOfWeek holds the string denoting the frequency_day_of_week field in the database.
	FieldFrequencyDayOfWeek = "frequency_day_of_week"
	// FieldFrequencyDayOfMonth holds the string denoting the frequency_day_of_month field in the database.
	FieldFrequencyDayOfMonth = "frequency_day_of_month"
	// FieldFrequencyMonthOfYear holds the string denoting the frequency_month_of_year field in the database.
	FieldFrequencyMonthOfYear = "frequency_month_of_year"
	// FieldVenueType holds the string denoting the venue_type field in the database.
	FieldVenueType = "venue_type"
	// FieldVenueName holds the string denoting the venue_name field in the database.
	FieldVenueName = "venue_name"
	// FieldVenueAddress holds the string denoting the venue_address field in the database.
	FieldVenueAddress = "venue_address"
	// FieldVenueCity holds the string denoting the venue_city field in the database.
	FieldVenueCity = "venue_city"
	// FieldVenueState holds the string denoting the venue_state field in the database.
	FieldVenueState = "venue_state"
	// FieldVenueCountry holds the string denoting the venue_country field in the database.
	FieldVenueCountry = "venue_country"
	// FieldVenueZip holds the string denoting the venue_zip field in the database.
	FieldVenueZip = "venue_zip"
	// FieldVenueLat holds the string denoting the venue_lat field in the database.
	FieldVenueLat = "venue_lat"
	// FieldVenueLon holds the string denoting the venue_lon field in the database.
	FieldVenueLon = "venue_lon"
	// FieldVenueURL holds the string denoting the venue_url field in the database.
	FieldVenueURL = "venue_url"
	// FieldVenuePhone holds the string denoting the venue_phone field in the database.
	FieldVenuePhone = "venue_phone"
	// FieldVenueEmail holds the string denoting the venue_email field in the database.
	FieldVenueEmail = "venue_email"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldEventSettings holds the string denoting the event_settings field in the database.
	FieldEventSettings = "event_settings"
	// FieldCoverImage holds the string denoting the cover_image field in the database.
	FieldCoverImage = "cover_image"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldMapCoordinates holds the string denoting the map_coordinates field in the database.
	FieldMapCoordinates = "map_coordinates"
	// FieldLongitude holds the string denoting the longitude field in the database.
	FieldLongitude = "longitude"
	// FieldLatitude holds the string denoting the latitude field in the database.
	FieldLatitude = "latitude"
	// FieldSearchText holds the string denoting the search_text field in the database.
	FieldSearchText = "search_text"
	// FieldRelevanceScore holds the string denoting the relevance_score field in the database.
	FieldRelevanceScore = "relevance_score"
	// FieldFollowerCount holds the string denoting the follower_count field in the database.
	FieldFollowerCount = "follower_count"
	// FieldFollowingCount holds the string denoting the following_count field in the database.
	FieldFollowingCount = "following_count"
	// FieldIsPremium holds the string denoting the is_premium field in the database.
	FieldIsPremium = "is_premium"
	// FieldIsPublished holds the string denoting the is_published field in the database.
	FieldIsPublished = "is_published"
	// FieldIsOnline holds the string denoting the is_online field in the database.
	FieldIsOnline = "is_online"
	// FieldIsFree holds the string denoting the is_free field in the database.
	FieldIsFree = "is_free"
	// FieldIsPaid holds the string denoting the is_paid field in the database.
	FieldIsPaid = "is_paid"
	// FieldIsOnlineOnly holds the string denoting the is_online_only field in the database.
	FieldIsOnlineOnly = "is_online_only"
	// FieldIsInPersonOnly holds the string denoting the is_in_person_only field in the database.
	FieldIsInPersonOnly = "is_in_person_only"
	// FieldIsHybrid holds the string denoting the is_hybrid field in the database.
	FieldIsHybrid = "is_hybrid"
	// FieldIsOnlineAndInPerson holds the string denoting the is_online_and_in_person field in the database.
	FieldIsOnlineAndInPerson = "is_online_and_in_person"
	// FieldIsOnlineAndInPersonOnly holds the string denoting the is_online_and_in_person_only field in the database.
	FieldIsOnlineAndInPersonOnly = "is_online_and_in_person_only"
	// FieldIsOnlineAndInPersonOrHybrid holds the string denoting the is_online_and_in_person_or_hybrid field in the database.
	FieldIsOnlineAndInPersonOrHybrid = "is_online_and_in_person_or_hybrid"
	// FieldLikedByCurrentUser holds the string denoting the likedbycurrentuser field in the database.
	FieldLikedByCurrentUser = "liked_by_current_user"
	// FieldFollowedByCurrentUser holds the string denoting the followedbycurrentuser field in the database.
	FieldFollowedByCurrentUser = "followed_by_current_user"
	// EdgeTickets holds the string denoting the tickets edge name in mutations.
	EdgeTickets = "tickets"
	// EdgeTicketOptions holds the string denoting the ticket_options edge name in mutations.
	EdgeTicketOptions = "ticket_options"
	// EdgePlace holds the string denoting the place edge name in mutations.
	EdgePlace = "place"
	// EdgeEventCategories holds the string denoting the event_categories edge name in mutations.
	EdgeEventCategories = "event_categories"
	// EdgeEventCategoryAssignments holds the string denoting the event_category_assignments edge name in mutations.
	EdgeEventCategoryAssignments = "event_category_assignments"
	// EdgeOwnerUser holds the string denoting the owneruser edge name in mutations.
	EdgeOwnerUser = "ownerUser"
	// EdgeOwnerBusiness holds the string denoting the ownerbusiness edge name in mutations.
	EdgeOwnerBusiness = "ownerBusiness"
	// EdgeUserFollowers holds the string denoting the userfollowers edge name in mutations.
	EdgeUserFollowers = "userFollowers"
	// EdgeBusinessFollowers holds the string denoting the businessfollowers edge name in mutations.
	EdgeBusinessFollowers = "businessFollowers"
	// EdgeFaqs holds the string denoting the faqs edge name in mutations.
	EdgeFaqs = "faqs"
	// EdgeRatings holds the string denoting the ratings edge name in mutations.
	EdgeRatings = "ratings"
	// Table holds the table name of the event in the database.
	Table = "events"
	// TicketsTable is the table that holds the tickets relation/edge.
	TicketsTable = "tickets"
	// TicketsInverseTable is the table name for the Ticket entity.
	// It exists in this package in order to avoid circular dependency with the "ticket" package.
	TicketsInverseTable = "tickets"
	// TicketsColumn is the table column denoting the tickets relation/edge.
	TicketsColumn = "event_tickets"
	// TicketOptionsTable is the table that holds the ticket_options relation/edge.
	TicketOptionsTable = "ticket_options"
	// TicketOptionsInverseTable is the table name for the TicketOption entity.
	// It exists in this package in order to avoid circular dependency with the "ticketoption" package.
	TicketOptionsInverseTable = "ticket_options"
	// TicketOptionsColumn is the table column denoting the ticket_options relation/edge.
	TicketOptionsColumn = "event_ticket_options"
	// PlaceTable is the table that holds the place relation/edge.
	PlaceTable = "places"
	// PlaceInverseTable is the table name for the Place entity.
	// It exists in this package in order to avoid circular dependency with the "place" package.
	PlaceInverseTable = "places"
	// PlaceColumn is the table column denoting the place relation/edge.
	PlaceColumn = "event_place"
	// EventCategoriesTable is the table that holds the event_categories relation/edge.
	EventCategoriesTable = "categories"
	// EventCategoriesInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	EventCategoriesInverseTable = "categories"
	// EventCategoriesColumn is the table column denoting the event_categories relation/edge.
	EventCategoriesColumn = "event_event_categories"
	// EventCategoryAssignmentsTable is the table that holds the event_category_assignments relation/edge.
	EventCategoryAssignmentsTable = "category_assignments"
	// EventCategoryAssignmentsInverseTable is the table name for the CategoryAssignment entity.
	// It exists in this package in order to avoid circular dependency with the "categoryassignment" package.
	EventCategoryAssignmentsInverseTable = "category_assignments"
	// EventCategoryAssignmentsColumn is the table column denoting the event_category_assignments relation/edge.
	EventCategoryAssignmentsColumn = "event_event_category_assignments"
	// OwnerUserTable is the table that holds the ownerUser relation/edge.
	OwnerUserTable = "events"
	// OwnerUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerUserInverseTable = "users"
	// OwnerUserColumn is the table column denoting the ownerUser relation/edge.
	OwnerUserColumn = "user_owned_events"
	// OwnerBusinessTable is the table that holds the ownerBusiness relation/edge.
	OwnerBusinessTable = "events"
	// OwnerBusinessInverseTable is the table name for the Business entity.
	// It exists in this package in order to avoid circular dependency with the "business" package.
	OwnerBusinessInverseTable = "businesses"
	// OwnerBusinessColumn is the table column denoting the ownerBusiness relation/edge.
	OwnerBusinessColumn = "business_events"
	// UserFollowersTable is the table that holds the userFollowers relation/edge.
	UserFollowersTable = "user_follow_events"
	// UserFollowersInverseTable is the table name for the UserFollowEvent entity.
	// It exists in this package in order to avoid circular dependency with the "userfollowevent" package.
	UserFollowersInverseTable = "user_follow_events"
	// UserFollowersColumn is the table column denoting the userFollowers relation/edge.
	UserFollowersColumn = "user_follow_event_event"
	// BusinessFollowersTable is the table that holds the businessFollowers relation/edge.
	BusinessFollowersTable = "business_follow_events"
	// BusinessFollowersInverseTable is the table name for the BusinessFollowEvent entity.
	// It exists in this package in order to avoid circular dependency with the "businessfollowevent" package.
	BusinessFollowersInverseTable = "business_follow_events"
	// BusinessFollowersColumn is the table column denoting the businessFollowers relation/edge.
	BusinessFollowersColumn = "business_follow_event_event"
	// FaqsTable is the table that holds the faqs relation/edge. The primary key declared below.
	FaqsTable = "faq_event"
	// FaqsInverseTable is the table name for the FAQ entity.
	// It exists in this package in order to avoid circular dependency with the "faq" package.
	FaqsInverseTable = "fa_qs"
	// RatingsTable is the table that holds the ratings relation/edge.
	RatingsTable = "ratings"
	// RatingsInverseTable is the table name for the Rating entity.
	// It exists in this package in order to avoid circular dependency with the "rating" package.
	RatingsInverseTable = "ratings"
	// RatingsColumn is the table column denoting the ratings relation/edge.
	RatingsColumn = "event_ratings"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEventType,
	FieldStatus,
	FieldLocation,
	FieldURL,
	FieldTitle,
	FieldTimeZone,
	FieldStartTime,
	FieldEndTime,
	FieldStartDate,
	FieldEndDate,
	FieldFrequency,
	FieldFrequencyInterval,
	FieldFrequencyDayOfWeek,
	FieldFrequencyDayOfMonth,
	FieldFrequencyMonthOfYear,
	FieldVenueType,
	FieldVenueName,
	FieldVenueAddress,
	FieldVenueCity,
	FieldVenueState,
	FieldVenueCountry,
	FieldVenueZip,
	FieldVenueLat,
	FieldVenueLon,
	FieldVenueURL,
	FieldVenuePhone,
	FieldVenueEmail,
	FieldTags,
	FieldDescription,
	FieldEventSettings,
	FieldCoverImage,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldMapCoordinates,
	FieldLongitude,
	FieldLatitude,
	FieldSearchText,
	FieldRelevanceScore,
	FieldFollowerCount,
	FieldFollowingCount,
	FieldIsPremium,
	FieldIsPublished,
	FieldIsOnline,
	FieldIsFree,
	FieldIsPaid,
	FieldIsOnlineOnly,
	FieldIsInPersonOnly,
	FieldIsHybrid,
	FieldIsOnlineAndInPerson,
	FieldIsOnlineAndInPersonOnly,
	FieldIsOnlineAndInPersonOrHybrid,
	FieldLikedByCurrentUser,
	FieldFollowedByCurrentUser,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "events"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"business_events",
	"place_events",
	"user_owned_events",
}

var (
	// FaqsPrimaryKey and FaqsColumn2 are the table columns denoting the
	// primary key for the faqs relation (M2M).
	FaqsPrimaryKey = []string{"faq_id", "event_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "placio-app/ent/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCoverImage holds the default value on creation for the "cover_image" field.
	DefaultCoverImage string
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultFollowerCount holds the default value on creation for the "follower_count" field.
	DefaultFollowerCount int
	// DefaultFollowingCount holds the default value on creation for the "following_count" field.
	DefaultFollowingCount int
	// DefaultIsPremium holds the default value on creation for the "is_Premium" field.
	DefaultIsPremium bool
	// DefaultIsPublished holds the default value on creation for the "is_published" field.
	DefaultIsPublished bool
	// DefaultIsOnline holds the default value on creation for the "is_Online" field.
	DefaultIsOnline bool
	// DefaultIsFree holds the default value on creation for the "is_Free" field.
	DefaultIsFree bool
	// DefaultIsPaid holds the default value on creation for the "is_Paid" field.
	DefaultIsPaid bool
	// DefaultIsOnlineOnly holds the default value on creation for the "is_Online_Only" field.
	DefaultIsOnlineOnly bool
	// DefaultIsInPersonOnly holds the default value on creation for the "is_In_Person_Only" field.
	DefaultIsInPersonOnly bool
	// DefaultIsHybrid holds the default value on creation for the "is_Hybrid" field.
	DefaultIsHybrid bool
	// DefaultIsOnlineAndInPerson holds the default value on creation for the "is_Online_And_In_Person" field.
	DefaultIsOnlineAndInPerson bool
	// DefaultIsOnlineAndInPersonOnly holds the default value on creation for the "is_Online_And_In_Person_Only" field.
	DefaultIsOnlineAndInPersonOnly bool
	// DefaultIsOnlineAndInPersonOrHybrid holds the default value on creation for the "is_Online_And_In_Person_Or_Hybrid" field.
	DefaultIsOnlineAndInPersonOrHybrid bool
	// DefaultLikedByCurrentUser holds the default value on creation for the "likedByCurrentUser" field.
	DefaultLikedByCurrentUser bool
	// DefaultFollowedByCurrentUser holds the default value on creation for the "followedByCurrentUser" field.
	DefaultFollowedByCurrentUser bool
)

// EventType defines the type for the "EventType" enum field.
type EventType string

// EventType values.
const (
	EventTypeEvent    EventType = "event"
	EventTypePlace    EventType = "place"
	EventTypeBusiness EventType = "business"
	EventTypeFree     EventType = "free"
	EventTypePaid     EventType = "paid"
)

func (_eventtype EventType) String() string {
	return string(_eventtype)
}

// EventTypeValidator is a validator for the "EventType" field enum values. It is called by the builders before save.
func EventTypeValidator(_eventtype EventType) error {
	switch _eventtype {
	case EventTypeEvent, EventTypePlace, EventTypeBusiness, EventTypeFree, EventTypePaid:
		return nil
	default:
		return fmt.Errorf("event: invalid enum value for EventType field: %q", _eventtype)
	}
}

// Frequency defines the type for the "frequency" enum field.
type Frequency string

// Frequency values.
const (
	FrequencyOnce    Frequency = "once"
	FrequencyDaily   Frequency = "daily"
	FrequencyWeekly  Frequency = "weekly"
	FrequencyMonthly Frequency = "monthly"
	FrequencyYearly  Frequency = "yearly"
)

func (f Frequency) String() string {
	return string(f)
}

// FrequencyValidator is a validator for the "frequency" field enum values. It is called by the builders before save.
func FrequencyValidator(f Frequency) error {
	switch f {
	case FrequencyOnce, FrequencyDaily, FrequencyWeekly, FrequencyMonthly, FrequencyYearly:
		return nil
	default:
		return fmt.Errorf("event: invalid enum value for frequency field: %q", f)
	}
}

// VenueType defines the type for the "venue_type" enum field.
type VenueType string

// VenueType values.
const (
	VenueTypeOnline   VenueType = "online"
	VenueTypeInPerson VenueType = "in_person"
	VenueTypeHybrid   VenueType = "hybrid"
)

func (vt VenueType) String() string {
	return string(vt)
}

// VenueTypeValidator is a validator for the "venue_type" field enum values. It is called by the builders before save.
func VenueTypeValidator(vt VenueType) error {
	switch vt {
	case VenueTypeOnline, VenueTypeInPerson, VenueTypeHybrid:
		return nil
	default:
		return fmt.Errorf("event: invalid enum value for venue_type field: %q", vt)
	}
}

// OrderOption defines the ordering options for the Event queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByEventType orders the results by the EventType field.
func ByEventType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEventType, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByLocation orders the results by the location field.
func ByLocation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocation, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByTimeZone orders the results by the time_zone field.
func ByTimeZone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimeZone, opts...).ToFunc()
}

// ByStartTime orders the results by the start_time field.
func ByStartTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartTime, opts...).ToFunc()
}

// ByEndTime orders the results by the end_time field.
func ByEndTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndTime, opts...).ToFunc()
}

// ByStartDate orders the results by the start_date field.
func ByStartDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartDate, opts...).ToFunc()
}

// ByEndDate orders the results by the end_date field.
func ByEndDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndDate, opts...).ToFunc()
}

// ByFrequency orders the results by the frequency field.
func ByFrequency(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFrequency, opts...).ToFunc()
}

// ByFrequencyInterval orders the results by the frequency_interval field.
func ByFrequencyInterval(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFrequencyInterval, opts...).ToFunc()
}

// ByFrequencyDayOfWeek orders the results by the frequency_day_of_week field.
func ByFrequencyDayOfWeek(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFrequencyDayOfWeek, opts...).ToFunc()
}

// ByFrequencyDayOfMonth orders the results by the frequency_day_of_month field.
func ByFrequencyDayOfMonth(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFrequencyDayOfMonth, opts...).ToFunc()
}

// ByFrequencyMonthOfYear orders the results by the frequency_month_of_year field.
func ByFrequencyMonthOfYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFrequencyMonthOfYear, opts...).ToFunc()
}

// ByVenueType orders the results by the venue_type field.
func ByVenueType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueType, opts...).ToFunc()
}

// ByVenueName orders the results by the venue_name field.
func ByVenueName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueName, opts...).ToFunc()
}

// ByVenueAddress orders the results by the venue_address field.
func ByVenueAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueAddress, opts...).ToFunc()
}

// ByVenueCity orders the results by the venue_city field.
func ByVenueCity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueCity, opts...).ToFunc()
}

// ByVenueState orders the results by the venue_state field.
func ByVenueState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueState, opts...).ToFunc()
}

// ByVenueCountry orders the results by the venue_country field.
func ByVenueCountry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueCountry, opts...).ToFunc()
}

// ByVenueZip orders the results by the venue_zip field.
func ByVenueZip(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueZip, opts...).ToFunc()
}

// ByVenueLat orders the results by the venue_lat field.
func ByVenueLat(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueLat, opts...).ToFunc()
}

// ByVenueLon orders the results by the venue_lon field.
func ByVenueLon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueLon, opts...).ToFunc()
}

// ByVenueURL orders the results by the venue_url field.
func ByVenueURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueURL, opts...).ToFunc()
}

// ByVenuePhone orders the results by the venue_phone field.
func ByVenuePhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenuePhone, opts...).ToFunc()
}

// ByVenueEmail orders the results by the venue_email field.
func ByVenueEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueEmail, opts...).ToFunc()
}

// ByTags orders the results by the tags field.
func ByTags(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTags, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCoverImage orders the results by the cover_image field.
func ByCoverImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoverImage, opts...).ToFunc()
}

// ByCreatedAt orders the results by the createdAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByLongitude orders the results by the longitude field.
func ByLongitude(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLongitude, opts...).ToFunc()
}

// ByLatitude orders the results by the latitude field.
func ByLatitude(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLatitude, opts...).ToFunc()
}

// BySearchText orders the results by the search_text field.
func BySearchText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSearchText, opts...).ToFunc()
}

// ByRelevanceScore orders the results by the relevance_score field.
func ByRelevanceScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRelevanceScore, opts...).ToFunc()
}

// ByFollowerCount orders the results by the follower_count field.
func ByFollowerCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFollowerCount, opts...).ToFunc()
}

// ByFollowingCount orders the results by the following_count field.
func ByFollowingCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFollowingCount, opts...).ToFunc()
}

// ByIsPremium orders the results by the is_Premium field.
func ByIsPremium(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsPremium, opts...).ToFunc()
}

// ByIsPublished orders the results by the is_published field.
func ByIsPublished(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsPublished, opts...).ToFunc()
}

// ByIsOnline orders the results by the is_Online field.
func ByIsOnline(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsOnline, opts...).ToFunc()
}

// ByIsFree orders the results by the is_Free field.
func ByIsFree(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsFree, opts...).ToFunc()
}

// ByIsPaid orders the results by the is_Paid field.
func ByIsPaid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsPaid, opts...).ToFunc()
}

// ByIsOnlineOnly orders the results by the is_Online_Only field.
func ByIsOnlineOnly(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsOnlineOnly, opts...).ToFunc()
}

// ByIsInPersonOnly orders the results by the is_In_Person_Only field.
func ByIsInPersonOnly(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsInPersonOnly, opts...).ToFunc()
}

// ByIsHybrid orders the results by the is_Hybrid field.
func ByIsHybrid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsHybrid, opts...).ToFunc()
}

// ByIsOnlineAndInPerson orders the results by the is_Online_And_In_Person field.
func ByIsOnlineAndInPerson(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsOnlineAndInPerson, opts...).ToFunc()
}

// ByIsOnlineAndInPersonOnly orders the results by the is_Online_And_In_Person_Only field.
func ByIsOnlineAndInPersonOnly(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsOnlineAndInPersonOnly, opts...).ToFunc()
}

// ByIsOnlineAndInPersonOrHybrid orders the results by the is_Online_And_In_Person_Or_Hybrid field.
func ByIsOnlineAndInPersonOrHybrid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsOnlineAndInPersonOrHybrid, opts...).ToFunc()
}

// ByLikedByCurrentUser orders the results by the likedByCurrentUser field.
func ByLikedByCurrentUser(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLikedByCurrentUser, opts...).ToFunc()
}

// ByFollowedByCurrentUser orders the results by the followedByCurrentUser field.
func ByFollowedByCurrentUser(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFollowedByCurrentUser, opts...).ToFunc()
}

// ByTicketsCount orders the results by tickets count.
func ByTicketsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTicketsStep(), opts...)
	}
}

// ByTickets orders the results by tickets terms.
func ByTickets(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTicketsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTicketOptionsCount orders the results by ticket_options count.
func ByTicketOptionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTicketOptionsStep(), opts...)
	}
}

// ByTicketOptions orders the results by ticket_options terms.
func ByTicketOptions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTicketOptionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPlaceCount orders the results by place count.
func ByPlaceCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPlaceStep(), opts...)
	}
}

// ByPlace orders the results by place terms.
func ByPlace(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlaceStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByEventCategoriesCount orders the results by event_categories count.
func ByEventCategoriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEventCategoriesStep(), opts...)
	}
}

// ByEventCategories orders the results by event_categories terms.
func ByEventCategories(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEventCategoriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByEventCategoryAssignmentsCount orders the results by event_category_assignments count.
func ByEventCategoryAssignmentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEventCategoryAssignmentsStep(), opts...)
	}
}

// ByEventCategoryAssignments orders the results by event_category_assignments terms.
func ByEventCategoryAssignments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEventCategoryAssignmentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOwnerUserField orders the results by ownerUser field.
func ByOwnerUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByOwnerBusinessField orders the results by ownerBusiness field.
func ByOwnerBusinessField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerBusinessStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserFollowersCount orders the results by userFollowers count.
func ByUserFollowersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserFollowersStep(), opts...)
	}
}

// ByUserFollowers orders the results by userFollowers terms.
func ByUserFollowers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserFollowersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBusinessFollowersCount orders the results by businessFollowers count.
func ByBusinessFollowersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBusinessFollowersStep(), opts...)
	}
}

// ByBusinessFollowers orders the results by businessFollowers terms.
func ByBusinessFollowers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessFollowersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFaqsCount orders the results by faqs count.
func ByFaqsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFaqsStep(), opts...)
	}
}

// ByFaqs orders the results by faqs terms.
func ByFaqs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFaqsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRatingsCount orders the results by ratings count.
func ByRatingsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRatingsStep(), opts...)
	}
}

// ByRatings orders the results by ratings terms.
func ByRatings(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRatingsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTicketsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TicketsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TicketsTable, TicketsColumn),
	)
}
func newTicketOptionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TicketOptionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TicketOptionsTable, TicketOptionsColumn),
	)
}
func newPlaceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlaceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PlaceTable, PlaceColumn),
	)
}
func newEventCategoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EventCategoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EventCategoriesTable, EventCategoriesColumn),
	)
}
func newEventCategoryAssignmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EventCategoryAssignmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EventCategoryAssignmentsTable, EventCategoryAssignmentsColumn),
	)
}
func newOwnerUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, OwnerUserTable, OwnerUserColumn),
	)
}
func newOwnerBusinessStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerBusinessInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerBusinessTable, OwnerBusinessColumn),
	)
}
func newUserFollowersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserFollowersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, UserFollowersTable, UserFollowersColumn),
	)
}
func newBusinessFollowersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessFollowersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, BusinessFollowersTable, BusinessFollowersColumn),
	)
}
func newFaqsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FaqsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, FaqsTable, FaqsPrimaryKey...),
	)
}
func newRatingsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RatingsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RatingsTable, RatingsColumn),
	)
}

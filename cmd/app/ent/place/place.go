// Code generated by ent, DO NOT EDIT.

package place

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the place type in the database.
	Label = "place"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldWebsite holds the string denoting the website field in the database.
	FieldWebsite = "website"
	// FieldCoverImage holds the string denoting the cover_image field in the database.
	FieldCoverImage = "cover_image"
	// FieldPicture holds the string denoting the picture field in the database.
	FieldPicture = "picture"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldPlaceSettings holds the string denoting the place_settings field in the database.
	FieldPlaceSettings = "place_settings"
	// FieldOpeningHours holds the string denoting the opening_hours field in the database.
	FieldOpeningHours = "opening_hours"
	// FieldSocialMedia holds the string denoting the social_media field in the database.
	FieldSocialMedia = "social_media"
	// FieldPaymentOptions holds the string denoting the payment_options field in the database.
	FieldPaymentOptions = "payment_options"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldFeatures holds the string denoting the features field in the database.
	FieldFeatures = "features"
	// FieldAdditionalInfo holds the string denoting the additional_info field in the database.
	FieldAdditionalInfo = "additional_info"
	// FieldImages holds the string denoting the images field in the database.
	FieldImages = "images"
	// FieldAvailability holds the string denoting the availability field in the database.
	FieldAvailability = "availability"
	// FieldSpecialOffers holds the string denoting the special_offers field in the database.
	FieldSpecialOffers = "special_offers"
	// FieldSustainabilityScore holds the string denoting the sustainability_score field in the database.
	FieldSustainabilityScore = "sustainability_score"
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
	// EdgeBusiness holds the string denoting the business edge name in mutations.
	EdgeBusiness = "business"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeReviews holds the string denoting the reviews edge name in mutations.
	EdgeReviews = "reviews"
	// EdgeEvents holds the string denoting the events edge name in mutations.
	EdgeEvents = "events"
	// EdgeAmenities holds the string denoting the amenities edge name in mutations.
	EdgeAmenities = "amenities"
	// EdgeMenus holds the string denoting the menus edge name in mutations.
	EdgeMenus = "menus"
	// EdgeRooms holds the string denoting the rooms edge name in mutations.
	EdgeRooms = "rooms"
	// EdgeReservations holds the string denoting the reservations edge name in mutations.
	EdgeReservations = "reservations"
	// EdgeBookings holds the string denoting the bookings edge name in mutations.
	EdgeBookings = "bookings"
	// EdgeCategories holds the string denoting the categories edge name in mutations.
	EdgeCategories = "categories"
	// EdgeCategoryAssignments holds the string denoting the categoryassignments edge name in mutations.
	EdgeCategoryAssignments = "categoryAssignments"
	// EdgeFaqs holds the string denoting the faqs edge name in mutations.
	EdgeFaqs = "faqs"
	// EdgeLikedByUsers holds the string denoting the likedbyusers edge name in mutations.
	EdgeLikedByUsers = "likedByUsers"
	// EdgeFollowerUsers holds the string denoting the followerusers edge name in mutations.
	EdgeFollowerUsers = "followerUsers"
	// EdgeRatings holds the string denoting the ratings edge name in mutations.
	EdgeRatings = "ratings"
	// Table holds the table name of the place in the database.
	Table = "places"
	// BusinessTable is the table that holds the business relation/edge.
	BusinessTable = "places"
	// BusinessInverseTable is the table name for the Business entity.
	// It exists in this package in order to avoid circular dependency with the "business" package.
	BusinessInverseTable = "businesses"
	// BusinessColumn is the table column denoting the business relation/edge.
	BusinessColumn = "business_places"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "user_places"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// ReviewsTable is the table that holds the reviews relation/edge.
	ReviewsTable = "reviews"
	// ReviewsInverseTable is the table name for the Review entity.
	// It exists in this package in order to avoid circular dependency with the "review" package.
	ReviewsInverseTable = "reviews"
	// ReviewsColumn is the table column denoting the reviews relation/edge.
	ReviewsColumn = "place_reviews"
	// EventsTable is the table that holds the events relation/edge.
	EventsTable = "events"
	// EventsInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventsInverseTable = "events"
	// EventsColumn is the table column denoting the events relation/edge.
	EventsColumn = "place_events"
	// AmenitiesTable is the table that holds the amenities relation/edge. The primary key declared below.
	AmenitiesTable = "amenity_places"
	// AmenitiesInverseTable is the table name for the Amenity entity.
	// It exists in this package in order to avoid circular dependency with the "amenity" package.
	AmenitiesInverseTable = "amenities"
	// MenusTable is the table that holds the menus relation/edge.
	MenusTable = "menus"
	// MenusInverseTable is the table name for the Menu entity.
	// It exists in this package in order to avoid circular dependency with the "menu" package.
	MenusInverseTable = "menus"
	// MenusColumn is the table column denoting the menus relation/edge.
	MenusColumn = "place_menus"
	// RoomsTable is the table that holds the rooms relation/edge.
	RoomsTable = "rooms"
	// RoomsInverseTable is the table name for the Room entity.
	// It exists in this package in order to avoid circular dependency with the "room" package.
	RoomsInverseTable = "rooms"
	// RoomsColumn is the table column denoting the rooms relation/edge.
	RoomsColumn = "place_rooms"
	// ReservationsTable is the table that holds the reservations relation/edge.
	ReservationsTable = "reservations"
	// ReservationsInverseTable is the table name for the Reservation entity.
	// It exists in this package in order to avoid circular dependency with the "reservation" package.
	ReservationsInverseTable = "reservations"
	// ReservationsColumn is the table column denoting the reservations relation/edge.
	ReservationsColumn = "place_reservations"
	// BookingsTable is the table that holds the bookings relation/edge.
	BookingsTable = "bookings"
	// BookingsInverseTable is the table name for the Booking entity.
	// It exists in this package in order to avoid circular dependency with the "booking" package.
	BookingsInverseTable = "bookings"
	// BookingsColumn is the table column denoting the bookings relation/edge.
	BookingsColumn = "place_bookings"
	// CategoriesTable is the table that holds the categories relation/edge.
	CategoriesTable = "categories"
	// CategoriesInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoriesInverseTable = "categories"
	// CategoriesColumn is the table column denoting the categories relation/edge.
	CategoriesColumn = "place_categories"
	// CategoryAssignmentsTable is the table that holds the categoryAssignments relation/edge.
	CategoryAssignmentsTable = "category_assignments"
	// CategoryAssignmentsInverseTable is the table name for the CategoryAssignment entity.
	// It exists in this package in order to avoid circular dependency with the "categoryassignment" package.
	CategoryAssignmentsInverseTable = "category_assignments"
	// CategoryAssignmentsColumn is the table column denoting the categoryAssignments relation/edge.
	CategoryAssignmentsColumn = "entity_id"
	// FaqsTable is the table that holds the faqs relation/edge. The primary key declared below.
	FaqsTable = "faq_place"
	// FaqsInverseTable is the table name for the FAQ entity.
	// It exists in this package in order to avoid circular dependency with the "faq" package.
	FaqsInverseTable = "fa_qs"
	// LikedByUsersTable is the table that holds the likedByUsers relation/edge.
	LikedByUsersTable = "user_like_places"
	// LikedByUsersInverseTable is the table name for the UserLikePlace entity.
	// It exists in this package in order to avoid circular dependency with the "userlikeplace" package.
	LikedByUsersInverseTable = "user_like_places"
	// LikedByUsersColumn is the table column denoting the likedByUsers relation/edge.
	LikedByUsersColumn = "user_like_place_place"
	// FollowerUsersTable is the table that holds the followerUsers relation/edge.
	FollowerUsersTable = "user_follow_places"
	// FollowerUsersInverseTable is the table name for the UserFollowPlace entity.
	// It exists in this package in order to avoid circular dependency with the "userfollowplace" package.
	FollowerUsersInverseTable = "user_follow_places"
	// FollowerUsersColumn is the table column denoting the followerUsers relation/edge.
	FollowerUsersColumn = "place_follower_users"
	// RatingsTable is the table that holds the ratings relation/edge.
	RatingsTable = "ratings"
	// RatingsInverseTable is the table name for the Rating entity.
	// It exists in this package in order to avoid circular dependency with the "rating" package.
	RatingsInverseTable = "ratings"
	// RatingsColumn is the table column denoting the ratings relation/edge.
	RatingsColumn = "place_ratings"
)

// Columns holds all SQL columns for place fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldType,
	FieldDescription,
	FieldLocation,
	FieldEmail,
	FieldPhone,
	FieldWebsite,
	FieldCoverImage,
	FieldPicture,
	FieldCountry,
	FieldCity,
	FieldState,
	FieldPlaceSettings,
	FieldOpeningHours,
	FieldSocialMedia,
	FieldPaymentOptions,
	FieldTags,
	FieldFeatures,
	FieldAdditionalInfo,
	FieldImages,
	FieldAvailability,
	FieldSpecialOffers,
	FieldSustainabilityScore,
	FieldMapCoordinates,
	FieldLongitude,
	FieldLatitude,
	FieldSearchText,
	FieldRelevanceScore,
	FieldFollowerCount,
	FieldFollowingCount,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "places"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"business_places",
	"event_place",
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "place_id"}
	// AmenitiesPrimaryKey and AmenitiesColumn2 are the table columns denoting the
	// primary key for the amenities relation (M2M).
	AmenitiesPrimaryKey = []string{"amenity_id", "place_id"}
	// FaqsPrimaryKey and FaqsColumn2 are the table columns denoting the
	// primary key for the faqs relation (M2M).
	FaqsPrimaryKey = []string{"faq_id", "place_id"}
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
	// DefaultFollowerCount holds the default value on creation for the "follower_count" field.
	DefaultFollowerCount int
	// DefaultFollowingCount holds the default value on creation for the "following_count" field.
	DefaultFollowingCount int
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Place queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByLocation orders the results by the location field.
func ByLocation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocation, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByWebsite orders the results by the website field.
func ByWebsite(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWebsite, opts...).ToFunc()
}

// ByCoverImage orders the results by the cover_image field.
func ByCoverImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoverImage, opts...).ToFunc()
}

// ByPicture orders the results by the picture field.
func ByPicture(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPicture, opts...).ToFunc()
}

// ByCountry orders the results by the country field.
func ByCountry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountry, opts...).ToFunc()
}

// ByCity orders the results by the city field.
func ByCity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCity, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// BySpecialOffers orders the results by the special_offers field.
func BySpecialOffers(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSpecialOffers, opts...).ToFunc()
}

// BySustainabilityScore orders the results by the sustainability_score field.
func BySustainabilityScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSustainabilityScore, opts...).ToFunc()
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

// ByBusinessField orders the results by business field.
func ByBusinessField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessStep(), sql.OrderByField(field, opts...))
	}
}

// ByUsersCount orders the results by users count.
func ByUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersStep(), opts...)
	}
}

// ByUsers orders the results by users terms.
func ByUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByReviewsCount orders the results by reviews count.
func ByReviewsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newReviewsStep(), opts...)
	}
}

// ByReviews orders the results by reviews terms.
func ByReviews(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReviewsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByEventsCount orders the results by events count.
func ByEventsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEventsStep(), opts...)
	}
}

// ByEvents orders the results by events terms.
func ByEvents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEventsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAmenitiesCount orders the results by amenities count.
func ByAmenitiesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAmenitiesStep(), opts...)
	}
}

// ByAmenities orders the results by amenities terms.
func ByAmenities(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAmenitiesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMenusCount orders the results by menus count.
func ByMenusCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMenusStep(), opts...)
	}
}

// ByMenus orders the results by menus terms.
func ByMenus(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMenusStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRoomsCount orders the results by rooms count.
func ByRoomsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRoomsStep(), opts...)
	}
}

// ByRooms orders the results by rooms terms.
func ByRooms(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoomsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByReservationsCount orders the results by reservations count.
func ByReservationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newReservationsStep(), opts...)
	}
}

// ByReservations orders the results by reservations terms.
func ByReservations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReservationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBookingsCount orders the results by bookings count.
func ByBookingsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBookingsStep(), opts...)
	}
}

// ByBookings orders the results by bookings terms.
func ByBookings(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBookingsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCategoriesCount orders the results by categories count.
func ByCategoriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCategoriesStep(), opts...)
	}
}

// ByCategories orders the results by categories terms.
func ByCategories(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCategoriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCategoryAssignmentsCount orders the results by categoryAssignments count.
func ByCategoryAssignmentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCategoryAssignmentsStep(), opts...)
	}
}

// ByCategoryAssignments orders the results by categoryAssignments terms.
func ByCategoryAssignments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCategoryAssignmentsStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByLikedByUsersCount orders the results by likedByUsers count.
func ByLikedByUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLikedByUsersStep(), opts...)
	}
}

// ByLikedByUsers orders the results by likedByUsers terms.
func ByLikedByUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLikedByUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFollowerUsersCount orders the results by followerUsers count.
func ByFollowerUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFollowerUsersStep(), opts...)
	}
}

// ByFollowerUsers orders the results by followerUsers terms.
func ByFollowerUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFollowerUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newBusinessStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, BusinessTable, BusinessColumn),
	)
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
	)
}
func newReviewsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReviewsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ReviewsTable, ReviewsColumn),
	)
}
func newEventsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EventsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EventsTable, EventsColumn),
	)
}
func newAmenitiesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AmenitiesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, AmenitiesTable, AmenitiesPrimaryKey...),
	)
}
func newMenusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MenusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MenusTable, MenusColumn),
	)
}
func newRoomsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoomsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RoomsTable, RoomsColumn),
	)
}
func newReservationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReservationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ReservationsTable, ReservationsColumn),
	)
}
func newBookingsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BookingsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BookingsTable, BookingsColumn),
	)
}
func newCategoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CategoriesTable, CategoriesColumn),
	)
}
func newCategoryAssignmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoryAssignmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CategoryAssignmentsTable, CategoryAssignmentsColumn),
	)
}
func newFaqsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FaqsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, FaqsTable, FaqsPrimaryKey...),
	)
}
func newLikedByUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LikedByUsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, LikedByUsersTable, LikedByUsersColumn),
	)
}
func newFollowerUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FollowerUsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FollowerUsersTable, FollowerUsersColumn),
	)
}
func newRatingsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RatingsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RatingsTable, RatingsColumn),
	)
}

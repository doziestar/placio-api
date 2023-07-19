// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAuth0ID holds the string denoting the auth0_id field in the database.
	FieldAuth0ID = "auth0_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPicture holds the string denoting the picture field in the database.
	FieldPicture = "picture"
	// FieldCoverImage holds the string denoting the cover_image field in the database.
	FieldCoverImage = "cover_image"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldWebsite holds the string denoting the website field in the database.
	FieldWebsite = "website"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldMapCoordinates holds the string denoting the map_coordinates field in the database.
	FieldMapCoordinates = "map_coordinates"
	// FieldLongitude holds the string denoting the longitude field in the database.
	FieldLongitude = "longitude"
	// FieldLatitude holds the string denoting the latitude field in the database.
	FieldLatitude = "latitude"
	// FieldBio holds the string denoting the bio field in the database.
	FieldBio = "bio"
	// FieldAuth0Data holds the string denoting the auth0_data field in the database.
	FieldAuth0Data = "auth0_data"
	// FieldAppSettings holds the string denoting the app_settings field in the database.
	FieldAppSettings = "app_settings"
	// FieldUserSettings holds the string denoting the user_settings field in the database.
	FieldUserSettings = "user_settings"
	// FieldSearchText holds the string denoting the search_text field in the database.
	FieldSearchText = "search_text"
	// FieldRelevanceScore holds the string denoting the relevance_score field in the database.
	FieldRelevanceScore = "relevance_score"
	// FieldFollowersCount holds the string denoting the followers_count field in the database.
	FieldFollowersCount = "followers_count"
	// FieldFollowingCount holds the string denoting the following_count field in the database.
	FieldFollowingCount = "following_count"
	// FieldPostsCount holds the string denoting the posts_count field in the database.
	FieldPostsCount = "posts_count"
	// FieldReviewsCount holds the string denoting the reviews_count field in the database.
	FieldReviewsCount = "reviews_count"
	// FieldLikesCount holds the string denoting the likes_count field in the database.
	FieldLikesCount = "likes_count"
	// EdgeUserBusinesses holds the string denoting the userbusinesses edge name in mutations.
	EdgeUserBusinesses = "userBusinesses"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// EdgeLikes holds the string denoting the likes edge name in mutations.
	EdgeLikes = "likes"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// EdgeFollowedUsers holds the string denoting the followedusers edge name in mutations.
	EdgeFollowedUsers = "followedUsers"
	// EdgeFollowerUsers holds the string denoting the followerusers edge name in mutations.
	EdgeFollowerUsers = "followerUsers"
	// EdgeFollowedBusinesses holds the string denoting the followedbusinesses edge name in mutations.
	EdgeFollowedBusinesses = "followedBusinesses"
	// EdgeFollowerBusinesses holds the string denoting the followerbusinesses edge name in mutations.
	EdgeFollowerBusinesses = "followerBusinesses"
	// EdgeReviews holds the string denoting the reviews edge name in mutations.
	EdgeReviews = "reviews"
	// EdgeBookings holds the string denoting the bookings edge name in mutations.
	EdgeBookings = "bookings"
	// EdgeReservations holds the string denoting the reservations edge name in mutations.
	EdgeReservations = "reservations"
	// EdgeHelps holds the string denoting the helps edge name in mutations.
	EdgeHelps = "helps"
	// EdgeCategories holds the string denoting the categories edge name in mutations.
	EdgeCategories = "categories"
	// EdgePlaces holds the string denoting the places edge name in mutations.
	EdgePlaces = "places"
	// EdgeCategoryAssignments holds the string denoting the categoryassignments edge name in mutations.
	EdgeCategoryAssignments = "categoryAssignments"
	// EdgeOwnedEvents holds the string denoting the ownedevents edge name in mutations.
	EdgeOwnedEvents = "ownedEvents"
	// EdgeUserFollowEvents holds the string denoting the userfollowevents edge name in mutations.
	EdgeUserFollowEvents = "userFollowEvents"
	// EdgeFollowedPlaces holds the string denoting the followedplaces edge name in mutations.
	EdgeFollowedPlaces = "followedPlaces"
	// EdgeLikedPlaces holds the string denoting the likedplaces edge name in mutations.
	EdgeLikedPlaces = "likedPlaces"
	// EdgeRatings holds the string denoting the ratings edge name in mutations.
	EdgeRatings = "ratings"
	// Table holds the table name of the user in the database.
	Table = "users"
	// UserBusinessesTable is the table that holds the userBusinesses relation/edge.
	UserBusinessesTable = "user_businesses"
	// UserBusinessesInverseTable is the table name for the UserBusiness entity.
	// It exists in this package in order to avoid circular dependency with the "userbusiness" package.
	UserBusinessesInverseTable = "user_businesses"
	// UserBusinessesColumn is the table column denoting the userBusinesses relation/edge.
	UserBusinessesColumn = "user_user_businesses"
	// CommentsTable is the table that holds the comments relation/edge.
	CommentsTable = "comments"
	// CommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentsInverseTable = "comments"
	// CommentsColumn is the table column denoting the comments relation/edge.
	CommentsColumn = "user_comments"
	// LikesTable is the table that holds the likes relation/edge.
	LikesTable = "likes"
	// LikesInverseTable is the table name for the Like entity.
	// It exists in this package in order to avoid circular dependency with the "like" package.
	LikesInverseTable = "likes"
	// LikesColumn is the table column denoting the likes relation/edge.
	LikesColumn = "user_likes"
	// PostsTable is the table that holds the posts relation/edge.
	PostsTable = "posts"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "posts"
	// PostsColumn is the table column denoting the posts relation/edge.
	PostsColumn = "user_posts"
	// FollowedUsersTable is the table that holds the followedUsers relation/edge.
	FollowedUsersTable = "user_follow_users"
	// FollowedUsersInverseTable is the table name for the UserFollowUser entity.
	// It exists in this package in order to avoid circular dependency with the "userfollowuser" package.
	FollowedUsersInverseTable = "user_follow_users"
	// FollowedUsersColumn is the table column denoting the followedUsers relation/edge.
	FollowedUsersColumn = "user_followed_users"
	// FollowerUsersTable is the table that holds the followerUsers relation/edge.
	FollowerUsersTable = "user_follow_users"
	// FollowerUsersInverseTable is the table name for the UserFollowUser entity.
	// It exists in this package in order to avoid circular dependency with the "userfollowuser" package.
	FollowerUsersInverseTable = "user_follow_users"
	// FollowerUsersColumn is the table column denoting the followerUsers relation/edge.
	FollowerUsersColumn = "user_follower_users"
	// FollowedBusinessesTable is the table that holds the followedBusinesses relation/edge.
	FollowedBusinessesTable = "user_follow_businesses"
	// FollowedBusinessesInverseTable is the table name for the UserFollowBusiness entity.
	// It exists in this package in order to avoid circular dependency with the "userfollowbusiness" package.
	FollowedBusinessesInverseTable = "user_follow_businesses"
	// FollowedBusinessesColumn is the table column denoting the followedBusinesses relation/edge.
	FollowedBusinessesColumn = "user_followed_businesses"
	// FollowerBusinessesTable is the table that holds the followerBusinesses relation/edge.
	FollowerBusinessesTable = "business_follow_users"
	// FollowerBusinessesInverseTable is the table name for the BusinessFollowUser entity.
	// It exists in this package in order to avoid circular dependency with the "businessfollowuser" package.
	FollowerBusinessesInverseTable = "business_follow_users"
	// FollowerBusinessesColumn is the table column denoting the followerBusinesses relation/edge.
	FollowerBusinessesColumn = "user_follower_businesses"
	// ReviewsTable is the table that holds the reviews relation/edge.
	ReviewsTable = "reviews"
	// ReviewsInverseTable is the table name for the Review entity.
	// It exists in this package in order to avoid circular dependency with the "review" package.
	ReviewsInverseTable = "reviews"
	// ReviewsColumn is the table column denoting the reviews relation/edge.
	ReviewsColumn = "user_reviews"
	// BookingsTable is the table that holds the bookings relation/edge.
	BookingsTable = "bookings"
	// BookingsInverseTable is the table name for the Booking entity.
	// It exists in this package in order to avoid circular dependency with the "booking" package.
	BookingsInverseTable = "bookings"
	// BookingsColumn is the table column denoting the bookings relation/edge.
	BookingsColumn = "user_bookings"
	// ReservationsTable is the table that holds the reservations relation/edge.
	ReservationsTable = "reservations"
	// ReservationsInverseTable is the table name for the Reservation entity.
	// It exists in this package in order to avoid circular dependency with the "reservation" package.
	ReservationsInverseTable = "reservations"
	// ReservationsColumn is the table column denoting the reservations relation/edge.
	ReservationsColumn = "user_reservations"
	// HelpsTable is the table that holds the helps relation/edge.
	HelpsTable = "helps"
	// HelpsInverseTable is the table name for the Help entity.
	// It exists in this package in order to avoid circular dependency with the "help" package.
	HelpsInverseTable = "helps"
	// HelpsColumn is the table column denoting the helps relation/edge.
	HelpsColumn = "user_id"
	// CategoriesTable is the table that holds the categories relation/edge.
	CategoriesTable = "categories"
	// CategoriesInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoriesInverseTable = "categories"
	// CategoriesColumn is the table column denoting the categories relation/edge.
	CategoriesColumn = "user_categories"
	// PlacesTable is the table that holds the places relation/edge. The primary key declared below.
	PlacesTable = "user_places"
	// PlacesInverseTable is the table name for the Place entity.
	// It exists in this package in order to avoid circular dependency with the "place" package.
	PlacesInverseTable = "places"
	// CategoryAssignmentsTable is the table that holds the categoryAssignments relation/edge.
	CategoryAssignmentsTable = "category_assignments"
	// CategoryAssignmentsInverseTable is the table name for the CategoryAssignment entity.
	// It exists in this package in order to avoid circular dependency with the "categoryassignment" package.
	CategoryAssignmentsInverseTable = "category_assignments"
	// CategoryAssignmentsColumn is the table column denoting the categoryAssignments relation/edge.
	CategoryAssignmentsColumn = "entity_id"
	// OwnedEventsTable is the table that holds the ownedEvents relation/edge.
	OwnedEventsTable = "events"
	// OwnedEventsInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	OwnedEventsInverseTable = "events"
	// OwnedEventsColumn is the table column denoting the ownedEvents relation/edge.
	OwnedEventsColumn = "user_owned_events"
	// UserFollowEventsTable is the table that holds the userFollowEvents relation/edge.
	UserFollowEventsTable = "user_follow_events"
	// UserFollowEventsInverseTable is the table name for the UserFollowEvent entity.
	// It exists in this package in order to avoid circular dependency with the "userfollowevent" package.
	UserFollowEventsInverseTable = "user_follow_events"
	// UserFollowEventsColumn is the table column denoting the userFollowEvents relation/edge.
	UserFollowEventsColumn = "user_user_follow_events"
	// FollowedPlacesTable is the table that holds the followedPlaces relation/edge.
	FollowedPlacesTable = "user_follow_places"
	// FollowedPlacesInverseTable is the table name for the UserFollowPlace entity.
	// It exists in this package in order to avoid circular dependency with the "userfollowplace" package.
	FollowedPlacesInverseTable = "user_follow_places"
	// FollowedPlacesColumn is the table column denoting the followedPlaces relation/edge.
	FollowedPlacesColumn = "user_followed_places"
	// LikedPlacesTable is the table that holds the likedPlaces relation/edge.
	LikedPlacesTable = "user_like_places"
	// LikedPlacesInverseTable is the table name for the UserLikePlace entity.
	// It exists in this package in order to avoid circular dependency with the "userlikeplace" package.
	LikedPlacesInverseTable = "user_like_places"
	// LikedPlacesColumn is the table column denoting the likedPlaces relation/edge.
	LikedPlacesColumn = "user_liked_places"
	// RatingsTable is the table that holds the ratings relation/edge.
	RatingsTable = "ratings"
	// RatingsInverseTable is the table name for the Rating entity.
	// It exists in this package in order to avoid circular dependency with the "rating" package.
	RatingsInverseTable = "ratings"
	// RatingsColumn is the table column denoting the ratings relation/edge.
	RatingsColumn = "user_ratings"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldAuth0ID,
	FieldName,
	FieldPicture,
	FieldCoverImage,
	FieldUsername,
	FieldWebsite,
	FieldLocation,
	FieldMapCoordinates,
	FieldLongitude,
	FieldLatitude,
	FieldBio,
	FieldAuth0Data,
	FieldAppSettings,
	FieldUserSettings,
	FieldSearchText,
	FieldRelevanceScore,
	FieldFollowersCount,
	FieldFollowingCount,
	FieldPostsCount,
	FieldReviewsCount,
	FieldLikesCount,
}

var (
	// PlacesPrimaryKey and PlacesColumn2 are the table columns denoting the
	// primary key for the places relation (M2M).
	PlacesPrimaryKey = []string{"user_id", "place_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
	// DefaultBio holds the default value on creation for the "bio" field.
	DefaultBio string
	// DefaultFollowersCount holds the default value on creation for the "followers_count" field.
	DefaultFollowersCount int
	// DefaultFollowingCount holds the default value on creation for the "following_count" field.
	DefaultFollowingCount int
	// DefaultPostsCount holds the default value on creation for the "posts_count" field.
	DefaultPostsCount int
	// DefaultReviewsCount holds the default value on creation for the "reviews_count" field.
	DefaultReviewsCount int
	// DefaultLikesCount holds the default value on creation for the "likes_count" field.
	DefaultLikesCount int
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAuth0ID orders the results by the auth0_id field.
func ByAuth0ID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuth0ID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPicture orders the results by the picture field.
func ByPicture(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPicture, opts...).ToFunc()
}

// ByCoverImage orders the results by the cover_image field.
func ByCoverImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoverImage, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByWebsite orders the results by the website field.
func ByWebsite(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWebsite, opts...).ToFunc()
}

// ByLocation orders the results by the location field.
func ByLocation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocation, opts...).ToFunc()
}

// ByLongitude orders the results by the longitude field.
func ByLongitude(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLongitude, opts...).ToFunc()
}

// ByLatitude orders the results by the latitude field.
func ByLatitude(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLatitude, opts...).ToFunc()
}

// ByBio orders the results by the bio field.
func ByBio(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBio, opts...).ToFunc()
}

// BySearchText orders the results by the search_text field.
func BySearchText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSearchText, opts...).ToFunc()
}

// ByRelevanceScore orders the results by the relevance_score field.
func ByRelevanceScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRelevanceScore, opts...).ToFunc()
}

// ByFollowersCount orders the results by the followers_count field.
func ByFollowersCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFollowersCount, opts...).ToFunc()
}

// ByFollowingCount orders the results by the following_count field.
func ByFollowingCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFollowingCount, opts...).ToFunc()
}

// ByPostsCount orders the results by the posts_count field.
func ByPostsCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostsCount, opts...).ToFunc()
}

// ByReviewsCount orders the results by the reviews_count field.
func ByReviewsCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReviewsCount, opts...).ToFunc()
}

// ByLikesCount orders the results by the likes_count field.
func ByLikesCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLikesCount, opts...).ToFunc()
}

// ByUserBusinessesCount orders the results by userBusinesses count.
func ByUserBusinessesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserBusinessesStep(), opts...)
	}
}

// ByUserBusinesses orders the results by userBusinesses terms.
func ByUserBusinesses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserBusinessesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCommentsCount orders the results by comments count.
func ByCommentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCommentsStep(), opts...)
	}
}

// ByComments orders the results by comments terms.
func ByComments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLikesCount orders the results by likes count.
func ByLikesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLikesStep(), opts...)
	}
}

// ByLikes orders the results by likes terms.
func ByLikes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLikesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPostsCount orders the results by posts count.
func ByPostsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPostsStep(), opts...)
	}
}

// ByPosts orders the results by posts terms.
func ByPosts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFollowedUsersCount orders the results by followedUsers count.
func ByFollowedUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFollowedUsersStep(), opts...)
	}
}

// ByFollowedUsers orders the results by followedUsers terms.
func ByFollowedUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFollowedUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByFollowedBusinessesCount orders the results by followedBusinesses count.
func ByFollowedBusinessesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFollowedBusinessesStep(), opts...)
	}
}

// ByFollowedBusinesses orders the results by followedBusinesses terms.
func ByFollowedBusinesses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFollowedBusinessesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFollowerBusinessesCount orders the results by followerBusinesses count.
func ByFollowerBusinessesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFollowerBusinessesStep(), opts...)
	}
}

// ByFollowerBusinesses orders the results by followerBusinesses terms.
func ByFollowerBusinesses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFollowerBusinessesStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByHelpsCount orders the results by helps count.
func ByHelpsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newHelpsStep(), opts...)
	}
}

// ByHelps orders the results by helps terms.
func ByHelps(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHelpsStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByPlacesCount orders the results by places count.
func ByPlacesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPlacesStep(), opts...)
	}
}

// ByPlaces orders the results by places terms.
func ByPlaces(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlacesStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByOwnedEventsField orders the results by ownedEvents field.
func ByOwnedEventsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnedEventsStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserFollowEventsCount orders the results by userFollowEvents count.
func ByUserFollowEventsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserFollowEventsStep(), opts...)
	}
}

// ByUserFollowEvents orders the results by userFollowEvents terms.
func ByUserFollowEvents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserFollowEventsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFollowedPlacesCount orders the results by followedPlaces count.
func ByFollowedPlacesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFollowedPlacesStep(), opts...)
	}
}

// ByFollowedPlaces orders the results by followedPlaces terms.
func ByFollowedPlaces(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFollowedPlacesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLikedPlacesCount orders the results by likedPlaces count.
func ByLikedPlacesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLikedPlacesStep(), opts...)
	}
}

// ByLikedPlaces orders the results by likedPlaces terms.
func ByLikedPlaces(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLikedPlacesStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newUserBusinessesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserBusinessesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserBusinessesTable, UserBusinessesColumn),
	)
}
func newCommentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
	)
}
func newLikesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LikesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LikesTable, LikesColumn),
	)
}
func newPostsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PostsTable, PostsColumn),
	)
}
func newFollowedUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FollowedUsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FollowedUsersTable, FollowedUsersColumn),
	)
}
func newFollowerUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FollowerUsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FollowerUsersTable, FollowerUsersColumn),
	)
}
func newFollowedBusinessesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FollowedBusinessesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FollowedBusinessesTable, FollowedBusinessesColumn),
	)
}
func newFollowerBusinessesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FollowerBusinessesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FollowerBusinessesTable, FollowerBusinessesColumn),
	)
}
func newReviewsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReviewsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ReviewsTable, ReviewsColumn),
	)
}
func newBookingsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BookingsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BookingsTable, BookingsColumn),
	)
}
func newReservationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReservationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ReservationsTable, ReservationsColumn),
	)
}
func newHelpsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HelpsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, HelpsTable, HelpsColumn),
	)
}
func newCategoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CategoriesTable, CategoriesColumn),
	)
}
func newPlacesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlacesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PlacesTable, PlacesPrimaryKey...),
	)
}
func newCategoryAssignmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoryAssignmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CategoryAssignmentsTable, CategoryAssignmentsColumn),
	)
}
func newOwnedEventsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnedEventsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, OwnedEventsTable, OwnedEventsColumn),
	)
}
func newUserFollowEventsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserFollowEventsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserFollowEventsTable, UserFollowEventsColumn),
	)
}
func newFollowedPlacesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FollowedPlacesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FollowedPlacesTable, FollowedPlacesColumn),
	)
}
func newLikedPlacesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LikedPlacesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LikedPlacesTable, LikedPlacesColumn),
	)
}
func newRatingsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RatingsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RatingsTable, RatingsColumn),
	)
}

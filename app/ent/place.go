// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/place"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Place is the model entity for the Place schema.
type Place struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Website holds the value of the "website" field.
	Website string `json:"website,omitempty"`
	// CoverImage holds the value of the "cover_image" field.
	CoverImage string `json:"cover_image,omitempty"`
	// Picture holds the value of the "picture" field.
	Picture string `json:"picture,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// PlaceSettings holds the value of the "place_settings" field.
	PlaceSettings map[string]interface{} `json:"place_settings,omitempty"`
	// OpeningHours holds the value of the "opening_hours" field.
	OpeningHours map[string]interface{} `json:"opening_hours,omitempty"`
	// SocialMedia holds the value of the "social_media" field.
	SocialMedia map[string]interface{} `json:"social_media,omitempty"`
	// PaymentOptions holds the value of the "payment_options" field.
	PaymentOptions map[string]interface{} `json:"payment_options,omitempty"`
	// Tags holds the value of the "tags" field.
	Tags []string `json:"tags,omitempty"`
	// Features holds the value of the "features" field.
	Features []string `json:"features,omitempty"`
	// AdditionalInfo holds the value of the "additional_info" field.
	AdditionalInfo map[string]interface{} `json:"additional_info,omitempty"`
	// Images holds the value of the "images" field.
	Images []string `json:"images,omitempty"`
	// Availability holds the value of the "availability" field.
	Availability map[string]interface{} `json:"availability,omitempty"`
	// SpecialOffers holds the value of the "special_offers" field.
	SpecialOffers string `json:"special_offers,omitempty"`
	// SustainabilityScore holds the value of the "sustainability_score" field.
	SustainabilityScore float64 `json:"sustainability_score,omitempty"`
	// MapCoordinates holds the value of the "map_coordinates" field.
	MapCoordinates map[string]interface{} `json:"map_coordinates,omitempty"`
	// Longitude holds the value of the "longitude" field.
	Longitude string `json:"longitude,omitempty"`
	// Latitude holds the value of the "latitude" field.
	Latitude string `json:"latitude,omitempty"`
	// SearchText holds the value of the "search_text" field.
	SearchText string `json:"search_text,omitempty"`
	// RelevanceScore holds the value of the "relevance_score" field.
	RelevanceScore float64 `json:"relevance_score,omitempty"`
	// FollowerCount holds the value of the "follower_count" field.
	FollowerCount int `json:"follower_count,omitempty"`
	// FollowingCount holds the value of the "following_count" field.
	FollowingCount int `json:"following_count,omitempty"`
	// IsPremium holds the value of the "is_Premium" field.
	IsPremium bool `json:"is_Premium,omitempty"`
	// IsPublished holds the value of the "is_published" field.
	IsPublished bool `json:"is_published,omitempty"`
	// LikedByCurrentUser holds the value of the "likedByCurrentUser" field.
	LikedByCurrentUser bool `json:"likedByCurrentUser,omitempty"`
	// FollowedByCurrentUser holds the value of the "followedByCurrentUser" field.
	FollowedByCurrentUser bool `json:"followedByCurrentUser,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlaceQuery when eager-loading is set.
	Edges           PlaceEdges `json:"edges"`
	business_places *string
	event_place     *string
	selectValues    sql.SelectValues
}

// PlaceEdges holds the relations/edges for other nodes in the graph.
type PlaceEdges struct {
	// Business holds the value of the business edge.
	Business *Business `json:"business,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Reviews holds the value of the reviews edge.
	Reviews []*Review `json:"reviews,omitempty"`
	// Events holds the value of the events edge.
	Events []*Event `json:"events,omitempty"`
	// Amenities holds the value of the amenities edge.
	Amenities []*Amenity `json:"amenities,omitempty"`
	// Menus holds the value of the menus edge.
	Menus []*Menu `json:"menus,omitempty"`
	// Medias holds the value of the medias edge.
	Medias []*Media `json:"medias,omitempty"`
	// Rooms holds the value of the rooms edge.
	Rooms []*Room `json:"rooms,omitempty"`
	// Reservations holds the value of the reservations edge.
	Reservations []*Reservation `json:"reservations,omitempty"`
	// Bookings holds the value of the bookings edge.
	Bookings []*Booking `json:"bookings,omitempty"`
	// Categories holds the value of the categories edge.
	Categories []*Category `json:"categories,omitempty"`
	// CategoryAssignments holds the value of the categoryAssignments edge.
	CategoryAssignments []*CategoryAssignment `json:"categoryAssignments,omitempty"`
	// Faqs holds the value of the faqs edge.
	Faqs []*FAQ `json:"faqs,omitempty"`
	// LikedByUsers holds the value of the likedByUsers edge.
	LikedByUsers []*UserLikePlace `json:"likedByUsers,omitempty"`
	// FollowerUsers holds the value of the followerUsers edge.
	FollowerUsers []*UserFollowPlace `json:"followerUsers,omitempty"`
	// Ratings holds the value of the ratings edge.
	Ratings []*Rating `json:"ratings,omitempty"`
	// Inventories holds the value of the inventories edge.
	Inventories []*PlaceInventory `json:"inventories,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [17]bool
}

// BusinessOrErr returns the Business value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlaceEdges) BusinessOrErr() (*Business, error) {
	if e.loadedTypes[0] {
		if e.Business == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: business.Label}
		}
		return e.Business, nil
	}
	return nil, &NotLoadedError{edge: "business"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// ReviewsOrErr returns the Reviews value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) ReviewsOrErr() ([]*Review, error) {
	if e.loadedTypes[2] {
		return e.Reviews, nil
	}
	return nil, &NotLoadedError{edge: "reviews"}
}

// EventsOrErr returns the Events value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) EventsOrErr() ([]*Event, error) {
	if e.loadedTypes[3] {
		return e.Events, nil
	}
	return nil, &NotLoadedError{edge: "events"}
}

// AmenitiesOrErr returns the Amenities value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) AmenitiesOrErr() ([]*Amenity, error) {
	if e.loadedTypes[4] {
		return e.Amenities, nil
	}
	return nil, &NotLoadedError{edge: "amenities"}
}

// MenusOrErr returns the Menus value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) MenusOrErr() ([]*Menu, error) {
	if e.loadedTypes[5] {
		return e.Menus, nil
	}
	return nil, &NotLoadedError{edge: "menus"}
}

// MediasOrErr returns the Medias value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) MediasOrErr() ([]*Media, error) {
	if e.loadedTypes[6] {
		return e.Medias, nil
	}
	return nil, &NotLoadedError{edge: "medias"}
}

// RoomsOrErr returns the Rooms value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) RoomsOrErr() ([]*Room, error) {
	if e.loadedTypes[7] {
		return e.Rooms, nil
	}
	return nil, &NotLoadedError{edge: "rooms"}
}

// ReservationsOrErr returns the Reservations value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) ReservationsOrErr() ([]*Reservation, error) {
	if e.loadedTypes[8] {
		return e.Reservations, nil
	}
	return nil, &NotLoadedError{edge: "reservations"}
}

// BookingsOrErr returns the Bookings value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) BookingsOrErr() ([]*Booking, error) {
	if e.loadedTypes[9] {
		return e.Bookings, nil
	}
	return nil, &NotLoadedError{edge: "bookings"}
}

// CategoriesOrErr returns the Categories value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) CategoriesOrErr() ([]*Category, error) {
	if e.loadedTypes[10] {
		return e.Categories, nil
	}
	return nil, &NotLoadedError{edge: "categories"}
}

// CategoryAssignmentsOrErr returns the CategoryAssignments value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) CategoryAssignmentsOrErr() ([]*CategoryAssignment, error) {
	if e.loadedTypes[11] {
		return e.CategoryAssignments, nil
	}
	return nil, &NotLoadedError{edge: "categoryAssignments"}
}

// FaqsOrErr returns the Faqs value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) FaqsOrErr() ([]*FAQ, error) {
	if e.loadedTypes[12] {
		return e.Faqs, nil
	}
	return nil, &NotLoadedError{edge: "faqs"}
}

// LikedByUsersOrErr returns the LikedByUsers value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) LikedByUsersOrErr() ([]*UserLikePlace, error) {
	if e.loadedTypes[13] {
		return e.LikedByUsers, nil
	}
	return nil, &NotLoadedError{edge: "likedByUsers"}
}

// FollowerUsersOrErr returns the FollowerUsers value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) FollowerUsersOrErr() ([]*UserFollowPlace, error) {
	if e.loadedTypes[14] {
		return e.FollowerUsers, nil
	}
	return nil, &NotLoadedError{edge: "followerUsers"}
}

// RatingsOrErr returns the Ratings value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) RatingsOrErr() ([]*Rating, error) {
	if e.loadedTypes[15] {
		return e.Ratings, nil
	}
	return nil, &NotLoadedError{edge: "ratings"}
}

// InventoriesOrErr returns the Inventories value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceEdges) InventoriesOrErr() ([]*PlaceInventory, error) {
	if e.loadedTypes[16] {
		return e.Inventories, nil
	}
	return nil, &NotLoadedError{edge: "inventories"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Place) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case place.FieldPlaceSettings, place.FieldOpeningHours, place.FieldSocialMedia, place.FieldPaymentOptions, place.FieldTags, place.FieldFeatures, place.FieldAdditionalInfo, place.FieldImages, place.FieldAvailability, place.FieldMapCoordinates:
			values[i] = new([]byte)
		case place.FieldIsPremium, place.FieldIsPublished, place.FieldLikedByCurrentUser, place.FieldFollowedByCurrentUser:
			values[i] = new(sql.NullBool)
		case place.FieldSustainabilityScore, place.FieldRelevanceScore:
			values[i] = new(sql.NullFloat64)
		case place.FieldFollowerCount, place.FieldFollowingCount:
			values[i] = new(sql.NullInt64)
		case place.FieldID, place.FieldName, place.FieldType, place.FieldDescription, place.FieldLocation, place.FieldEmail, place.FieldPhone, place.FieldWebsite, place.FieldCoverImage, place.FieldPicture, place.FieldCountry, place.FieldCity, place.FieldState, place.FieldSpecialOffers, place.FieldLongitude, place.FieldLatitude, place.FieldSearchText:
			values[i] = new(sql.NullString)
		case place.ForeignKeys[0]: // business_places
			values[i] = new(sql.NullString)
		case place.ForeignKeys[1]: // event_place
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Place fields.
func (pl *Place) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case place.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pl.ID = value.String
			}
		case place.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pl.Name = value.String
			}
		case place.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				pl.Type = value.String
			}
		case place.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pl.Description = value.String
			}
		case place.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				pl.Location = value.String
			}
		case place.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				pl.Email = value.String
			}
		case place.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				pl.Phone = value.String
			}
		case place.FieldWebsite:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field website", values[i])
			} else if value.Valid {
				pl.Website = value.String
			}
		case place.FieldCoverImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cover_image", values[i])
			} else if value.Valid {
				pl.CoverImage = value.String
			}
		case place.FieldPicture:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field picture", values[i])
			} else if value.Valid {
				pl.Picture = value.String
			}
		case place.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				pl.Country = value.String
			}
		case place.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				pl.City = value.String
			}
		case place.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				pl.State = value.String
			}
		case place.FieldPlaceSettings:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field place_settings", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pl.PlaceSettings); err != nil {
					return fmt.Errorf("unmarshal field place_settings: %w", err)
				}
			}
		case place.FieldOpeningHours:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field opening_hours", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pl.OpeningHours); err != nil {
					return fmt.Errorf("unmarshal field opening_hours: %w", err)
				}
			}
		case place.FieldSocialMedia:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field social_media", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pl.SocialMedia); err != nil {
					return fmt.Errorf("unmarshal field social_media: %w", err)
				}
			}
		case place.FieldPaymentOptions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field payment_options", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pl.PaymentOptions); err != nil {
					return fmt.Errorf("unmarshal field payment_options: %w", err)
				}
			}
		case place.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pl.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case place.FieldFeatures:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field features", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pl.Features); err != nil {
					return fmt.Errorf("unmarshal field features: %w", err)
				}
			}
		case place.FieldAdditionalInfo:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field additional_info", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pl.AdditionalInfo); err != nil {
					return fmt.Errorf("unmarshal field additional_info: %w", err)
				}
			}
		case place.FieldImages:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field images", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pl.Images); err != nil {
					return fmt.Errorf("unmarshal field images: %w", err)
				}
			}
		case place.FieldAvailability:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field availability", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pl.Availability); err != nil {
					return fmt.Errorf("unmarshal field availability: %w", err)
				}
			}
		case place.FieldSpecialOffers:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field special_offers", values[i])
			} else if value.Valid {
				pl.SpecialOffers = value.String
			}
		case place.FieldSustainabilityScore:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field sustainability_score", values[i])
			} else if value.Valid {
				pl.SustainabilityScore = value.Float64
			}
		case place.FieldMapCoordinates:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field map_coordinates", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pl.MapCoordinates); err != nil {
					return fmt.Errorf("unmarshal field map_coordinates: %w", err)
				}
			}
		case place.FieldLongitude:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field longitude", values[i])
			} else if value.Valid {
				pl.Longitude = value.String
			}
		case place.FieldLatitude:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field latitude", values[i])
			} else if value.Valid {
				pl.Latitude = value.String
			}
		case place.FieldSearchText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field search_text", values[i])
			} else if value.Valid {
				pl.SearchText = value.String
			}
		case place.FieldRelevanceScore:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field relevance_score", values[i])
			} else if value.Valid {
				pl.RelevanceScore = value.Float64
			}
		case place.FieldFollowerCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field follower_count", values[i])
			} else if value.Valid {
				pl.FollowerCount = int(value.Int64)
			}
		case place.FieldFollowingCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field following_count", values[i])
			} else if value.Valid {
				pl.FollowingCount = int(value.Int64)
			}
		case place.FieldIsPremium:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_Premium", values[i])
			} else if value.Valid {
				pl.IsPremium = value.Bool
			}
		case place.FieldIsPublished:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_published", values[i])
			} else if value.Valid {
				pl.IsPublished = value.Bool
			}
		case place.FieldLikedByCurrentUser:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field likedByCurrentUser", values[i])
			} else if value.Valid {
				pl.LikedByCurrentUser = value.Bool
			}
		case place.FieldFollowedByCurrentUser:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field followedByCurrentUser", values[i])
			} else if value.Valid {
				pl.FollowedByCurrentUser = value.Bool
			}
		case place.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field business_places", values[i])
			} else if value.Valid {
				pl.business_places = new(string)
				*pl.business_places = value.String
			}
		case place.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_place", values[i])
			} else if value.Valid {
				pl.event_place = new(string)
				*pl.event_place = value.String
			}
		default:
			pl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Place.
// This includes values selected through modifiers, order, etc.
func (pl *Place) Value(name string) (ent.Value, error) {
	return pl.selectValues.Get(name)
}

// QueryBusiness queries the "business" edge of the Place entity.
func (pl *Place) QueryBusiness() *BusinessQuery {
	return NewPlaceClient(pl.config).QueryBusiness(pl)
}

// QueryUsers queries the "users" edge of the Place entity.
func (pl *Place) QueryUsers() *UserQuery {
	return NewPlaceClient(pl.config).QueryUsers(pl)
}

// QueryReviews queries the "reviews" edge of the Place entity.
func (pl *Place) QueryReviews() *ReviewQuery {
	return NewPlaceClient(pl.config).QueryReviews(pl)
}

// QueryEvents queries the "events" edge of the Place entity.
func (pl *Place) QueryEvents() *EventQuery {
	return NewPlaceClient(pl.config).QueryEvents(pl)
}

// QueryAmenities queries the "amenities" edge of the Place entity.
func (pl *Place) QueryAmenities() *AmenityQuery {
	return NewPlaceClient(pl.config).QueryAmenities(pl)
}

// QueryMenus queries the "menus" edge of the Place entity.
func (pl *Place) QueryMenus() *MenuQuery {
	return NewPlaceClient(pl.config).QueryMenus(pl)
}

// QueryMedias queries the "medias" edge of the Place entity.
func (pl *Place) QueryMedias() *MediaQuery {
	return NewPlaceClient(pl.config).QueryMedias(pl)
}

// QueryRooms queries the "rooms" edge of the Place entity.
func (pl *Place) QueryRooms() *RoomQuery {
	return NewPlaceClient(pl.config).QueryRooms(pl)
}

// QueryReservations queries the "reservations" edge of the Place entity.
func (pl *Place) QueryReservations() *ReservationQuery {
	return NewPlaceClient(pl.config).QueryReservations(pl)
}

// QueryBookings queries the "bookings" edge of the Place entity.
func (pl *Place) QueryBookings() *BookingQuery {
	return NewPlaceClient(pl.config).QueryBookings(pl)
}

// QueryCategories queries the "categories" edge of the Place entity.
func (pl *Place) QueryCategories() *CategoryQuery {
	return NewPlaceClient(pl.config).QueryCategories(pl)
}

// QueryCategoryAssignments queries the "categoryAssignments" edge of the Place entity.
func (pl *Place) QueryCategoryAssignments() *CategoryAssignmentQuery {
	return NewPlaceClient(pl.config).QueryCategoryAssignments(pl)
}

// QueryFaqs queries the "faqs" edge of the Place entity.
func (pl *Place) QueryFaqs() *FAQQuery {
	return NewPlaceClient(pl.config).QueryFaqs(pl)
}

// QueryLikedByUsers queries the "likedByUsers" edge of the Place entity.
func (pl *Place) QueryLikedByUsers() *UserLikePlaceQuery {
	return NewPlaceClient(pl.config).QueryLikedByUsers(pl)
}

// QueryFollowerUsers queries the "followerUsers" edge of the Place entity.
func (pl *Place) QueryFollowerUsers() *UserFollowPlaceQuery {
	return NewPlaceClient(pl.config).QueryFollowerUsers(pl)
}

// QueryRatings queries the "ratings" edge of the Place entity.
func (pl *Place) QueryRatings() *RatingQuery {
	return NewPlaceClient(pl.config).QueryRatings(pl)
}

// QueryInventories queries the "inventories" edge of the Place entity.
func (pl *Place) QueryInventories() *PlaceInventoryQuery {
	return NewPlaceClient(pl.config).QueryInventories(pl)
}

// Update returns a builder for updating this Place.
// Note that you need to call Place.Unwrap() before calling this method if this Place
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Place) Update() *PlaceUpdateOne {
	return NewPlaceClient(pl.config).UpdateOne(pl)
}

// Unwrap unwraps the Place entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pl *Place) Unwrap() *Place {
	_tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Place is not a transactional entity")
	}
	pl.config.driver = _tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Place) String() string {
	var builder strings.Builder
	builder.WriteString("Place(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pl.ID))
	builder.WriteString("name=")
	builder.WriteString(pl.Name)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(pl.Type)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pl.Description)
	builder.WriteString(", ")
	builder.WriteString("location=")
	builder.WriteString(pl.Location)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(pl.Email)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(pl.Phone)
	builder.WriteString(", ")
	builder.WriteString("website=")
	builder.WriteString(pl.Website)
	builder.WriteString(", ")
	builder.WriteString("cover_image=")
	builder.WriteString(pl.CoverImage)
	builder.WriteString(", ")
	builder.WriteString("picture=")
	builder.WriteString(pl.Picture)
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(pl.Country)
	builder.WriteString(", ")
	builder.WriteString("city=")
	builder.WriteString(pl.City)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(pl.State)
	builder.WriteString(", ")
	builder.WriteString("place_settings=")
	builder.WriteString(fmt.Sprintf("%v", pl.PlaceSettings))
	builder.WriteString(", ")
	builder.WriteString("opening_hours=")
	builder.WriteString(fmt.Sprintf("%v", pl.OpeningHours))
	builder.WriteString(", ")
	builder.WriteString("social_media=")
	builder.WriteString(fmt.Sprintf("%v", pl.SocialMedia))
	builder.WriteString(", ")
	builder.WriteString("payment_options=")
	builder.WriteString(fmt.Sprintf("%v", pl.PaymentOptions))
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", pl.Tags))
	builder.WriteString(", ")
	builder.WriteString("features=")
	builder.WriteString(fmt.Sprintf("%v", pl.Features))
	builder.WriteString(", ")
	builder.WriteString("additional_info=")
	builder.WriteString(fmt.Sprintf("%v", pl.AdditionalInfo))
	builder.WriteString(", ")
	builder.WriteString("images=")
	builder.WriteString(fmt.Sprintf("%v", pl.Images))
	builder.WriteString(", ")
	builder.WriteString("availability=")
	builder.WriteString(fmt.Sprintf("%v", pl.Availability))
	builder.WriteString(", ")
	builder.WriteString("special_offers=")
	builder.WriteString(pl.SpecialOffers)
	builder.WriteString(", ")
	builder.WriteString("sustainability_score=")
	builder.WriteString(fmt.Sprintf("%v", pl.SustainabilityScore))
	builder.WriteString(", ")
	builder.WriteString("map_coordinates=")
	builder.WriteString(fmt.Sprintf("%v", pl.MapCoordinates))
	builder.WriteString(", ")
	builder.WriteString("longitude=")
	builder.WriteString(pl.Longitude)
	builder.WriteString(", ")
	builder.WriteString("latitude=")
	builder.WriteString(pl.Latitude)
	builder.WriteString(", ")
	builder.WriteString("search_text=")
	builder.WriteString(pl.SearchText)
	builder.WriteString(", ")
	builder.WriteString("relevance_score=")
	builder.WriteString(fmt.Sprintf("%v", pl.RelevanceScore))
	builder.WriteString(", ")
	builder.WriteString("follower_count=")
	builder.WriteString(fmt.Sprintf("%v", pl.FollowerCount))
	builder.WriteString(", ")
	builder.WriteString("following_count=")
	builder.WriteString(fmt.Sprintf("%v", pl.FollowingCount))
	builder.WriteString(", ")
	builder.WriteString("is_Premium=")
	builder.WriteString(fmt.Sprintf("%v", pl.IsPremium))
	builder.WriteString(", ")
	builder.WriteString("is_published=")
	builder.WriteString(fmt.Sprintf("%v", pl.IsPublished))
	builder.WriteString(", ")
	builder.WriteString("likedByCurrentUser=")
	builder.WriteString(fmt.Sprintf("%v", pl.LikedByCurrentUser))
	builder.WriteString(", ")
	builder.WriteString("followedByCurrentUser=")
	builder.WriteString(fmt.Sprintf("%v", pl.FollowedByCurrentUser))
	builder.WriteByte(')')
	return builder.String()
}

// Places is a parsable slice of Place.
type Places []*Place
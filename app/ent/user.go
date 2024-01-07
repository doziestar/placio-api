// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"placio-app/ent/accountwallet"
	"placio-app/ent/event"
	"placio-app/ent/user"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/auth0/go-auth0/management"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Auth0ID holds the value of the "auth0_id" field.
	Auth0ID string `json:"auth0_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Picture holds the value of the "picture" field.
	Picture string `json:"picture,omitempty"`
	// CoverImage holds the value of the "cover_image" field.
	CoverImage string `json:"cover_image,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Website holds the value of the "website" field.
	Website string `json:"website,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// MapCoordinates holds the value of the "map_coordinates" field.
	MapCoordinates map[string]interface{} `json:"map_coordinates,omitempty"`
	// Longitude holds the value of the "longitude" field.
	Longitude string `json:"longitude,omitempty"`
	// Latitude holds the value of the "latitude" field.
	Latitude string `json:"latitude,omitempty"`
	// Bio holds the value of the "bio" field.
	Bio string `json:"bio,omitempty"`
	// Auth0Data holds the value of the "auth0_data" field.
	Auth0Data *management.User `json:"auth0_data,omitempty"`
	// AppSettings holds the value of the "app_settings" field.
	AppSettings map[string]interface{} `json:"app_settings,omitempty"`
	// UserSettings holds the value of the "user_settings" field.
	UserSettings map[string]interface{} `json:"user_settings,omitempty"`
	// SearchText holds the value of the "search_text" field.
	SearchText string `json:"search_text,omitempty"`
	// RelevanceScore holds the value of the "relevance_score" field.
	RelevanceScore float64 `json:"relevance_score,omitempty"`
	// FollowerCount holds the value of the "follower_count" field.
	FollowerCount int `json:"follower_count,omitempty"`
	// FollowingCount holds the value of the "following_count" field.
	FollowingCount int `json:"following_count,omitempty"`
	// Role holds the value of the "role" field.
	Role user.Role `json:"role,omitempty"`
	// Permissions holds the value of the "permissions" field.
	Permissions []string `json:"permissions,omitempty"`
	// IsPremium holds the value of the "is_premium" field.
	IsPremium bool `json:"is_premium,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// UserBusinesses holds the value of the userBusinesses edge.
	UserBusinesses []*UserBusiness `json:"userBusinesses,omitempty"`
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// Likes holds the value of the likes edge.
	Likes []*Like `json:"likes,omitempty"`
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// FollowedUsers holds the value of the followedUsers edge.
	FollowedUsers []*UserFollowUser `json:"followedUsers,omitempty"`
	// FollowerUsers holds the value of the followerUsers edge.
	FollowerUsers []*UserFollowUser `json:"followerUsers,omitempty"`
	// FollowedBusinesses holds the value of the followedBusinesses edge.
	FollowedBusinesses []*UserFollowBusiness `json:"followedBusinesses,omitempty"`
	// FollowerBusinesses holds the value of the followerBusinesses edge.
	FollowerBusinesses []*BusinessFollowUser `json:"followerBusinesses,omitempty"`
	// Reviews holds the value of the reviews edge.
	Reviews []*Review `json:"reviews,omitempty"`
	// Bookings holds the value of the bookings edge.
	Bookings []*Booking `json:"bookings,omitempty"`
	// Reservations holds the value of the reservations edge.
	Reservations []*Reservation `json:"reservations,omitempty"`
	// Helps holds the value of the helps edge.
	Helps []*Help `json:"helps,omitempty"`
	// Categories holds the value of the categories edge.
	Categories []*Category `json:"categories,omitempty"`
	// Places holds the value of the places edge.
	Places []*Place `json:"places,omitempty"`
	// CategoryAssignments holds the value of the categoryAssignments edge.
	CategoryAssignments []*CategoryAssignment `json:"categoryAssignments,omitempty"`
	// OwnedEvents holds the value of the ownedEvents edge.
	OwnedEvents *Event `json:"ownedEvents,omitempty"`
	// UserFollowEvents holds the value of the userFollowEvents edge.
	UserFollowEvents []*UserFollowEvent `json:"userFollowEvents,omitempty"`
	// FollowedPlaces holds the value of the followedPlaces edge.
	FollowedPlaces []*UserFollowPlace `json:"followedPlaces,omitempty"`
	// LikedPlaces holds the value of the likedPlaces edge.
	LikedPlaces []*UserLikePlace `json:"likedPlaces,omitempty"`
	// Ratings holds the value of the ratings edge.
	Ratings []*Rating `json:"ratings,omitempty"`
	// TransactionHistories holds the value of the transaction_histories edge.
	TransactionHistories []*TransactionHistory `json:"transaction_histories,omitempty"`
	// ReservationBlocks holds the value of the reservation_blocks edge.
	ReservationBlocks []*ReservationBlock `json:"reservation_blocks,omitempty"`
	// Notifications holds the value of the notifications edge.
	Notifications []*Notification `json:"notifications,omitempty"`
	// Wallet holds the value of the wallet edge.
	Wallet *AccountWallet `json:"wallet,omitempty"`
	// Orders holds the value of the orders edge.
	Orders []*Order `json:"orders,omitempty"`
	// TablesCreated holds the value of the tables_created edge.
	TablesCreated []*PlaceTable `json:"tables_created,omitempty"`
	// TablesUpdated holds the value of the tables_updated edge.
	TablesUpdated []*PlaceTable `json:"tables_updated,omitempty"`
	// TablesDeleted holds the value of the tables_deleted edge.
	TablesDeleted []*PlaceTable `json:"tables_deleted,omitempty"`
	// TablesReserved holds the value of the tables_reserved edge.
	TablesReserved []*PlaceTable `json:"tables_reserved,omitempty"`
	// TablesWaited holds the value of the tables_waited edge.
	TablesWaited []*PlaceTable `json:"tables_waited,omitempty"`
	// Staffs holds the value of the staffs edge.
	Staffs []*Staff `json:"staffs,omitempty"`
	// CreatedMenus holds the value of the created_menus edge.
	CreatedMenus []*Menu `json:"created_menus,omitempty"`
	// UpdatedMenus holds the value of the updated_menus edge.
	UpdatedMenus []*Menu `json:"updated_menus,omitempty"`
	// Plans holds the value of the plans edge.
	Plans []*Plan `json:"plans,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [34]bool
}

// UserBusinessesOrErr returns the UserBusinesses value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserBusinessesOrErr() ([]*UserBusiness, error) {
	if e.loadedTypes[0] {
		return e.UserBusinesses, nil
	}
	return nil, &NotLoadedError{edge: "userBusinesses"}
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[1] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// LikesOrErr returns the Likes value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) LikesOrErr() ([]*Like, error) {
	if e.loadedTypes[2] {
		return e.Likes, nil
	}
	return nil, &NotLoadedError{edge: "likes"}
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[3] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// FollowedUsersOrErr returns the FollowedUsers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowedUsersOrErr() ([]*UserFollowUser, error) {
	if e.loadedTypes[4] {
		return e.FollowedUsers, nil
	}
	return nil, &NotLoadedError{edge: "followedUsers"}
}

// FollowerUsersOrErr returns the FollowerUsers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowerUsersOrErr() ([]*UserFollowUser, error) {
	if e.loadedTypes[5] {
		return e.FollowerUsers, nil
	}
	return nil, &NotLoadedError{edge: "followerUsers"}
}

// FollowedBusinessesOrErr returns the FollowedBusinesses value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowedBusinessesOrErr() ([]*UserFollowBusiness, error) {
	if e.loadedTypes[6] {
		return e.FollowedBusinesses, nil
	}
	return nil, &NotLoadedError{edge: "followedBusinesses"}
}

// FollowerBusinessesOrErr returns the FollowerBusinesses value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowerBusinessesOrErr() ([]*BusinessFollowUser, error) {
	if e.loadedTypes[7] {
		return e.FollowerBusinesses, nil
	}
	return nil, &NotLoadedError{edge: "followerBusinesses"}
}

// ReviewsOrErr returns the Reviews value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ReviewsOrErr() ([]*Review, error) {
	if e.loadedTypes[8] {
		return e.Reviews, nil
	}
	return nil, &NotLoadedError{edge: "reviews"}
}

// BookingsOrErr returns the Bookings value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) BookingsOrErr() ([]*Booking, error) {
	if e.loadedTypes[9] {
		return e.Bookings, nil
	}
	return nil, &NotLoadedError{edge: "bookings"}
}

// ReservationsOrErr returns the Reservations value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ReservationsOrErr() ([]*Reservation, error) {
	if e.loadedTypes[10] {
		return e.Reservations, nil
	}
	return nil, &NotLoadedError{edge: "reservations"}
}

// HelpsOrErr returns the Helps value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) HelpsOrErr() ([]*Help, error) {
	if e.loadedTypes[11] {
		return e.Helps, nil
	}
	return nil, &NotLoadedError{edge: "helps"}
}

// CategoriesOrErr returns the Categories value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CategoriesOrErr() ([]*Category, error) {
	if e.loadedTypes[12] {
		return e.Categories, nil
	}
	return nil, &NotLoadedError{edge: "categories"}
}

// PlacesOrErr returns the Places value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PlacesOrErr() ([]*Place, error) {
	if e.loadedTypes[13] {
		return e.Places, nil
	}
	return nil, &NotLoadedError{edge: "places"}
}

// CategoryAssignmentsOrErr returns the CategoryAssignments value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CategoryAssignmentsOrErr() ([]*CategoryAssignment, error) {
	if e.loadedTypes[14] {
		return e.CategoryAssignments, nil
	}
	return nil, &NotLoadedError{edge: "categoryAssignments"}
}

// OwnedEventsOrErr returns the OwnedEvents value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) OwnedEventsOrErr() (*Event, error) {
	if e.loadedTypes[15] {
		if e.OwnedEvents == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: event.Label}
		}
		return e.OwnedEvents, nil
	}
	return nil, &NotLoadedError{edge: "ownedEvents"}
}

// UserFollowEventsOrErr returns the UserFollowEvents value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserFollowEventsOrErr() ([]*UserFollowEvent, error) {
	if e.loadedTypes[16] {
		return e.UserFollowEvents, nil
	}
	return nil, &NotLoadedError{edge: "userFollowEvents"}
}

// FollowedPlacesOrErr returns the FollowedPlaces value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowedPlacesOrErr() ([]*UserFollowPlace, error) {
	if e.loadedTypes[17] {
		return e.FollowedPlaces, nil
	}
	return nil, &NotLoadedError{edge: "followedPlaces"}
}

// LikedPlacesOrErr returns the LikedPlaces value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) LikedPlacesOrErr() ([]*UserLikePlace, error) {
	if e.loadedTypes[18] {
		return e.LikedPlaces, nil
	}
	return nil, &NotLoadedError{edge: "likedPlaces"}
}

// RatingsOrErr returns the Ratings value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RatingsOrErr() ([]*Rating, error) {
	if e.loadedTypes[19] {
		return e.Ratings, nil
	}
	return nil, &NotLoadedError{edge: "ratings"}
}

// TransactionHistoriesOrErr returns the TransactionHistories value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TransactionHistoriesOrErr() ([]*TransactionHistory, error) {
	if e.loadedTypes[20] {
		return e.TransactionHistories, nil
	}
	return nil, &NotLoadedError{edge: "transaction_histories"}
}

// ReservationBlocksOrErr returns the ReservationBlocks value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ReservationBlocksOrErr() ([]*ReservationBlock, error) {
	if e.loadedTypes[21] {
		return e.ReservationBlocks, nil
	}
	return nil, &NotLoadedError{edge: "reservation_blocks"}
}

// NotificationsOrErr returns the Notifications value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) NotificationsOrErr() ([]*Notification, error) {
	if e.loadedTypes[22] {
		return e.Notifications, nil
	}
	return nil, &NotLoadedError{edge: "notifications"}
}

// WalletOrErr returns the Wallet value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) WalletOrErr() (*AccountWallet, error) {
	if e.loadedTypes[23] {
		if e.Wallet == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: accountwallet.Label}
		}
		return e.Wallet, nil
	}
	return nil, &NotLoadedError{edge: "wallet"}
}

// OrdersOrErr returns the Orders value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OrdersOrErr() ([]*Order, error) {
	if e.loadedTypes[24] {
		return e.Orders, nil
	}
	return nil, &NotLoadedError{edge: "orders"}
}

// TablesCreatedOrErr returns the TablesCreated value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TablesCreatedOrErr() ([]*PlaceTable, error) {
	if e.loadedTypes[25] {
		return e.TablesCreated, nil
	}
	return nil, &NotLoadedError{edge: "tables_created"}
}

// TablesUpdatedOrErr returns the TablesUpdated value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TablesUpdatedOrErr() ([]*PlaceTable, error) {
	if e.loadedTypes[26] {
		return e.TablesUpdated, nil
	}
	return nil, &NotLoadedError{edge: "tables_updated"}
}

// TablesDeletedOrErr returns the TablesDeleted value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TablesDeletedOrErr() ([]*PlaceTable, error) {
	if e.loadedTypes[27] {
		return e.TablesDeleted, nil
	}
	return nil, &NotLoadedError{edge: "tables_deleted"}
}

// TablesReservedOrErr returns the TablesReserved value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TablesReservedOrErr() ([]*PlaceTable, error) {
	if e.loadedTypes[28] {
		return e.TablesReserved, nil
	}
	return nil, &NotLoadedError{edge: "tables_reserved"}
}

// TablesWaitedOrErr returns the TablesWaited value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TablesWaitedOrErr() ([]*PlaceTable, error) {
	if e.loadedTypes[29] {
		return e.TablesWaited, nil
	}
	return nil, &NotLoadedError{edge: "tables_waited"}
}

// StaffsOrErr returns the Staffs value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) StaffsOrErr() ([]*Staff, error) {
	if e.loadedTypes[30] {
		return e.Staffs, nil
	}
	return nil, &NotLoadedError{edge: "staffs"}
}

// CreatedMenusOrErr returns the CreatedMenus value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CreatedMenusOrErr() ([]*Menu, error) {
	if e.loadedTypes[31] {
		return e.CreatedMenus, nil
	}
	return nil, &NotLoadedError{edge: "created_menus"}
}

// UpdatedMenusOrErr returns the UpdatedMenus value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UpdatedMenusOrErr() ([]*Menu, error) {
	if e.loadedTypes[32] {
		return e.UpdatedMenus, nil
	}
	return nil, &NotLoadedError{edge: "updated_menus"}
}

// PlansOrErr returns the Plans value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PlansOrErr() ([]*Plan, error) {
	if e.loadedTypes[33] {
		return e.Plans, nil
	}
	return nil, &NotLoadedError{edge: "plans"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldMapCoordinates, user.FieldAuth0Data, user.FieldAppSettings, user.FieldUserSettings, user.FieldPermissions:
			values[i] = new([]byte)
		case user.FieldIsPremium:
			values[i] = new(sql.NullBool)
		case user.FieldRelevanceScore:
			values[i] = new(sql.NullFloat64)
		case user.FieldFollowerCount, user.FieldFollowingCount:
			values[i] = new(sql.NullInt64)
		case user.FieldID, user.FieldAuth0ID, user.FieldName, user.FieldPicture, user.FieldCoverImage, user.FieldUsername, user.FieldWebsite, user.FieldLocation, user.FieldLongitude, user.FieldLatitude, user.FieldBio, user.FieldSearchText, user.FieldRole:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				u.ID = value.String
			}
		case user.FieldAuth0ID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field auth0_id", values[i])
			} else if value.Valid {
				u.Auth0ID = value.String
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldPicture:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field picture", values[i])
			} else if value.Valid {
				u.Picture = value.String
			}
		case user.FieldCoverImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cover_image", values[i])
			} else if value.Valid {
				u.CoverImage = value.String
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldWebsite:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field website", values[i])
			} else if value.Valid {
				u.Website = value.String
			}
		case user.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				u.Location = value.String
			}
		case user.FieldMapCoordinates:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field map_coordinates", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &u.MapCoordinates); err != nil {
					return fmt.Errorf("unmarshal field map_coordinates: %w", err)
				}
			}
		case user.FieldLongitude:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field longitude", values[i])
			} else if value.Valid {
				u.Longitude = value.String
			}
		case user.FieldLatitude:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field latitude", values[i])
			} else if value.Valid {
				u.Latitude = value.String
			}
		case user.FieldBio:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bio", values[i])
			} else if value.Valid {
				u.Bio = value.String
			}
		case user.FieldAuth0Data:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field auth0_data", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &u.Auth0Data); err != nil {
					return fmt.Errorf("unmarshal field auth0_data: %w", err)
				}
			}
		case user.FieldAppSettings:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field app_settings", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &u.AppSettings); err != nil {
					return fmt.Errorf("unmarshal field app_settings: %w", err)
				}
			}
		case user.FieldUserSettings:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field user_settings", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &u.UserSettings); err != nil {
					return fmt.Errorf("unmarshal field user_settings: %w", err)
				}
			}
		case user.FieldSearchText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field search_text", values[i])
			} else if value.Valid {
				u.SearchText = value.String
			}
		case user.FieldRelevanceScore:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field relevance_score", values[i])
			} else if value.Valid {
				u.RelevanceScore = value.Float64
			}
		case user.FieldFollowerCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field follower_count", values[i])
			} else if value.Valid {
				u.FollowerCount = int(value.Int64)
			}
		case user.FieldFollowingCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field following_count", values[i])
			} else if value.Valid {
				u.FollowingCount = int(value.Int64)
			}
		case user.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				u.Role = user.Role(value.String)
			}
		case user.FieldPermissions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field permissions", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &u.Permissions); err != nil {
					return fmt.Errorf("unmarshal field permissions: %w", err)
				}
			}
		case user.FieldIsPremium:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_premium", values[i])
			} else if value.Valid {
				u.IsPremium = value.Bool
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryUserBusinesses queries the "userBusinesses" edge of the User entity.
func (u *User) QueryUserBusinesses() *UserBusinessQuery {
	return NewUserClient(u.config).QueryUserBusinesses(u)
}

// QueryComments queries the "comments" edge of the User entity.
func (u *User) QueryComments() *CommentQuery {
	return NewUserClient(u.config).QueryComments(u)
}

// QueryLikes queries the "likes" edge of the User entity.
func (u *User) QueryLikes() *LikeQuery {
	return NewUserClient(u.config).QueryLikes(u)
}

// QueryPosts queries the "posts" edge of the User entity.
func (u *User) QueryPosts() *PostQuery {
	return NewUserClient(u.config).QueryPosts(u)
}

// QueryFollowedUsers queries the "followedUsers" edge of the User entity.
func (u *User) QueryFollowedUsers() *UserFollowUserQuery {
	return NewUserClient(u.config).QueryFollowedUsers(u)
}

// QueryFollowerUsers queries the "followerUsers" edge of the User entity.
func (u *User) QueryFollowerUsers() *UserFollowUserQuery {
	return NewUserClient(u.config).QueryFollowerUsers(u)
}

// QueryFollowedBusinesses queries the "followedBusinesses" edge of the User entity.
func (u *User) QueryFollowedBusinesses() *UserFollowBusinessQuery {
	return NewUserClient(u.config).QueryFollowedBusinesses(u)
}

// QueryFollowerBusinesses queries the "followerBusinesses" edge of the User entity.
func (u *User) QueryFollowerBusinesses() *BusinessFollowUserQuery {
	return NewUserClient(u.config).QueryFollowerBusinesses(u)
}

// QueryReviews queries the "reviews" edge of the User entity.
func (u *User) QueryReviews() *ReviewQuery {
	return NewUserClient(u.config).QueryReviews(u)
}

// QueryBookings queries the "bookings" edge of the User entity.
func (u *User) QueryBookings() *BookingQuery {
	return NewUserClient(u.config).QueryBookings(u)
}

// QueryReservations queries the "reservations" edge of the User entity.
func (u *User) QueryReservations() *ReservationQuery {
	return NewUserClient(u.config).QueryReservations(u)
}

// QueryHelps queries the "helps" edge of the User entity.
func (u *User) QueryHelps() *HelpQuery {
	return NewUserClient(u.config).QueryHelps(u)
}

// QueryCategories queries the "categories" edge of the User entity.
func (u *User) QueryCategories() *CategoryQuery {
	return NewUserClient(u.config).QueryCategories(u)
}

// QueryPlaces queries the "places" edge of the User entity.
func (u *User) QueryPlaces() *PlaceQuery {
	return NewUserClient(u.config).QueryPlaces(u)
}

// QueryCategoryAssignments queries the "categoryAssignments" edge of the User entity.
func (u *User) QueryCategoryAssignments() *CategoryAssignmentQuery {
	return NewUserClient(u.config).QueryCategoryAssignments(u)
}

// QueryOwnedEvents queries the "ownedEvents" edge of the User entity.
func (u *User) QueryOwnedEvents() *EventQuery {
	return NewUserClient(u.config).QueryOwnedEvents(u)
}

// QueryUserFollowEvents queries the "userFollowEvents" edge of the User entity.
func (u *User) QueryUserFollowEvents() *UserFollowEventQuery {
	return NewUserClient(u.config).QueryUserFollowEvents(u)
}

// QueryFollowedPlaces queries the "followedPlaces" edge of the User entity.
func (u *User) QueryFollowedPlaces() *UserFollowPlaceQuery {
	return NewUserClient(u.config).QueryFollowedPlaces(u)
}

// QueryLikedPlaces queries the "likedPlaces" edge of the User entity.
func (u *User) QueryLikedPlaces() *UserLikePlaceQuery {
	return NewUserClient(u.config).QueryLikedPlaces(u)
}

// QueryRatings queries the "ratings" edge of the User entity.
func (u *User) QueryRatings() *RatingQuery {
	return NewUserClient(u.config).QueryRatings(u)
}

// QueryTransactionHistories queries the "transaction_histories" edge of the User entity.
func (u *User) QueryTransactionHistories() *TransactionHistoryQuery {
	return NewUserClient(u.config).QueryTransactionHistories(u)
}

// QueryReservationBlocks queries the "reservation_blocks" edge of the User entity.
func (u *User) QueryReservationBlocks() *ReservationBlockQuery {
	return NewUserClient(u.config).QueryReservationBlocks(u)
}

// QueryNotifications queries the "notifications" edge of the User entity.
func (u *User) QueryNotifications() *NotificationQuery {
	return NewUserClient(u.config).QueryNotifications(u)
}

// QueryWallet queries the "wallet" edge of the User entity.
func (u *User) QueryWallet() *AccountWalletQuery {
	return NewUserClient(u.config).QueryWallet(u)
}

// QueryOrders queries the "orders" edge of the User entity.
func (u *User) QueryOrders() *OrderQuery {
	return NewUserClient(u.config).QueryOrders(u)
}

// QueryTablesCreated queries the "tables_created" edge of the User entity.
func (u *User) QueryTablesCreated() *PlaceTableQuery {
	return NewUserClient(u.config).QueryTablesCreated(u)
}

// QueryTablesUpdated queries the "tables_updated" edge of the User entity.
func (u *User) QueryTablesUpdated() *PlaceTableQuery {
	return NewUserClient(u.config).QueryTablesUpdated(u)
}

// QueryTablesDeleted queries the "tables_deleted" edge of the User entity.
func (u *User) QueryTablesDeleted() *PlaceTableQuery {
	return NewUserClient(u.config).QueryTablesDeleted(u)
}

// QueryTablesReserved queries the "tables_reserved" edge of the User entity.
func (u *User) QueryTablesReserved() *PlaceTableQuery {
	return NewUserClient(u.config).QueryTablesReserved(u)
}

// QueryTablesWaited queries the "tables_waited" edge of the User entity.
func (u *User) QueryTablesWaited() *PlaceTableQuery {
	return NewUserClient(u.config).QueryTablesWaited(u)
}

// QueryStaffs queries the "staffs" edge of the User entity.
func (u *User) QueryStaffs() *StaffQuery {
	return NewUserClient(u.config).QueryStaffs(u)
}

// QueryCreatedMenus queries the "created_menus" edge of the User entity.
func (u *User) QueryCreatedMenus() *MenuQuery {
	return NewUserClient(u.config).QueryCreatedMenus(u)
}

// QueryUpdatedMenus queries the "updated_menus" edge of the User entity.
func (u *User) QueryUpdatedMenus() *MenuQuery {
	return NewUserClient(u.config).QueryUpdatedMenus(u)
}

// QueryPlans queries the "plans" edge of the User entity.
func (u *User) QueryPlans() *PlanQuery {
	return NewUserClient(u.config).QueryPlans(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("auth0_id=")
	builder.WriteString(u.Auth0ID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("picture=")
	builder.WriteString(u.Picture)
	builder.WriteString(", ")
	builder.WriteString("cover_image=")
	builder.WriteString(u.CoverImage)
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("website=")
	builder.WriteString(u.Website)
	builder.WriteString(", ")
	builder.WriteString("location=")
	builder.WriteString(u.Location)
	builder.WriteString(", ")
	builder.WriteString("map_coordinates=")
	builder.WriteString(fmt.Sprintf("%v", u.MapCoordinates))
	builder.WriteString(", ")
	builder.WriteString("longitude=")
	builder.WriteString(u.Longitude)
	builder.WriteString(", ")
	builder.WriteString("latitude=")
	builder.WriteString(u.Latitude)
	builder.WriteString(", ")
	builder.WriteString("bio=")
	builder.WriteString(u.Bio)
	builder.WriteString(", ")
	builder.WriteString("auth0_data=")
	builder.WriteString(fmt.Sprintf("%v", u.Auth0Data))
	builder.WriteString(", ")
	builder.WriteString("app_settings=")
	builder.WriteString(fmt.Sprintf("%v", u.AppSettings))
	builder.WriteString(", ")
	builder.WriteString("user_settings=")
	builder.WriteString(fmt.Sprintf("%v", u.UserSettings))
	builder.WriteString(", ")
	builder.WriteString("search_text=")
	builder.WriteString(u.SearchText)
	builder.WriteString(", ")
	builder.WriteString("relevance_score=")
	builder.WriteString(fmt.Sprintf("%v", u.RelevanceScore))
	builder.WriteString(", ")
	builder.WriteString("follower_count=")
	builder.WriteString(fmt.Sprintf("%v", u.FollowerCount))
	builder.WriteString(", ")
	builder.WriteString("following_count=")
	builder.WriteString(fmt.Sprintf("%v", u.FollowingCount))
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", u.Role))
	builder.WriteString(", ")
	builder.WriteString("permissions=")
	builder.WriteString(fmt.Sprintf("%v", u.Permissions))
	builder.WriteString(", ")
	builder.WriteString("is_premium=")
	builder.WriteString(fmt.Sprintf("%v", u.IsPremium))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

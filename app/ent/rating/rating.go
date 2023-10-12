// Code generated by ent, DO NOT EDIT.

package rating

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the rating type in the database.
	Label = "rating"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldScore holds the string denoting the score field in the database.
	FieldScore = "score"
	// FieldReview holds the string denoting the review field in the database.
	FieldReview = "review"
	// FieldRatedAt holds the string denoting the ratedat field in the database.
	FieldRatedAt = "rated_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeBusiness holds the string denoting the business edge name in mutations.
	EdgeBusiness = "business"
	// EdgePlace holds the string denoting the place edge name in mutations.
	EdgePlace = "place"
	// EdgeEvent holds the string denoting the event edge name in mutations.
	EdgeEvent = "event"
	// Table holds the table name of the rating in the database.
	Table = "ratings"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "ratings"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_ratings"
	// BusinessTable is the table that holds the business relation/edge.
	BusinessTable = "ratings"
	// BusinessInverseTable is the table name for the Business entity.
	// It exists in this package in order to avoid circular dependency with the "business" package.
	BusinessInverseTable = "businesses"
	// BusinessColumn is the table column denoting the business relation/edge.
	BusinessColumn = "rating_business"
	// PlaceTable is the table that holds the place relation/edge.
	PlaceTable = "ratings"
	// PlaceInverseTable is the table name for the Place entity.
	// It exists in this package in order to avoid circular dependency with the "place" package.
	PlaceInverseTable = "places"
	// PlaceColumn is the table column denoting the place relation/edge.
	PlaceColumn = "rating_place"
	// EventTable is the table that holds the event relation/edge.
	EventTable = "ratings"
	// EventInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventInverseTable = "events"
	// EventColumn is the table column denoting the event relation/edge.
	EventColumn = "rating_event"
)

// Columns holds all SQL columns for rating fields.
var Columns = []string{
	FieldID,
	FieldScore,
	FieldReview,
	FieldRatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "ratings"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"business_ratings",
	"event_ratings",
	"place_ratings",
	"rating_business",
	"rating_place",
	"rating_event",
	"user_ratings",
}

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
	// ScoreValidator is a validator for the "score" field. It is called by the builders before save.
	ScoreValidator func(int) error
	// DefaultRatedAt holds the default value on creation for the "ratedAt" field.
	DefaultRatedAt func() time.Time
)

// OrderOption defines the ordering options for the Rating queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByScore orders the results by the score field.
func ByScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScore, opts...).ToFunc()
}

// ByReview orders the results by the review field.
func ByReview(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReview, opts...).ToFunc()
}

// ByRatedAt orders the results by the ratedAt field.
func ByRatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRatedAt, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByBusinessField orders the results by business field.
func ByBusinessField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessStep(), sql.OrderByField(field, opts...))
	}
}

// ByPlaceField orders the results by place field.
func ByPlaceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlaceStep(), sql.OrderByField(field, opts...))
	}
}

// ByEventField orders the results by event field.
func ByEventField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEventStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newBusinessStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, BusinessTable, BusinessColumn),
	)
}
func newPlaceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlaceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, PlaceTable, PlaceColumn),
	)
}
func newEventStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EventInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, EventTable, EventColumn),
	)
}
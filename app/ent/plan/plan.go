// Code generated by ent, DO NOT EDIT.

package plan

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the plan type in the database.
	Label = "plan"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldOverview holds the string denoting the overview field in the database.
	FieldOverview = "overview"
	// FieldFeatures holds the string denoting the features field in the database.
	FieldFeatures = "features"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeBusinesses holds the string denoting the businesses edge name in mutations.
	EdgeBusinesses = "businesses"
	// EdgePlaces holds the string denoting the places edge name in mutations.
	EdgePlaces = "places"
	// EdgeMedia holds the string denoting the media edge name in mutations.
	EdgeMedia = "media"
	// Table holds the table name of the plan in the database.
	Table = "plans"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "plan_users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// BusinessesTable is the table that holds the businesses relation/edge. The primary key declared below.
	BusinessesTable = "plan_businesses"
	// BusinessesInverseTable is the table name for the Business entity.
	// It exists in this package in order to avoid circular dependency with the "business" package.
	BusinessesInverseTable = "businesses"
	// PlacesTable is the table that holds the places relation/edge. The primary key declared below.
	PlacesTable = "plan_places"
	// PlacesInverseTable is the table name for the Place entity.
	// It exists in this package in order to avoid circular dependency with the "place" package.
	PlacesInverseTable = "places"
	// MediaTable is the table that holds the media relation/edge.
	MediaTable = "media"
	// MediaInverseTable is the table name for the Media entity.
	// It exists in this package in order to avoid circular dependency with the "media" package.
	MediaInverseTable = "media"
	// MediaColumn is the table column denoting the media relation/edge.
	MediaColumn = "plan_media"
)

// Columns holds all SQL columns for plan fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldOverview,
	FieldFeatures,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"plan_id", "user_id"}
	// BusinessesPrimaryKey and BusinessesColumn2 are the table columns denoting the
	// primary key for the businesses relation (M2M).
	BusinessesPrimaryKey = []string{"plan_id", "business_id"}
	// PlacesPrimaryKey and PlacesColumn2 are the table columns denoting the
	// primary key for the places relation (M2M).
	PlacesPrimaryKey = []string{"plan_id", "place_id"}
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

var (
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Plan queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByOverview orders the results by the overview field.
func ByOverview(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOverview, opts...).ToFunc()
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

// ByBusinessesCount orders the results by businesses count.
func ByBusinessesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBusinessesStep(), opts...)
	}
}

// ByBusinesses orders the results by businesses terms.
func ByBusinesses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessesStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByMediaCount orders the results by media count.
func ByMediaCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMediaStep(), opts...)
	}
}

// ByMedia orders the results by media terms.
func ByMedia(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMediaStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, UsersTable, UsersPrimaryKey...),
	)
}
func newBusinessesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, BusinessesTable, BusinessesPrimaryKey...),
	)
}
func newPlacesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlacesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PlacesTable, PlacesPrimaryKey...),
	)
}
func newMediaStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MediaInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MediaTable, MediaColumn),
	)
}

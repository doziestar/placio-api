// Code generated by ent, DO NOT EDIT.

package userbusiness

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the userbusiness type in the database.
	Label = "user_business"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeBusiness holds the string denoting the business edge name in mutations.
	EdgeBusiness = "business"
	// Table holds the table name of the userbusiness in the database.
	Table = "user_businesses"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "user_businesses"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_user_businesses"
	// BusinessTable is the table that holds the business relation/edge.
	BusinessTable = "user_businesses"
	// BusinessInverseTable is the table name for the Business entity.
	// It exists in this package in order to avoid circular dependency with the "business" package.
	BusinessInverseTable = "businesses"
	// BusinessColumn is the table column denoting the business relation/edge.
	BusinessColumn = "business_user_businesses"
)

// Columns holds all SQL columns for userbusiness fields.
var Columns = []string{
	FieldID,
	FieldRole,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "user_businesses"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"business_user_businesses",
	"user_user_businesses",
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

// OrderOption defines the ordering options for the UserBusiness queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
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
		sqlgraph.Edge(sqlgraph.M2O, true, BusinessTable, BusinessColumn),
	)
}

// Code generated by ent, DO NOT EDIT.

package transactionhistory

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the transactionhistory type in the database.
	Label = "transaction_history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTransactionType holds the string denoting the transaction_type field in the database.
	FieldTransactionType = "transaction_type"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// EdgePlaceInventory holds the string denoting the place_inventory edge name in mutations.
	EdgePlaceInventory = "place_inventory"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the transactionhistory in the database.
	Table = "transaction_histories"
	// PlaceInventoryTable is the table that holds the place_inventory relation/edge.
	PlaceInventoryTable = "transaction_histories"
	// PlaceInventoryInverseTable is the table name for the PlaceInventory entity.
	// It exists in this package in order to avoid circular dependency with the "placeinventory" package.
	PlaceInventoryInverseTable = "place_inventories"
	// PlaceInventoryColumn is the table column denoting the place_inventory relation/edge.
	PlaceInventoryColumn = "place_inventory_transaction_histories"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "transaction_histories"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_transaction_histories"
)

// Columns holds all SQL columns for transactionhistory fields.
var Columns = []string{
	FieldID,
	FieldTransactionType,
	FieldQuantity,
	FieldDate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "transaction_histories"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"place_inventory_transaction_histories",
	"user_transaction_histories",
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

var (
	// DefaultDate holds the default value on creation for the "date" field.
	DefaultDate func() time.Time
)

// TransactionType defines the type for the "transaction_type" enum field.
type TransactionType string

// TransactionType values.
const (
	TransactionTypePurchase TransactionType = "purchase"
	TransactionTypeSale     TransactionType = "sale"
	TransactionTypeReturn   TransactionType = "return"
	TransactionTypeUsage    TransactionType = "usage"
)

func (tt TransactionType) String() string {
	return string(tt)
}

// TransactionTypeValidator is a validator for the "transaction_type" field enum values. It is called by the builders before save.
func TransactionTypeValidator(tt TransactionType) error {
	switch tt {
	case TransactionTypePurchase, TransactionTypeSale, TransactionTypeReturn, TransactionTypeUsage:
		return nil
	default:
		return fmt.Errorf("transactionhistory: invalid enum value for transaction_type field: %q", tt)
	}
}

// OrderOption defines the ordering options for the TransactionHistory queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTransactionType orders the results by the transaction_type field.
func ByTransactionType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTransactionType, opts...).ToFunc()
}

// ByQuantity orders the results by the quantity field.
func ByQuantity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantity, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}

// ByPlaceInventoryField orders the results by place_inventory field.
func ByPlaceInventoryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlaceInventoryStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newPlaceInventoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlaceInventoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PlaceInventoryTable, PlaceInventoryColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}

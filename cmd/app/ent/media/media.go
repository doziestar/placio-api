// Code generated by ent, DO NOT EDIT.

package media

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the media type in the database.
	Label = "media"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldMediaType holds the string denoting the mediatype field in the database.
	FieldMediaType = "media_type"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgePost holds the string denoting the post edge name in mutations.
	EdgePost = "post"
	// EdgeCategories holds the string denoting the categories edge name in mutations.
	EdgeCategories = "categories"
	// Table holds the table name of the media in the database.
	Table = "media"
	// PostTable is the table that holds the post relation/edge.
	PostTable = "media"
	// PostInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostInverseTable = "posts"
	// PostColumn is the table column denoting the post relation/edge.
	PostColumn = "post_medias"
	// CategoriesTable is the table that holds the categories relation/edge.
	CategoriesTable = "categories"
	// CategoriesInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoriesInverseTable = "categories"
	// CategoriesColumn is the table column denoting the categories relation/edge.
	CategoriesColumn = "media_categories"
)

// Columns holds all SQL columns for media fields.
var Columns = []string{
	FieldID,
	FieldURL,
	FieldMediaType,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "media"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"post_medias",
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
	// DefaultCreatedAt holds the default value on creation for the "CreatedAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "UpdatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "UpdatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Media queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByURL orders the results by the URL field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByMediaType orders the results by the MediaType field.
func ByMediaType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMediaType, opts...).ToFunc()
}

// ByCreatedAt orders the results by the CreatedAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the UpdatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByPostField orders the results by post field.
func ByPostField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostStep(), sql.OrderByField(field, opts...))
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
func newPostStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PostTable, PostColumn),
	)
}
func newCategoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CategoriesTable, CategoriesColumn),
	)
}

// Code generated by ent, DO NOT EDIT.

package inventorytype

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the inventorytype type in the database.
	Label = "inventory_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldImageURL holds the string denoting the image_url field in the database.
	FieldImageURL = "image_url"
	// FieldIconURL holds the string denoting the icon_url field in the database.
	FieldIconURL = "icon_url"
	// FieldIndustryType holds the string denoting the industry_type field in the database.
	FieldIndustryType = "industry_type"
	// FieldMeasurementUnit holds the string denoting the measurement_unit field in the database.
	FieldMeasurementUnit = "measurement_unit"
	// EdgeAttributes holds the string denoting the attributes edge name in mutations.
	EdgeAttributes = "attributes"
	// EdgePlaceInventories holds the string denoting the place_inventories edge name in mutations.
	EdgePlaceInventories = "place_inventories"
	// Table holds the table name of the inventorytype in the database.
	Table = "inventory_types"
	// AttributesTable is the table that holds the attributes relation/edge.
	AttributesTable = "inventory_attributes"
	// AttributesInverseTable is the table name for the InventoryAttribute entity.
	// It exists in this package in order to avoid circular dependency with the "inventoryattribute" package.
	AttributesInverseTable = "inventory_attributes"
	// AttributesColumn is the table column denoting the attributes relation/edge.
	AttributesColumn = "inventory_type_attributes"
	// PlaceInventoriesTable is the table that holds the place_inventories relation/edge.
	PlaceInventoriesTable = "place_inventories"
	// PlaceInventoriesInverseTable is the table name for the PlaceInventory entity.
	// It exists in this package in order to avoid circular dependency with the "placeinventory" package.
	PlaceInventoriesInverseTable = "place_inventories"
	// PlaceInventoriesColumn is the table column denoting the place_inventories relation/edge.
	PlaceInventoriesColumn = "inventory_type_place_inventories"
)

// Columns holds all SQL columns for inventorytype fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldImageURL,
	FieldIconURL,
	FieldIndustryType,
	FieldMeasurementUnit,
}

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

// IndustryType defines the type for the "industry_type" enum field.
type IndustryType string

// IndustryType values.
const (
	IndustryTypeHotel      IndustryType = "hotel"
	IndustryTypeRestaurant IndustryType = "restaurant"
	IndustryTypeBar        IndustryType = "bar"
	IndustryTypeClub       IndustryType = "club"
	IndustryTypeGym        IndustryType = "gym"
	IndustryTypeEvents     IndustryType = "events"
	IndustryTypeRetail     IndustryType = "retail"
	IndustryTypeOther      IndustryType = "other"
)

func (it IndustryType) String() string {
	return string(it)
}

// IndustryTypeValidator is a validator for the "industry_type" field enum values. It is called by the builders before save.
func IndustryTypeValidator(it IndustryType) error {
	switch it {
	case IndustryTypeHotel, IndustryTypeRestaurant, IndustryTypeBar, IndustryTypeClub, IndustryTypeGym, IndustryTypeEvents, IndustryTypeRetail, IndustryTypeOther:
		return nil
	default:
		return fmt.Errorf("inventorytype: invalid enum value for industry_type field: %q", it)
	}
}

// OrderOption defines the ordering options for the InventoryType queries.
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

// ByImageURL orders the results by the image_url field.
func ByImageURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImageURL, opts...).ToFunc()
}

// ByIconURL orders the results by the icon_url field.
func ByIconURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIconURL, opts...).ToFunc()
}

// ByIndustryType orders the results by the industry_type field.
func ByIndustryType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIndustryType, opts...).ToFunc()
}

// ByMeasurementUnit orders the results by the measurement_unit field.
func ByMeasurementUnit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMeasurementUnit, opts...).ToFunc()
}

// ByAttributesCount orders the results by attributes count.
func ByAttributesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAttributesStep(), opts...)
	}
}

// ByAttributes orders the results by attributes terms.
func ByAttributes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAttributesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPlaceInventoriesCount orders the results by place_inventories count.
func ByPlaceInventoriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPlaceInventoriesStep(), opts...)
	}
}

// ByPlaceInventories orders the results by place_inventories terms.
func ByPlaceInventories(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlaceInventoriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAttributesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AttributesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AttributesTable, AttributesColumn),
	)
}
func newPlaceInventoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlaceInventoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PlaceInventoriesTable, PlaceInventoriesColumn),
	)
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"placio_api/inventoryattribute"
	"placio_api/inventorytype"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// InventoryAttribute is the model entity for the InventoryAttribute schema.
type InventoryAttribute struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// IsMandatory holds the value of the "is_mandatory" field.
	IsMandatory bool `json:"is_mandatory,omitempty"`
	// DataType holds the value of the "data_type" field.
	DataType inventoryattribute.DataType `json:"data_type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InventoryAttributeQuery when eager-loading is set.
	Edges                     InventoryAttributeEdges `json:"edges"`
	inventory_type_attributes *string
	selectValues              sql.SelectValues
}

// InventoryAttributeEdges holds the relations/edges for other nodes in the graph.
type InventoryAttributeEdges struct {
	// InventoryType holds the value of the inventory_type edge.
	InventoryType *InventoryType `json:"inventory_type,omitempty"`
	// PlaceInventoryAttributes holds the value of the place_inventory_attributes edge.
	PlaceInventoryAttributes []*PlaceInventoryAttribute `json:"place_inventory_attributes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// InventoryTypeOrErr returns the InventoryType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InventoryAttributeEdges) InventoryTypeOrErr() (*InventoryType, error) {
	if e.loadedTypes[0] {
		if e.InventoryType == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: inventorytype.Label}
		}
		return e.InventoryType, nil
	}
	return nil, &NotLoadedError{edge: "inventory_type"}
}

// PlaceInventoryAttributesOrErr returns the PlaceInventoryAttributes value or an error if the edge
// was not loaded in eager-loading.
func (e InventoryAttributeEdges) PlaceInventoryAttributesOrErr() ([]*PlaceInventoryAttribute, error) {
	if e.loadedTypes[1] {
		return e.PlaceInventoryAttributes, nil
	}
	return nil, &NotLoadedError{edge: "place_inventory_attributes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*InventoryAttribute) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case inventoryattribute.FieldIsMandatory:
			values[i] = new(sql.NullBool)
		case inventoryattribute.FieldID, inventoryattribute.FieldName, inventoryattribute.FieldDataType:
			values[i] = new(sql.NullString)
		case inventoryattribute.ForeignKeys[0]: // inventory_type_attributes
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the InventoryAttribute fields.
func (ia *InventoryAttribute) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case inventoryattribute.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ia.ID = value.String
			}
		case inventoryattribute.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ia.Name = value.String
			}
		case inventoryattribute.FieldIsMandatory:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_mandatory", values[i])
			} else if value.Valid {
				ia.IsMandatory = value.Bool
			}
		case inventoryattribute.FieldDataType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field data_type", values[i])
			} else if value.Valid {
				ia.DataType = inventoryattribute.DataType(value.String)
			}
		case inventoryattribute.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field inventory_type_attributes", values[i])
			} else if value.Valid {
				ia.inventory_type_attributes = new(string)
				*ia.inventory_type_attributes = value.String
			}
		default:
			ia.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the InventoryAttribute.
// This includes values selected through modifiers, order, etc.
func (ia *InventoryAttribute) Value(name string) (ent.Value, error) {
	return ia.selectValues.Get(name)
}

// QueryInventoryType queries the "inventory_type" edge of the InventoryAttribute entity.
func (ia *InventoryAttribute) QueryInventoryType() *InventoryTypeQuery {
	return NewInventoryAttributeClient(ia.config).QueryInventoryType(ia)
}

// QueryPlaceInventoryAttributes queries the "place_inventory_attributes" edge of the InventoryAttribute entity.
func (ia *InventoryAttribute) QueryPlaceInventoryAttributes() *PlaceInventoryAttributeQuery {
	return NewInventoryAttributeClient(ia.config).QueryPlaceInventoryAttributes(ia)
}

// Update returns a builder for updating this InventoryAttribute.
// Note that you need to call InventoryAttribute.Unwrap() before calling this method if this InventoryAttribute
// was returned from a transaction, and the transaction was committed or rolled back.
func (ia *InventoryAttribute) Update() *InventoryAttributeUpdateOne {
	return NewInventoryAttributeClient(ia.config).UpdateOne(ia)
}

// Unwrap unwraps the InventoryAttribute entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ia *InventoryAttribute) Unwrap() *InventoryAttribute {
	_tx, ok := ia.config.driver.(*txDriver)
	if !ok {
		panic("placio_api: InventoryAttribute is not a transactional entity")
	}
	ia.config.driver = _tx.drv
	return ia
}

// String implements the fmt.Stringer.
func (ia *InventoryAttribute) String() string {
	var builder strings.Builder
	builder.WriteString("InventoryAttribute(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ia.ID))
	builder.WriteString("name=")
	builder.WriteString(ia.Name)
	builder.WriteString(", ")
	builder.WriteString("is_mandatory=")
	builder.WriteString(fmt.Sprintf("%v", ia.IsMandatory))
	builder.WriteString(", ")
	builder.WriteString("data_type=")
	builder.WriteString(fmt.Sprintf("%v", ia.DataType))
	builder.WriteByte(')')
	return builder.String()
}

// InventoryAttributes is a parsable slice of InventoryAttribute.
type InventoryAttributes []*InventoryAttribute

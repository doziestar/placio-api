



// Code generated by ent, DO NOT EDIT.



package ent



	
import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"
		"placio-app/ent/predicate"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
			"database/sql/driver"
			"entgo.io/ent/dialect/sql"
			"entgo.io/ent/dialect/sql/sqlgraph"
			"entgo.io/ent/dialect/sql/sqljson"
			"entgo.io/ent/schema/field"
			 "placio-app/ent/inventorytype"
			 "placio-app/ent/inventoryattribute"
			 "placio-app/ent/placeinventory"

)




		// InventoryType is the model entity for the InventoryType schema.
type InventoryType struct {
	config `json:"-"`
		// ID of the ent.
		ID string `json:"id,omitempty"`
		// Name holds the value of the "name" field.
		Name string `json:"name,omitempty"`
		// Description holds the value of the "description" field.
		Description string `json:"description,omitempty"`
		// ImageURL holds the value of the "image_url" field.
		ImageURL string `json:"image_url,omitempty"`
		// IconURL holds the value of the "icon_url" field.
		IconURL string `json:"icon_url,omitempty"`
		// IndustryType holds the value of the "industry_type" field.
		IndustryType inventorytype.IndustryType `json:"industry_type,omitempty"`
		// MeasurementUnit holds the value of the "measurement_unit" field.
		MeasurementUnit string `json:"measurement_unit,omitempty"`
		// Edges holds the relations/edges for other nodes in the graph.
		// The values are being populated by the InventoryTypeQuery when eager-loading is set.
		Edges InventoryTypeEdges `json:"edges"`
	selectValues sql.SelectValues

}
// InventoryTypeEdges holds the relations/edges for other nodes in the graph.
type InventoryTypeEdges struct {
		// Attributes holds the value of the attributes edge.
		Attributes []*InventoryAttribute `json:"attributes,omitempty"`
		// PlaceInventories holds the value of the place_inventories edge.
		PlaceInventories []*PlaceInventory `json:"place_inventories,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool


}
	// AttributesOrErr returns the Attributes value or an error if the edge
	// was not loaded in eager-loading.
	func (e InventoryTypeEdges) AttributesOrErr() ([]*InventoryAttribute, error) {
		if e.loadedTypes[0] {
			return e.Attributes, nil
		}
		return nil, &NotLoadedError{edge: "attributes"}
	}
	// PlaceInventoriesOrErr returns the PlaceInventories value or an error if the edge
	// was not loaded in eager-loading.
	func (e InventoryTypeEdges) PlaceInventoriesOrErr() ([]*PlaceInventory, error) {
		if e.loadedTypes[1] {
			return e.PlaceInventories, nil
		}
		return nil, &NotLoadedError{edge: "place_inventories"}
	}







	
	


	
	
		
	
	
	

	
	
		
	
	
	

	
	
		
	
	
	

	
	
		
	
	
	

	
	
		
	
	
	

	
	
		
	
	
	


// scanValues returns the types for scanning values from sql.Rows.
func (*InventoryType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
				case inventorytype.FieldID,inventorytype.FieldName,inventorytype.FieldDescription,inventorytype.FieldImageURL,inventorytype.FieldIconURL,inventorytype.FieldIndustryType,inventorytype.FieldMeasurementUnit:
					values[i] = new(sql.NullString)
			default:
				values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the InventoryType fields.
func (it *InventoryType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
			case inventorytype.FieldID:
						if value, ok := values[i].(*sql.NullString); !ok {
			return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
					it.ID = value.String
		}
			case inventorytype.FieldName:
					if value, ok := values[i].(*sql.NullString); !ok {
			return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
					it.Name = value.String
		}
			case inventorytype.FieldDescription:
					if value, ok := values[i].(*sql.NullString); !ok {
			return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
					it.Description = value.String
		}
			case inventorytype.FieldImageURL:
					if value, ok := values[i].(*sql.NullString); !ok {
			return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
					it.ImageURL = value.String
		}
			case inventorytype.FieldIconURL:
					if value, ok := values[i].(*sql.NullString); !ok {
			return fmt.Errorf("unexpected type %T for field icon_url", values[i])
			} else if value.Valid {
					it.IconURL = value.String
		}
			case inventorytype.FieldIndustryType:
					if value, ok := values[i].(*sql.NullString); !ok {
			return fmt.Errorf("unexpected type %T for field industry_type", values[i])
			} else if value.Valid {
					it.IndustryType = inventorytype.IndustryType(value.String)
		}
			case inventorytype.FieldMeasurementUnit:
					if value, ok := values[i].(*sql.NullString); !ok {
			return fmt.Errorf("unexpected type %T for field measurement_unit", values[i])
			} else if value.Valid {
					it.MeasurementUnit = value.String
		}
		default:
			it.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the InventoryType.
// This includes values selected through modifiers, order, etc.
func (it *InventoryType) Value(name string) (ent.Value, error) {
	return it.selectValues.Get(name)
}





	
	// QueryAttributes queries the "attributes" edge of the InventoryType entity.
	func (it *InventoryType) QueryAttributes() *InventoryAttributeQuery {
		return NewInventoryTypeClient(it.config).QueryAttributes(it)
	}

	
	// QueryPlaceInventories queries the "place_inventories" edge of the InventoryType entity.
	func (it *InventoryType) QueryPlaceInventories() *PlaceInventoryQuery {
		return NewInventoryTypeClient(it.config).QueryPlaceInventories(it)
	}


// Update returns a builder for updating this InventoryType.
// Note that you need to call InventoryType.Unwrap() before calling this method if this InventoryType
// was returned from a transaction, and the transaction was committed or rolled back.
func (it *InventoryType) Update() *InventoryTypeUpdateOne {
	return NewInventoryTypeClient(it.config).UpdateOne(it)
}

// Unwrap unwraps the InventoryType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (it *InventoryType) Unwrap() *InventoryType {
	_tx, ok := it.config.driver.(*txDriver)
	if !ok {
		panic("ent: InventoryType is not a transactional entity")
	}
	it.config.driver = _tx.drv
	return it
}


	

	// String implements the fmt.Stringer.
	func (it *InventoryType) String() string {
		var builder strings.Builder
		builder.WriteString("InventoryType(")
			builder.WriteString(fmt.Sprintf("id=%v, ", it.ID))
					builder.WriteString("name=")
						builder.WriteString(it.Name)
				builder.WriteString(", ")
					builder.WriteString("description=")
						builder.WriteString(it.Description)
				builder.WriteString(", ")
					builder.WriteString("image_url=")
						builder.WriteString(it.ImageURL)
				builder.WriteString(", ")
					builder.WriteString("icon_url=")
						builder.WriteString(it.IconURL)
				builder.WriteString(", ")
					builder.WriteString("industry_type=")
						builder.WriteString(fmt.Sprintf("%v", it.IndustryType))
				builder.WriteString(", ")
					builder.WriteString("measurement_unit=")
						builder.WriteString(it.MeasurementUnit)
		builder.WriteByte(')')
		return builder.String()
	}







// InventoryTypes is a parsable slice of InventoryType.
type InventoryTypes []*InventoryType


	
	



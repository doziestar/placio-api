// Code generated by ent, DO NOT EDIT.

package placio_api

import (
	"fmt"
	"placio_api/amenity"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Amenity is the model entity for the Amenity schema.
type Amenity struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Icon holds the value of the "icon" field.
	Icon string `json:"icon,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AmenityQuery when eager-loading is set.
	Edges        AmenityEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AmenityEdges holds the relations/edges for other nodes in the graph.
type AmenityEdges struct {
	// Places holds the value of the places edge.
	Places []*Place `json:"places,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PlacesOrErr returns the Places value or an error if the edge
// was not loaded in eager-loading.
func (e AmenityEdges) PlacesOrErr() ([]*Place, error) {
	if e.loadedTypes[0] {
		return e.Places, nil
	}
	return nil, &NotLoadedError{edge: "places"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Amenity) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case amenity.FieldID, amenity.FieldName, amenity.FieldIcon:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Amenity fields.
func (a *Amenity) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case amenity.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				a.ID = value.String
			}
		case amenity.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case amenity.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				a.Icon = value.String
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Amenity.
// This includes values selected through modifiers, order, etc.
func (a *Amenity) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryPlaces queries the "places" edge of the Amenity entity.
func (a *Amenity) QueryPlaces() *PlaceQuery {
	return NewAmenityClient(a.config).QueryPlaces(a)
}

// Update returns a builder for updating this Amenity.
// Note that you need to call Amenity.Unwrap() before calling this method if this Amenity
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Amenity) Update() *AmenityUpdateOne {
	return NewAmenityClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Amenity entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Amenity) Unwrap() *Amenity {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("placio_api: Amenity is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Amenity) String() string {
	var builder strings.Builder
	builder.WriteString("Amenity(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("icon=")
	builder.WriteString(a.Icon)
	builder.WriteByte(')')
	return builder.String()
}

// Amenities is a parsable slice of Amenity.
type Amenities []*Amenity

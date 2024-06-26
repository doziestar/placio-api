// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"placio-app/ent/place"
	"placio-app/ent/placetable"
	"placio-app/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PlaceTable is the model entity for the PlaceTable schema.
type PlaceTable struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Number holds the value of the "number" field.
	Number int `json:"number,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Capacity holds the value of the "capacity" field.
	Capacity int `json:"capacity,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt string `json:"deleted_at,omitempty"`
	// IsDeleted holds the value of the "is_deleted" field.
	IsDeleted bool `json:"is_deleted,omitempty"`
	// QrCode holds the value of the "qr_code" field.
	QrCode string `json:"qr_code,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Type holds the value of the "type" field.
	Type placetable.Type `json:"type,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// IsReserved holds the value of the "is_reserved" field.
	IsReserved bool `json:"is_reserved,omitempty"`
	// IsVip holds the value of the "is_vip" field.
	IsVip bool `json:"is_vip,omitempty"`
	// IsPremium holds the value of the "is_premium" field.
	IsPremium bool `json:"is_premium,omitempty"`
	// LocationDescription holds the value of the "location_description" field.
	LocationDescription string `json:"location_description,omitempty"`
	// MinimumSpend holds the value of the "minimum_spend" field.
	MinimumSpend float64 `json:"minimum_spend,omitempty"`
	// ReservationTime holds the value of the "reservation_time" field.
	ReservationTime *time.Time `json:"reservation_time,omitempty"`
	// NextAvailableTime holds the value of the "next_available_time" field.
	NextAvailableTime *time.Time `json:"next_available_time,omitempty"`
	// SpecialRequirements holds the value of the "special_requirements" field.
	SpecialRequirements []string `json:"special_requirements,omitempty"`
	// Layout holds the value of the "layout" field.
	Layout string `json:"layout,omitempty"`
	// ServiceArea holds the value of the "service_area" field.
	ServiceArea string `json:"service_area,omitempty"`
	// Ambient holds the value of the "ambient" field.
	Ambient string `json:"ambient,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty"`
	// Rating holds the value of the "rating" field.
	Rating *float64 `json:"rating,omitempty"`
	// Tags holds the value of the "tags" field.
	Tags []string `json:"tags,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlaceTableQuery when eager-loading is set.
	Edges                PlaceTableEdges `json:"edges"`
	place_tables         *string
	user_tables_created  *string
	user_tables_updated  *string
	user_tables_deleted  *string
	user_tables_reserved *string
	user_tables_waited   *string
	selectValues         sql.SelectValues
}

// PlaceTableEdges holds the relations/edges for other nodes in the graph.
type PlaceTableEdges struct {
	// Place holds the value of the place edge.
	Place *Place `json:"place,omitempty"`
	// CreatedBy holds the value of the created_by edge.
	CreatedBy *User `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the updated_by edge.
	UpdatedBy *User `json:"updated_by,omitempty"`
	// DeletedBy holds the value of the deleted_by edge.
	DeletedBy *User `json:"deleted_by,omitempty"`
	// ReservedBy holds the value of the reserved_by edge.
	ReservedBy *User `json:"reserved_by,omitempty"`
	// Waiter holds the value of the waiter edge.
	Waiter *User `json:"waiter,omitempty"`
	// Orders holds the value of the orders edge.
	Orders []*Order `json:"orders,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// PlaceOrErr returns the Place value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlaceTableEdges) PlaceOrErr() (*Place, error) {
	if e.loadedTypes[0] {
		if e.Place == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: place.Label}
		}
		return e.Place, nil
	}
	return nil, &NotLoadedError{edge: "place"}
}

// CreatedByOrErr returns the CreatedBy value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlaceTableEdges) CreatedByOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.CreatedBy == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.CreatedBy, nil
	}
	return nil, &NotLoadedError{edge: "created_by"}
}

// UpdatedByOrErr returns the UpdatedBy value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlaceTableEdges) UpdatedByOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.UpdatedBy == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.UpdatedBy, nil
	}
	return nil, &NotLoadedError{edge: "updated_by"}
}

// DeletedByOrErr returns the DeletedBy value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlaceTableEdges) DeletedByOrErr() (*User, error) {
	if e.loadedTypes[3] {
		if e.DeletedBy == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.DeletedBy, nil
	}
	return nil, &NotLoadedError{edge: "deleted_by"}
}

// ReservedByOrErr returns the ReservedBy value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlaceTableEdges) ReservedByOrErr() (*User, error) {
	if e.loadedTypes[4] {
		if e.ReservedBy == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.ReservedBy, nil
	}
	return nil, &NotLoadedError{edge: "reserved_by"}
}

// WaiterOrErr returns the Waiter value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlaceTableEdges) WaiterOrErr() (*User, error) {
	if e.loadedTypes[5] {
		if e.Waiter == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Waiter, nil
	}
	return nil, &NotLoadedError{edge: "waiter"}
}

// OrdersOrErr returns the Orders value or an error if the edge
// was not loaded in eager-loading.
func (e PlaceTableEdges) OrdersOrErr() ([]*Order, error) {
	if e.loadedTypes[6] {
		return e.Orders, nil
	}
	return nil, &NotLoadedError{edge: "orders"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PlaceTable) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case placetable.FieldSpecialRequirements, placetable.FieldTags, placetable.FieldMetadata:
			values[i] = new([]byte)
		case placetable.FieldIsDeleted, placetable.FieldIsActive, placetable.FieldIsReserved, placetable.FieldIsVip, placetable.FieldIsPremium:
			values[i] = new(sql.NullBool)
		case placetable.FieldMinimumSpend, placetable.FieldRating:
			values[i] = new(sql.NullFloat64)
		case placetable.FieldNumber, placetable.FieldCapacity:
			values[i] = new(sql.NullInt64)
		case placetable.FieldID, placetable.FieldName, placetable.FieldDeletedAt, placetable.FieldQrCode, placetable.FieldDescription, placetable.FieldStatus, placetable.FieldType, placetable.FieldLocationDescription, placetable.FieldLayout, placetable.FieldServiceArea, placetable.FieldAmbient, placetable.FieldImageURL:
			values[i] = new(sql.NullString)
		case placetable.FieldReservationTime, placetable.FieldNextAvailableTime:
			values[i] = new(sql.NullTime)
		case placetable.ForeignKeys[0]: // place_tables
			values[i] = new(sql.NullString)
		case placetable.ForeignKeys[1]: // user_tables_created
			values[i] = new(sql.NullString)
		case placetable.ForeignKeys[2]: // user_tables_updated
			values[i] = new(sql.NullString)
		case placetable.ForeignKeys[3]: // user_tables_deleted
			values[i] = new(sql.NullString)
		case placetable.ForeignKeys[4]: // user_tables_reserved
			values[i] = new(sql.NullString)
		case placetable.ForeignKeys[5]: // user_tables_waited
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PlaceTable fields.
func (pt *PlaceTable) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case placetable.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pt.ID = value.String
			}
		case placetable.FieldNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				pt.Number = int(value.Int64)
			}
		case placetable.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pt.Name = value.String
			}
		case placetable.FieldCapacity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field capacity", values[i])
			} else if value.Valid {
				pt.Capacity = int(value.Int64)
			}
		case placetable.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pt.DeletedAt = value.String
			}
		case placetable.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				pt.IsDeleted = value.Bool
			}
		case placetable.FieldQrCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field qr_code", values[i])
			} else if value.Valid {
				pt.QrCode = value.String
			}
		case placetable.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pt.Description = value.String
			}
		case placetable.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pt.Status = value.String
			}
		case placetable.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				pt.Type = placetable.Type(value.String)
			}
		case placetable.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				pt.IsActive = value.Bool
			}
		case placetable.FieldIsReserved:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_reserved", values[i])
			} else if value.Valid {
				pt.IsReserved = value.Bool
			}
		case placetable.FieldIsVip:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_vip", values[i])
			} else if value.Valid {
				pt.IsVip = value.Bool
			}
		case placetable.FieldIsPremium:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_premium", values[i])
			} else if value.Valid {
				pt.IsPremium = value.Bool
			}
		case placetable.FieldLocationDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location_description", values[i])
			} else if value.Valid {
				pt.LocationDescription = value.String
			}
		case placetable.FieldMinimumSpend:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field minimum_spend", values[i])
			} else if value.Valid {
				pt.MinimumSpend = value.Float64
			}
		case placetable.FieldReservationTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field reservation_time", values[i])
			} else if value.Valid {
				pt.ReservationTime = new(time.Time)
				*pt.ReservationTime = value.Time
			}
		case placetable.FieldNextAvailableTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field next_available_time", values[i])
			} else if value.Valid {
				pt.NextAvailableTime = new(time.Time)
				*pt.NextAvailableTime = value.Time
			}
		case placetable.FieldSpecialRequirements:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field special_requirements", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pt.SpecialRequirements); err != nil {
					return fmt.Errorf("unmarshal field special_requirements: %w", err)
				}
			}
		case placetable.FieldLayout:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field layout", values[i])
			} else if value.Valid {
				pt.Layout = value.String
			}
		case placetable.FieldServiceArea:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field service_area", values[i])
			} else if value.Valid {
				pt.ServiceArea = value.String
			}
		case placetable.FieldAmbient:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ambient", values[i])
			} else if value.Valid {
				pt.Ambient = value.String
			}
		case placetable.FieldImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
				pt.ImageURL = value.String
			}
		case placetable.FieldRating:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field rating", values[i])
			} else if value.Valid {
				pt.Rating = new(float64)
				*pt.Rating = value.Float64
			}
		case placetable.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pt.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case placetable.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pt.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		case placetable.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field place_tables", values[i])
			} else if value.Valid {
				pt.place_tables = new(string)
				*pt.place_tables = value.String
			}
		case placetable.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_tables_created", values[i])
			} else if value.Valid {
				pt.user_tables_created = new(string)
				*pt.user_tables_created = value.String
			}
		case placetable.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_tables_updated", values[i])
			} else if value.Valid {
				pt.user_tables_updated = new(string)
				*pt.user_tables_updated = value.String
			}
		case placetable.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_tables_deleted", values[i])
			} else if value.Valid {
				pt.user_tables_deleted = new(string)
				*pt.user_tables_deleted = value.String
			}
		case placetable.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_tables_reserved", values[i])
			} else if value.Valid {
				pt.user_tables_reserved = new(string)
				*pt.user_tables_reserved = value.String
			}
		case placetable.ForeignKeys[5]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_tables_waited", values[i])
			} else if value.Valid {
				pt.user_tables_waited = new(string)
				*pt.user_tables_waited = value.String
			}
		default:
			pt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PlaceTable.
// This includes values selected through modifiers, order, etc.
func (pt *PlaceTable) Value(name string) (ent.Value, error) {
	return pt.selectValues.Get(name)
}

// QueryPlace queries the "place" edge of the PlaceTable entity.
func (pt *PlaceTable) QueryPlace() *PlaceQuery {
	return NewPlaceTableClient(pt.config).QueryPlace(pt)
}

// QueryCreatedBy queries the "created_by" edge of the PlaceTable entity.
func (pt *PlaceTable) QueryCreatedBy() *UserQuery {
	return NewPlaceTableClient(pt.config).QueryCreatedBy(pt)
}

// QueryUpdatedBy queries the "updated_by" edge of the PlaceTable entity.
func (pt *PlaceTable) QueryUpdatedBy() *UserQuery {
	return NewPlaceTableClient(pt.config).QueryUpdatedBy(pt)
}

// QueryDeletedBy queries the "deleted_by" edge of the PlaceTable entity.
func (pt *PlaceTable) QueryDeletedBy() *UserQuery {
	return NewPlaceTableClient(pt.config).QueryDeletedBy(pt)
}

// QueryReservedBy queries the "reserved_by" edge of the PlaceTable entity.
func (pt *PlaceTable) QueryReservedBy() *UserQuery {
	return NewPlaceTableClient(pt.config).QueryReservedBy(pt)
}

// QueryWaiter queries the "waiter" edge of the PlaceTable entity.
func (pt *PlaceTable) QueryWaiter() *UserQuery {
	return NewPlaceTableClient(pt.config).QueryWaiter(pt)
}

// QueryOrders queries the "orders" edge of the PlaceTable entity.
func (pt *PlaceTable) QueryOrders() *OrderQuery {
	return NewPlaceTableClient(pt.config).QueryOrders(pt)
}

// Update returns a builder for updating this PlaceTable.
// Note that you need to call PlaceTable.Unwrap() before calling this method if this PlaceTable
// was returned from a transaction, and the transaction was committed or rolled back.
func (pt *PlaceTable) Update() *PlaceTableUpdateOne {
	return NewPlaceTableClient(pt.config).UpdateOne(pt)
}

// Unwrap unwraps the PlaceTable entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pt *PlaceTable) Unwrap() *PlaceTable {
	_tx, ok := pt.config.driver.(*txDriver)
	if !ok {
		panic("ent: PlaceTable is not a transactional entity")
	}
	pt.config.driver = _tx.drv
	return pt
}

// String implements the fmt.Stringer.
func (pt *PlaceTable) String() string {
	var builder strings.Builder
	builder.WriteString("PlaceTable(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pt.ID))
	builder.WriteString("number=")
	builder.WriteString(fmt.Sprintf("%v", pt.Number))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pt.Name)
	builder.WriteString(", ")
	builder.WriteString("capacity=")
	builder.WriteString(fmt.Sprintf("%v", pt.Capacity))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(pt.DeletedAt)
	builder.WriteString(", ")
	builder.WriteString("is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", pt.IsDeleted))
	builder.WriteString(", ")
	builder.WriteString("qr_code=")
	builder.WriteString(pt.QrCode)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pt.Description)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(pt.Status)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", pt.Type))
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", pt.IsActive))
	builder.WriteString(", ")
	builder.WriteString("is_reserved=")
	builder.WriteString(fmt.Sprintf("%v", pt.IsReserved))
	builder.WriteString(", ")
	builder.WriteString("is_vip=")
	builder.WriteString(fmt.Sprintf("%v", pt.IsVip))
	builder.WriteString(", ")
	builder.WriteString("is_premium=")
	builder.WriteString(fmt.Sprintf("%v", pt.IsPremium))
	builder.WriteString(", ")
	builder.WriteString("location_description=")
	builder.WriteString(pt.LocationDescription)
	builder.WriteString(", ")
	builder.WriteString("minimum_spend=")
	builder.WriteString(fmt.Sprintf("%v", pt.MinimumSpend))
	builder.WriteString(", ")
	if v := pt.ReservationTime; v != nil {
		builder.WriteString("reservation_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := pt.NextAvailableTime; v != nil {
		builder.WriteString("next_available_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("special_requirements=")
	builder.WriteString(fmt.Sprintf("%v", pt.SpecialRequirements))
	builder.WriteString(", ")
	builder.WriteString("layout=")
	builder.WriteString(pt.Layout)
	builder.WriteString(", ")
	builder.WriteString("service_area=")
	builder.WriteString(pt.ServiceArea)
	builder.WriteString(", ")
	builder.WriteString("ambient=")
	builder.WriteString(pt.Ambient)
	builder.WriteString(", ")
	builder.WriteString("image_url=")
	builder.WriteString(pt.ImageURL)
	builder.WriteString(", ")
	if v := pt.Rating; v != nil {
		builder.WriteString("rating=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", pt.Tags))
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", pt.Metadata))
	builder.WriteByte(')')
	return builder.String()
}

// PlaceTables is a parsable slice of PlaceTable.
type PlaceTables []*PlaceTable

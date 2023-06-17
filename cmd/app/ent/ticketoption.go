// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"placio-app/ent/event"
	"placio-app/ent/ticketoption"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// TicketOption is the model entity for the TicketOption schema.
type TicketOption struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TicketOptionQuery when eager-loading is set.
	Edges                 TicketOptionEdges `json:"edges"`
	event_ticket_options  *string
	ticket_ticket_options *string
	selectValues          sql.SelectValues
}

// TicketOptionEdges holds the relations/edges for other nodes in the graph.
type TicketOptionEdges struct {
	// Event holds the value of the event edge.
	Event *Event `json:"event,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EventOrErr returns the Event value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TicketOptionEdges) EventOrErr() (*Event, error) {
	if e.loadedTypes[0] {
		if e.Event == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: event.Label}
		}
		return e.Event, nil
	}
	return nil, &NotLoadedError{edge: "event"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TicketOption) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ticketoption.FieldID:
			values[i] = new(sql.NullString)
		case ticketoption.FieldCreatedAt, ticketoption.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case ticketoption.ForeignKeys[0]: // event_ticket_options
			values[i] = new(sql.NullString)
		case ticketoption.ForeignKeys[1]: // ticket_ticket_options
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TicketOption fields.
func (to *TicketOption) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ticketoption.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				to.ID = value.String
			}
		case ticketoption.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				to.CreatedAt = value.Time
			}
		case ticketoption.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				to.UpdatedAt = value.Time
			}
		case ticketoption.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_ticket_options", values[i])
			} else if value.Valid {
				to.event_ticket_options = new(string)
				*to.event_ticket_options = value.String
			}
		case ticketoption.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ticket_ticket_options", values[i])
			} else if value.Valid {
				to.ticket_ticket_options = new(string)
				*to.ticket_ticket_options = value.String
			}
		default:
			to.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TicketOption.
// This includes values selected through modifiers, order, etc.
func (to *TicketOption) Value(name string) (ent.Value, error) {
	return to.selectValues.Get(name)
}

// QueryEvent queries the "event" edge of the TicketOption entity.
func (to *TicketOption) QueryEvent() *EventQuery {
	return NewTicketOptionClient(to.config).QueryEvent(to)
}

// Update returns a builder for updating this TicketOption.
// Note that you need to call TicketOption.Unwrap() before calling this method if this TicketOption
// was returned from a transaction, and the transaction was committed or rolled back.
func (to *TicketOption) Update() *TicketOptionUpdateOne {
	return NewTicketOptionClient(to.config).UpdateOne(to)
}

// Unwrap unwraps the TicketOption entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (to *TicketOption) Unwrap() *TicketOption {
	_tx, ok := to.config.driver.(*txDriver)
	if !ok {
		panic("ent: TicketOption is not a transactional entity")
	}
	to.config.driver = _tx.drv
	return to
}

// String implements the fmt.Stringer.
func (to *TicketOption) String() string {
	var builder strings.Builder
	builder.WriteString("TicketOption(")
	builder.WriteString(fmt.Sprintf("id=%v, ", to.ID))
	builder.WriteString("createdAt=")
	builder.WriteString(to.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(to.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TicketOptions is a parsable slice of TicketOption.
type TicketOptions []*TicketOption

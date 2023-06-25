// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"placio-app/ent/event"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Event is the model entity for the Event schema.
type Event struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EventQuery when eager-loading is set.
	Edges        EventEdges `json:"edges"`
	place_events *string
	user_events  *string
	selectValues sql.SelectValues
}

// EventEdges holds the relations/edges for other nodes in the graph.
type EventEdges struct {
	// Tickets holds the value of the tickets edge.
	Tickets []*Ticket `json:"tickets,omitempty"`
	// TicketOptions holds the value of the ticket_options edge.
	TicketOptions []*TicketOption `json:"ticket_options,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TicketsOrErr returns the Tickets value or an error if the edge
// was not loaded in eager-loading.
func (e EventEdges) TicketsOrErr() ([]*Ticket, error) {
	if e.loadedTypes[0] {
		return e.Tickets, nil
	}
	return nil, &NotLoadedError{edge: "tickets"}
}

// TicketOptionsOrErr returns the TicketOptions value or an error if the edge
// was not loaded in eager-loading.
func (e EventEdges) TicketOptionsOrErr() ([]*TicketOption, error) {
	if e.loadedTypes[1] {
		return e.TicketOptions, nil
	}
	return nil, &NotLoadedError{edge: "ticket_options"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Event) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case event.FieldID, event.FieldName:
			values[i] = new(sql.NullString)
		case event.FieldCreatedAt, event.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case event.ForeignKeys[0]: // place_events
			values[i] = new(sql.NullString)
		case event.ForeignKeys[1]: // user_events
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Event fields.
func (e *Event) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case event.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				e.ID = value.String
			}
		case event.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		case event.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		case event.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				e.UpdatedAt = value.Time
			}
		case event.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field place_events", values[i])
			} else if value.Valid {
				e.place_events = new(string)
				*e.place_events = value.String
			}
		case event.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_events", values[i])
			} else if value.Valid {
				e.user_events = new(string)
				*e.user_events = value.String
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Event.
// This includes values selected through modifiers, order, etc.
func (e *Event) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// QueryTickets queries the "tickets" edge of the Event entity.
func (e *Event) QueryTickets() *TicketQuery {
	return NewEventClient(e.config).QueryTickets(e)
}

// QueryTicketOptions queries the "ticket_options" edge of the Event entity.
func (e *Event) QueryTicketOptions() *TicketOptionQuery {
	return NewEventClient(e.config).QueryTicketOptions(e)
}

// Update returns a builder for updating this Event.
// Note that you need to call Event.Unwrap() before calling this method if this Event
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Event) Update() *EventUpdateOne {
	return NewEventClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Event entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Event) Unwrap() *Event {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Event is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Event) String() string {
	var builder strings.Builder
	builder.WriteString("Event(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("name=")
	builder.WriteString(e.Name)
	builder.WriteString(", ")
	builder.WriteString("createdAt=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(e.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Events is a parsable slice of Event.
type Events []*Event

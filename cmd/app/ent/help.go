// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"placio-app/ent/help"
	"placio-app/ent/user"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Help is the model entity for the Help schema.
type Help struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Category holds the value of the "category" field.
	Category string `json:"category,omitempty"`
	// Subject holds the value of the "subject" field.
	Subject string `json:"subject,omitempty"`
	// Body holds the value of the "body" field.
	Body string `json:"body,omitempty"`
	// Media holds the value of the "media" field.
	Media string `json:"media,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HelpQuery when eager-loading is set.
	Edges        HelpEdges `json:"edges"`
	selectValues sql.SelectValues
}

// HelpEdges holds the relations/edges for other nodes in the graph.
type HelpEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HelpEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Help) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case help.FieldID, help.FieldCategory, help.FieldSubject, help.FieldBody, help.FieldMedia, help.FieldStatus, help.FieldUserID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Help fields.
func (h *Help) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case help.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				h.ID = value.String
			}
		case help.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				h.Category = value.String
			}
		case help.FieldSubject:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subject", values[i])
			} else if value.Valid {
				h.Subject = value.String
			}
		case help.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				h.Body = value.String
			}
		case help.FieldMedia:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field media", values[i])
			} else if value.Valid {
				h.Media = value.String
			}
		case help.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				h.Status = value.String
			}
		case help.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				h.UserID = value.String
			}
		default:
			h.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Help.
// This includes values selected through modifiers, order, etc.
func (h *Help) Value(name string) (ent.Value, error) {
	return h.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Help entity.
func (h *Help) QueryUser() *UserQuery {
	return NewHelpClient(h.config).QueryUser(h)
}

// Update returns a builder for updating this Help.
// Note that you need to call Help.Unwrap() before calling this method if this Help
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Help) Update() *HelpUpdateOne {
	return NewHelpClient(h.config).UpdateOne(h)
}

// Unwrap unwraps the Help entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *Help) Unwrap() *Help {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Help is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Help) String() string {
	var builder strings.Builder
	builder.WriteString("Help(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("category=")
	builder.WriteString(h.Category)
	builder.WriteString(", ")
	builder.WriteString("subject=")
	builder.WriteString(h.Subject)
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(h.Body)
	builder.WriteString(", ")
	builder.WriteString("media=")
	builder.WriteString(h.Media)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(h.Status)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(h.UserID)
	builder.WriteByte(')')
	return builder.String()
}

// Helps is a parsable slice of Help.
type Helps []*Help

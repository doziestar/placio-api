// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"placio_api/notification"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Notification is the model entity for the Notification schema.
type Notification struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// Link holds the value of the "link" field.
	Link string `json:"link,omitempty"`
	// IsRead holds the value of the "is_read" field.
	IsRead bool `json:"is_read,omitempty"`
	// Type holds the value of the "type" field.
	Type int `json:"type,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// NotifiableType holds the value of the "notifiable_type" field.
	NotifiableType string `json:"notifiable_type,omitempty"`
	// NotifiableID holds the value of the "notifiable_id" field.
	NotifiableID string `json:"notifiable_id,omitempty"`
	// TriggeredBy holds the value of the "triggered_by" field.
	TriggeredBy string `json:"triggered_by,omitempty"`
	// TriggeredTo holds the value of the "triggered_to" field.
	TriggeredTo string `json:"triggered_to,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NotificationQuery when eager-loading is set.
	Edges        NotificationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// NotificationEdges holds the relations/edges for other nodes in the graph.
type NotificationEdges struct {
	// User holds the value of the user edge.
	User []*User `json:"user,omitempty"`
	// BusinessAccount holds the value of the business_account edge.
	BusinessAccount []*Business `json:"business_account,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e NotificationEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// BusinessAccountOrErr returns the BusinessAccount value or an error if the edge
// was not loaded in eager-loading.
func (e NotificationEdges) BusinessAccountOrErr() ([]*Business, error) {
	if e.loadedTypes[1] {
		return e.BusinessAccount, nil
	}
	return nil, &NotLoadedError{edge: "business_account"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Notification) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case notification.FieldIsRead:
			values[i] = new(sql.NullBool)
		case notification.FieldType:
			values[i] = new(sql.NullInt64)
		case notification.FieldID, notification.FieldTitle, notification.FieldMessage, notification.FieldLink, notification.FieldNotifiableType, notification.FieldNotifiableID, notification.FieldTriggeredBy, notification.FieldTriggeredTo:
			values[i] = new(sql.NullString)
		case notification.FieldCreatedAt, notification.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Notification fields.
func (n *Notification) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case notification.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				n.ID = value.String
			}
		case notification.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				n.Title = value.String
			}
		case notification.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				n.Message = value.String
			}
		case notification.FieldLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field link", values[i])
			} else if value.Valid {
				n.Link = value.String
			}
		case notification.FieldIsRead:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_read", values[i])
			} else if value.Valid {
				n.IsRead = value.Bool
			}
		case notification.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				n.Type = int(value.Int64)
			}
		case notification.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				n.CreatedAt = value.Time
			}
		case notification.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				n.UpdatedAt = value.Time
			}
		case notification.FieldNotifiableType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notifiable_type", values[i])
			} else if value.Valid {
				n.NotifiableType = value.String
			}
		case notification.FieldNotifiableID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notifiable_id", values[i])
			} else if value.Valid {
				n.NotifiableID = value.String
			}
		case notification.FieldTriggeredBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field triggered_by", values[i])
			} else if value.Valid {
				n.TriggeredBy = value.String
			}
		case notification.FieldTriggeredTo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field triggered_to", values[i])
			} else if value.Valid {
				n.TriggeredTo = value.String
			}
		default:
			n.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Notification.
// This includes values selected through modifiers, order, etc.
func (n *Notification) Value(name string) (ent.Value, error) {
	return n.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Notification entity.
func (n *Notification) QueryUser() *UserQuery {
	return NewNotificationClient(n.config).QueryUser(n)
}

// QueryBusinessAccount queries the "business_account" edge of the Notification entity.
func (n *Notification) QueryBusinessAccount() *BusinessQuery {
	return NewNotificationClient(n.config).QueryBusinessAccount(n)
}

// Update returns a builder for updating this Notification.
// Note that you need to call Notification.Unwrap() before calling this method if this Notification
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Notification) Update() *NotificationUpdateOne {
	return NewNotificationClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Notification entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Notification) Unwrap() *Notification {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("placio_api: Notification is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Notification) String() string {
	var builder strings.Builder
	builder.WriteString("Notification(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("title=")
	builder.WriteString(n.Title)
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(n.Message)
	builder.WriteString(", ")
	builder.WriteString("link=")
	builder.WriteString(n.Link)
	builder.WriteString(", ")
	builder.WriteString("is_read=")
	builder.WriteString(fmt.Sprintf("%v", n.IsRead))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", n.Type))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(n.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(n.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("notifiable_type=")
	builder.WriteString(n.NotifiableType)
	builder.WriteString(", ")
	builder.WriteString("notifiable_id=")
	builder.WriteString(n.NotifiableID)
	builder.WriteString(", ")
	builder.WriteString("triggered_by=")
	builder.WriteString(n.TriggeredBy)
	builder.WriteString(", ")
	builder.WriteString("triggered_to=")
	builder.WriteString(n.TriggeredTo)
	builder.WriteByte(')')
	return builder.String()
}

// Notifications is a parsable slice of Notification.
type Notifications []*Notification

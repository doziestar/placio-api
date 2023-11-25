// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/businessfollowuser"
	"placio-app/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// BusinessFollowUser is the model entity for the BusinessFollowUser schema.
type BusinessFollowUser struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "CreatedAt" field.
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
	// UpdatedAt holds the value of the "UpdatedAt" field.
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BusinessFollowUserQuery when eager-loading is set.
	Edges                    BusinessFollowUserEdges `json:"edges"`
	business_followed_users  *string
	user_follower_businesses *string
	selectValues             sql.SelectValues
}

// BusinessFollowUserEdges holds the relations/edges for other nodes in the graph.
type BusinessFollowUserEdges struct {
	// Business holds the value of the business edge.
	Business *Business `json:"business,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BusinessOrErr returns the Business value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BusinessFollowUserEdges) BusinessOrErr() (*Business, error) {
	if e.loadedTypes[0] {
		if e.Business == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: business.Label}
		}
		return e.Business, nil
	}
	return nil, &NotLoadedError{edge: "business"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BusinessFollowUserEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BusinessFollowUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case businessfollowuser.FieldID:
			values[i] = new(sql.NullString)
		case businessfollowuser.FieldCreatedAt, businessfollowuser.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case businessfollowuser.ForeignKeys[0]: // business_followed_users
			values[i] = new(sql.NullString)
		case businessfollowuser.ForeignKeys[1]: // user_follower_businesses
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BusinessFollowUser fields.
func (bfu *BusinessFollowUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case businessfollowuser.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				bfu.ID = value.String
			}
		case businessfollowuser.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedAt", values[i])
			} else if value.Valid {
				bfu.CreatedAt = value.Time
			}
		case businessfollowuser.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
				bfu.UpdatedAt = value.Time
			}
		case businessfollowuser.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field business_followed_users", values[i])
			} else if value.Valid {
				bfu.business_followed_users = new(string)
				*bfu.business_followed_users = value.String
			}
		case businessfollowuser.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_follower_businesses", values[i])
			} else if value.Valid {
				bfu.user_follower_businesses = new(string)
				*bfu.user_follower_businesses = value.String
			}
		default:
			bfu.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BusinessFollowUser.
// This includes values selected through modifiers, order, etc.
func (bfu *BusinessFollowUser) Value(name string) (ent.Value, error) {
	return bfu.selectValues.Get(name)
}

// QueryBusiness queries the "business" edge of the BusinessFollowUser entity.
func (bfu *BusinessFollowUser) QueryBusiness() *BusinessQuery {
	return NewBusinessFollowUserClient(bfu.config).QueryBusiness(bfu)
}

// QueryUser queries the "user" edge of the BusinessFollowUser entity.
func (bfu *BusinessFollowUser) QueryUser() *UserQuery {
	return NewBusinessFollowUserClient(bfu.config).QueryUser(bfu)
}

// Update returns a builder for updating this BusinessFollowUser.
// Note that you need to call BusinessFollowUser.Unwrap() before calling this method if this BusinessFollowUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (bfu *BusinessFollowUser) Update() *BusinessFollowUserUpdateOne {
	return NewBusinessFollowUserClient(bfu.config).UpdateOne(bfu)
}

// Unwrap unwraps the BusinessFollowUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bfu *BusinessFollowUser) Unwrap() *BusinessFollowUser {
	_tx, ok := bfu.config.driver.(*txDriver)
	if !ok {
		panic("ent: BusinessFollowUser is not a transactional entity")
	}
	bfu.config.driver = _tx.drv
	return bfu
}

// String implements the fmt.Stringer.
func (bfu *BusinessFollowUser) String() string {
	var builder strings.Builder
	builder.WriteString("BusinessFollowUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bfu.ID))
	builder.WriteString("CreatedAt=")
	builder.WriteString(bfu.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("UpdatedAt=")
	builder.WriteString(bfu.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// BusinessFollowUsers is a parsable slice of BusinessFollowUser.
type BusinessFollowUsers []*BusinessFollowUser

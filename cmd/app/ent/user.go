// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"placio-app/ent/user"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Auth0ID holds the value of the "auth0_id" field.
	Auth0ID string `json:"auth0_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// UserBusinesses holds the value of the userBusinesses edge.
	UserBusinesses []*UserBusiness `json:"userBusinesses,omitempty"`
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// Likes holds the value of the likes edge.
	Likes []*Like `json:"likes,omitempty"`
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// UserBusinessesOrErr returns the UserBusinesses value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserBusinessesOrErr() ([]*UserBusiness, error) {
	if e.loadedTypes[0] {
		return e.UserBusinesses, nil
	}
	return nil, &NotLoadedError{edge: "userBusinesses"}
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[1] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// LikesOrErr returns the Likes value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) LikesOrErr() ([]*Like, error) {
	if e.loadedTypes[2] {
		return e.Likes, nil
	}
	return nil, &NotLoadedError{edge: "likes"}
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[3] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldAuth0ID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldAuth0ID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field auth0_id", values[i])
			} else if value.Valid {
				u.Auth0ID = value.String
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryUserBusinesses queries the "userBusinesses" edge of the User entity.
func (u *User) QueryUserBusinesses() *UserBusinessQuery {
	return NewUserClient(u.config).QueryUserBusinesses(u)
}

// QueryComments queries the "comments" edge of the User entity.
func (u *User) QueryComments() *CommentQuery {
	return NewUserClient(u.config).QueryComments(u)
}

// QueryLikes queries the "likes" edge of the User entity.
func (u *User) QueryLikes() *LikeQuery {
	return NewUserClient(u.config).QueryLikes(u)
}

// QueryPosts queries the "posts" edge of the User entity.
func (u *User) QueryPosts() *PostQuery {
	return NewUserClient(u.config).QueryPosts(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("auth0_id=")
	builder.WriteString(u.Auth0ID)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

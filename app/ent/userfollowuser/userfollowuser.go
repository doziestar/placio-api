
// Code generated by ent, DO NOT EDIT.



	

package userfollowuser




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

)


const (
	// Label holds the string label denoting the userfollowuser type in the database.
	Label = "user_follow_user"
		// FieldID holds the string denoting the id field in the database.
		FieldID = "id"
		// FieldCreatedAt holds the string denoting the createdat field in the database.
		FieldCreatedAt = "created_at"
		// FieldUpdatedAt holds the string denoting the updatedat field in the database.
		FieldUpdatedAt = "updated_at"
		// EdgeFollower holds the string denoting the follower edge name in mutations.
		EdgeFollower = "follower"
		// EdgeFollowed holds the string denoting the followed edge name in mutations.
		EdgeFollowed = "followed"
	// Table holds the table name of the userfollowuser in the database.
	Table = "user_follow_users"
		// FollowerTable is the table that holds the follower relation/edge.
		FollowerTable = "user_follow_users"
			// FollowerInverseTable is the table name for the User entity.
			// It exists in this package in order to avoid circular dependency with the "user" package.
			FollowerInverseTable = "users"
			// FollowerColumn is the table column denoting the follower relation/edge.
			FollowerColumn = "user_followed_users"
		// FollowedTable is the table that holds the followed relation/edge.
		FollowedTable = "user_follow_users"
			// FollowedInverseTable is the table name for the User entity.
			// It exists in this package in order to avoid circular dependency with the "user" package.
			FollowedInverseTable = "users"
			// FollowedColumn is the table column denoting the followed relation/edge.
			FollowedColumn = "user_follower_users"

)



	
	// Columns holds all SQL columns for userfollowuser fields.
	var Columns = []string{
			FieldID,
			FieldCreatedAt,
			FieldUpdatedAt,
	}
	
	
		// ForeignKeys holds the SQL foreign-keys that are owned by the "user_follow_users"
		// table and are not defined as standalone fields in the schema.
		var ForeignKeys = []string{
				"user_followed_users",
				"user_follower_users",
		}
	

	





	
// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
		for i := range ForeignKeys {
			if column == ForeignKeys[i] {
				return true
			}
		}
	return false
}









	var (
				// DefaultCreatedAt holds the default value on creation for the "CreatedAt" field.
				DefaultCreatedAt func() time.Time
				// UpdateDefaultUpdatedAt holds the default value on update for the "UpdatedAt" field.
				UpdateDefaultUpdatedAt func() time.Time
				// IDValidator is a validator for the "id" field. It is called by the builders before save.
				IDValidator func (string) error
	)






// OrderOption defines the ordering options for the UserFollowUser queries.
type OrderOption func(*sql.Selector)

	
	
		// ByID orders the results by the id field.
		func ByID(opts ...sql.OrderTermOption) OrderOption {
			return sql.OrderByField(FieldID, opts...).ToFunc()
		}
	
			// ByCreatedAt orders the results by the CreatedAt field.
			func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
				return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
			}
			// ByUpdatedAt orders the results by the UpdatedAt field.
			func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
				return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
			}
			// ByFollowerField orders the results by follower field.
			func ByFollowerField(field string, opts ...sql.OrderTermOption) OrderOption {
				return func(s *sql.Selector) {
					sqlgraph.OrderByNeighborTerms(s, newFollowerStep(), sql.OrderByField(field, opts...))
				}
			}
			// ByFollowedField orders the results by followed field.
			func ByFollowedField(field string, opts ...sql.OrderTermOption) OrderOption {
				return func(s *sql.Selector) {
					sqlgraph.OrderByNeighborTerms(s, newFollowedStep(), sql.OrderByField(field, opts...))
				}
			}
		func newFollowerStep() *sqlgraph.Step {
			return sqlgraph.NewStep(
					sqlgraph.From(Table, FieldID),
						sqlgraph.To(FollowerInverseTable, FieldID),
				sqlgraph.Edge(sqlgraph.M2O, true, FollowerTable,FollowerColumn),
			)
		}
		func newFollowedStep() *sqlgraph.Step {
			return sqlgraph.NewStep(
					sqlgraph.From(Table, FieldID),
						sqlgraph.To(FollowedInverseTable, FieldID),
				sqlgraph.Edge(sqlgraph.M2O, true, FollowedTable,FollowedColumn),
			)
		}








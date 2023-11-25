



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
			 "placio-app/ent/businessfollowbusiness"
			 "placio-app/ent/business"

)




		// BusinessFollowBusiness is the model entity for the BusinessFollowBusiness schema.
type BusinessFollowBusiness struct {
	config `json:"-"`
		// ID of the ent.
		ID string `json:"id,omitempty"`
		// CreatedAt holds the value of the "CreatedAt" field.
		CreatedAt time.Time `json:"CreatedAt,omitempty"`
		// UpdatedAt holds the value of the "UpdatedAt" field.
		UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
		// Edges holds the relations/edges for other nodes in the graph.
		// The values are being populated by the BusinessFollowBusinessQuery when eager-loading is set.
		Edges BusinessFollowBusinessEdges `json:"edges"`
		business_followed_businesses *string
		business_follower_businesses *string
	selectValues sql.SelectValues

}
// BusinessFollowBusinessEdges holds the relations/edges for other nodes in the graph.
type BusinessFollowBusinessEdges struct {
		// Follower holds the value of the follower edge.
		Follower *Business `json:"follower,omitempty"`
		// Followed holds the value of the followed edge.
		Followed *Business `json:"followed,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool


}
	// FollowerOrErr returns the Follower value or an error if the edge
	// was not loaded in eager-loading, or loaded but was not found.
	func (e BusinessFollowBusinessEdges) FollowerOrErr() (*Business, error) {
		if e.loadedTypes[0] {
				if e.Follower == nil {
					// Edge was loaded but was not found.
					return nil, &NotFoundError{label: business.Label}
				}
			return e.Follower, nil
		}
		return nil, &NotLoadedError{edge: "follower"}
	}
	// FollowedOrErr returns the Followed value or an error if the edge
	// was not loaded in eager-loading, or loaded but was not found.
	func (e BusinessFollowBusinessEdges) FollowedOrErr() (*Business, error) {
		if e.loadedTypes[1] {
				if e.Followed == nil {
					// Edge was loaded but was not found.
					return nil, &NotFoundError{label: business.Label}
				}
			return e.Followed, nil
		}
		return nil, &NotLoadedError{edge: "followed"}
	}







	
	


	
	
	
	

	
	
		
	
	
	


// scanValues returns the types for scanning values from sql.Rows.
func (*BusinessFollowBusiness) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
				case businessfollowbusiness.FieldID:
					values[i] = new(sql.NullString)
				case businessfollowbusiness.FieldCreatedAt,businessfollowbusiness.FieldUpdatedAt:
					values[i] = new(sql.NullTime)
				case businessfollowbusiness.ForeignKeys[0]: // business_followed_businesses
					values[i] = new(sql.NullString)
				case businessfollowbusiness.ForeignKeys[1]: // business_follower_businesses
					values[i] = new(sql.NullString)
			default:
				values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BusinessFollowBusiness fields.
func (bfb *BusinessFollowBusiness) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
			case businessfollowbusiness.FieldID:
						if value, ok := values[i].(*sql.NullString); !ok {
			return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
					bfb.ID = value.String
		}
			case businessfollowbusiness.FieldCreatedAt:
					if value, ok := values[i].(*sql.NullTime); !ok {
			return fmt.Errorf("unexpected type %T for field CreatedAt", values[i])
			} else if value.Valid {
					bfb.CreatedAt = value.Time
		}
			case businessfollowbusiness.FieldUpdatedAt:
					if value, ok := values[i].(*sql.NullTime); !ok {
			return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
					bfb.UpdatedAt = value.Time
		}
			case businessfollowbusiness.ForeignKeys[0]:
						if value, ok := values[i].(*sql.NullString); !ok {
			return fmt.Errorf("unexpected type %T for field business_followed_businesses", values[i])
			} else if value.Valid {
					bfb.business_followed_businesses = new(string)
					*bfb.business_followed_businesses = value.String
		}
			case businessfollowbusiness.ForeignKeys[1]:
						if value, ok := values[i].(*sql.NullString); !ok {
			return fmt.Errorf("unexpected type %T for field business_follower_businesses", values[i])
			} else if value.Valid {
					bfb.business_follower_businesses = new(string)
					*bfb.business_follower_businesses = value.String
		}
		default:
			bfb.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BusinessFollowBusiness.
// This includes values selected through modifiers, order, etc.
func (bfb *BusinessFollowBusiness) Value(name string) (ent.Value, error) {
	return bfb.selectValues.Get(name)
}





	
	// QueryFollower queries the "follower" edge of the BusinessFollowBusiness entity.
	func (bfb *BusinessFollowBusiness) QueryFollower() *BusinessQuery {
		return NewBusinessFollowBusinessClient(bfb.config).QueryFollower(bfb)
	}

	
	// QueryFollowed queries the "followed" edge of the BusinessFollowBusiness entity.
	func (bfb *BusinessFollowBusiness) QueryFollowed() *BusinessQuery {
		return NewBusinessFollowBusinessClient(bfb.config).QueryFollowed(bfb)
	}


// Update returns a builder for updating this BusinessFollowBusiness.
// Note that you need to call BusinessFollowBusiness.Unwrap() before calling this method if this BusinessFollowBusiness
// was returned from a transaction, and the transaction was committed or rolled back.
func (bfb *BusinessFollowBusiness) Update() *BusinessFollowBusinessUpdateOne {
	return NewBusinessFollowBusinessClient(bfb.config).UpdateOne(bfb)
}

// Unwrap unwraps the BusinessFollowBusiness entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bfb *BusinessFollowBusiness) Unwrap() *BusinessFollowBusiness {
	_tx, ok := bfb.config.driver.(*txDriver)
	if !ok {
		panic("ent: BusinessFollowBusiness is not a transactional entity")
	}
	bfb.config.driver = _tx.drv
	return bfb
}


	

	// String implements the fmt.Stringer.
	func (bfb *BusinessFollowBusiness) String() string {
		var builder strings.Builder
		builder.WriteString("BusinessFollowBusiness(")
			builder.WriteString(fmt.Sprintf("id=%v, ", bfb.ID))
					builder.WriteString("CreatedAt=")
						builder.WriteString(bfb.CreatedAt.Format(time.ANSIC))
				builder.WriteString(", ")
					builder.WriteString("UpdatedAt=")
						builder.WriteString(bfb.UpdatedAt.Format(time.ANSIC))
		builder.WriteByte(')')
		return builder.String()
	}







// BusinessFollowBusinesses is a parsable slice of BusinessFollowBusiness.
type BusinessFollowBusinesses []*BusinessFollowBusiness


	
	




// Code generated by ent, DO NOT EDIT.



	

package comment




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



	// ID filters vertices based on their ID field.
	func ID(id string) predicate.Comment {
		return predicate.Comment(sql.FieldEQ(FieldID, id))
	}

	
		
		
		// IDEQ applies the EQ predicate on the ID field.
		func IDEQ(id string) predicate.Comment {
			return predicate.Comment(sql.FieldEQ(FieldID,id))
		}
	
		
		
		// IDNEQ applies the NEQ predicate on the ID field.
		func IDNEQ(id string) predicate.Comment {
			return predicate.Comment(sql.FieldNEQ(FieldID,id))
		}
	
		
		
		// IDIn applies the In predicate on the ID field.
		func IDIn(ids ...string) predicate.Comment {
			return predicate.Comment(sql.FieldIn(FieldID,ids...))
		}
	
		
		
		// IDNotIn applies the NotIn predicate on the ID field.
		func IDNotIn(ids ...string) predicate.Comment {
			return predicate.Comment(sql.FieldNotIn(FieldID,ids...))
		}
	
		
		
		// IDGT applies the GT predicate on the ID field.
		func IDGT(id string) predicate.Comment {
			return predicate.Comment(sql.FieldGT(FieldID,id))
		}
	
		
		
		// IDGTE applies the GTE predicate on the ID field.
		func IDGTE(id string) predicate.Comment {
			return predicate.Comment(sql.FieldGTE(FieldID,id))
		}
	
		
		
		// IDLT applies the LT predicate on the ID field.
		func IDLT(id string) predicate.Comment {
			return predicate.Comment(sql.FieldLT(FieldID,id))
		}
	
		
		
		// IDLTE applies the LTE predicate on the ID field.
		func IDLTE(id string) predicate.Comment {
			return predicate.Comment(sql.FieldLTE(FieldID,id))
		}
	
		
		
		// IDEqualFold applies the EqualFold predicate on the ID field.
		func IDEqualFold(id string) predicate.Comment {
			return predicate.Comment(sql.FieldEqualFold(FieldID,id))
		}
	
		
		
		// IDContainsFold applies the ContainsFold predicate on the ID field.
		func IDContainsFold(id string) predicate.Comment {
			return predicate.Comment(sql.FieldContainsFold(FieldID,id))
		}
	



	
	
	
	
	
		
		// Content applies equality check predicate on the "Content" field. It's identical to ContentEQ.
		func Content(v string) predicate.Comment {
				return predicate.Comment(sql.FieldEQ(FieldContent, v))
		}

	
	
	
	
	
		
		// CreatedAt applies equality check predicate on the "CreatedAt" field. It's identical to CreatedAtEQ.
		func CreatedAt(v time.Time) predicate.Comment {
				return predicate.Comment(sql.FieldEQ(FieldCreatedAt, v))
		}

	
	
	
	
	
		
		// UpdatedAt applies equality check predicate on the "UpdatedAt" field. It's identical to UpdatedAtEQ.
		func UpdatedAt(v time.Time) predicate.Comment {
				return predicate.Comment(sql.FieldEQ(FieldUpdatedAt, v))
		}

	
	
	
	
	
		
		// ParentCommentID applies equality check predicate on the "parentCommentID" field. It's identical to ParentCommentIDEQ.
		func ParentCommentID(v string) predicate.Comment {
				return predicate.Comment(sql.FieldEQ(FieldParentCommentID, v))
		}



	
		
		
		
		
		// ContentEQ applies the EQ predicate on the "Content" field.
		func ContentEQ(v string) predicate.Comment {
				return predicate.Comment(sql.FieldEQ(FieldContent, v))
		}
	
		
		
		
		
		// ContentNEQ applies the NEQ predicate on the "Content" field.
		func ContentNEQ(v string) predicate.Comment {
				return predicate.Comment(sql.FieldNEQ(FieldContent, v))
		}
	
		
		
		
		
		// ContentIn applies the In predicate on the "Content" field.
		func ContentIn(vs ...string) predicate.Comment {
				return predicate.Comment(sql.FieldIn(FieldContent, vs...))
		}
	
		
		
		
		
		// ContentNotIn applies the NotIn predicate on the "Content" field.
		func ContentNotIn(vs ...string) predicate.Comment {
				return predicate.Comment(sql.FieldNotIn(FieldContent, vs...))
		}
	
		
		
		
		
		// ContentGT applies the GT predicate on the "Content" field.
		func ContentGT(v string) predicate.Comment {
				return predicate.Comment(sql.FieldGT(FieldContent, v))
		}
	
		
		
		
		
		// ContentGTE applies the GTE predicate on the "Content" field.
		func ContentGTE(v string) predicate.Comment {
				return predicate.Comment(sql.FieldGTE(FieldContent, v))
		}
	
		
		
		
		
		// ContentLT applies the LT predicate on the "Content" field.
		func ContentLT(v string) predicate.Comment {
				return predicate.Comment(sql.FieldLT(FieldContent, v))
		}
	
		
		
		
		
		// ContentLTE applies the LTE predicate on the "Content" field.
		func ContentLTE(v string) predicate.Comment {
				return predicate.Comment(sql.FieldLTE(FieldContent, v))
		}
	
		
		
		
		
		// ContentContains applies the Contains predicate on the "Content" field.
		func ContentContains(v string) predicate.Comment {
				return predicate.Comment(sql.FieldContains(FieldContent, v))
		}
	
		
		
		
		
		// ContentHasPrefix applies the HasPrefix predicate on the "Content" field.
		func ContentHasPrefix(v string) predicate.Comment {
				return predicate.Comment(sql.FieldHasPrefix(FieldContent, v))
		}
	
		
		
		
		
		// ContentHasSuffix applies the HasSuffix predicate on the "Content" field.
		func ContentHasSuffix(v string) predicate.Comment {
				return predicate.Comment(sql.FieldHasSuffix(FieldContent, v))
		}
	
		
		
		
		
		// ContentEqualFold applies the EqualFold predicate on the "Content" field.
		func ContentEqualFold(v string) predicate.Comment {
				return predicate.Comment(sql.FieldEqualFold(FieldContent, v))
		}
	
		
		
		
		
		// ContentContainsFold applies the ContainsFold predicate on the "Content" field.
		func ContentContainsFold(v string) predicate.Comment {
				return predicate.Comment(sql.FieldContainsFold(FieldContent, v))
		}
	

	
		
		
		
		
		// CreatedAtEQ applies the EQ predicate on the "CreatedAt" field.
		func CreatedAtEQ(v time.Time) predicate.Comment {
				return predicate.Comment(sql.FieldEQ(FieldCreatedAt, v))
		}
	
		
		
		
		
		// CreatedAtNEQ applies the NEQ predicate on the "CreatedAt" field.
		func CreatedAtNEQ(v time.Time) predicate.Comment {
				return predicate.Comment(sql.FieldNEQ(FieldCreatedAt, v))
		}
	
		
		
		
		
		// CreatedAtIn applies the In predicate on the "CreatedAt" field.
		func CreatedAtIn(vs ...time.Time) predicate.Comment {
				return predicate.Comment(sql.FieldIn(FieldCreatedAt, vs...))
		}
	
		
		
		
		
		// CreatedAtNotIn applies the NotIn predicate on the "CreatedAt" field.
		func CreatedAtNotIn(vs ...time.Time) predicate.Comment {
				return predicate.Comment(sql.FieldNotIn(FieldCreatedAt, vs...))
		}
	
		
		
		
		
		// CreatedAtGT applies the GT predicate on the "CreatedAt" field.
		func CreatedAtGT(v time.Time) predicate.Comment {
				return predicate.Comment(sql.FieldGT(FieldCreatedAt, v))
		}
	
		
		
		
		
		// CreatedAtGTE applies the GTE predicate on the "CreatedAt" field.
		func CreatedAtGTE(v time.Time) predicate.Comment {
				return predicate.Comment(sql.FieldGTE(FieldCreatedAt, v))
		}
	
		
		
		
		
		// CreatedAtLT applies the LT predicate on the "CreatedAt" field.
		func CreatedAtLT(v time.Time) predicate.Comment {
				return predicate.Comment(sql.FieldLT(FieldCreatedAt, v))
		}
	
		
		
		
		
		// CreatedAtLTE applies the LTE predicate on the "CreatedAt" field.
		func CreatedAtLTE(v time.Time) predicate.Comment {
				return predicate.Comment(sql.FieldLTE(FieldCreatedAt, v))
		}
	

	
		
		
		
		
		// UpdatedAtEQ applies the EQ predicate on the "UpdatedAt" field.
		func UpdatedAtEQ(v time.Time) predicate.Comment {
				return predicate.Comment(sql.FieldEQ(FieldUpdatedAt, v))
		}
	
		
		
		
		
		// UpdatedAtNEQ applies the NEQ predicate on the "UpdatedAt" field.
		func UpdatedAtNEQ(v time.Time) predicate.Comment {
				return predicate.Comment(sql.FieldNEQ(FieldUpdatedAt, v))
		}
	
		
		
		
		
		// UpdatedAtIn applies the In predicate on the "UpdatedAt" field.
		func UpdatedAtIn(vs ...time.Time) predicate.Comment {
				return predicate.Comment(sql.FieldIn(FieldUpdatedAt, vs...))
		}
	
		
		
		
		
		// UpdatedAtNotIn applies the NotIn predicate on the "UpdatedAt" field.
		func UpdatedAtNotIn(vs ...time.Time) predicate.Comment {
				return predicate.Comment(sql.FieldNotIn(FieldUpdatedAt, vs...))
		}
	
		
		
		
		
		// UpdatedAtGT applies the GT predicate on the "UpdatedAt" field.
		func UpdatedAtGT(v time.Time) predicate.Comment {
				return predicate.Comment(sql.FieldGT(FieldUpdatedAt, v))
		}
	
		
		
		
		
		// UpdatedAtGTE applies the GTE predicate on the "UpdatedAt" field.
		func UpdatedAtGTE(v time.Time) predicate.Comment {
				return predicate.Comment(sql.FieldGTE(FieldUpdatedAt, v))
		}
	
		
		
		
		
		// UpdatedAtLT applies the LT predicate on the "UpdatedAt" field.
		func UpdatedAtLT(v time.Time) predicate.Comment {
				return predicate.Comment(sql.FieldLT(FieldUpdatedAt, v))
		}
	
		
		
		
		
		// UpdatedAtLTE applies the LTE predicate on the "UpdatedAt" field.
		func UpdatedAtLTE(v time.Time) predicate.Comment {
				return predicate.Comment(sql.FieldLTE(FieldUpdatedAt, v))
		}
	

	
		
		
		
		
		// ParentCommentIDEQ applies the EQ predicate on the "parentCommentID" field.
		func ParentCommentIDEQ(v string) predicate.Comment {
				return predicate.Comment(sql.FieldEQ(FieldParentCommentID, v))
		}
	
		
		
		
		
		// ParentCommentIDNEQ applies the NEQ predicate on the "parentCommentID" field.
		func ParentCommentIDNEQ(v string) predicate.Comment {
				return predicate.Comment(sql.FieldNEQ(FieldParentCommentID, v))
		}
	
		
		
		
		
		// ParentCommentIDIn applies the In predicate on the "parentCommentID" field.
		func ParentCommentIDIn(vs ...string) predicate.Comment {
				return predicate.Comment(sql.FieldIn(FieldParentCommentID, vs...))
		}
	
		
		
		
		
		// ParentCommentIDNotIn applies the NotIn predicate on the "parentCommentID" field.
		func ParentCommentIDNotIn(vs ...string) predicate.Comment {
				return predicate.Comment(sql.FieldNotIn(FieldParentCommentID, vs...))
		}
	
		
		
		
		
		// ParentCommentIDGT applies the GT predicate on the "parentCommentID" field.
		func ParentCommentIDGT(v string) predicate.Comment {
				return predicate.Comment(sql.FieldGT(FieldParentCommentID, v))
		}
	
		
		
		
		
		// ParentCommentIDGTE applies the GTE predicate on the "parentCommentID" field.
		func ParentCommentIDGTE(v string) predicate.Comment {
				return predicate.Comment(sql.FieldGTE(FieldParentCommentID, v))
		}
	
		
		
		
		
		// ParentCommentIDLT applies the LT predicate on the "parentCommentID" field.
		func ParentCommentIDLT(v string) predicate.Comment {
				return predicate.Comment(sql.FieldLT(FieldParentCommentID, v))
		}
	
		
		
		
		
		// ParentCommentIDLTE applies the LTE predicate on the "parentCommentID" field.
		func ParentCommentIDLTE(v string) predicate.Comment {
				return predicate.Comment(sql.FieldLTE(FieldParentCommentID, v))
		}
	
		
		
		
		
		// ParentCommentIDContains applies the Contains predicate on the "parentCommentID" field.
		func ParentCommentIDContains(v string) predicate.Comment {
				return predicate.Comment(sql.FieldContains(FieldParentCommentID, v))
		}
	
		
		
		
		
		// ParentCommentIDHasPrefix applies the HasPrefix predicate on the "parentCommentID" field.
		func ParentCommentIDHasPrefix(v string) predicate.Comment {
				return predicate.Comment(sql.FieldHasPrefix(FieldParentCommentID, v))
		}
	
		
		
		
		
		// ParentCommentIDHasSuffix applies the HasSuffix predicate on the "parentCommentID" field.
		func ParentCommentIDHasSuffix(v string) predicate.Comment {
				return predicate.Comment(sql.FieldHasSuffix(FieldParentCommentID, v))
		}
	
		
		
		
		
		// ParentCommentIDIsNil applies the IsNil predicate on the "parentCommentID" field.
		func ParentCommentIDIsNil() predicate.Comment {
				return predicate.Comment(sql.FieldIsNull(FieldParentCommentID))
		}
	
		
		
		
		
		// ParentCommentIDNotNil applies the NotNil predicate on the "parentCommentID" field.
		func ParentCommentIDNotNil() predicate.Comment {
				return predicate.Comment(sql.FieldNotNull(FieldParentCommentID))
		}
	
		
		
		
		
		// ParentCommentIDEqualFold applies the EqualFold predicate on the "parentCommentID" field.
		func ParentCommentIDEqualFold(v string) predicate.Comment {
				return predicate.Comment(sql.FieldEqualFold(FieldParentCommentID, v))
		}
	
		
		
		
		
		// ParentCommentIDContainsFold applies the ContainsFold predicate on the "parentCommentID" field.
		func ParentCommentIDContainsFold(v string) predicate.Comment {
				return predicate.Comment(sql.FieldContainsFold(FieldParentCommentID, v))
		}
	



	
	// HasUser applies the HasEdge predicate on the "user" edge.
	func HasUser() predicate.Comment {
		return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
				sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable,UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
	}
	
	// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
	func HasUserWith(preds ...predicate.User) predicate.Comment {
		return predicate.Comment(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
	}

	
	// HasPost applies the HasEdge predicate on the "post" edge.
	func HasPost() predicate.Comment {
		return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
				sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PostTable,PostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
	}
	
	// HasPostWith applies the HasEdge predicate on the "post" edge with a given conditions (other predicates).
	func HasPostWith(preds ...predicate.Post) predicate.Comment {
		return predicate.Comment(func(s *sql.Selector) {
		step := newPostStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
	}

	
	// HasParentComment applies the HasEdge predicate on the "parentComment" edge.
	func HasParentComment() predicate.Comment {
		return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
				sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentCommentTable,ParentCommentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
	}
	
	// HasParentCommentWith applies the HasEdge predicate on the "parentComment" edge with a given conditions (other predicates).
	func HasParentCommentWith(preds ...predicate.Comment) predicate.Comment {
		return predicate.Comment(func(s *sql.Selector) {
		step := newParentCommentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
	}

	
	// HasReplies applies the HasEdge predicate on the "replies" edge.
	func HasReplies() predicate.Comment {
		return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
				sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RepliesTable,RepliesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
	}
	
	// HasRepliesWith applies the HasEdge predicate on the "replies" edge with a given conditions (other predicates).
	func HasRepliesWith(preds ...predicate.Comment) predicate.Comment {
		return predicate.Comment(func(s *sql.Selector) {
		step := newRepliesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
	}

	
	// HasNotifications applies the HasEdge predicate on the "notifications" edge.
	func HasNotifications() predicate.Comment {
		return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
				sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, NotificationsTable,NotificationsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
	}
	
	// HasNotificationsWith applies the HasEdge predicate on the "notifications" edge with a given conditions (other predicates).
	func HasNotificationsWith(preds ...predicate.Notification) predicate.Comment {
		return predicate.Comment(func(s *sql.Selector) {
		step := newNotificationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
	}


// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Comment) predicate.Comment {
	return predicate.Comment(sql.NotPredicates(p))
}






// Code generated by ent, DO NOT EDIT.

package businessfollowevent

import (
	"placio-app/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasBusiness applies the HasEdge predicate on the "business" edge.
func HasBusiness() predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BusinessTable, BusinessColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessWith applies the HasEdge predicate on the "business" edge with a given conditions (other predicates).
func HasBusinessWith(preds ...predicate.Business) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(func(s *sql.Selector) {
		step := newBusinessStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEvent applies the HasEdge predicate on the "event" edge.
func HasEvent() predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, EventTable, EventColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventWith applies the HasEdge predicate on the "event" edge with a given conditions (other predicates).
func HasEventWith(preds ...predicate.Event) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(func(s *sql.Selector) {
		step := newEventStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BusinessFollowEvent) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BusinessFollowEvent) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BusinessFollowEvent) predicate.BusinessFollowEvent {
	return predicate.BusinessFollowEvent(func(s *sql.Selector) {
		p(s.Not())
	})
}

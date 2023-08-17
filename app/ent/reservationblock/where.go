// Code generated by ent, DO NOT EDIT.

package reservationblock

import (
	"placio-app/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldContainsFold(FieldID, id))
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldEQ(FieldStartTime, v))
}

// EndTime applies equality check predicate on the "end_time" field. It's identical to EndTimeEQ.
func EndTime(v time.Time) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldEQ(FieldEndTime, v))
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v time.Time) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldEQ(FieldStartTime, v))
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v time.Time) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldNEQ(FieldStartTime, v))
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...time.Time) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldIn(FieldStartTime, vs...))
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...time.Time) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldNotIn(FieldStartTime, vs...))
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v time.Time) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldGT(FieldStartTime, v))
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v time.Time) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldGTE(FieldStartTime, v))
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v time.Time) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldLT(FieldStartTime, v))
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v time.Time) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldLTE(FieldStartTime, v))
}

// EndTimeEQ applies the EQ predicate on the "end_time" field.
func EndTimeEQ(v time.Time) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldEQ(FieldEndTime, v))
}

// EndTimeNEQ applies the NEQ predicate on the "end_time" field.
func EndTimeNEQ(v time.Time) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldNEQ(FieldEndTime, v))
}

// EndTimeIn applies the In predicate on the "end_time" field.
func EndTimeIn(vs ...time.Time) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldIn(FieldEndTime, vs...))
}

// EndTimeNotIn applies the NotIn predicate on the "end_time" field.
func EndTimeNotIn(vs ...time.Time) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldNotIn(FieldEndTime, vs...))
}

// EndTimeGT applies the GT predicate on the "end_time" field.
func EndTimeGT(v time.Time) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldGT(FieldEndTime, v))
}

// EndTimeGTE applies the GTE predicate on the "end_time" field.
func EndTimeGTE(v time.Time) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldGTE(FieldEndTime, v))
}

// EndTimeLT applies the LT predicate on the "end_time" field.
func EndTimeLT(v time.Time) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldLT(FieldEndTime, v))
}

// EndTimeLTE applies the LTE predicate on the "end_time" field.
func EndTimeLTE(v time.Time) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldLTE(FieldEndTime, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.ReservationBlock {
	return predicate.ReservationBlock(sql.FieldNotIn(FieldStatus, vs...))
}

// HasPlaceInventory applies the HasEdge predicate on the "place_inventory" edge.
func HasPlaceInventory() predicate.ReservationBlock {
	return predicate.ReservationBlock(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlaceInventoryTable, PlaceInventoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlaceInventoryWith applies the HasEdge predicate on the "place_inventory" edge with a given conditions (other predicates).
func HasPlaceInventoryWith(preds ...predicate.PlaceInventory) predicate.ReservationBlock {
	return predicate.ReservationBlock(func(s *sql.Selector) {
		step := newPlaceInventoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.ReservationBlock {
	return predicate.ReservationBlock(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.ReservationBlock {
	return predicate.ReservationBlock(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ReservationBlock) predicate.ReservationBlock {
	return predicate.ReservationBlock(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ReservationBlock) predicate.ReservationBlock {
	return predicate.ReservationBlock(func(s *sql.Selector) {
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
func Not(p predicate.ReservationBlock) predicate.ReservationBlock {
	return predicate.ReservationBlock(func(s *sql.Selector) {
		p(s.Not())
	})
}

// Code generated by ent, DO NOT EDIT.

package booking

import (
	"placio-app/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Booking {
	return predicate.Booking(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Booking {
	return predicate.Booking(sql.FieldContainsFold(FieldID, id))
}

// StartDate applies equality check predicate on the "startDate" field. It's identical to StartDateEQ.
func StartDate(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldStartDate, v))
}

// EndDate applies equality check predicate on the "endDate" field. It's identical to EndDateEQ.
func EndDate(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldEndDate, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldStatus, v))
}

// BookingDate applies equality check predicate on the "bookingDate" field. It's identical to BookingDateEQ.
func BookingDate(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldBookingDate, v))
}

// StartDateEQ applies the EQ predicate on the "startDate" field.
func StartDateEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldStartDate, v))
}

// StartDateNEQ applies the NEQ predicate on the "startDate" field.
func StartDateNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldStartDate, v))
}

// StartDateIn applies the In predicate on the "startDate" field.
func StartDateIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldStartDate, vs...))
}

// StartDateNotIn applies the NotIn predicate on the "startDate" field.
func StartDateNotIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldStartDate, vs...))
}

// StartDateGT applies the GT predicate on the "startDate" field.
func StartDateGT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldStartDate, v))
}

// StartDateGTE applies the GTE predicate on the "startDate" field.
func StartDateGTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldStartDate, v))
}

// StartDateLT applies the LT predicate on the "startDate" field.
func StartDateLT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldStartDate, v))
}

// StartDateLTE applies the LTE predicate on the "startDate" field.
func StartDateLTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldStartDate, v))
}

// EndDateEQ applies the EQ predicate on the "endDate" field.
func EndDateEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldEndDate, v))
}

// EndDateNEQ applies the NEQ predicate on the "endDate" field.
func EndDateNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldEndDate, v))
}

// EndDateIn applies the In predicate on the "endDate" field.
func EndDateIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldEndDate, vs...))
}

// EndDateNotIn applies the NotIn predicate on the "endDate" field.
func EndDateNotIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldEndDate, vs...))
}

// EndDateGT applies the GT predicate on the "endDate" field.
func EndDateGT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldEndDate, v))
}

// EndDateGTE applies the GTE predicate on the "endDate" field.
func EndDateGTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldEndDate, v))
}

// EndDateLT applies the LT predicate on the "endDate" field.
func EndDateLT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldEndDate, v))
}

// EndDateLTE applies the LTE predicate on the "endDate" field.
func EndDateLTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldEndDate, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Booking {
	return predicate.Booking(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Booking {
	return predicate.Booking(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Booking {
	return predicate.Booking(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Booking {
	return predicate.Booking(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Booking {
	return predicate.Booking(sql.FieldContainsFold(FieldStatus, v))
}

// BookingDateEQ applies the EQ predicate on the "bookingDate" field.
func BookingDateEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldBookingDate, v))
}

// BookingDateNEQ applies the NEQ predicate on the "bookingDate" field.
func BookingDateNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldBookingDate, v))
}

// BookingDateIn applies the In predicate on the "bookingDate" field.
func BookingDateIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldBookingDate, vs...))
}

// BookingDateNotIn applies the NotIn predicate on the "bookingDate" field.
func BookingDateNotIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldBookingDate, vs...))
}

// BookingDateGT applies the GT predicate on the "bookingDate" field.
func BookingDateGT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldBookingDate, v))
}

// BookingDateGTE applies the GTE predicate on the "bookingDate" field.
func BookingDateGTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldBookingDate, v))
}

// BookingDateLT applies the LT predicate on the "bookingDate" field.
func BookingDateLT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldBookingDate, v))
}

// BookingDateLTE applies the LTE predicate on the "bookingDate" field.
func BookingDateLTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldBookingDate, v))
}

// HasRoom applies the HasEdge predicate on the "room" edge.
func HasRoom() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoomTable, RoomColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomWith applies the HasEdge predicate on the "room" edge with a given conditions (other predicates).
func HasRoomWith(preds ...predicate.Room) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := newRoomStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Booking) predicate.Booking {
	return predicate.Booking(sql.NotPredicates(p))
}

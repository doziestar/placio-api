// Code generated by ent, DO NOT EDIT.

package transactionhistory

import (
	"placio-app/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldContainsFold(FieldID, id))
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v int) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldEQ(FieldQuantity, v))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldEQ(FieldDate, v))
}

// TransactionTypeEQ applies the EQ predicate on the "transaction_type" field.
func TransactionTypeEQ(v TransactionType) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldEQ(FieldTransactionType, v))
}

// TransactionTypeNEQ applies the NEQ predicate on the "transaction_type" field.
func TransactionTypeNEQ(v TransactionType) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldNEQ(FieldTransactionType, v))
}

// TransactionTypeIn applies the In predicate on the "transaction_type" field.
func TransactionTypeIn(vs ...TransactionType) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldIn(FieldTransactionType, vs...))
}

// TransactionTypeNotIn applies the NotIn predicate on the "transaction_type" field.
func TransactionTypeNotIn(vs ...TransactionType) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldNotIn(FieldTransactionType, vs...))
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v int) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldEQ(FieldQuantity, v))
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v int) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldNEQ(FieldQuantity, v))
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...int) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldIn(FieldQuantity, vs...))
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...int) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldNotIn(FieldQuantity, vs...))
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v int) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldGT(FieldQuantity, v))
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v int) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldGTE(FieldQuantity, v))
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v int) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldLT(FieldQuantity, v))
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v int) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldLTE(FieldQuantity, v))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.TransactionHistory {
	return predicate.TransactionHistory(sql.FieldLTE(FieldDate, v))
}

// HasPlaceInventory applies the HasEdge predicate on the "place_inventory" edge.
func HasPlaceInventory() predicate.TransactionHistory {
	return predicate.TransactionHistory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlaceInventoryTable, PlaceInventoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlaceInventoryWith applies the HasEdge predicate on the "place_inventory" edge with a given conditions (other predicates).
func HasPlaceInventoryWith(preds ...predicate.PlaceInventory) predicate.TransactionHistory {
	return predicate.TransactionHistory(func(s *sql.Selector) {
		step := newPlaceInventoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.TransactionHistory {
	return predicate.TransactionHistory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.TransactionHistory {
	return predicate.TransactionHistory(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TransactionHistory) predicate.TransactionHistory {
	return predicate.TransactionHistory(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TransactionHistory) predicate.TransactionHistory {
	return predicate.TransactionHistory(func(s *sql.Selector) {
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
func Not(p predicate.TransactionHistory) predicate.TransactionHistory {
	return predicate.TransactionHistory(func(s *sql.Selector) {
		p(s.Not())
	})
}

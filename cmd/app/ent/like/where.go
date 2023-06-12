// Code generated by ent, DO NOT EDIT.

package like

import (
	"placio-app/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Like {
	return predicate.Like(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Like {
	return predicate.Like(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Like {
	return predicate.Like(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Like {
	return predicate.Like(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Like {
	return predicate.Like(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Like {
	return predicate.Like(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Like {
	return predicate.Like(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "CreatedAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "UpdatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "CreatedAt" field.
func CreatedAtEQ(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "CreatedAt" field.
func CreatedAtNEQ(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "CreatedAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Like {
	return predicate.Like(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "CreatedAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Like {
	return predicate.Like(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "CreatedAt" field.
func CreatedAtGT(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "CreatedAt" field.
func CreatedAtGTE(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "CreatedAt" field.
func CreatedAtLT(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "CreatedAt" field.
func CreatedAtLTE(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "UpdatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "UpdatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "UpdatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.Like {
	return predicate.Like(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "UpdatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Like {
	return predicate.Like(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "UpdatedAt" field.
func UpdatedAtGT(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "UpdatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "UpdatedAt" field.
func UpdatedAtLT(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "UpdatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Like {
	return predicate.Like(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Like {
	return predicate.Like(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPost applies the HasEdge predicate on the "post" edge.
func HasPost() predicate.Like {
	return predicate.Like(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PostTable, PostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostWith applies the HasEdge predicate on the "post" edge with a given conditions (other predicates).
func HasPostWith(preds ...predicate.Post) predicate.Like {
	return predicate.Like(func(s *sql.Selector) {
		step := newPostStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Like) predicate.Like {
	return predicate.Like(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Like) predicate.Like {
	return predicate.Like(func(s *sql.Selector) {
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
func Not(p predicate.Like) predicate.Like {
	return predicate.Like(func(s *sql.Selector) {
		p(s.Not())
	})
}

// Code generated by ent, DO NOT EDIT.

package businessfollowuser

import (
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.BusinessFollowUser {
	return predicate.BusinessFollowUser(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.BusinessFollowUser {
	return predicate.BusinessFollowUser(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.BusinessFollowUser {
	return predicate.BusinessFollowUser(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.BusinessFollowUser {
	return predicate.BusinessFollowUser(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.BusinessFollowUser {
	return predicate.BusinessFollowUser(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.BusinessFollowUser {
	return predicate.BusinessFollowUser(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.BusinessFollowUser {
	return predicate.BusinessFollowUser(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.BusinessFollowUser {
	return predicate.BusinessFollowUser(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.BusinessFollowUser {
	return predicate.BusinessFollowUser(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.BusinessFollowUser {
	return predicate.BusinessFollowUser(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.BusinessFollowUser {
	return predicate.BusinessFollowUser(sql.FieldContainsFold(FieldID, id))
}

// HasBusiness applies the HasEdge predicate on the "business" edge.
func HasBusiness() predicate.BusinessFollowUser {
	return predicate.BusinessFollowUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BusinessTable, BusinessColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessWith applies the HasEdge predicate on the "business" edge with a given conditions (other predicates).
func HasBusinessWith(preds ...predicate.Business) predicate.BusinessFollowUser {
	return predicate.BusinessFollowUser(func(s *sql.Selector) {
		step := newBusinessStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.BusinessFollowUser {
	return predicate.BusinessFollowUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.BusinessFollowUser {
	return predicate.BusinessFollowUser(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BusinessFollowUser) predicate.BusinessFollowUser {
	return predicate.BusinessFollowUser(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BusinessFollowUser) predicate.BusinessFollowUser {
	return predicate.BusinessFollowUser(func(s *sql.Selector) {
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
func Not(p predicate.BusinessFollowUser) predicate.BusinessFollowUser {
	return predicate.BusinessFollowUser(func(s *sql.Selector) {
		p(s.Not())
	})
}

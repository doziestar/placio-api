// Code generated by ent, DO NOT EDIT.

package customblock

import (
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldContainsFold(FieldID, id))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldEQ(FieldContent, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.CustomBlock {
	return predicate.CustomBlock(sql.FieldContainsFold(FieldContent, v))
}

// HasWebsite applies the HasEdge predicate on the "website" edge.
func HasWebsite() predicate.CustomBlock {
	return predicate.CustomBlock(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WebsiteTable, WebsiteColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWebsiteWith applies the HasEdge predicate on the "website" edge with a given conditions (other predicates).
func HasWebsiteWith(preds ...predicate.Website) predicate.CustomBlock {
	return predicate.CustomBlock(func(s *sql.Selector) {
		step := newWebsiteStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CustomBlock) predicate.CustomBlock {
	return predicate.CustomBlock(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CustomBlock) predicate.CustomBlock {
	return predicate.CustomBlock(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CustomBlock) predicate.CustomBlock {
	return predicate.CustomBlock(sql.NotPredicates(p))
}

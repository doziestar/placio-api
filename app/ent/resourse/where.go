// Code generated by ent, DO NOT EDIT.

package resourse

import (
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Resourse {
	return predicate.Resourse(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Resourse {
	return predicate.Resourse(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Resourse {
	return predicate.Resourse(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Resourse {
	return predicate.Resourse(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Resourse {
	return predicate.Resourse(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Resourse {
	return predicate.Resourse(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Resourse {
	return predicate.Resourse(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Resourse {
	return predicate.Resourse(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Resourse {
	return predicate.Resourse(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Resourse {
	return predicate.Resourse(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Resourse {
	return predicate.Resourse(sql.FieldContainsFold(FieldID, id))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Resourse) predicate.Resourse {
	return predicate.Resourse(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Resourse) predicate.Resourse {
	return predicate.Resourse(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Resourse) predicate.Resourse {
	return predicate.Resourse(sql.NotPredicates(p))
}

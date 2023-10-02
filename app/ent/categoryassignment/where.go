// Code generated by ent, DO NOT EDIT.

package categoryassignment

import (
	"placio_api/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldContainsFold(FieldID, id))
}

// EntityID applies equality check predicate on the "entity_id" field. It's identical to EntityIDEQ.
func EntityID(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldEQ(FieldEntityID, v))
}

// EntityType applies equality check predicate on the "entity_type" field. It's identical to EntityTypeEQ.
func EntityType(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldEQ(FieldEntityType, v))
}

// CategoryID applies equality check predicate on the "category_id" field. It's identical to CategoryIDEQ.
func CategoryID(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldEQ(FieldCategoryID, v))
}

// EntityIDEQ applies the EQ predicate on the "entity_id" field.
func EntityIDEQ(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldEQ(FieldEntityID, v))
}

// EntityIDNEQ applies the NEQ predicate on the "entity_id" field.
func EntityIDNEQ(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldNEQ(FieldEntityID, v))
}

// EntityIDIn applies the In predicate on the "entity_id" field.
func EntityIDIn(vs ...string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldIn(FieldEntityID, vs...))
}

// EntityIDNotIn applies the NotIn predicate on the "entity_id" field.
func EntityIDNotIn(vs ...string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldNotIn(FieldEntityID, vs...))
}

// EntityIDGT applies the GT predicate on the "entity_id" field.
func EntityIDGT(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldGT(FieldEntityID, v))
}

// EntityIDGTE applies the GTE predicate on the "entity_id" field.
func EntityIDGTE(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldGTE(FieldEntityID, v))
}

// EntityIDLT applies the LT predicate on the "entity_id" field.
func EntityIDLT(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldLT(FieldEntityID, v))
}

// EntityIDLTE applies the LTE predicate on the "entity_id" field.
func EntityIDLTE(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldLTE(FieldEntityID, v))
}

// EntityIDContains applies the Contains predicate on the "entity_id" field.
func EntityIDContains(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldContains(FieldEntityID, v))
}

// EntityIDHasPrefix applies the HasPrefix predicate on the "entity_id" field.
func EntityIDHasPrefix(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldHasPrefix(FieldEntityID, v))
}

// EntityIDHasSuffix applies the HasSuffix predicate on the "entity_id" field.
func EntityIDHasSuffix(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldHasSuffix(FieldEntityID, v))
}

// EntityIDIsNil applies the IsNil predicate on the "entity_id" field.
func EntityIDIsNil() predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldIsNull(FieldEntityID))
}

// EntityIDNotNil applies the NotNil predicate on the "entity_id" field.
func EntityIDNotNil() predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldNotNull(FieldEntityID))
}

// EntityIDEqualFold applies the EqualFold predicate on the "entity_id" field.
func EntityIDEqualFold(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldEqualFold(FieldEntityID, v))
}

// EntityIDContainsFold applies the ContainsFold predicate on the "entity_id" field.
func EntityIDContainsFold(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldContainsFold(FieldEntityID, v))
}

// EntityTypeEQ applies the EQ predicate on the "entity_type" field.
func EntityTypeEQ(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldEQ(FieldEntityType, v))
}

// EntityTypeNEQ applies the NEQ predicate on the "entity_type" field.
func EntityTypeNEQ(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldNEQ(FieldEntityType, v))
}

// EntityTypeIn applies the In predicate on the "entity_type" field.
func EntityTypeIn(vs ...string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldIn(FieldEntityType, vs...))
}

// EntityTypeNotIn applies the NotIn predicate on the "entity_type" field.
func EntityTypeNotIn(vs ...string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldNotIn(FieldEntityType, vs...))
}

// EntityTypeGT applies the GT predicate on the "entity_type" field.
func EntityTypeGT(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldGT(FieldEntityType, v))
}

// EntityTypeGTE applies the GTE predicate on the "entity_type" field.
func EntityTypeGTE(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldGTE(FieldEntityType, v))
}

// EntityTypeLT applies the LT predicate on the "entity_type" field.
func EntityTypeLT(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldLT(FieldEntityType, v))
}

// EntityTypeLTE applies the LTE predicate on the "entity_type" field.
func EntityTypeLTE(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldLTE(FieldEntityType, v))
}

// EntityTypeContains applies the Contains predicate on the "entity_type" field.
func EntityTypeContains(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldContains(FieldEntityType, v))
}

// EntityTypeHasPrefix applies the HasPrefix predicate on the "entity_type" field.
func EntityTypeHasPrefix(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldHasPrefix(FieldEntityType, v))
}

// EntityTypeHasSuffix applies the HasSuffix predicate on the "entity_type" field.
func EntityTypeHasSuffix(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldHasSuffix(FieldEntityType, v))
}

// EntityTypeIsNil applies the IsNil predicate on the "entity_type" field.
func EntityTypeIsNil() predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldIsNull(FieldEntityType))
}

// EntityTypeNotNil applies the NotNil predicate on the "entity_type" field.
func EntityTypeNotNil() predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldNotNull(FieldEntityType))
}

// EntityTypeEqualFold applies the EqualFold predicate on the "entity_type" field.
func EntityTypeEqualFold(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldEqualFold(FieldEntityType, v))
}

// EntityTypeContainsFold applies the ContainsFold predicate on the "entity_type" field.
func EntityTypeContainsFold(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldContainsFold(FieldEntityType, v))
}

// CategoryIDEQ applies the EQ predicate on the "category_id" field.
func CategoryIDEQ(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldEQ(FieldCategoryID, v))
}

// CategoryIDNEQ applies the NEQ predicate on the "category_id" field.
func CategoryIDNEQ(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldNEQ(FieldCategoryID, v))
}

// CategoryIDIn applies the In predicate on the "category_id" field.
func CategoryIDIn(vs ...string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldIn(FieldCategoryID, vs...))
}

// CategoryIDNotIn applies the NotIn predicate on the "category_id" field.
func CategoryIDNotIn(vs ...string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldNotIn(FieldCategoryID, vs...))
}

// CategoryIDGT applies the GT predicate on the "category_id" field.
func CategoryIDGT(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldGT(FieldCategoryID, v))
}

// CategoryIDGTE applies the GTE predicate on the "category_id" field.
func CategoryIDGTE(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldGTE(FieldCategoryID, v))
}

// CategoryIDLT applies the LT predicate on the "category_id" field.
func CategoryIDLT(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldLT(FieldCategoryID, v))
}

// CategoryIDLTE applies the LTE predicate on the "category_id" field.
func CategoryIDLTE(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldLTE(FieldCategoryID, v))
}

// CategoryIDContains applies the Contains predicate on the "category_id" field.
func CategoryIDContains(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldContains(FieldCategoryID, v))
}

// CategoryIDHasPrefix applies the HasPrefix predicate on the "category_id" field.
func CategoryIDHasPrefix(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldHasPrefix(FieldCategoryID, v))
}

// CategoryIDHasSuffix applies the HasSuffix predicate on the "category_id" field.
func CategoryIDHasSuffix(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldHasSuffix(FieldCategoryID, v))
}

// CategoryIDIsNil applies the IsNil predicate on the "category_id" field.
func CategoryIDIsNil() predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldIsNull(FieldCategoryID))
}

// CategoryIDNotNil applies the NotNil predicate on the "category_id" field.
func CategoryIDNotNil() predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldNotNull(FieldCategoryID))
}

// CategoryIDEqualFold applies the EqualFold predicate on the "category_id" field.
func CategoryIDEqualFold(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldEqualFold(FieldCategoryID, v))
}

// CategoryIDContainsFold applies the ContainsFold predicate on the "category_id" field.
func CategoryIDContainsFold(v string) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.FieldContainsFold(FieldCategoryID, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.CategoryAssignment {
	return predicate.CategoryAssignment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBusiness applies the HasEdge predicate on the "business" edge.
func HasBusiness() predicate.CategoryAssignment {
	return predicate.CategoryAssignment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BusinessTable, BusinessColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessWith applies the HasEdge predicate on the "business" edge with a given conditions (other predicates).
func HasBusinessWith(preds ...predicate.Business) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(func(s *sql.Selector) {
		step := newBusinessStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlace applies the HasEdge predicate on the "place" edge.
func HasPlace() predicate.CategoryAssignment {
	return predicate.CategoryAssignment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlaceTable, PlaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlaceWith applies the HasEdge predicate on the "place" edge with a given conditions (other predicates).
func HasPlaceWith(preds ...predicate.Place) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(func(s *sql.Selector) {
		step := newPlaceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCategory applies the HasEdge predicate on the "category" edge.
func HasCategory() predicate.CategoryAssignment {
	return predicate.CategoryAssignment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryWith applies the HasEdge predicate on the "category" edge with a given conditions (other predicates).
func HasCategoryWith(preds ...predicate.Category) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(func(s *sql.Selector) {
		step := newCategoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CategoryAssignment) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CategoryAssignment) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CategoryAssignment) predicate.CategoryAssignment {
	return predicate.CategoryAssignment(sql.NotPredicates(p))
}

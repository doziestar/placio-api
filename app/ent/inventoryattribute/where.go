// Code generated by ent, DO NOT EDIT.

package inventoryattribute

import (
	"placio_api/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldContainsFold(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldEQ(FieldName, v))
}

// IsMandatory applies equality check predicate on the "is_mandatory" field. It's identical to IsMandatoryEQ.
func IsMandatory(v bool) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldEQ(FieldIsMandatory, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldContainsFold(FieldName, v))
}

// IsMandatoryEQ applies the EQ predicate on the "is_mandatory" field.
func IsMandatoryEQ(v bool) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldEQ(FieldIsMandatory, v))
}

// IsMandatoryNEQ applies the NEQ predicate on the "is_mandatory" field.
func IsMandatoryNEQ(v bool) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldNEQ(FieldIsMandatory, v))
}

// DataTypeEQ applies the EQ predicate on the "data_type" field.
func DataTypeEQ(v DataType) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldEQ(FieldDataType, v))
}

// DataTypeNEQ applies the NEQ predicate on the "data_type" field.
func DataTypeNEQ(v DataType) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldNEQ(FieldDataType, v))
}

// DataTypeIn applies the In predicate on the "data_type" field.
func DataTypeIn(vs ...DataType) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldIn(FieldDataType, vs...))
}

// DataTypeNotIn applies the NotIn predicate on the "data_type" field.
func DataTypeNotIn(vs ...DataType) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldNotIn(FieldDataType, vs...))
}

// DataTypeIsNil applies the IsNil predicate on the "data_type" field.
func DataTypeIsNil() predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldIsNull(FieldDataType))
}

// DataTypeNotNil applies the NotNil predicate on the "data_type" field.
func DataTypeNotNil() predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.FieldNotNull(FieldDataType))
}

// HasInventoryType applies the HasEdge predicate on the "inventory_type" edge.
func HasInventoryType() predicate.InventoryAttribute {
	return predicate.InventoryAttribute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InventoryTypeTable, InventoryTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInventoryTypeWith applies the HasEdge predicate on the "inventory_type" edge with a given conditions (other predicates).
func HasInventoryTypeWith(preds ...predicate.InventoryType) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(func(s *sql.Selector) {
		step := newInventoryTypeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlaceInventoryAttributes applies the HasEdge predicate on the "place_inventory_attributes" edge.
func HasPlaceInventoryAttributes() predicate.InventoryAttribute {
	return predicate.InventoryAttribute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PlaceInventoryAttributesTable, PlaceInventoryAttributesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlaceInventoryAttributesWith applies the HasEdge predicate on the "place_inventory_attributes" edge with a given conditions (other predicates).
func HasPlaceInventoryAttributesWith(preds ...predicate.PlaceInventoryAttribute) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(func(s *sql.Selector) {
		step := newPlaceInventoryAttributesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.InventoryAttribute) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.InventoryAttribute) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.InventoryAttribute) predicate.InventoryAttribute {
	return predicate.InventoryAttribute(sql.NotPredicates(p))
}

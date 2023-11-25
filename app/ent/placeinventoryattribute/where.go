
// Code generated by ent, DO NOT EDIT.



	

package placeinventoryattribute




import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"
		"placio-app/ent/predicate"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
			"database/sql/driver"
			"entgo.io/ent/dialect/sql"
			"entgo.io/ent/dialect/sql/sqlgraph"
			"entgo.io/ent/dialect/sql/sqljson"
			"entgo.io/ent/schema/field"

)



	// ID filters vertices based on their ID field.
	func ID(id string) predicate.PlaceInventoryAttribute {
		return predicate.PlaceInventoryAttribute(sql.FieldEQ(FieldID, id))
	}

	
		
		
		// IDEQ applies the EQ predicate on the ID field.
		func IDEQ(id string) predicate.PlaceInventoryAttribute {
			return predicate.PlaceInventoryAttribute(sql.FieldEQ(FieldID,id))
		}
	
		
		
		// IDNEQ applies the NEQ predicate on the ID field.
		func IDNEQ(id string) predicate.PlaceInventoryAttribute {
			return predicate.PlaceInventoryAttribute(sql.FieldNEQ(FieldID,id))
		}
	
		
		
		// IDIn applies the In predicate on the ID field.
		func IDIn(ids ...string) predicate.PlaceInventoryAttribute {
			return predicate.PlaceInventoryAttribute(sql.FieldIn(FieldID,ids...))
		}
	
		
		
		// IDNotIn applies the NotIn predicate on the ID field.
		func IDNotIn(ids ...string) predicate.PlaceInventoryAttribute {
			return predicate.PlaceInventoryAttribute(sql.FieldNotIn(FieldID,ids...))
		}
	
		
		
		// IDGT applies the GT predicate on the ID field.
		func IDGT(id string) predicate.PlaceInventoryAttribute {
			return predicate.PlaceInventoryAttribute(sql.FieldGT(FieldID,id))
		}
	
		
		
		// IDGTE applies the GTE predicate on the ID field.
		func IDGTE(id string) predicate.PlaceInventoryAttribute {
			return predicate.PlaceInventoryAttribute(sql.FieldGTE(FieldID,id))
		}
	
		
		
		// IDLT applies the LT predicate on the ID field.
		func IDLT(id string) predicate.PlaceInventoryAttribute {
			return predicate.PlaceInventoryAttribute(sql.FieldLT(FieldID,id))
		}
	
		
		
		// IDLTE applies the LTE predicate on the ID field.
		func IDLTE(id string) predicate.PlaceInventoryAttribute {
			return predicate.PlaceInventoryAttribute(sql.FieldLTE(FieldID,id))
		}
	
		
		
		// IDEqualFold applies the EqualFold predicate on the ID field.
		func IDEqualFold(id string) predicate.PlaceInventoryAttribute {
			return predicate.PlaceInventoryAttribute(sql.FieldEqualFold(FieldID,id))
		}
	
		
		
		// IDContainsFold applies the ContainsFold predicate on the ID field.
		func IDContainsFold(id string) predicate.PlaceInventoryAttribute {
			return predicate.PlaceInventoryAttribute(sql.FieldContainsFold(FieldID,id))
		}
	



	
	
	
	
	
		
		// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
		func Value(v string) predicate.PlaceInventoryAttribute {
				return predicate.PlaceInventoryAttribute(sql.FieldEQ(FieldValue, v))
		}

	
	
	
	
	



	
		
		
		
		
		// ValueEQ applies the EQ predicate on the "value" field.
		func ValueEQ(v string) predicate.PlaceInventoryAttribute {
				return predicate.PlaceInventoryAttribute(sql.FieldEQ(FieldValue, v))
		}
	
		
		
		
		
		// ValueNEQ applies the NEQ predicate on the "value" field.
		func ValueNEQ(v string) predicate.PlaceInventoryAttribute {
				return predicate.PlaceInventoryAttribute(sql.FieldNEQ(FieldValue, v))
		}
	
		
		
		
		
		// ValueIn applies the In predicate on the "value" field.
		func ValueIn(vs ...string) predicate.PlaceInventoryAttribute {
				return predicate.PlaceInventoryAttribute(sql.FieldIn(FieldValue, vs...))
		}
	
		
		
		
		
		// ValueNotIn applies the NotIn predicate on the "value" field.
		func ValueNotIn(vs ...string) predicate.PlaceInventoryAttribute {
				return predicate.PlaceInventoryAttribute(sql.FieldNotIn(FieldValue, vs...))
		}
	
		
		
		
		
		// ValueGT applies the GT predicate on the "value" field.
		func ValueGT(v string) predicate.PlaceInventoryAttribute {
				return predicate.PlaceInventoryAttribute(sql.FieldGT(FieldValue, v))
		}
	
		
		
		
		
		// ValueGTE applies the GTE predicate on the "value" field.
		func ValueGTE(v string) predicate.PlaceInventoryAttribute {
				return predicate.PlaceInventoryAttribute(sql.FieldGTE(FieldValue, v))
		}
	
		
		
		
		
		// ValueLT applies the LT predicate on the "value" field.
		func ValueLT(v string) predicate.PlaceInventoryAttribute {
				return predicate.PlaceInventoryAttribute(sql.FieldLT(FieldValue, v))
		}
	
		
		
		
		
		// ValueLTE applies the LTE predicate on the "value" field.
		func ValueLTE(v string) predicate.PlaceInventoryAttribute {
				return predicate.PlaceInventoryAttribute(sql.FieldLTE(FieldValue, v))
		}
	
		
		
		
		
		// ValueContains applies the Contains predicate on the "value" field.
		func ValueContains(v string) predicate.PlaceInventoryAttribute {
				return predicate.PlaceInventoryAttribute(sql.FieldContains(FieldValue, v))
		}
	
		
		
		
		
		// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
		func ValueHasPrefix(v string) predicate.PlaceInventoryAttribute {
				return predicate.PlaceInventoryAttribute(sql.FieldHasPrefix(FieldValue, v))
		}
	
		
		
		
		
		// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
		func ValueHasSuffix(v string) predicate.PlaceInventoryAttribute {
				return predicate.PlaceInventoryAttribute(sql.FieldHasSuffix(FieldValue, v))
		}
	
		
		
		
		
		// ValueEqualFold applies the EqualFold predicate on the "value" field.
		func ValueEqualFold(v string) predicate.PlaceInventoryAttribute {
				return predicate.PlaceInventoryAttribute(sql.FieldEqualFold(FieldValue, v))
		}
	
		
		
		
		
		// ValueContainsFold applies the ContainsFold predicate on the "value" field.
		func ValueContainsFold(v string) predicate.PlaceInventoryAttribute {
				return predicate.PlaceInventoryAttribute(sql.FieldContainsFold(FieldValue, v))
		}
	

	
		
		
		
		
		// CategorySpecificValueIsNil applies the IsNil predicate on the "category_specific_value" field.
		func CategorySpecificValueIsNil() predicate.PlaceInventoryAttribute {
				return predicate.PlaceInventoryAttribute(sql.FieldIsNull(FieldCategorySpecificValue))
		}
	
		
		
		
		
		// CategorySpecificValueNotNil applies the NotNil predicate on the "category_specific_value" field.
		func CategorySpecificValueNotNil() predicate.PlaceInventoryAttribute {
				return predicate.PlaceInventoryAttribute(sql.FieldNotNull(FieldCategorySpecificValue))
		}
	



	
	// HasInventory applies the HasEdge predicate on the "inventory" edge.
	func HasInventory() predicate.PlaceInventoryAttribute {
		return predicate.PlaceInventoryAttribute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
				sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InventoryTable,InventoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
	}
	
	// HasInventoryWith applies the HasEdge predicate on the "inventory" edge with a given conditions (other predicates).
	func HasInventoryWith(preds ...predicate.PlaceInventory) predicate.PlaceInventoryAttribute {
		return predicate.PlaceInventoryAttribute(func(s *sql.Selector) {
		step := newInventoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
	}

	
	// HasAttributeType applies the HasEdge predicate on the "attribute_type" edge.
	func HasAttributeType() predicate.PlaceInventoryAttribute {
		return predicate.PlaceInventoryAttribute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
				sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AttributeTypeTable,AttributeTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
	}
	
	// HasAttributeTypeWith applies the HasEdge predicate on the "attribute_type" edge with a given conditions (other predicates).
	func HasAttributeTypeWith(preds ...predicate.InventoryAttribute) predicate.PlaceInventoryAttribute {
		return predicate.PlaceInventoryAttribute(func(s *sql.Selector) {
		step := newAttributeTypeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
	}


// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PlaceInventoryAttribute) predicate.PlaceInventoryAttribute {
	return predicate.PlaceInventoryAttribute(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PlaceInventoryAttribute) predicate.PlaceInventoryAttribute {
	return predicate.PlaceInventoryAttribute(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PlaceInventoryAttribute) predicate.PlaceInventoryAttribute {
	return predicate.PlaceInventoryAttribute(sql.NotPredicates(p))
}






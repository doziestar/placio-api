
// Code generated by ent, DO NOT EDIT.



	

package reaction




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
	func ID(id string) predicate.Reaction {
		return predicate.Reaction(sql.FieldEQ(FieldID, id))
	}

	
		
		
		// IDEQ applies the EQ predicate on the ID field.
		func IDEQ(id string) predicate.Reaction {
			return predicate.Reaction(sql.FieldEQ(FieldID,id))
		}
	
		
		
		// IDNEQ applies the NEQ predicate on the ID field.
		func IDNEQ(id string) predicate.Reaction {
			return predicate.Reaction(sql.FieldNEQ(FieldID,id))
		}
	
		
		
		// IDIn applies the In predicate on the ID field.
		func IDIn(ids ...string) predicate.Reaction {
			return predicate.Reaction(sql.FieldIn(FieldID,ids...))
		}
	
		
		
		// IDNotIn applies the NotIn predicate on the ID field.
		func IDNotIn(ids ...string) predicate.Reaction {
			return predicate.Reaction(sql.FieldNotIn(FieldID,ids...))
		}
	
		
		
		// IDGT applies the GT predicate on the ID field.
		func IDGT(id string) predicate.Reaction {
			return predicate.Reaction(sql.FieldGT(FieldID,id))
		}
	
		
		
		// IDGTE applies the GTE predicate on the ID field.
		func IDGTE(id string) predicate.Reaction {
			return predicate.Reaction(sql.FieldGTE(FieldID,id))
		}
	
		
		
		// IDLT applies the LT predicate on the ID field.
		func IDLT(id string) predicate.Reaction {
			return predicate.Reaction(sql.FieldLT(FieldID,id))
		}
	
		
		
		// IDLTE applies the LTE predicate on the ID field.
		func IDLTE(id string) predicate.Reaction {
			return predicate.Reaction(sql.FieldLTE(FieldID,id))
		}
	
		
		
		// IDEqualFold applies the EqualFold predicate on the ID field.
		func IDEqualFold(id string) predicate.Reaction {
			return predicate.Reaction(sql.FieldEqualFold(FieldID,id))
		}
	
		
		
		// IDContainsFold applies the ContainsFold predicate on the ID field.
		func IDContainsFold(id string) predicate.Reaction {
			return predicate.Reaction(sql.FieldContainsFold(FieldID,id))
		}
	








// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Reaction) predicate.Reaction {
	return predicate.Reaction(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Reaction) predicate.Reaction {
	return predicate.Reaction(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Reaction) predicate.Reaction {
	return predicate.Reaction(sql.NotPredicates(p))
}






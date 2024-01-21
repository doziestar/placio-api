// Code generated by ent, DO NOT EDIT.

package price

import (
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Price {
	return predicate.Price(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Price {
	return predicate.Price(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Price {
	return predicate.Price(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Price {
	return predicate.Price(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Price {
	return predicate.Price(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Price {
	return predicate.Price(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Price {
	return predicate.Price(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Price {
	return predicate.Price(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Price {
	return predicate.Price(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Price {
	return predicate.Price(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Price {
	return predicate.Price(sql.FieldContainsFold(FieldID, id))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.Price {
	return predicate.Price(sql.FieldEQ(FieldPrice, v))
}

// Currency applies equality check predicate on the "currency" field. It's identical to CurrencyEQ.
func Currency(v string) predicate.Price {
	return predicate.Price(sql.FieldEQ(FieldCurrency, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Price {
	return predicate.Price(sql.FieldEQ(FieldDescription, v))
}

// Duration applies equality check predicate on the "duration" field. It's identical to DurationEQ.
func Duration(v int) predicate.Price {
	return predicate.Price(sql.FieldEQ(FieldDuration, v))
}

// Session applies equality check predicate on the "session" field. It's identical to SessionEQ.
func Session(v int) predicate.Price {
	return predicate.Price(sql.FieldEQ(FieldSession, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.Price {
	return predicate.Price(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.Price {
	return predicate.Price(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.Price {
	return predicate.Price(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.Price {
	return predicate.Price(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.Price {
	return predicate.Price(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.Price {
	return predicate.Price(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.Price {
	return predicate.Price(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.Price {
	return predicate.Price(sql.FieldLTE(FieldPrice, v))
}

// CurrencyEQ applies the EQ predicate on the "currency" field.
func CurrencyEQ(v string) predicate.Price {
	return predicate.Price(sql.FieldEQ(FieldCurrency, v))
}

// CurrencyNEQ applies the NEQ predicate on the "currency" field.
func CurrencyNEQ(v string) predicate.Price {
	return predicate.Price(sql.FieldNEQ(FieldCurrency, v))
}

// CurrencyIn applies the In predicate on the "currency" field.
func CurrencyIn(vs ...string) predicate.Price {
	return predicate.Price(sql.FieldIn(FieldCurrency, vs...))
}

// CurrencyNotIn applies the NotIn predicate on the "currency" field.
func CurrencyNotIn(vs ...string) predicate.Price {
	return predicate.Price(sql.FieldNotIn(FieldCurrency, vs...))
}

// CurrencyGT applies the GT predicate on the "currency" field.
func CurrencyGT(v string) predicate.Price {
	return predicate.Price(sql.FieldGT(FieldCurrency, v))
}

// CurrencyGTE applies the GTE predicate on the "currency" field.
func CurrencyGTE(v string) predicate.Price {
	return predicate.Price(sql.FieldGTE(FieldCurrency, v))
}

// CurrencyLT applies the LT predicate on the "currency" field.
func CurrencyLT(v string) predicate.Price {
	return predicate.Price(sql.FieldLT(FieldCurrency, v))
}

// CurrencyLTE applies the LTE predicate on the "currency" field.
func CurrencyLTE(v string) predicate.Price {
	return predicate.Price(sql.FieldLTE(FieldCurrency, v))
}

// CurrencyContains applies the Contains predicate on the "currency" field.
func CurrencyContains(v string) predicate.Price {
	return predicate.Price(sql.FieldContains(FieldCurrency, v))
}

// CurrencyHasPrefix applies the HasPrefix predicate on the "currency" field.
func CurrencyHasPrefix(v string) predicate.Price {
	return predicate.Price(sql.FieldHasPrefix(FieldCurrency, v))
}

// CurrencyHasSuffix applies the HasSuffix predicate on the "currency" field.
func CurrencyHasSuffix(v string) predicate.Price {
	return predicate.Price(sql.FieldHasSuffix(FieldCurrency, v))
}

// CurrencyIsNil applies the IsNil predicate on the "currency" field.
func CurrencyIsNil() predicate.Price {
	return predicate.Price(sql.FieldIsNull(FieldCurrency))
}

// CurrencyNotNil applies the NotNil predicate on the "currency" field.
func CurrencyNotNil() predicate.Price {
	return predicate.Price(sql.FieldNotNull(FieldCurrency))
}

// CurrencyEqualFold applies the EqualFold predicate on the "currency" field.
func CurrencyEqualFold(v string) predicate.Price {
	return predicate.Price(sql.FieldEqualFold(FieldCurrency, v))
}

// CurrencyContainsFold applies the ContainsFold predicate on the "currency" field.
func CurrencyContainsFold(v string) predicate.Price {
	return predicate.Price(sql.FieldContainsFold(FieldCurrency, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Price {
	return predicate.Price(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Price {
	return predicate.Price(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Price {
	return predicate.Price(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Price {
	return predicate.Price(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Price {
	return predicate.Price(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Price {
	return predicate.Price(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Price {
	return predicate.Price(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Price {
	return predicate.Price(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Price {
	return predicate.Price(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Price {
	return predicate.Price(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Price {
	return predicate.Price(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Price {
	return predicate.Price(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Price {
	return predicate.Price(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Price {
	return predicate.Price(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Price {
	return predicate.Price(sql.FieldContainsFold(FieldDescription, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Price {
	return predicate.Price(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Price {
	return predicate.Price(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Price {
	return predicate.Price(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Price {
	return predicate.Price(sql.FieldNotIn(FieldType, vs...))
}

// DurationEQ applies the EQ predicate on the "duration" field.
func DurationEQ(v int) predicate.Price {
	return predicate.Price(sql.FieldEQ(FieldDuration, v))
}

// DurationNEQ applies the NEQ predicate on the "duration" field.
func DurationNEQ(v int) predicate.Price {
	return predicate.Price(sql.FieldNEQ(FieldDuration, v))
}

// DurationIn applies the In predicate on the "duration" field.
func DurationIn(vs ...int) predicate.Price {
	return predicate.Price(sql.FieldIn(FieldDuration, vs...))
}

// DurationNotIn applies the NotIn predicate on the "duration" field.
func DurationNotIn(vs ...int) predicate.Price {
	return predicate.Price(sql.FieldNotIn(FieldDuration, vs...))
}

// DurationGT applies the GT predicate on the "duration" field.
func DurationGT(v int) predicate.Price {
	return predicate.Price(sql.FieldGT(FieldDuration, v))
}

// DurationGTE applies the GTE predicate on the "duration" field.
func DurationGTE(v int) predicate.Price {
	return predicate.Price(sql.FieldGTE(FieldDuration, v))
}

// DurationLT applies the LT predicate on the "duration" field.
func DurationLT(v int) predicate.Price {
	return predicate.Price(sql.FieldLT(FieldDuration, v))
}

// DurationLTE applies the LTE predicate on the "duration" field.
func DurationLTE(v int) predicate.Price {
	return predicate.Price(sql.FieldLTE(FieldDuration, v))
}

// DurationIsNil applies the IsNil predicate on the "duration" field.
func DurationIsNil() predicate.Price {
	return predicate.Price(sql.FieldIsNull(FieldDuration))
}

// DurationNotNil applies the NotNil predicate on the "duration" field.
func DurationNotNil() predicate.Price {
	return predicate.Price(sql.FieldNotNull(FieldDuration))
}

// SessionEQ applies the EQ predicate on the "session" field.
func SessionEQ(v int) predicate.Price {
	return predicate.Price(sql.FieldEQ(FieldSession, v))
}

// SessionNEQ applies the NEQ predicate on the "session" field.
func SessionNEQ(v int) predicate.Price {
	return predicate.Price(sql.FieldNEQ(FieldSession, v))
}

// SessionIn applies the In predicate on the "session" field.
func SessionIn(vs ...int) predicate.Price {
	return predicate.Price(sql.FieldIn(FieldSession, vs...))
}

// SessionNotIn applies the NotIn predicate on the "session" field.
func SessionNotIn(vs ...int) predicate.Price {
	return predicate.Price(sql.FieldNotIn(FieldSession, vs...))
}

// SessionGT applies the GT predicate on the "session" field.
func SessionGT(v int) predicate.Price {
	return predicate.Price(sql.FieldGT(FieldSession, v))
}

// SessionGTE applies the GTE predicate on the "session" field.
func SessionGTE(v int) predicate.Price {
	return predicate.Price(sql.FieldGTE(FieldSession, v))
}

// SessionLT applies the LT predicate on the "session" field.
func SessionLT(v int) predicate.Price {
	return predicate.Price(sql.FieldLT(FieldSession, v))
}

// SessionLTE applies the LTE predicate on the "session" field.
func SessionLTE(v int) predicate.Price {
	return predicate.Price(sql.FieldLTE(FieldSession, v))
}

// SessionIsNil applies the IsNil predicate on the "session" field.
func SessionIsNil() predicate.Price {
	return predicate.Price(sql.FieldIsNull(FieldSession))
}

// SessionNotNil applies the NotNil predicate on the "session" field.
func SessionNotNil() predicate.Price {
	return predicate.Price(sql.FieldNotNull(FieldSession))
}

// HasPlan applies the HasEdge predicate on the "plan" edge.
func HasPlan() predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlanTable, PlanColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlanWith applies the HasEdge predicate on the "plan" edge with a given conditions (other predicates).
func HasPlanWith(preds ...predicate.Plan) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		step := newPlanStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubscriptions applies the HasEdge predicate on the "subscriptions" edge.
func HasSubscriptions() predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SubscriptionsTable, SubscriptionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubscriptionsWith applies the HasEdge predicate on the "subscriptions" edge with a given conditions (other predicates).
func HasSubscriptionsWith(preds ...predicate.Subscription) predicate.Price {
	return predicate.Price(func(s *sql.Selector) {
		step := newSubscriptionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Price) predicate.Price {
	return predicate.Price(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Price) predicate.Price {
	return predicate.Price(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Price) predicate.Price {
	return predicate.Price(sql.NotPredicates(p))
}

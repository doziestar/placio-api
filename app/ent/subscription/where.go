// Code generated by ent, DO NOT EDIT.

package subscription

import (
	"placio-app/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Subscription {
	return predicate.Subscription(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Subscription {
	return predicate.Subscription(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Subscription {
	return predicate.Subscription(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Subscription {
	return predicate.Subscription(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Subscription {
	return predicate.Subscription(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Subscription {
	return predicate.Subscription(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Subscription {
	return predicate.Subscription(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Subscription {
	return predicate.Subscription(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Subscription {
	return predicate.Subscription(sql.FieldContainsFold(FieldID, id))
}

// StartDate applies equality check predicate on the "start_date" field. It's identical to StartDateEQ.
func StartDate(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldStartDate, v))
}

// EndDate applies equality check predicate on the "end_date" field. It's identical to EndDateEQ.
func EndDate(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldEndDate, v))
}

// FlutterwaveSubscriptionID applies equality check predicate on the "flutterwave_subscription_id" field. It's identical to FlutterwaveSubscriptionIDEQ.
func FlutterwaveSubscriptionID(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldFlutterwaveSubscriptionID, v))
}

// StartDateEQ applies the EQ predicate on the "start_date" field.
func StartDateEQ(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldStartDate, v))
}

// StartDateNEQ applies the NEQ predicate on the "start_date" field.
func StartDateNEQ(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldNEQ(FieldStartDate, v))
}

// StartDateIn applies the In predicate on the "start_date" field.
func StartDateIn(vs ...time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldIn(FieldStartDate, vs...))
}

// StartDateNotIn applies the NotIn predicate on the "start_date" field.
func StartDateNotIn(vs ...time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldNotIn(FieldStartDate, vs...))
}

// StartDateGT applies the GT predicate on the "start_date" field.
func StartDateGT(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldGT(FieldStartDate, v))
}

// StartDateGTE applies the GTE predicate on the "start_date" field.
func StartDateGTE(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldGTE(FieldStartDate, v))
}

// StartDateLT applies the LT predicate on the "start_date" field.
func StartDateLT(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldLT(FieldStartDate, v))
}

// StartDateLTE applies the LTE predicate on the "start_date" field.
func StartDateLTE(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldLTE(FieldStartDate, v))
}

// EndDateEQ applies the EQ predicate on the "end_date" field.
func EndDateEQ(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldEndDate, v))
}

// EndDateNEQ applies the NEQ predicate on the "end_date" field.
func EndDateNEQ(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldNEQ(FieldEndDate, v))
}

// EndDateIn applies the In predicate on the "end_date" field.
func EndDateIn(vs ...time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldIn(FieldEndDate, vs...))
}

// EndDateNotIn applies the NotIn predicate on the "end_date" field.
func EndDateNotIn(vs ...time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldNotIn(FieldEndDate, vs...))
}

// EndDateGT applies the GT predicate on the "end_date" field.
func EndDateGT(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldGT(FieldEndDate, v))
}

// EndDateGTE applies the GTE predicate on the "end_date" field.
func EndDateGTE(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldGTE(FieldEndDate, v))
}

// EndDateLT applies the LT predicate on the "end_date" field.
func EndDateLT(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldLT(FieldEndDate, v))
}

// EndDateLTE applies the LTE predicate on the "end_date" field.
func EndDateLTE(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldLTE(FieldEndDate, v))
}

// FlutterwaveSubscriptionIDEQ applies the EQ predicate on the "flutterwave_subscription_id" field.
func FlutterwaveSubscriptionIDEQ(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldFlutterwaveSubscriptionID, v))
}

// FlutterwaveSubscriptionIDNEQ applies the NEQ predicate on the "flutterwave_subscription_id" field.
func FlutterwaveSubscriptionIDNEQ(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldNEQ(FieldFlutterwaveSubscriptionID, v))
}

// FlutterwaveSubscriptionIDIn applies the In predicate on the "flutterwave_subscription_id" field.
func FlutterwaveSubscriptionIDIn(vs ...string) predicate.Subscription {
	return predicate.Subscription(sql.FieldIn(FieldFlutterwaveSubscriptionID, vs...))
}

// FlutterwaveSubscriptionIDNotIn applies the NotIn predicate on the "flutterwave_subscription_id" field.
func FlutterwaveSubscriptionIDNotIn(vs ...string) predicate.Subscription {
	return predicate.Subscription(sql.FieldNotIn(FieldFlutterwaveSubscriptionID, vs...))
}

// FlutterwaveSubscriptionIDGT applies the GT predicate on the "flutterwave_subscription_id" field.
func FlutterwaveSubscriptionIDGT(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldGT(FieldFlutterwaveSubscriptionID, v))
}

// FlutterwaveSubscriptionIDGTE applies the GTE predicate on the "flutterwave_subscription_id" field.
func FlutterwaveSubscriptionIDGTE(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldGTE(FieldFlutterwaveSubscriptionID, v))
}

// FlutterwaveSubscriptionIDLT applies the LT predicate on the "flutterwave_subscription_id" field.
func FlutterwaveSubscriptionIDLT(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldLT(FieldFlutterwaveSubscriptionID, v))
}

// FlutterwaveSubscriptionIDLTE applies the LTE predicate on the "flutterwave_subscription_id" field.
func FlutterwaveSubscriptionIDLTE(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldLTE(FieldFlutterwaveSubscriptionID, v))
}

// FlutterwaveSubscriptionIDContains applies the Contains predicate on the "flutterwave_subscription_id" field.
func FlutterwaveSubscriptionIDContains(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldContains(FieldFlutterwaveSubscriptionID, v))
}

// FlutterwaveSubscriptionIDHasPrefix applies the HasPrefix predicate on the "flutterwave_subscription_id" field.
func FlutterwaveSubscriptionIDHasPrefix(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldHasPrefix(FieldFlutterwaveSubscriptionID, v))
}

// FlutterwaveSubscriptionIDHasSuffix applies the HasSuffix predicate on the "flutterwave_subscription_id" field.
func FlutterwaveSubscriptionIDHasSuffix(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldHasSuffix(FieldFlutterwaveSubscriptionID, v))
}

// FlutterwaveSubscriptionIDIsNil applies the IsNil predicate on the "flutterwave_subscription_id" field.
func FlutterwaveSubscriptionIDIsNil() predicate.Subscription {
	return predicate.Subscription(sql.FieldIsNull(FieldFlutterwaveSubscriptionID))
}

// FlutterwaveSubscriptionIDNotNil applies the NotNil predicate on the "flutterwave_subscription_id" field.
func FlutterwaveSubscriptionIDNotNil() predicate.Subscription {
	return predicate.Subscription(sql.FieldNotNull(FieldFlutterwaveSubscriptionID))
}

// FlutterwaveSubscriptionIDEqualFold applies the EqualFold predicate on the "flutterwave_subscription_id" field.
func FlutterwaveSubscriptionIDEqualFold(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldEqualFold(FieldFlutterwaveSubscriptionID, v))
}

// FlutterwaveSubscriptionIDContainsFold applies the ContainsFold predicate on the "flutterwave_subscription_id" field.
func FlutterwaveSubscriptionIDContainsFold(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldContainsFold(FieldFlutterwaveSubscriptionID, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlan applies the HasEdge predicate on the "plan" edge.
func HasPlan() predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlanTable, PlanColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlanWith applies the HasEdge predicate on the "plan" edge with a given conditions (other predicates).
func HasPlanWith(preds ...predicate.Plan) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		step := newPlanStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPrice applies the HasEdge predicate on the "price" edge.
func HasPrice() predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PriceTable, PriceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPriceWith applies the HasEdge predicate on the "price" edge with a given conditions (other predicates).
func HasPriceWith(preds ...predicate.Price) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		step := newPriceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Subscription) predicate.Subscription {
	return predicate.Subscription(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Subscription) predicate.Subscription {
	return predicate.Subscription(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Subscription) predicate.Subscription {
	return predicate.Subscription(sql.NotPredicates(p))
}
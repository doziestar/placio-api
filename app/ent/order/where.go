// Code generated by ent, DO NOT EDIT.

package order

import (
	"placio-app/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldUpdatedAt, v))
}

// TotalAmount applies equality check predicate on the "total_amount" field. It's identical to TotalAmountEQ.
func TotalAmount(v float64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldTotalAmount, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldDeletedAt, v))
}

// IsDeleted applies equality check predicate on the "is_deleted" field. It's identical to IsDeletedEQ.
func IsDeleted(v bool) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldIsDeleted, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldStatus, vs...))
}

// TotalAmountEQ applies the EQ predicate on the "total_amount" field.
func TotalAmountEQ(v float64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldTotalAmount, v))
}

// TotalAmountNEQ applies the NEQ predicate on the "total_amount" field.
func TotalAmountNEQ(v float64) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldTotalAmount, v))
}

// TotalAmountIn applies the In predicate on the "total_amount" field.
func TotalAmountIn(vs ...float64) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldTotalAmount, vs...))
}

// TotalAmountNotIn applies the NotIn predicate on the "total_amount" field.
func TotalAmountNotIn(vs ...float64) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldTotalAmount, vs...))
}

// TotalAmountGT applies the GT predicate on the "total_amount" field.
func TotalAmountGT(v float64) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldTotalAmount, v))
}

// TotalAmountGTE applies the GTE predicate on the "total_amount" field.
func TotalAmountGTE(v float64) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldTotalAmount, v))
}

// TotalAmountLT applies the LT predicate on the "total_amount" field.
func TotalAmountLT(v float64) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldTotalAmount, v))
}

// TotalAmountLTE applies the LTE predicate on the "total_amount" field.
func TotalAmountLTE(v float64) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldTotalAmount, v))
}

// AdditionalInfoIsNil applies the IsNil predicate on the "additional_info" field.
func AdditionalInfoIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldAdditionalInfo))
}

// AdditionalInfoNotNil applies the NotNil predicate on the "additional_info" field.
func AdditionalInfoNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldAdditionalInfo))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtContains applies the Contains predicate on the "deleted_at" field.
func DeletedAtContains(v string) predicate.Order {
	return predicate.Order(sql.FieldContains(FieldDeletedAt, v))
}

// DeletedAtHasPrefix applies the HasPrefix predicate on the "deleted_at" field.
func DeletedAtHasPrefix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasPrefix(FieldDeletedAt, v))
}

// DeletedAtHasSuffix applies the HasSuffix predicate on the "deleted_at" field.
func DeletedAtHasSuffix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasSuffix(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldDeletedAt))
}

// DeletedAtEqualFold applies the EqualFold predicate on the "deleted_at" field.
func DeletedAtEqualFold(v string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldDeletedAt, v))
}

// DeletedAtContainsFold applies the ContainsFold predicate on the "deleted_at" field.
func DeletedAtContainsFold(v string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldDeletedAt, v))
}

// IsDeletedEQ applies the EQ predicate on the "is_deleted" field.
func IsDeletedEQ(v bool) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldIsDeleted, v))
}

// IsDeletedNEQ applies the NEQ predicate on the "is_deleted" field.
func IsDeletedNEQ(v bool) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldIsDeleted, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrderItems applies the HasEdge predicate on the "order_items" edge.
func HasOrderItems() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, OrderItemsTable, OrderItemsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderItemsWith applies the HasEdge predicate on the "order_items" edge with a given conditions (other predicates).
func HasOrderItemsWith(preds ...predicate.OrderItem) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newOrderItemsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTable applies the HasEdge predicate on the "table" edge.
func HasTable() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TableTable, TablePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTableWith applies the HasEdge predicate on the "table" edge with a given conditions (other predicates).
func HasTableWith(preds ...predicate.PlaceTable) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newTableStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Order) predicate.Order {
	return predicate.Order(sql.NotPredicates(p))
}

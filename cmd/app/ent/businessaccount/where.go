// Code generated by ent, DO NOT EDIT.

package businessaccount

import (
	"placio-app/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldLTE(FieldID, id))
}

// ID applies equality check predicate on the "ID" field. It's identical to IDEQ.
func ID(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldEQ(FieldID, v))
}

// Name applies equality check predicate on the "Name" field. It's identical to NameEQ.
func Name(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldEQ(FieldName, v))
}

// Active applies equality check predicate on the "Active" field. It's identical to ActiveEQ.
func Active(v bool) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldEQ(FieldActive, v))
}

// CreatedAt applies equality check predicate on the "CreatedAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "UpdatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldEQ(FieldUpdatedAt, v))
}

// IDEQ applies the EQ predicate on the "ID" field.
func IDEQ(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldEQ(FieldID, v))
}

// IDNEQ applies the NEQ predicate on the "ID" field.
func IDNEQ(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldNEQ(FieldID, v))
}

// IDIn applies the In predicate on the "ID" field.
func IDIn(vs ...string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldIn(FieldID, vs...))
}

// IDNotIn applies the NotIn predicate on the "ID" field.
func IDNotIn(vs ...string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldNotIn(FieldID, vs...))
}

// IDGT applies the GT predicate on the "ID" field.
func IDGT(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldGT(FieldID, v))
}

// IDGTE applies the GTE predicate on the "ID" field.
func IDGTE(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldGTE(FieldID, v))
}

// IDLT applies the LT predicate on the "ID" field.
func IDLT(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldLT(FieldID, v))
}

// IDLTE applies the LTE predicate on the "ID" field.
func IDLTE(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldLTE(FieldID, v))
}

// IDEqualFold applies the EqualFold predicate on the "ID" field.
func IDEqualFold(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldEqualFold(FieldID, v))
}

// IDContainsFold applies the ContainsFold predicate on the "ID" field.
func IDContainsFold(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldContainsFold(FieldID, v))
}

// NameEQ applies the EQ predicate on the "Name" field.
func NameEQ(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "Name" field.
func NameNEQ(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "Name" field.
func NameIn(vs ...string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "Name" field.
func NameNotIn(vs ...string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "Name" field.
func NameGT(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "Name" field.
func NameGTE(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "Name" field.
func NameLT(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "Name" field.
func NameLTE(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "Name" field.
func NameContains(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "Name" field.
func NameHasPrefix(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "Name" field.
func NameHasSuffix(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "Name" field.
func NameEqualFold(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "Name" field.
func NameContainsFold(v string) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldContainsFold(FieldName, v))
}

// ActiveEQ applies the EQ predicate on the "Active" field.
func ActiveEQ(v bool) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldEQ(FieldActive, v))
}

// ActiveNEQ applies the NEQ predicate on the "Active" field.
func ActiveNEQ(v bool) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldNEQ(FieldActive, v))
}

// CreatedAtEQ applies the EQ predicate on the "CreatedAt" field.
func CreatedAtEQ(v time.Time) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "CreatedAt" field.
func CreatedAtNEQ(v time.Time) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "CreatedAt" field.
func CreatedAtIn(vs ...time.Time) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "CreatedAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "CreatedAt" field.
func CreatedAtGT(v time.Time) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "CreatedAt" field.
func CreatedAtGTE(v time.Time) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "CreatedAt" field.
func CreatedAtLT(v time.Time) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "CreatedAt" field.
func CreatedAtLTE(v time.Time) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "UpdatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "UpdatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "UpdatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "UpdatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "UpdatedAt" field.
func UpdatedAtGT(v time.Time) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "UpdatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "UpdatedAt" field.
func UpdatedAtLT(v time.Time) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "UpdatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.BusinessAccount {
	return predicate.BusinessAccount(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasPosts applies the HasEdge predicate on the "posts" edge.
func HasPosts() predicate.BusinessAccount {
	return predicate.BusinessAccount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PostsTable, PostsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostsWith applies the HasEdge predicate on the "posts" edge with a given conditions (other predicates).
func HasPostsWith(preds ...predicate.Post) predicate.BusinessAccount {
	return predicate.BusinessAccount(func(s *sql.Selector) {
		step := newPostsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRelationships applies the HasEdge predicate on the "relationships" edge.
func HasRelationships() predicate.BusinessAccount {
	return predicate.BusinessAccount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RelationshipsTable, RelationshipsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRelationshipsWith applies the HasEdge predicate on the "relationships" edge with a given conditions (other predicates).
func HasRelationshipsWith(preds ...predicate.UserBusinessRelationship) predicate.BusinessAccount {
	return predicate.BusinessAccount(func(s *sql.Selector) {
		step := newRelationshipsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAccountSettings applies the HasEdge predicate on the "account_settings" edge.
func HasAccountSettings() predicate.BusinessAccount {
	return predicate.BusinessAccount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AccountSettingsTable, AccountSettingsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountSettingsWith applies the HasEdge predicate on the "account_settings" edge with a given conditions (other predicates).
func HasAccountSettingsWith(preds ...predicate.AccountSettings) predicate.BusinessAccount {
	return predicate.BusinessAccount(func(s *sql.Selector) {
		step := newAccountSettingsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInvitations applies the HasEdge predicate on the "invitations" edge.
func HasInvitations() predicate.BusinessAccount {
	return predicate.BusinessAccount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, InvitationsTable, InvitationsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInvitationsWith applies the HasEdge predicate on the "invitations" edge with a given conditions (other predicates).
func HasInvitationsWith(preds ...predicate.Invitation) predicate.BusinessAccount {
	return predicate.BusinessAccount(func(s *sql.Selector) {
		step := newInvitationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BusinessAccount) predicate.BusinessAccount {
	return predicate.BusinessAccount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BusinessAccount) predicate.BusinessAccount {
	return predicate.BusinessAccount(func(s *sql.Selector) {
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
func Not(p predicate.BusinessAccount) predicate.BusinessAccount {
	return predicate.BusinessAccount(func(s *sql.Selector) {
		p(s.Not())
	})
}

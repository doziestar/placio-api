// Code generated by ent, DO NOT EDIT.

package help

import (
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Help {
	return predicate.Help(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Help {
	return predicate.Help(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Help {
	return predicate.Help(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Help {
	return predicate.Help(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Help {
	return predicate.Help(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Help {
	return predicate.Help(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Help {
	return predicate.Help(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Help {
	return predicate.Help(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Help {
	return predicate.Help(sql.FieldContainsFold(FieldID, id))
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v string) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldCategory, v))
}

// Subject applies equality check predicate on the "subject" field. It's identical to SubjectEQ.
func Subject(v string) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldSubject, v))
}

// Body applies equality check predicate on the "body" field. It's identical to BodyEQ.
func Body(v string) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldBody, v))
}

// Media applies equality check predicate on the "media" field. It's identical to MediaEQ.
func Media(v string) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldMedia, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldStatus, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldUserID, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v string) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v string) predicate.Help {
	return predicate.Help(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...string) predicate.Help {
	return predicate.Help(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...string) predicate.Help {
	return predicate.Help(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v string) predicate.Help {
	return predicate.Help(sql.FieldGT(FieldCategory, v))
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v string) predicate.Help {
	return predicate.Help(sql.FieldGTE(FieldCategory, v))
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v string) predicate.Help {
	return predicate.Help(sql.FieldLT(FieldCategory, v))
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v string) predicate.Help {
	return predicate.Help(sql.FieldLTE(FieldCategory, v))
}

// CategoryContains applies the Contains predicate on the "category" field.
func CategoryContains(v string) predicate.Help {
	return predicate.Help(sql.FieldContains(FieldCategory, v))
}

// CategoryHasPrefix applies the HasPrefix predicate on the "category" field.
func CategoryHasPrefix(v string) predicate.Help {
	return predicate.Help(sql.FieldHasPrefix(FieldCategory, v))
}

// CategoryHasSuffix applies the HasSuffix predicate on the "category" field.
func CategoryHasSuffix(v string) predicate.Help {
	return predicate.Help(sql.FieldHasSuffix(FieldCategory, v))
}

// CategoryEqualFold applies the EqualFold predicate on the "category" field.
func CategoryEqualFold(v string) predicate.Help {
	return predicate.Help(sql.FieldEqualFold(FieldCategory, v))
}

// CategoryContainsFold applies the ContainsFold predicate on the "category" field.
func CategoryContainsFold(v string) predicate.Help {
	return predicate.Help(sql.FieldContainsFold(FieldCategory, v))
}

// SubjectEQ applies the EQ predicate on the "subject" field.
func SubjectEQ(v string) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldSubject, v))
}

// SubjectNEQ applies the NEQ predicate on the "subject" field.
func SubjectNEQ(v string) predicate.Help {
	return predicate.Help(sql.FieldNEQ(FieldSubject, v))
}

// SubjectIn applies the In predicate on the "subject" field.
func SubjectIn(vs ...string) predicate.Help {
	return predicate.Help(sql.FieldIn(FieldSubject, vs...))
}

// SubjectNotIn applies the NotIn predicate on the "subject" field.
func SubjectNotIn(vs ...string) predicate.Help {
	return predicate.Help(sql.FieldNotIn(FieldSubject, vs...))
}

// SubjectGT applies the GT predicate on the "subject" field.
func SubjectGT(v string) predicate.Help {
	return predicate.Help(sql.FieldGT(FieldSubject, v))
}

// SubjectGTE applies the GTE predicate on the "subject" field.
func SubjectGTE(v string) predicate.Help {
	return predicate.Help(sql.FieldGTE(FieldSubject, v))
}

// SubjectLT applies the LT predicate on the "subject" field.
func SubjectLT(v string) predicate.Help {
	return predicate.Help(sql.FieldLT(FieldSubject, v))
}

// SubjectLTE applies the LTE predicate on the "subject" field.
func SubjectLTE(v string) predicate.Help {
	return predicate.Help(sql.FieldLTE(FieldSubject, v))
}

// SubjectContains applies the Contains predicate on the "subject" field.
func SubjectContains(v string) predicate.Help {
	return predicate.Help(sql.FieldContains(FieldSubject, v))
}

// SubjectHasPrefix applies the HasPrefix predicate on the "subject" field.
func SubjectHasPrefix(v string) predicate.Help {
	return predicate.Help(sql.FieldHasPrefix(FieldSubject, v))
}

// SubjectHasSuffix applies the HasSuffix predicate on the "subject" field.
func SubjectHasSuffix(v string) predicate.Help {
	return predicate.Help(sql.FieldHasSuffix(FieldSubject, v))
}

// SubjectEqualFold applies the EqualFold predicate on the "subject" field.
func SubjectEqualFold(v string) predicate.Help {
	return predicate.Help(sql.FieldEqualFold(FieldSubject, v))
}

// SubjectContainsFold applies the ContainsFold predicate on the "subject" field.
func SubjectContainsFold(v string) predicate.Help {
	return predicate.Help(sql.FieldContainsFold(FieldSubject, v))
}

// BodyEQ applies the EQ predicate on the "body" field.
func BodyEQ(v string) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldBody, v))
}

// BodyNEQ applies the NEQ predicate on the "body" field.
func BodyNEQ(v string) predicate.Help {
	return predicate.Help(sql.FieldNEQ(FieldBody, v))
}

// BodyIn applies the In predicate on the "body" field.
func BodyIn(vs ...string) predicate.Help {
	return predicate.Help(sql.FieldIn(FieldBody, vs...))
}

// BodyNotIn applies the NotIn predicate on the "body" field.
func BodyNotIn(vs ...string) predicate.Help {
	return predicate.Help(sql.FieldNotIn(FieldBody, vs...))
}

// BodyGT applies the GT predicate on the "body" field.
func BodyGT(v string) predicate.Help {
	return predicate.Help(sql.FieldGT(FieldBody, v))
}

// BodyGTE applies the GTE predicate on the "body" field.
func BodyGTE(v string) predicate.Help {
	return predicate.Help(sql.FieldGTE(FieldBody, v))
}

// BodyLT applies the LT predicate on the "body" field.
func BodyLT(v string) predicate.Help {
	return predicate.Help(sql.FieldLT(FieldBody, v))
}

// BodyLTE applies the LTE predicate on the "body" field.
func BodyLTE(v string) predicate.Help {
	return predicate.Help(sql.FieldLTE(FieldBody, v))
}

// BodyContains applies the Contains predicate on the "body" field.
func BodyContains(v string) predicate.Help {
	return predicate.Help(sql.FieldContains(FieldBody, v))
}

// BodyHasPrefix applies the HasPrefix predicate on the "body" field.
func BodyHasPrefix(v string) predicate.Help {
	return predicate.Help(sql.FieldHasPrefix(FieldBody, v))
}

// BodyHasSuffix applies the HasSuffix predicate on the "body" field.
func BodyHasSuffix(v string) predicate.Help {
	return predicate.Help(sql.FieldHasSuffix(FieldBody, v))
}

// BodyEqualFold applies the EqualFold predicate on the "body" field.
func BodyEqualFold(v string) predicate.Help {
	return predicate.Help(sql.FieldEqualFold(FieldBody, v))
}

// BodyContainsFold applies the ContainsFold predicate on the "body" field.
func BodyContainsFold(v string) predicate.Help {
	return predicate.Help(sql.FieldContainsFold(FieldBody, v))
}

// MediaEQ applies the EQ predicate on the "media" field.
func MediaEQ(v string) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldMedia, v))
}

// MediaNEQ applies the NEQ predicate on the "media" field.
func MediaNEQ(v string) predicate.Help {
	return predicate.Help(sql.FieldNEQ(FieldMedia, v))
}

// MediaIn applies the In predicate on the "media" field.
func MediaIn(vs ...string) predicate.Help {
	return predicate.Help(sql.FieldIn(FieldMedia, vs...))
}

// MediaNotIn applies the NotIn predicate on the "media" field.
func MediaNotIn(vs ...string) predicate.Help {
	return predicate.Help(sql.FieldNotIn(FieldMedia, vs...))
}

// MediaGT applies the GT predicate on the "media" field.
func MediaGT(v string) predicate.Help {
	return predicate.Help(sql.FieldGT(FieldMedia, v))
}

// MediaGTE applies the GTE predicate on the "media" field.
func MediaGTE(v string) predicate.Help {
	return predicate.Help(sql.FieldGTE(FieldMedia, v))
}

// MediaLT applies the LT predicate on the "media" field.
func MediaLT(v string) predicate.Help {
	return predicate.Help(sql.FieldLT(FieldMedia, v))
}

// MediaLTE applies the LTE predicate on the "media" field.
func MediaLTE(v string) predicate.Help {
	return predicate.Help(sql.FieldLTE(FieldMedia, v))
}

// MediaContains applies the Contains predicate on the "media" field.
func MediaContains(v string) predicate.Help {
	return predicate.Help(sql.FieldContains(FieldMedia, v))
}

// MediaHasPrefix applies the HasPrefix predicate on the "media" field.
func MediaHasPrefix(v string) predicate.Help {
	return predicate.Help(sql.FieldHasPrefix(FieldMedia, v))
}

// MediaHasSuffix applies the HasSuffix predicate on the "media" field.
func MediaHasSuffix(v string) predicate.Help {
	return predicate.Help(sql.FieldHasSuffix(FieldMedia, v))
}

// MediaIsNil applies the IsNil predicate on the "media" field.
func MediaIsNil() predicate.Help {
	return predicate.Help(sql.FieldIsNull(FieldMedia))
}

// MediaNotNil applies the NotNil predicate on the "media" field.
func MediaNotNil() predicate.Help {
	return predicate.Help(sql.FieldNotNull(FieldMedia))
}

// MediaEqualFold applies the EqualFold predicate on the "media" field.
func MediaEqualFold(v string) predicate.Help {
	return predicate.Help(sql.FieldEqualFold(FieldMedia, v))
}

// MediaContainsFold applies the ContainsFold predicate on the "media" field.
func MediaContainsFold(v string) predicate.Help {
	return predicate.Help(sql.FieldContainsFold(FieldMedia, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Help {
	return predicate.Help(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Help {
	return predicate.Help(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Help {
	return predicate.Help(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Help {
	return predicate.Help(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Help {
	return predicate.Help(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Help {
	return predicate.Help(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Help {
	return predicate.Help(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Help {
	return predicate.Help(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Help {
	return predicate.Help(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Help {
	return predicate.Help(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Help {
	return predicate.Help(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Help {
	return predicate.Help(sql.FieldContainsFold(FieldStatus, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Help {
	return predicate.Help(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Help {
	return predicate.Help(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Help {
	return predicate.Help(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Help {
	return predicate.Help(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Help {
	return predicate.Help(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Help {
	return predicate.Help(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Help {
	return predicate.Help(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Help {
	return predicate.Help(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Help {
	return predicate.Help(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Help {
	return predicate.Help(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Help {
	return predicate.Help(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Help {
	return predicate.Help(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Help {
	return predicate.Help(sql.FieldContainsFold(FieldUserID, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Help {
	return predicate.Help(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Help {
	return predicate.Help(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Help) predicate.Help {
	return predicate.Help(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Help) predicate.Help {
	return predicate.Help(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Help) predicate.Help {
	return predicate.Help(sql.NotPredicates(p))
}

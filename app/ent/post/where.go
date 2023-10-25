// Code generated by ent, DO NOT EDIT.

package post

import (
	"placio-app/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldID, id))
}

// Content applies equality check predicate on the "Content" field. It's identical to ContentEQ.
func Content(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldContent, v))
}

// CreatedAt applies equality check predicate on the "CreatedAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "UpdatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUpdatedAt, v))
}

// LikedByMe applies equality check predicate on the "LikedByMe" field. It's identical to LikedByMeEQ.
func LikedByMe(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldLikedByMe, v))
}

// LikeCount applies equality check predicate on the "LikeCount" field. It's identical to LikeCountEQ.
func LikeCount(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldLikeCount, v))
}

// CommentCount applies equality check predicate on the "CommentCount" field. It's identical to CommentCountEQ.
func CommentCount(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCommentCount, v))
}

// ShareCount applies equality check predicate on the "ShareCount" field. It's identical to ShareCountEQ.
func ShareCount(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldShareCount, v))
}

// ViewCount applies equality check predicate on the "ViewCount" field. It's identical to ViewCountEQ.
func ViewCount(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldViewCount, v))
}

// IsSponsored applies equality check predicate on the "IsSponsored" field. It's identical to IsSponsoredEQ.
func IsSponsored(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldIsSponsored, v))
}

// IsPromoted applies equality check predicate on the "IsPromoted" field. It's identical to IsPromotedEQ.
func IsPromoted(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldIsPromoted, v))
}

// IsBoosted applies equality check predicate on the "IsBoosted" field. It's identical to IsBoostedEQ.
func IsBoosted(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldIsBoosted, v))
}

// IsPinned applies equality check predicate on the "IsPinned" field. It's identical to IsPinnedEQ.
func IsPinned(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldIsPinned, v))
}

// IsHidden applies equality check predicate on the "IsHidden" field. It's identical to IsHiddenEQ.
func IsHidden(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldIsHidden, v))
}

// RelevanceScore applies equality check predicate on the "RelevanceScore" field. It's identical to RelevanceScoreEQ.
func RelevanceScore(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldRelevanceScore, v))
}

// SearchText applies equality check predicate on the "SearchText" field. It's identical to SearchTextEQ.
func SearchText(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldSearchText, v))
}

// ContentEQ applies the EQ predicate on the "Content" field.
func ContentEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "Content" field.
func ContentNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "Content" field.
func ContentIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "Content" field.
func ContentNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "Content" field.
func ContentGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "Content" field.
func ContentGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "Content" field.
func ContentLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "Content" field.
func ContentLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "Content" field.
func ContentContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "Content" field.
func ContentHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "Content" field.
func ContentHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "Content" field.
func ContentEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "Content" field.
func ContentContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldContent, v))
}

// CreatedAtEQ applies the EQ predicate on the "CreatedAt" field.
func CreatedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "CreatedAt" field.
func CreatedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "CreatedAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "CreatedAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "CreatedAt" field.
func CreatedAtGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "CreatedAt" field.
func CreatedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "CreatedAt" field.
func CreatedAtLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "CreatedAt" field.
func CreatedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "UpdatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "UpdatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "UpdatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "UpdatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "UpdatedAt" field.
func UpdatedAtGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "UpdatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "UpdatedAt" field.
func UpdatedAtLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "UpdatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldUpdatedAt, v))
}

// PrivacyEQ applies the EQ predicate on the "Privacy" field.
func PrivacyEQ(v Privacy) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldPrivacy, v))
}

// PrivacyNEQ applies the NEQ predicate on the "Privacy" field.
func PrivacyNEQ(v Privacy) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldPrivacy, v))
}

// PrivacyIn applies the In predicate on the "Privacy" field.
func PrivacyIn(vs ...Privacy) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldPrivacy, vs...))
}

// PrivacyNotIn applies the NotIn predicate on the "Privacy" field.
func PrivacyNotIn(vs ...Privacy) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldPrivacy, vs...))
}

// LikedByMeEQ applies the EQ predicate on the "LikedByMe" field.
func LikedByMeEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldLikedByMe, v))
}

// LikedByMeNEQ applies the NEQ predicate on the "LikedByMe" field.
func LikedByMeNEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldLikedByMe, v))
}

// LikeCountEQ applies the EQ predicate on the "LikeCount" field.
func LikeCountEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldLikeCount, v))
}

// LikeCountNEQ applies the NEQ predicate on the "LikeCount" field.
func LikeCountNEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldLikeCount, v))
}

// LikeCountIn applies the In predicate on the "LikeCount" field.
func LikeCountIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldLikeCount, vs...))
}

// LikeCountNotIn applies the NotIn predicate on the "LikeCount" field.
func LikeCountNotIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldLikeCount, vs...))
}

// LikeCountGT applies the GT predicate on the "LikeCount" field.
func LikeCountGT(v int) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldLikeCount, v))
}

// LikeCountGTE applies the GTE predicate on the "LikeCount" field.
func LikeCountGTE(v int) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldLikeCount, v))
}

// LikeCountLT applies the LT predicate on the "LikeCount" field.
func LikeCountLT(v int) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldLikeCount, v))
}

// LikeCountLTE applies the LTE predicate on the "LikeCount" field.
func LikeCountLTE(v int) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldLikeCount, v))
}

// CommentCountEQ applies the EQ predicate on the "CommentCount" field.
func CommentCountEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCommentCount, v))
}

// CommentCountNEQ applies the NEQ predicate on the "CommentCount" field.
func CommentCountNEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldCommentCount, v))
}

// CommentCountIn applies the In predicate on the "CommentCount" field.
func CommentCountIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldCommentCount, vs...))
}

// CommentCountNotIn applies the NotIn predicate on the "CommentCount" field.
func CommentCountNotIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldCommentCount, vs...))
}

// CommentCountGT applies the GT predicate on the "CommentCount" field.
func CommentCountGT(v int) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldCommentCount, v))
}

// CommentCountGTE applies the GTE predicate on the "CommentCount" field.
func CommentCountGTE(v int) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldCommentCount, v))
}

// CommentCountLT applies the LT predicate on the "CommentCount" field.
func CommentCountLT(v int) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldCommentCount, v))
}

// CommentCountLTE applies the LTE predicate on the "CommentCount" field.
func CommentCountLTE(v int) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldCommentCount, v))
}

// ShareCountEQ applies the EQ predicate on the "ShareCount" field.
func ShareCountEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldShareCount, v))
}

// ShareCountNEQ applies the NEQ predicate on the "ShareCount" field.
func ShareCountNEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldShareCount, v))
}

// ShareCountIn applies the In predicate on the "ShareCount" field.
func ShareCountIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldShareCount, vs...))
}

// ShareCountNotIn applies the NotIn predicate on the "ShareCount" field.
func ShareCountNotIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldShareCount, vs...))
}

// ShareCountGT applies the GT predicate on the "ShareCount" field.
func ShareCountGT(v int) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldShareCount, v))
}

// ShareCountGTE applies the GTE predicate on the "ShareCount" field.
func ShareCountGTE(v int) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldShareCount, v))
}

// ShareCountLT applies the LT predicate on the "ShareCount" field.
func ShareCountLT(v int) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldShareCount, v))
}

// ShareCountLTE applies the LTE predicate on the "ShareCount" field.
func ShareCountLTE(v int) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldShareCount, v))
}

// ViewCountEQ applies the EQ predicate on the "ViewCount" field.
func ViewCountEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldViewCount, v))
}

// ViewCountNEQ applies the NEQ predicate on the "ViewCount" field.
func ViewCountNEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldViewCount, v))
}

// ViewCountIn applies the In predicate on the "ViewCount" field.
func ViewCountIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldViewCount, vs...))
}

// ViewCountNotIn applies the NotIn predicate on the "ViewCount" field.
func ViewCountNotIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldViewCount, vs...))
}

// ViewCountGT applies the GT predicate on the "ViewCount" field.
func ViewCountGT(v int) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldViewCount, v))
}

// ViewCountGTE applies the GTE predicate on the "ViewCount" field.
func ViewCountGTE(v int) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldViewCount, v))
}

// ViewCountLT applies the LT predicate on the "ViewCount" field.
func ViewCountLT(v int) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldViewCount, v))
}

// ViewCountLTE applies the LTE predicate on the "ViewCount" field.
func ViewCountLTE(v int) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldViewCount, v))
}

// IsSponsoredEQ applies the EQ predicate on the "IsSponsored" field.
func IsSponsoredEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldIsSponsored, v))
}

// IsSponsoredNEQ applies the NEQ predicate on the "IsSponsored" field.
func IsSponsoredNEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldIsSponsored, v))
}

// IsPromotedEQ applies the EQ predicate on the "IsPromoted" field.
func IsPromotedEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldIsPromoted, v))
}

// IsPromotedNEQ applies the NEQ predicate on the "IsPromoted" field.
func IsPromotedNEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldIsPromoted, v))
}

// IsBoostedEQ applies the EQ predicate on the "IsBoosted" field.
func IsBoostedEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldIsBoosted, v))
}

// IsBoostedNEQ applies the NEQ predicate on the "IsBoosted" field.
func IsBoostedNEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldIsBoosted, v))
}

// IsPinnedEQ applies the EQ predicate on the "IsPinned" field.
func IsPinnedEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldIsPinned, v))
}

// IsPinnedNEQ applies the NEQ predicate on the "IsPinned" field.
func IsPinnedNEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldIsPinned, v))
}

// IsHiddenEQ applies the EQ predicate on the "IsHidden" field.
func IsHiddenEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldIsHidden, v))
}

// IsHiddenNEQ applies the NEQ predicate on the "IsHidden" field.
func IsHiddenNEQ(v bool) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldIsHidden, v))
}

// RelevanceScoreEQ applies the EQ predicate on the "RelevanceScore" field.
func RelevanceScoreEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldRelevanceScore, v))
}

// RelevanceScoreNEQ applies the NEQ predicate on the "RelevanceScore" field.
func RelevanceScoreNEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldRelevanceScore, v))
}

// RelevanceScoreIn applies the In predicate on the "RelevanceScore" field.
func RelevanceScoreIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldRelevanceScore, vs...))
}

// RelevanceScoreNotIn applies the NotIn predicate on the "RelevanceScore" field.
func RelevanceScoreNotIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldRelevanceScore, vs...))
}

// RelevanceScoreGT applies the GT predicate on the "RelevanceScore" field.
func RelevanceScoreGT(v int) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldRelevanceScore, v))
}

// RelevanceScoreGTE applies the GTE predicate on the "RelevanceScore" field.
func RelevanceScoreGTE(v int) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldRelevanceScore, v))
}

// RelevanceScoreLT applies the LT predicate on the "RelevanceScore" field.
func RelevanceScoreLT(v int) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldRelevanceScore, v))
}

// RelevanceScoreLTE applies the LTE predicate on the "RelevanceScore" field.
func RelevanceScoreLTE(v int) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldRelevanceScore, v))
}

// SearchTextEQ applies the EQ predicate on the "SearchText" field.
func SearchTextEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldSearchText, v))
}

// SearchTextNEQ applies the NEQ predicate on the "SearchText" field.
func SearchTextNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldSearchText, v))
}

// SearchTextIn applies the In predicate on the "SearchText" field.
func SearchTextIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldSearchText, vs...))
}

// SearchTextNotIn applies the NotIn predicate on the "SearchText" field.
func SearchTextNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldSearchText, vs...))
}

// SearchTextGT applies the GT predicate on the "SearchText" field.
func SearchTextGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldSearchText, v))
}

// SearchTextGTE applies the GTE predicate on the "SearchText" field.
func SearchTextGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldSearchText, v))
}

// SearchTextLT applies the LT predicate on the "SearchText" field.
func SearchTextLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldSearchText, v))
}

// SearchTextLTE applies the LTE predicate on the "SearchText" field.
func SearchTextLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldSearchText, v))
}

// SearchTextContains applies the Contains predicate on the "SearchText" field.
func SearchTextContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldSearchText, v))
}

// SearchTextHasPrefix applies the HasPrefix predicate on the "SearchText" field.
func SearchTextHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldSearchText, v))
}

// SearchTextHasSuffix applies the HasSuffix predicate on the "SearchText" field.
func SearchTextHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldSearchText, v))
}

// SearchTextIsNil applies the IsNil predicate on the "SearchText" field.
func SearchTextIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldSearchText))
}

// SearchTextNotNil applies the NotNil predicate on the "SearchText" field.
func SearchTextNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldSearchText))
}

// SearchTextEqualFold applies the EqualFold predicate on the "SearchText" field.
func SearchTextEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldSearchText, v))
}

// SearchTextContainsFold applies the ContainsFold predicate on the "SearchText" field.
func SearchTextContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldSearchText, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBusinessAccount applies the HasEdge predicate on the "business_account" edge.
func HasBusinessAccount() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BusinessAccountTable, BusinessAccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessAccountWith applies the HasEdge predicate on the "business_account" edge with a given conditions (other predicates).
func HasBusinessAccountWith(preds ...predicate.Business) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := newBusinessAccountStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMedias applies the HasEdge predicate on the "medias" edge.
func HasMedias() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MediasTable, MediasColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMediasWith applies the HasEdge predicate on the "medias" edge with a given conditions (other predicates).
func HasMediasWith(preds ...predicate.Media) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := newMediasStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComments applies the HasEdge predicate on the "comments" edge.
func HasComments() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentsWith applies the HasEdge predicate on the "comments" edge with a given conditions (other predicates).
func HasCommentsWith(preds ...predicate.Comment) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := newCommentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLikes applies the HasEdge predicate on the "likes" edge.
func HasLikes() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LikesTable, LikesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLikesWith applies the HasEdge predicate on the "likes" edge with a given conditions (other predicates).
func HasLikesWith(preds ...predicate.Like) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := newLikesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCategories applies the HasEdge predicate on the "categories" edge.
func HasCategories() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CategoriesTable, CategoriesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoriesWith applies the HasEdge predicate on the "categories" edge with a given conditions (other predicates).
func HasCategoriesWith(preds ...predicate.Category) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := newCategoriesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNotifications applies the HasEdge predicate on the "notifications" edge.
func HasNotifications() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, NotificationsTable, NotificationsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotificationsWith applies the HasEdge predicate on the "notifications" edge with a given conditions (other predicates).
func HasNotificationsWith(preds ...predicate.Notification) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := newNotificationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Post) predicate.Post {
	return predicate.Post(sql.NotPredicates(p))
}

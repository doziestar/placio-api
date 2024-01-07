// Code generated by ent, DO NOT EDIT.

package media

import (
	"placio-app/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Media {
	return predicate.Media(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Media {
	return predicate.Media(sql.FieldContainsFold(FieldID, id))
}

// URL applies equality check predicate on the "URL" field. It's identical to URLEQ.
func URL(v string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldURL, v))
}

// MediaType applies equality check predicate on the "MediaType" field. It's identical to MediaTypeEQ.
func MediaType(v string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldMediaType, v))
}

// CreatedAt applies equality check predicate on the "CreatedAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "UpdatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldUpdatedAt, v))
}

// LikeCount applies equality check predicate on the "likeCount" field. It's identical to LikeCountEQ.
func LikeCount(v int) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldLikeCount, v))
}

// DislikeCount applies equality check predicate on the "dislikeCount" field. It's identical to DislikeCountEQ.
func DislikeCount(v int) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldDislikeCount, v))
}

// URLEQ applies the EQ predicate on the "URL" field.
func URLEQ(v string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "URL" field.
func URLNEQ(v string) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "URL" field.
func URLIn(vs ...string) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "URL" field.
func URLNotIn(vs ...string) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "URL" field.
func URLGT(v string) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "URL" field.
func URLGTE(v string) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "URL" field.
func URLLT(v string) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "URL" field.
func URLLTE(v string) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "URL" field.
func URLContains(v string) predicate.Media {
	return predicate.Media(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "URL" field.
func URLHasPrefix(v string) predicate.Media {
	return predicate.Media(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "URL" field.
func URLHasSuffix(v string) predicate.Media {
	return predicate.Media(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "URL" field.
func URLEqualFold(v string) predicate.Media {
	return predicate.Media(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "URL" field.
func URLContainsFold(v string) predicate.Media {
	return predicate.Media(sql.FieldContainsFold(FieldURL, v))
}

// MediaTypeEQ applies the EQ predicate on the "MediaType" field.
func MediaTypeEQ(v string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldMediaType, v))
}

// MediaTypeNEQ applies the NEQ predicate on the "MediaType" field.
func MediaTypeNEQ(v string) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldMediaType, v))
}

// MediaTypeIn applies the In predicate on the "MediaType" field.
func MediaTypeIn(vs ...string) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldMediaType, vs...))
}

// MediaTypeNotIn applies the NotIn predicate on the "MediaType" field.
func MediaTypeNotIn(vs ...string) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldMediaType, vs...))
}

// MediaTypeGT applies the GT predicate on the "MediaType" field.
func MediaTypeGT(v string) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldMediaType, v))
}

// MediaTypeGTE applies the GTE predicate on the "MediaType" field.
func MediaTypeGTE(v string) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldMediaType, v))
}

// MediaTypeLT applies the LT predicate on the "MediaType" field.
func MediaTypeLT(v string) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldMediaType, v))
}

// MediaTypeLTE applies the LTE predicate on the "MediaType" field.
func MediaTypeLTE(v string) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldMediaType, v))
}

// MediaTypeContains applies the Contains predicate on the "MediaType" field.
func MediaTypeContains(v string) predicate.Media {
	return predicate.Media(sql.FieldContains(FieldMediaType, v))
}

// MediaTypeHasPrefix applies the HasPrefix predicate on the "MediaType" field.
func MediaTypeHasPrefix(v string) predicate.Media {
	return predicate.Media(sql.FieldHasPrefix(FieldMediaType, v))
}

// MediaTypeHasSuffix applies the HasSuffix predicate on the "MediaType" field.
func MediaTypeHasSuffix(v string) predicate.Media {
	return predicate.Media(sql.FieldHasSuffix(FieldMediaType, v))
}

// MediaTypeEqualFold applies the EqualFold predicate on the "MediaType" field.
func MediaTypeEqualFold(v string) predicate.Media {
	return predicate.Media(sql.FieldEqualFold(FieldMediaType, v))
}

// MediaTypeContainsFold applies the ContainsFold predicate on the "MediaType" field.
func MediaTypeContainsFold(v string) predicate.Media {
	return predicate.Media(sql.FieldContainsFold(FieldMediaType, v))
}

// CreatedAtEQ applies the EQ predicate on the "CreatedAt" field.
func CreatedAtEQ(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "CreatedAt" field.
func CreatedAtNEQ(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "CreatedAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "CreatedAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "CreatedAt" field.
func CreatedAtGT(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "CreatedAt" field.
func CreatedAtGTE(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "CreatedAt" field.
func CreatedAtLT(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "CreatedAt" field.
func CreatedAtLTE(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "UpdatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "UpdatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "UpdatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "UpdatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "UpdatedAt" field.
func UpdatedAtGT(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "UpdatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "UpdatedAt" field.
func UpdatedAtLT(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "UpdatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldUpdatedAt, v))
}

// LikeCountEQ applies the EQ predicate on the "likeCount" field.
func LikeCountEQ(v int) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldLikeCount, v))
}

// LikeCountNEQ applies the NEQ predicate on the "likeCount" field.
func LikeCountNEQ(v int) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldLikeCount, v))
}

// LikeCountIn applies the In predicate on the "likeCount" field.
func LikeCountIn(vs ...int) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldLikeCount, vs...))
}

// LikeCountNotIn applies the NotIn predicate on the "likeCount" field.
func LikeCountNotIn(vs ...int) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldLikeCount, vs...))
}

// LikeCountGT applies the GT predicate on the "likeCount" field.
func LikeCountGT(v int) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldLikeCount, v))
}

// LikeCountGTE applies the GTE predicate on the "likeCount" field.
func LikeCountGTE(v int) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldLikeCount, v))
}

// LikeCountLT applies the LT predicate on the "likeCount" field.
func LikeCountLT(v int) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldLikeCount, v))
}

// LikeCountLTE applies the LTE predicate on the "likeCount" field.
func LikeCountLTE(v int) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldLikeCount, v))
}

// DislikeCountEQ applies the EQ predicate on the "dislikeCount" field.
func DislikeCountEQ(v int) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldDislikeCount, v))
}

// DislikeCountNEQ applies the NEQ predicate on the "dislikeCount" field.
func DislikeCountNEQ(v int) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldDislikeCount, v))
}

// DislikeCountIn applies the In predicate on the "dislikeCount" field.
func DislikeCountIn(vs ...int) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldDislikeCount, vs...))
}

// DislikeCountNotIn applies the NotIn predicate on the "dislikeCount" field.
func DislikeCountNotIn(vs ...int) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldDislikeCount, vs...))
}

// DislikeCountGT applies the GT predicate on the "dislikeCount" field.
func DislikeCountGT(v int) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldDislikeCount, v))
}

// DislikeCountGTE applies the GTE predicate on the "dislikeCount" field.
func DislikeCountGTE(v int) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldDislikeCount, v))
}

// DislikeCountLT applies the LT predicate on the "dislikeCount" field.
func DislikeCountLT(v int) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldDislikeCount, v))
}

// DislikeCountLTE applies the LTE predicate on the "dislikeCount" field.
func DislikeCountLTE(v int) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldDislikeCount, v))
}

// HasPost applies the HasEdge predicate on the "post" edge.
func HasPost() predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PostTable, PostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostWith applies the HasEdge predicate on the "post" edge with a given conditions (other predicates).
func HasPostWith(preds ...predicate.Post) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := newPostStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReview applies the HasEdge predicate on the "review" edge.
func HasReview() predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ReviewTable, ReviewColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReviewWith applies the HasEdge predicate on the "review" edge with a given conditions (other predicates).
func HasReviewWith(preds ...predicate.Review) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := newReviewStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCategories applies the HasEdge predicate on the "categories" edge.
func HasCategories() predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CategoriesTable, CategoriesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoriesWith applies the HasEdge predicate on the "categories" edge with a given conditions (other predicates).
func HasCategoriesWith(preds ...predicate.Category) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := newCategoriesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlace applies the HasEdge predicate on the "place" edge.
func HasPlace() predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PlaceTable, PlacePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlaceWith applies the HasEdge predicate on the "place" edge with a given conditions (other predicates).
func HasPlaceWith(preds ...predicate.Place) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := newPlaceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlaceInventory applies the HasEdge predicate on the "place_inventory" edge.
func HasPlaceInventory() predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PlaceInventoryTable, PlaceInventoryPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlaceInventoryWith applies the HasEdge predicate on the "place_inventory" edge with a given conditions (other predicates).
func HasPlaceInventoryWith(preds ...predicate.PlaceInventory) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := newPlaceInventoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMenu applies the HasEdge predicate on the "menu" edge.
func HasMenu() predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, MenuTable, MenuPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMenuWith applies the HasEdge predicate on the "menu" edge with a given conditions (other predicates).
func HasMenuWith(preds ...predicate.Menu) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := newMenuStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRoomCategory applies the HasEdge predicate on the "room_category" edge.
func HasRoomCategory() predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RoomCategoryTable, RoomCategoryPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomCategoryWith applies the HasEdge predicate on the "room_category" edge with a given conditions (other predicates).
func HasRoomCategoryWith(preds ...predicate.RoomCategory) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := newRoomCategoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRoom applies the HasEdge predicate on the "room" edge.
func HasRoom() predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RoomTable, RoomPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomWith applies the HasEdge predicate on the "room" edge with a given conditions (other predicates).
func HasRoomWith(preds ...predicate.Room) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := newRoomStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlan applies the HasEdge predicate on the "plan" edge.
func HasPlan() predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlanTable, PlanColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlanWith applies the HasEdge predicate on the "plan" edge with a given conditions (other predicates).
func HasPlanWith(preds ...predicate.Plan) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		step := newPlanStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Media) predicate.Media {
	return predicate.Media(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Media) predicate.Media {
	return predicate.Media(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Media) predicate.Media {
	return predicate.Media(sql.NotPredicates(p))
}

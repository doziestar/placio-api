package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"time"
)

type Review struct {
	ent.Schema
}

func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.Float("score").
			Min(1).
			Max(5).
			Comment("Score should be between 1 and 5."),
		field.String("content").
			Comment("User's review to the business/place/event."),
		field.Time("createdAt").
			Default(time.Now).
			Comment("When the review was created."),
		field.Int("likeCount").
			Default(0).
			Comment("Count of likes for this review."),
		field.Int("dislikeCount").
			Default(0).
			Comment("Count of dislikes for this review."),
		field.String("flag").
			Default("").
			Comment("Flag for this review."),
	}
}

func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("reviews").
			Unique().
			Required().
			Comment("The user who wrote the review."),
		edge.To("business", Business.Type).
			Unique().
			Comment("The business that was reviewed."),
		edge.To("place", Place.Type).
			Unique().
			Comment("The place that was reviewed."),
		edge.To("event", Event.Type).
			Unique().
			Comment("The event that was reviewed."),
		edge.To("medias", Media.Type).
			Comment("The media content related to the review."),
		edge.To("comments", Comment.Type).
			Comment("The comments related to the review."),
		edge.To("likes", Like.Type).
			Comment("The likes related to the review."),
	}
}

//func (Review) Hooks() []ent.Hook {
//	return []ent.Hook{
//		hook.On(func(next ent.Mutator) ent.Mutator {
//			return hook.ReviewFunc(func(ctx context.Context, m *gen.ReviewMutation) (ent.Value, error) {
//				// check if the operation is not update
//				if !m.Op().Is(ent.OpUpdateOne) {
//					// check if the id is not provided
//					id, ok := m.ID()
//					if !ok || id == "" {
//						m.SetID(uuid.New().String())
//					}
//				}
//				return next.Mutate(ctx, m)
//			})
//		}, ent.OpCreate),
//	}
//}

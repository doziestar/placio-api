package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"time"
)

type Rating struct {
	ent.Schema
}

func (Rating) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.Int("score").
			Min(1).
			Max(5).
			Comment("Score should be between 1 and 5."),
		field.String("review").
			Optional().
			Comment("User's review to the business/place/event."),
		field.Time("ratedAt").
			Default(time.Now).
			Comment("When the rating was given."),
	}
}

func (Rating) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("ratings").
			Unique().
			Required().
			Comment("The user who gave the rating."),
		edge.To("business", Business.Type).
			Unique().
			Comment("The business that was rated."),
		edge.To("place", Place.Type).
			Unique().
			Comment("The place that was rated."),
		edge.To("event", Event.Type).
			Unique().
			Comment("The event that was rated."),
	}
}

//func (Rating) Hooks() []ent.Hook {
//	return []ent.Hook{
//		hook.On(
//			func(next ent.Mutator) ent.Mutator {
//				return hook.RatingFunc(func(ctx context.Context, m *gen.RatingMutation) (ent.Value, error) {
//					if !m.Op().Is(ent.OpCreate) {
//						return next.Mutate(ctx, m)
//					}
//
//					uid, err := uuid.NewRandom()
//					if err != nil {
//						return nil, fmt.Errorf("failed to generate uuid: %w", err)
//					}
//
//					m.SetID(uid.String())
//					return next.Mutate(ctx, m)
//				})
//			},
//			ent.OpCreate,
//		),
//	}
//}

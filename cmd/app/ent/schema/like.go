package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Like holds the schema definition for the Like entity.
type Like struct {
	ent.Schema
}

func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").UpdateDefault(time.Now),
		field.Bool("like").Comment("True for like, False for dislike"),
	}
}

func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("likes").
			Unique().
			Comment("The user who liked/disliked the review/media."),
		edge.To("review", Review.Type).Unique().
			Comment("The review that was liked/disliked."),
		edge.To("media", Media.Type).Unique().
			Comment("The media content that was liked/disliked."),
		edge.To("post", Post.Type).Unique(),
	}
}

//func (Like) Hooks() []ent.Hook {
//	return []ent.Hook{
//		hook.On(
//			func(next ent.Mutator) ent.Mutator {
//				return hook.LikeFunc(func(ctx context.Context, m *gen.LikeMutation) (ent.Value, error) {
//					if m.Op().Is(ent.OpCreate) {
//						reviewID, exists := m.ReviewID()
//						if !exists {
//							return nil, errors.New("review not associated with the like/dislike")
//						}
//						// Fetch the associated review
//						review, err := m.Client().Review.Get(ctx, reviewID)
//						if err != nil {
//							return nil, fmt.Errorf("failed to get review: %w", err)
//						}
//
//						// Determine whether this is a like or a dislike
//						isLike, exists := m.Like()
//						if !exists {
//							return nil, errors.New("like/dislike not specified")
//						}
//
//						// Increment the appropriate count
//						if isLike {
//							review.Update().AddLikeCount(1).Save(ctx)
//						} else {
//							review.Update().AddDislikeCount(1).Save(ctx)
//						}
//					}
//					return next.Mutate(ctx, m)
//				})
//			},
//			ent.OpCreate,
//		),
//	}
//}

type UserLikePlace struct {
	ent.Schema
}

func (UserLikePlace) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").UpdateDefault(time.Now),
	}
}

func (UserLikePlace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("likedPlaces").Unique(),
		edge.To("place", Place.Type).Unique(),
	}
}

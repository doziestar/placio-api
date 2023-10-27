package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("Content").MaxLen(2147483647).Optional(),
		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").UpdateDefault(time.Now),
		field.Enum("Privacy").Values("Public", "FollowersOnly", "OnlyMe").Default("Public"),
		field.Bool("LikedByMe").Default(false),
		field.Int("LikeCount").Default(0),
		field.Int("CommentCount").Default(0),
		field.Int("ShareCount").Default(0),
		field.Int("ViewCount").Default(0),
		field.Bool("IsSponsored").Default(false),
		field.Bool("IsPromoted").Default(false),
		field.Bool("IsBoosted").Default(false),
		field.Bool("IsPinned").Default(false),
		field.Bool("IsHidden").Default(false),
		field.Int("RepostCount").Default(0),
		field.Bool("IsRepost").Default(false),
		field.Int("RelevanceScore").Default(0),
		field.String("SearchText").Optional(),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("posts").
			Unique(),
		edge.From("business_account", Business.Type).
			Ref("posts").
			Unique(),
		edge.To("medias", Media.Type),
		edge.To("comments", Comment.Type),
		edge.To("likes", Like.Type),
		edge.To("categories", Category.Type),
		edge.To("notifications", Notification.Type),
		edge.To("original_post", Post.Type).
			From("reposts").
			Unique(),
	}
}

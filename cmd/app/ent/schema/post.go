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
		field.String("Content").MaxLen(2147483647), // equivalent to TEXT in SQL
		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").UpdateDefault(time.Now),
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
	}
}

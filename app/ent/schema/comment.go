package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("Content").MaxLen(2147483647),
		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").UpdateDefault(time.Now),
		field.String("parentCommentID").
			Optional().
			Nillable(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("comments").
			Unique(),
		edge.From("post", Post.Type).
			Ref("comments").
			Unique(),
		edge.To("replies", Comment.Type).
			From("parentComment").
			Field("parentCommentID").
			Unique(),
	}
}

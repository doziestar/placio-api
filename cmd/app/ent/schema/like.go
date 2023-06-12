package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Like holds the schema definition for the Like entity.
type Like struct {
	ent.Schema
}

// Fields of the Like.
func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.String("LikeID").Unique(),
		field.String("UserId"),
		field.String("PostID"),
		field.Time("CreatedAt"),
		field.Time("UpdatedAt"),
	}
}

// Edges of the Like.
func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("likes").
			Field("UserId").
			Unique(),
		edge.From("post", Post.Type).
			Ref("likes").
			Field("PostID").
			Unique(),
	}
}

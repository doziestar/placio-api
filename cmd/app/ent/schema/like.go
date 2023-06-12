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
		field.Int("UserID"),
		field.Int("PostID"),
		field.Time("CreatedAt"),
		field.Time("UpdatedAt"),
	}
}

// Edges of the Like.
// Edges of the Like.
func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Field("UserID").
			Unique(),
		edge.To("post", Post.Type).
			Field("PostID").
			Unique(),
	}
}

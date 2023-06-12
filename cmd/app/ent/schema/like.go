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

// Fields of the Like.
func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").UpdateDefault(time.Now),
	}
}

// Edges of the Like.
func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("likes").
			Unique(),
		edge.To("post", Post.Type).Unique(),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Review struct {
	ent.Schema
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.Float("rating"),
		field.String("comment").Optional(),
		field.JSON("images_videos", []string{}).Optional(),
		field.Time("timestamp"),
	}
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("reviews").
			Unique(),
		edge.From("place", Place.Type).
			Ref("reviews").
			Unique(),
	}
}

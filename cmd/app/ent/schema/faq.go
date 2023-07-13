package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type FAQ struct {
	ent.Schema
}

// Fields of the FAQ.
func (FAQ) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("question"),
		field.String("answer"),
	}
}

// Edges of the FAQ.
func (FAQ) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("business", Business.Type).
			Ref("faqs").
			Unique(),
		edge.To("place", Place.Type),
	}
}

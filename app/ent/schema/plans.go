package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Plan struct {
	ent.Schema
}

// Fields of the Plan.
func (Plan) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("name").Optional(),
		field.String("description").Optional(),
		field.String("overview").Optional(),
		field.JSON("features", []string{}).Optional(),
	}
}

// Edges of the Plan.
func (Plan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
		edge.To("businesses", Business.Type),
		edge.To("places", Place.Type),
		edge.To("media", Media.Type),
	}
}

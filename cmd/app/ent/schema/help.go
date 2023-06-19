package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Help struct {
	ent.Schema
}

// Fields of the Help.
func (Help) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.String("category"),
		field.String("subject"),
		field.Text("body"),
		field.String("media").Optional(),
		field.String("status").Default("open"),
		field.String("user_id").
			MaxLen(36).
			Immutable(),
	}
}

// Edges of the Help.
func (Help) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("helps").
			Unique().
			Required().
			Field("user_id").Immutable(),
	}
}

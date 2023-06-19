package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("name").Unique(),
		field.String("image").Optional(),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return nil
}

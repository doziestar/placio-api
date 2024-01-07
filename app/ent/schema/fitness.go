package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Fitness struct {
	ent.Schema
}

// Fields of the Fitness.
func (Fitness) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("name"),
	}
}

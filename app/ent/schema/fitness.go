package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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

type Trainer struct {
	ent.Schema
}

// Fields of the Trainer.
func (Trainer) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("name"),
		field.String("email"),
		field.String("phone"),
	}
}

func (Trainer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("trainers"),
		edge.From("place", Place.Type).Ref("trainers"),
	}
}

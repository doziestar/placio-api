package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Staff struct {
	ent.Schema
}

// Fields of the Staff.
func (Staff) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("position").Optional(),
		// Add other staff-specific fields if necessary
	}
}

// Edges of the Staff.
func (Staff) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("staffs").Unique(),
		edge.From("place", Place.Type).Ref("staffs"),
		edge.To("permissions", Permission.Type),
		edge.From("business", Business.Type).Ref("staffs"),
	}
}

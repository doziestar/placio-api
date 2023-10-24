package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Table struct {
	ent.Schema
}

// Fields of the Table.
func (Table) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.Int("number"),
		field.String("qr_code").Optional(),
	}
}

// Edges of the Table.
func (Table) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("place", Place.Type).
			Ref("tables").
			Unique(),
		edge.To("orders", Order.Type),
	}
}

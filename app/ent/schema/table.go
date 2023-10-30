package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PlaceTable struct {
	ent.Schema
}

// Fields of the Table.
func (PlaceTable) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.Int("number"),
		field.String("deleted_at").Optional(),
		field.Bool("is_deleted").Default(false),
		field.String("qr_code").Optional(),
		field.String("description").Optional(),
	}
}

// Edges of the Table.
func (PlaceTable) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("place", Place.Type).
			Ref("tables").
			Unique(),
		edge.To("orders", Order.Type),
	}
}

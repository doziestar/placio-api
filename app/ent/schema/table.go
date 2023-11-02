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
		field.String("status").Default("available"),
		field.String("type").Default("table"),
		field.Bool("is_active").Default(true),
		field.Bool("is_reserved").Default(false),
		field.Bool("is_vip").Default(false),
		field.Bool("is_premium").Default(false),
	}
}

// Edges of the Table.
func (PlaceTable) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("place", Place.Type).
			Ref("tables").
			Unique(),
		edge.From("created_by", User.Type).
			Ref("tables_created").
			Unique(),
		edge.From("updated_by", User.Type).
			Ref("tables_updated").
			Unique(),
		edge.From("deleted_by", User.Type).
			Ref("tables_deleted").
			Unique(),
		edge.From("reserved_by", User.Type).
			Ref("tables_reserved").
			Unique(),
		edge.From("waiter", User.Type).
			Ref("tables_waited").
			Unique(),
		edge.To("orders", Order.Type),
	}
}

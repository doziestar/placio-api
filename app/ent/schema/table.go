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
		field.String("name").Optional(),
		field.Int("capacity").Default(4),
		field.String("deleted_at").Optional(),
		field.Bool("is_deleted").Default(false),
		field.String("qr_code").Optional(),
		field.String("description").Optional(),
		field.String("status").Default("available"),
		field.Enum("type").Optional().Values("regular", "vip", "premium"),
		field.Bool("is_active").Default(true),
		field.Bool("is_reserved").Default(false),
		field.Bool("is_vip").Default(false),
		field.Bool("is_premium").Default(false),
		field.String("location_description").Optional(),
		field.Float("minimum_spend").Optional(),
		field.Time("reservation_time").Optional().Nillable(),
		field.Time("next_available_time").Optional().Nillable(),
		field.JSON("special_requirements", []string{}).Optional(),
		field.String("layout").Optional(),
		field.String("service_area").Optional(),
		field.String("ambient").Optional(),
		field.String("image_url").Optional(),
		field.Float("rating").Optional().Nillable(),
		field.JSON("tags", []string{}).Optional(),
		field.JSON("metadata", map[string]interface{}{}).Optional(),
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

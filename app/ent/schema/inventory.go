package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type InventoryType struct {
	ent.Schema
}

// Fields of the InventoryType.
func (InventoryType) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("name").
			Unique(),
		field.Enum("industry_type").Values("hotel", "restaurant", "bar", "club", "gym"),
		field.String("measurement_unit").Optional(),
	}
}

func (InventoryType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attributes", InventoryAttribute.Type),
		edge.To("place_inventories", PlaceInventory.Type),
		// ... other edges
	}
}

type InventoryAttribute struct {
	ent.Schema
}

func (InventoryAttribute) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("name"),
		field.Bool("is_mandatory").Default(false),
		field.Enum("data_type").Values("string", "number", "boolean", "date", "enum").Optional(),
	}
}

func (InventoryAttribute) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("inventory_type", InventoryType.Type).
			Ref("attributes").
			Unique(),
		edge.To("place_inventory_attributes", PlaceInventoryAttribute.Type),
	}
}

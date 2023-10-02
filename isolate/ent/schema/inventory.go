package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"placio-app/domains/inventory"
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
		field.String("description").Optional(),
		field.Enum("industry_type").Values(
			string(inventory.IndustryHotel),
			string(inventory.IndustryRestaurant),
			string(inventory.IndustryBar),
			string(inventory.IndustryClub),
			string(inventory.IndustryGym),
			string(inventory.IndustryEvents),
			string(inventory.IndustryRetail),
			string(inventory.IndustryOther)).Optional(),
		field.String("measurement_unit").Optional(),
	}
}

func (InventoryType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attributes", InventoryAttribute.Type),
		edge.To("place_inventories", PlaceInventory.Type),
		//edge.From("business", Business.Type).Ref("inventory_types").Unique(),
		//edge.From("category", Category.Type).Ref("inventory_types").Unique(),
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

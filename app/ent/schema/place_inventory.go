package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type PlaceInventory struct {
	ent.Schema
}

func (PlaceInventory) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("name"),
		field.Float("price"),
		field.Int("stock_quantity"),
		field.Int("min_stock_threshold").Optional(),
		field.String("sku").Optional(),
		field.Time("expiry_date").Optional(),
		field.String("size").Optional(),
		field.String("color").Optional(),
		field.String("brand").Optional(),
		field.Time("purchase_date").Optional(),
		field.Time("last_updated").Default(time.Now),
	}
}

func (PlaceInventory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("place", Place.Type).
			Ref("inventories").
			Unique(),
		edge.From("inventory_type", InventoryType.Type).
			Ref("place_inventories").
			Unique(),
		edge.To("attributes", PlaceInventoryAttribute.Type),
		edge.To("media", Media.Type),
		edge.To("transaction_histories", TransactionHistory.Type),
		edge.To("reservation_blocks", ReservationBlock.Type),
	}
}

type PlaceInventoryAttribute struct {
	ent.Schema
}

func (PlaceInventoryAttribute) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("value"),
		field.JSON("category_specific_value", map[string]interface{}{}).Optional(),
	}
}

func (PlaceInventoryAttribute) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("inventory", PlaceInventory.Type).
			Ref("attributes").
			Unique(),
		edge.From("attribute_type", InventoryAttribute.Type).
			Ref("place_inventory_attributes").
			Unique(),
	}
}

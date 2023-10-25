package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Enum("status").
			Values("pending", "confirmed", "completed", "cancelled").
			Default("pending"),
		field.Float("total_amount"),
		field.JSON("additional_info", map[string]interface{}{}).Optional(),
		field.String("deleted_at").Optional(),
		field.Bool("is_deleted").Default(false),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("orders").
			Unique(),
		edge.To("order_items", OrderItem.Type),
		edge.From("table", PlaceTable.Type).
			Ref("orders"),
	}
}

type OrderItem struct {
	ent.Schema
}

// Fields of the OrderItem.
func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.Int("quantity").
			Default(1),
		field.Float("price_per_item"),
		field.Float("total_price").
			Immutable().
			SchemaType(map[string]string{"postgres": "decimal"}),
		field.JSON("options", []string{}).Optional(),
	}
}

// Edges of the OrderItem.
func (OrderItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("order_items"),
		edge.From("menu_item", MenuItem.Type).
			Ref("order_items"),
	}
}

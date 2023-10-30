package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Menu struct {
	ent.Schema
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("name"),
		field.String("deleted_at").Optional(),
		field.Bool("is_deleted").Default(false),
		field.String("description").Optional(),
		field.String("preparation_time").Optional(),
		field.String("options").Optional(),
		field.String("price").Optional(),
		field.Bool("is_available").Default(true),

	}
}

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("place", Place.Type).
			Ref("menus"),
		edge.To("categories", Category.Type),
		edge.To("menu_items", MenuItem.Type),
		edge.To("media", Media.Type),
		// add any other edges that a Menu would have
	}
}

type MenuItem struct {
	ent.Schema
}

// Fields of the MenuItem.
func (MenuItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("name"),
		field.String("description").Optional(),
		field.Float("price"),
		field.Int("preparation_time").Optional(),
		field.JSON("options", []string{}).Optional(),
		field.String("deleted_at").Optional(),
		field.Bool("is_deleted").Default(false),
	}
}

// Edges of the MenuItem.
func (MenuItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("menu", Menu.Type).
			Ref("menu_items"),
		edge.To("inventory", PlaceInventory.Type).Unique(),
		edge.To("media", Media.Type),
		edge.To("order_items", OrderItem.Type),
	}
}

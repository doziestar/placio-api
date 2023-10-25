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
		field.String("description").Optional(),
	}
}

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("place", Place.Type).
			Ref("menus").
			Unique(),
		edge.To("categories", Category.Type),
		edge.To("menu_items", MenuItem.Type),
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

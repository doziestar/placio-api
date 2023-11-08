package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
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
		field.String("options").Optional(),
		field.Enum("foodType").Optional().Values("local", "intercontinental", "national", "regional", "continental"),
		field.Enum("menuItemType").Optional().Values("food", "drink"),
		field.Enum("drinkType").Optional().Values("alcoholic", "non-alcoholic", "both"),
		field.Enum("dietaryType").Optional().Values("vegan", "vegetarian", "non-vegetarian", "both"),
		field.Bool("is_available").Default(true),
		field.Time("updated_at").Default(time.Now().Local()).Optional(),
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
		edge.To("created_by", User.Type),
		edge.To("updated_by", User.Type),
	}
}

type Ingredient struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
	Unit     string `json:"unit"`
	Note     string `json:"note"` // Any special note about the ingredient, like "chopped" or "fresh".
}

// NutritionalInfo provides detailed information on the nutritional value of the menu item.
type NutritionalInfo struct {
	Calories     int                `json:"calories"`
	TotalFat     float64            `json:"total_fat"`           // In grams
	SaturatedFat float64            `json:"saturated_fat"`       // In grams
	TransFat     float64            `json:"trans_fat"`           // In grams
	Cholesterol  float64            `json:"cholesterol"`         // In milligrams
	Sodium       float64            `json:"sodium"`              // In milligrams
	TotalCarbs   float64            `json:"total_carbohydrates"` // In grams
	DietaryFiber float64            `json:"dietary_fiber"`       // In grams
	Sugars       float64            `json:"sugars"`              // In grams
	Protein      float64            `json:"protein"`             // In grams
	Vitamins     map[string]float64 `json:"vitamins"`            // Vitamins and their amounts
	Minerals     map[string]float64 `json:"minerals"`            // Minerals and their amounts
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
		field.String("currency").Optional(),
		field.Bool("is_available").Default(true),
		field.Int("preparation_time").Optional(),
		field.JSON("options", []string{}).Optional(),
		field.String("deleted_at").Optional(),
		field.Enum("type").Optional().Values("local", "intercontinental", "national", "regional", "continental"),
		field.Enum("status").Optional().Values("available", "unavailable"),
		field.Enum("DrinkType").Optional().Values("alcoholic", "non-alcoholic", "both"),
		field.Enum("DietaryType").Optional().Values("vegan", "vegetarian", "non-vegetarian", "both"),
		field.Bool("is_deleted").Default(false),
		//field.JSON("ingredients", []Ingredient{}).Optional(),
		//field.JSON("nutritional_info", NutritionalInfo{}).Optional(),
		field.Int("calories").Optional(),
		field.Int("serve_size").Optional(),
		field.Time("available_from").Optional(),
		field.Time("available_until").Optional(),
		field.String("image_url").Optional(),
		field.Enum("spiciness_level").Optional().Values("mild", "medium", "hot"),
		field.JSON("allergens", []string{}).Optional(),
		field.String("chef_special_note").Optional(),
		field.Int("rating").Optional(),                      // Average rating, out of 5
		field.Int("review_count").Optional(),                // Number of reviews
		field.String("category").Optional(),                 // Such as 'Entrees', 'Salads', etc.
		field.Int("order_count").Optional(),                 // How often the item has been ordered
		field.String("sku").Optional(),                      // Stock Keeping Unit for inventory management
		field.Bool("is_featured").Default(false).Optional(), // To highlight specific menu items
		field.Bool("is_new").Default(false).Optional(),      // To flag new items on the menu
		field.Bool("is_seasonal").Default(false).Optional(), // Seasonal availability
		field.String("season").Optional(),                   // If is_seasonal is true, specify which season
		field.Int("discount_percentage").Optional(),         // If any discounts are applied to the item
		field.String("promotion_description").Optional(),    // If the item is part of a promotion
		field.Time("promotion_start").Optional(),            // Promotion start time
		field.Time("promotion_end").Optional(),              // Promotion end time
		field.JSON("tags", []string{}).Optional(),           // Tags for searchability or characteristics like 'signature dish'
		field.JSON("related_items", []string{}).Optional(),
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
		//edge.To("reviews", Review.Type),
	}
}

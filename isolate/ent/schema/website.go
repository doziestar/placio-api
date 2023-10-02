package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Website struct {
	ent.Schema
}

func (Website) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.String("domainName").Unique(),
		field.String("heading_text"),
		field.String("business_logo"),
		field.String("business_name"),
		field.String("banner_section_background_image"),
		field.String("banner_section_background_color"),
		field.String("banner_section_text"),
		field.String("three_items_section_heading_text"),
		field.Text("three_items_section_details_text"),
		field.String("three_items_section_item_one_text"),
		field.String("three_items_section_item_two_text"),
		field.String("three_items_section_item_three_text"),
		field.String("banner_two_section_background_image"),
		field.String("banner_two_section_background_color"),
		field.String("banner_two_left_section_heading_text"),
		field.String("banner_two_left_section_details_text"),
		field.String("banner_two_left_section_button_text"),
		field.String("banner_two_left_section_button_link"),
		field.String("banner_two_right_side_image"),
		field.JSON("achievements_section", map[string]interface{}{}),
		field.String("Inventory_section_heading_text"),
		field.Time("creationDate").Default(time.Now),
		field.Time("lastUpdated").UpdateDefault(time.Now),
		field.String("title"),
		field.String("description"),
		field.String("keywords"),
		field.String("language"),
		field.String("logo"),
		field.String("favicon"),
		field.String("facebook"),
		field.String("twitter"),
		field.String("instagram"),
		field.String("youtube"),
		field.String("linkedin"),
		field.String("pinterest"),
		field.JSON("mapCoordinates", map[string]interface{}{}),
		field.String("longitude"),
		field.String("latitude"),
		field.String("address"),
		field.String("city"),
		field.String("state"),
		field.String("country"),
		field.String("zipCode"),
		field.String("phoneNumber"),
		field.String("email"),
		field.JSON("metaTags", map[string]interface{}{}),
	}
}

func (Website) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("business", Business.Type).Ref("websites").Unique().Required(),
		edge.From("template", Template.Type).Unique().Required().Ref("websites"),
		edge.To("customBlocks", CustomBlock.Type),
		edge.To("assets", Media.Type),
		//edge.From("users", User.Type).Ref("websites"),
	}
}

type Template struct {
	ent.Schema
}

func (Template) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("defaultHTML"),
		field.String("defaultCSS"),
	}
}

func (Template) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("websites", Website.Type),
	}
}

type CustomBlock struct {
	ent.Schema
}

func (CustomBlock) Fields() []ent.Field {
	return []ent.Field{
		field.String("content"),
	}
}

func (CustomBlock) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("website", Website.Type).Unique().Required().Ref("customBlocks"),
	}
}

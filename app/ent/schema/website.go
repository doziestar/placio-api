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
		field.String("heading_text").Optional(),
		field.String("business_logo").Optional(),
		field.String("business_name").Optional(),
		field.String("banner_section_background_image").Optional(),
		field.String("banner_section_background_color").Optional(),
		field.String("banner_section_text").Optional(),
		field.String("three_items_section_heading_text").Optional(),
		field.Text("three_items_section_details_text").Optional(),
		field.String("three_items_section_item_one_text").Optional(),
		field.String("three_items_section_item_two_text").Optional(),
		field.String("three_items_section_item_three_text").Optional(),
		field.String("banner_two_section_background_image").Optional(),
		field.String("banner_two_section_background_color").Optional(),
		field.String("banner_two_left_section_heading_text").Optional(),
		field.String("banner_two_left_section_details_text").Optional(),
		field.String("banner_two_left_section_button_text").Optional(),
		field.String("banner_two_left_section_button_link").Optional(),
		field.String("banner_two_right_side_image").Optional(),
		field.JSON("achievements_section", map[string]interface{}{}).Optional(),
		field.String("Inventory_section_heading_text").Optional(),
		field.Time("creationDate").Default(time.Now),
		field.Time("lastUpdated").UpdateDefault(time.Now),
		field.String("title").Optional(),
		field.String("description").Optional(),
		field.String("keywords").Optional(),
		field.String("language").Optional(),
		field.String("logo").Optional(),
		field.String("favicon").Optional(),
		field.String("facebook").Optional(),
		field.String("twitter").Optional(),
		field.String("instagram").Optional(),
		field.String("youtube").Optional(),
		field.String("linkedin").Optional(),
		field.String("pinterest").Optional(),
		field.JSON("mapCoordinates", map[string]interface{}{}).Optional(),
		field.String("longitude").Optional(),
		field.String("latitude").Optional(),
		field.String("address").Optional(),
		field.String("city").Optional(),
		field.String("state").Optional(),
		field.String("country").Optional(),
		field.String("zipCode").Optional(),
		field.String("phoneNumber").Optional(),
		field.String("email").Optional(),
		field.JSON("metaTags", map[string]interface{}{}).Optional(),
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

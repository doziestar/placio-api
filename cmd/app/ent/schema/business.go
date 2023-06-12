package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Business holds the schema definition for the Business entity.
type Business struct {
	ent.Schema
}

// Fields of the Business.
func (Business) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Business.
func (Business) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("userBusinesses", UserBusiness.Type),
		edge.To("business_account_settings", BusinessAccountSettings.Type).
			Unique(),
		edge.To("posts", Post.Type),
	}
}

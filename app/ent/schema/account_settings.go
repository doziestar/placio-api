package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BusinessAccountSettings holds the schema definition for the BusinessAccountSettings entity.
type AccountSettings struct {
	ent.Schema
}

// Fields of the BusinessAccountSettings.
func (AccountSettings) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.Bool("TwoFactorAuthentication"),
		field.JSON("BlockedUsers", []string{}),
		field.JSON("MutedUsers", []string{}),
	}
}

// Edges of the BusinessAccountSettings.
func (AccountSettings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("business_account", Business.Type).
			Ref("business_account_settings").
			Unique().
			Required(),
	}
}

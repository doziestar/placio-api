package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BusinessAccountSettings holds the schema definition for the BusinessAccountSettings entity.
type BusinessAccountSettings struct {
	ent.Schema
}

// Fields of the BusinessAccountSettings.
func (BusinessAccountSettings) Fields() []ent.Field {
	return []ent.Field{
		field.String("BusinessAccountSettingsID").Unique(),
		field.String("BusinessAccountID").Unique(),
		field.Bool("TwoFactorAuthentication"),
		field.JSON("BlockedUsers", []string{}),
		field.JSON("MutedUsers", []string{}),
	}
}

// Edges of the BusinessAccountSettings.
func (BusinessAccountSettings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("business_account", Business.Type).
			Ref("business_account_settings").
			Unique().
			Required(),
	}
}

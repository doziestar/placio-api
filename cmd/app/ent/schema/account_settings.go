package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AccountSettings holds the schema definition for the AccountSettings entity.
type AccountSettings struct {
	ent.Schema
}

// Fields of the AccountSettings.
func (AccountSettings) Fields() []ent.Field {
	return []ent.Field{
		field.String("ID").Unique(),
		field.String("AccountID").Unique(),
		field.Bool("TwoFactorAuthentication"),
		field.JSON("BlockedUsers", []string{}),
		field.JSON("MutedUsers", []string{}),
		field.String("BusinessAccountID"),
	}
}

// Edges of the AccountSettings.
func (AccountSettings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("business_account", BusinessAccount.Type).
			Ref("account_settings").
			Unique().
			Required(),
	}
}

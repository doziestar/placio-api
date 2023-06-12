package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BusinessAccount holds the schema definition for the BusinessAccount entity.
type BusinessAccount struct {
	ent.Schema
}

// Fields of the BusinessAccount.
func (BusinessAccount) Fields() []ent.Field {
	return []ent.Field{
		field.String("BusinessAccountID").Unique(),
		field.String("Name"),
		field.Bool("Active").Default(false),
		field.Time("CreatedAt"),
		field.Time("UpdatedAt"),
	}
}

// Edges of the BusinessAccount.
func (BusinessAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("relationships", UserBusinessRelationship.Type),
		edge.To("account_settings", AccountSettings.Type),
		edge.From("invitations", Invitation.Type).Ref("business_account"),
		edge.To("userBusinessRelationships", UserBusinessRelationship.Type),
	}
}

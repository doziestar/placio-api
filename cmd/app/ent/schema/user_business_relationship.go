package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserBusinessRelationship holds the schema definition for the UserBusinessRelationship entity.
type UserBusinessRelationship struct {
	ent.Schema
}

// Fields of the UserBusinessRelationship.
func (UserBusinessRelationship) Fields() []ent.Field {
	return []ent.Field{
		field.String("ID").Unique(),
		field.String("UserID"),
		field.String("BusinessAccountID"),
		field.String("Role"),
		field.Time("CreatedAt"),
		field.Time("UpdatedAt"),
	}
}

// Edges of the UserBusinessRelationship.
func (UserBusinessRelationship) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("relationships").
			Unique().
			Required(),
		edge.From("business_account", BusinessAccount.Type).
			Ref("relationships").
			Unique().
			Required(),
	}
}

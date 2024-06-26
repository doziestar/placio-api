package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserBusiness holds the schema definition for the UserBusiness entity.
type UserBusiness struct {
	ent.Schema
}

// Fields of the UserBusiness.
func (UserBusiness) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("role"),
		field.String("permissions"). // added this field
						Optional().
						NotEmpty(),
	}
}

// Edges of the UserBusiness.
func (UserBusiness) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("userBusinesses").
			Unique(),
		edge.From("business", Business.Type).
			Ref("userBusinesses").
			Unique(),
	}
}

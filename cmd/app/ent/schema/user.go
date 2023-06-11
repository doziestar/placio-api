package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/auth0/go-auth0/management"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("UserID").Unique(),
		field.String("Auth0ID"),
		field.Time("CreatedAt"),
		field.Time("UpdatedAt"),
		field.JSON("Auth0Data", &management.User{}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("relationships", UserBusinessRelationship.Type),
	}
}

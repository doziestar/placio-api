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
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("auth0_id").Unique(),
		field.JSON("auth0_data", &management.User{}).Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("userBusinesses", UserBusiness.Type),
		edge.To("comments", Comment.Type),
		edge.To("likes", Like.Type),
		edge.To("posts", Post.Type),
	}
}

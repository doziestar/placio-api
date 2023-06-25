package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

type UserFollowUser struct {
	ent.Schema
}

// Edges of the UserFollowUser.
func (UserFollowUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("follower", User.Type).
			Ref("followedUsers").
			Unique(),
		edge.From("followed", User.Type).
			Ref("followerUsers").
			Unique(),
	}
}

type UserFollowBusiness struct {
	ent.Schema
}

// Edges of the UserFollowBusiness.
func (UserFollowBusiness) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("followedBusinesses").
			Unique(),
		edge.From("business", Business.Type).
			Ref("followerUsers").
			Unique(),
	}
}

type BusinessFollowBusiness struct {
	ent.Schema
}

// Edges of the BusinessFollowBusiness.
func (BusinessFollowBusiness) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("follower", Business.Type).
			Ref("followedBusinesses").
			Unique(),
		edge.From("followed", Business.Type).
			Ref("followerBusinesses").
			Unique(),
	}
}

type BusinessFollowUser struct {
	ent.Schema
}

// Edges of the BusinessFollowUser.
func (BusinessFollowUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("business", Business.Type).
			Ref("followedUsers").
			Unique(),
		edge.From("user", User.Type).
			Ref("followerBusinesses").
			Unique(),
	}
}

// UserFollowPlace holds the schema definition for the UserFollowPlace entity.
type UserFollowPlace struct {
	ent.Schema
}

// Edges of the UserFollowPlace.
func (UserFollowPlace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("followedPlaces").
			Unique(),
		edge.From("place", Place.Type).
			Ref("followerUsers").
			Unique(),
	}
}

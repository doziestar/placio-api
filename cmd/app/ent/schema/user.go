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
		field.String("name").Optional(),
		field.String("picture").Optional(),
		field.String("cover_image").Optional().Default("https://res.cloudinary.com/placio/image/upload/v1686842319/mjl8stmbn5xmfsm50vbg.jpg"),
		field.String("username").Unique(),
		field.Text("bio").Optional().Default("Add a bio to your profile"),
		field.JSON("auth0_data", &management.User{}).Optional(),
		field.JSON("app_settings", map[string]interface{}{}).Optional(),
		field.JSON("user_settings", map[string]interface{}{}).Optional(),
		field.String("search_text").Optional(),
		field.Float("relevance_score").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("userBusinesses", UserBusiness.Type),
		edge.To("comments", Comment.Type),
		edge.To("likes", Like.Type),
		edge.To("posts", Post.Type),
		edge.To("followedUsers", UserFollowUser.Type),
		edge.To("followerUsers", UserFollowUser.Type),
		edge.To("followedBusinesses", UserFollowBusiness.Type),
		edge.To("followerBusinesses", BusinessFollowUser.Type),
		edge.To("reviews", Review.Type),
		edge.To("bookings", Booking.Type),
		edge.To("reservations", Reservation.Type),
		edge.To("helps", Help.Type).Immutable(),
		edge.To("categories", Category.Type),
		edge.To("events", Event.Type),
		edge.To("places", Place.Type),
		edge.To("categoryAssignments", CategoryAssignment.Type),
	}
}

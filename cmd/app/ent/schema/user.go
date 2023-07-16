package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/auth0/go-auth0/management"
	"github.com/getsentry/sentry-go"
	gen "placio-app/ent"
	"placio-app/ent/hook"
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
		field.String("website").Optional(),
		field.String("location").Optional(),
		field.JSON("map_coordinates", map[string]interface{}{}).Optional(),
		field.String("longitude").Optional(),
		field.String("latitude").Optional(),
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
		edge.To("places", Place.Type),
		edge.To("categoryAssignments", CategoryAssignment.Type),
		edge.To("ownedEvents", Event.Type).
			Unique(),
		edge.To("userFollowEvents", UserFollowEvent.Type),
		edge.To("followedPlaces", UserFollowPlace.Type),
		edge.To("likedPlaces", UserLikePlace.Type),
		edge.To("ratings", Rating.Type),
	}
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
					oldLocation, _ := m.OldLocation(ctx)
					err := ProcessLocation(m, oldLocation)
					if err != nil {
						sentry.CaptureEvent(&sentry.Event{
							Message: "Failed to process location",
							Level:   sentry.LevelError,
							Extra: map[string]interface{}{
								"error": err,
								"ctx":   ctx,
							},
						})
					}
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate|ent.OpUpdate,
		),
	}
}

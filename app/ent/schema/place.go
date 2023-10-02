package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	//"github.com/getsentry/sentry-go"
	//gen "placio-app/ent/ent"
	//"placio-app/ent/ent/hook"
)

type Place struct {
	ent.Schema
}

// Fields of the Place.
func (Place) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("name"),
		field.String("type").Optional(),
		field.String("description").Optional(),
		field.String("location").Optional(),
		field.String("email").Optional(),
		field.String("phone").Optional(),
		field.String("website").Optional(),
		field.String("cover_image").Optional().Default("https://res.cloudinary.com/placio/image/upload/v1686842319/mjl8stmbn5xmfsm50vbg.jpg"),
		field.String("picture").Optional(),
		field.String("country").Optional(),
		field.String("city").Optional(),
		field.String("state").Optional(),
		field.JSON("place_settings", map[string]interface{}{}).Optional(),
		field.JSON("opening_hours", map[string]interface{}{}).Optional(),
		field.JSON("social_media", map[string]interface{}{}).Optional(),
		field.JSON("payment_options", map[string]interface{}{}).Optional(),
		field.JSON("tags", []string{}).Optional(),
		field.JSON("features", []string{}).Optional(),
		field.JSON("additional_info", map[string]interface{}{}).Optional(),
		field.JSON("images", []string{}).Optional(),
		field.JSON("availability", map[string]interface{}{}).Optional(),
		field.String("special_offers").Optional(),
		field.Float("sustainability_score").Optional(),
		field.JSON("map_coordinates", map[string]interface{}{}).Optional(),
		field.String("longitude").Optional(),
		field.String("latitude").Optional(),
		field.String("search_text").Optional(),
		field.Float("relevance_score").Optional(),
		field.Int("follower_count").Default(0),
		field.Int("like_count").Default(0),
		field.Int("review_count").Default(0),
		field.Int("following_count").Default(0),
		field.Bool("is_Premium").Default(false),
		field.Bool("is_published").Default(false),
		field.Bool("likedByCurrentUser").Default(false),
		field.Bool("followedByCurrentUser").Default(false),
	}
}

// Edges of the Place.
func (Place) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("business", Business.Type).
			Ref("places").
			Unique(),
		edge.From("users", User.Type).Ref("places"),
		edge.To("reviews", Review.Type),
		edge.To("events", Event.Type),
		edge.From("amenities", Amenity.Type).
			Ref("places"),
		edge.To("menus", Menu.Type),
		edge.To("medias", Media.Type),
		edge.To("rooms", Room.Type),
		edge.To("reservations", Reservation.Type),
		edge.To("bookings", Booking.Type),
		edge.To("categories", Category.Type),
		edge.To("categoryAssignments", CategoryAssignment.Type),
		edge.From("faqs", FAQ.Type).Ref("place"),
		edge.From("likedByUsers", UserLikePlace.Type).Ref("place"),
		edge.To("followerUsers", UserFollowPlace.Type),
		edge.To("ratings", Rating.Type),
		edge.To("inventories", PlaceInventory.Type),
	}
}

//func (Place) Hooks() []ent.Hook {
//	return []ent.Hook{
//		hook.On(
//			func(next ent.Mutator) ent.Mutator {
//				return hook.PlaceFunc(func(ctx context.Context, m *gen.PlaceMutation) (ent.Value, error) {
//					oldLocation, _ := m.OldLocation(ctx)
//					err := ProcessLocation(m, oldLocation)
//					if err != nil {
//						sentry.CaptureEvent(&sentry.Event{
//							Message: "Failed to process location",
//							Level:   sentry.LevelError,
//							Extra: map[string]interface{}{
//								"error": err,
//								"ctx":   ctx,
//							},
//						})
//					}
//					return next.Mutate(ctx, m)
//				})
//			},
//			ent.OpCreate|ent.OpUpdate,
//		),
//	}
//}

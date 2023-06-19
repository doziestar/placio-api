package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
		field.String("type"),
		field.String("description").Optional(),
		field.String("location"),
		field.JSON("images", []string{}).Optional(),
		field.JSON("availability", map[string]interface{}{}).Optional(),
		field.String("special_offers").Optional(),
		field.Float("sustainability_score").Optional(),
	}
}

// Edges of the Place.
func (Place) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("business", Business.Type).
			Ref("places").
			Unique(),
		edge.To("reviews", Review.Type),
		edge.To("events", Event.Type),
		edge.From("amenities", Amenity.Type).
			Ref("places"),
		edge.To("menus", Menu.Type),
		edge.To("rooms", Room.Type),
		edge.To("reservations", Reservation.Type),
		edge.To("bookings", Booking.Type),
		edge.To("categories", Category.Type),
	}
}

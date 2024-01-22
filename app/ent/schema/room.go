package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type RoomCategory struct {
	ent.Schema
}

// Fields of the RoomCategory.
func (RoomCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("name"),
		field.String("description").Optional(),
		field.String("price").Optional(),
	}
}

// Edges of the RoomCategory.
func (RoomCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("place", Place.Type).
			Ref("room_categories"),
		edge.To("rooms", Room.Type),
		edge.To("media", Media.Type),
		edge.To("amenities", Amenity.Type),
	}
}

type Room struct {
	ent.Schema
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("name").Optional(),
		field.Int("room_number").Optional(),
		field.String("room_type").Optional(),
		field.String("room_status").Optional(),
		field.String("room_rating").Optional(),
		field.Float("room_price").Optional(),
		field.String("qr_code").Optional(),
		field.Enum("status").Values("available", "unavailable", "maintenance", "reserved").Default("available"),
		field.JSON("extras", map[string]interface{}{}).Optional(),
		field.String("description").Optional(),
		field.Bool("availability").Default(true),
		field.String("image").Optional(),
	}
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("place", Place.Type).
			Ref("rooms"),
		edge.From("room_category", RoomCategory.Type).
			Ref("rooms"),
		edge.To("bookings", Booking.Type),
		edge.To("amenities", Amenity.Type),
		edge.To("media", Media.Type),
		edge.To("reservations", Reservation.Type),
	}
}

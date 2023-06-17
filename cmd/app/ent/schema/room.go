package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

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
		field.String("number"),
		field.String("type"),
		field.Float("price"),
		field.String("description").Optional(),
		field.Bool("availability"),
		field.String("image").Optional(),
	}
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("place", Place.Type).
			Ref("rooms").
			Unique(),
		edge.To("bookings", Booking.Type),
	}
}

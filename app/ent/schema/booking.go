package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Booking holds the schema definition for the Booking, entity.
type Booking struct {
	ent.Schema
}

// Fields of the Booking.
func (Booking) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.Time("startDate"),
		field.Time("endDate"),
		field.String("status"),
		field.Time("bookingDate"),
	}
}

// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).
			Ref("bookings").
			Unique(),
		edge.From("user", User.Type).
			Ref("bookings").
			Unique(),
	}
}

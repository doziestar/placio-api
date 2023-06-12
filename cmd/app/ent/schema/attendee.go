package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Attendee holds the schema definition for the Attendee entity.
type Attendee struct {
	ent.Schema
}

// Fields of the Attendee.
func (Attendee) Fields() []ent.Field {
	return []ent.Field{
		field.String("AttendeeID").Unique(),
		field.String("EventID"),
		field.String("UserID"),
		field.String("TicketID"),
		field.Bool("Attended"),
		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Attendee.
func (Attendee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("attendees").
			Field("EventID").
			Unique(),
		edge.From("user", User.Type).
			Ref("attendees").
			Field("UserID").
			Unique(),
		edge.From("ticket", Ticket.Type).
			Ref("attendees").
			Field("TicketID").
			Unique(),
	}
}

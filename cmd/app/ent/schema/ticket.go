package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Ticket holds the schema definition for the Ticket entity.
type Ticket struct {
	ent.Schema
}

// Fields of the Ticket.
func (Ticket) Fields() []ent.Field {
	return []ent.Field{
		field.String("TicketID").Unique(),
		field.String("Name"),
		field.Int("Price"),
		field.String("EventID"),
		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Ticket.
func (Ticket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("tickets").
			Field("EventID").
			Unique(),
		edge.To("ticket_options", TicketOption.Type),
		//edge.To("attendees", Attendee.Type),
		edge.To("ratings", Rating.Type),
	}
}

// TicketOption holds the schema definition for the TicketOption entity.
type TicketOption struct {
	ent.Schema
}

// Fields of the TicketOption.
func (TicketOption) Fields() []ent.Field {
	return []ent.Field{
		field.String("TicketOptionID").Unique(),
		field.String("EventID"),
		field.String("Name"),
		field.Float("Price"),
		field.Int("Quantity"),
		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the TicketOption.
func (TicketOption) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("ticket_options").
			Field("EventID").
			Unique(),
	}
}

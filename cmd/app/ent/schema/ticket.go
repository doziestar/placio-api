package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Ticket holds the schema definition for the Ticket entity.
type Ticket struct {
	ent.Schema
}

// Fields of the Ticket.
func (Ticket) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		//field.String("name"),
		//field.Int("price"),
		//field.String("eventID"),
		//field.Time("createdAt").Default(time.Now),
		//field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Ticket.
func (Ticket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("tickets").
			Field("id").
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
		field.String("id").Unique().Immutable(),
		//field.String("event-id"),
		//field.String("name"),
		//field.Float("Price"),
		//field.Int("Quantity"),
		//field.Time("CreatedAt").Default(time.Now),
		//field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
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

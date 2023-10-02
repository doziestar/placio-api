package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Ticket struct {
	ent.Schema
}

func (Ticket) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Ticket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("tickets").Unique(),
		edge.To("ticket_options", TicketOption.Type),
	}
}

type TicketOption struct {
	ent.Schema
}

func (TicketOption) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (TicketOption) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("ticket_options").Unique(),
	}
}

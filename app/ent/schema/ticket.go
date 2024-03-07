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
		field.String("ticketCode").Unique(),
		field.Enum("status").Values("available", "reserved", "sold", "validated").Default("available"),
		field.Time("purchaseTime"),
		field.Time("validationTime").Optional(),
		field.String("purchaserEmail").Optional(),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Ticket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ticketOption", TicketOption.Type).Ref("tickets").Unique().Required(),
		edge.From("purchaser", User.Type).Ref("purchasedTickets").Unique(),
		edge.From("event", Event.Type).Ref("tickets").Unique().Required(),
	}
}

type TicketOption struct {
	ent.Schema
}

func (TicketOption) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.String("name").NotEmpty(),
		field.Text("description"),
		field.Float("price").Positive().Default(0),
		field.Int("quantityAvailable").Default(0),
		field.Int("quantitySold").Default(0),
		field.Enum("status").Values("active", "inactive", "sold_out").Default("active"),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (TicketOption) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("ticketOptions").Unique(),
		edge.To("tickets", Ticket.Type),
	}
}

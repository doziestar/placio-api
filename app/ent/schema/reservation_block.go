package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ReservationBlock struct {
	ent.Schema
}

func (ReservationBlock) Fields() []ent.Field {
	return []ent.Field{
		field.Time("start_time"),
		field.Time("end_time"),
		field.Enum("status").Values("confirmed", "pending", "canceled"),
	}
}

func (ReservationBlock) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("place_inventory", PlaceInventory.Type).Ref("reservation_blocks").Unique(),
		edge.From("user", User.Type).Ref("reservation_blocks").Unique(),
	}
}

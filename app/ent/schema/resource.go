package schema

import (
	"entgo.io/ent"
)

type Resourse struct {
	ent.Schema
}

//// Fields of the Reservation.
//func (Reservation) Fields() []ent.Field {
//	return []ent.Field{
//		field.String("id").
//			MaxLen(36).
//			Unique().
//			Immutable(),
//		field.Time("date"),
//		field.Time("time"),
//		field.Int("numberOfPeople"),
//		field.String("status"),
//	}
//}
//
//// Edges of the Reservation.
//func (Reservation) Edges() []ent.Edge {
//	return []ent.Edge{
//		edge.From("place", Place.Type).
//			Ref("reservations").
//			Unique(),
//		edge.From("user", User.Type).
//			Ref("reservations").
//			Unique(),
//	}
//}

package schema

import "entgo.io/ent"

// Booking holds the schema definition for the Booking, entity.
type Booking struct {
	ent.Schema
}

// Fields of the Booking,.
func (Booking) Fields() []ent.Field {
	return nil
}

// Edges of the Booking,.
func (Booking) Edges() []ent.Edge {
	return nil
}

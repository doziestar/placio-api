package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Amenity struct {
	ent.Schema
}

// Fields of the Amenity.
func (Amenity) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("name").
			Unique(),
		field.String("icon"),
	}
}

// Edges of the Amenity.
func (Amenity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("places", Place.Type),
		edge.From("rooms", Room.Type).
			Ref("amenities"),
		edge.From("room_categories", RoomCategory.Type).
			Ref("amenities"),
	}
}

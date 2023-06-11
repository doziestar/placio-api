package schema

import "entgo.io/ent"

// Rating holds the schema definition for the Rating, entity.
type Rating struct {
	ent.Schema
}

// Fields of the Rating,.
func (Rating) Fields() []ent.Field {
	return nil
}

// Edges of the Rating,.
func (Rating) Edges() []ent.Edge {
	return nil
}

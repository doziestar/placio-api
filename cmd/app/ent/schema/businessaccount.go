package schema

import "entgo.io/ent"

// BusinessAccount holds the schema definition for the BusinessAccount, entity.
type BusinessAccount struct {
	ent.Schema
}

// Fields of the BusinessAccount,.
func (BusinessAccount) Fields() []ent.Field {
	return nil
}

// Edges of the BusinessAccount,.
func (BusinessAccount) Edges() []ent.Edge {
	return nil
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Price struct {
	ent.Schema
}

// Fields of the Price.
func (Price) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.Float("price"),
		field.String("currency").Optional(),
		field.String("description").Optional(),
		field.Enum("type").Values("session", "day", "week", "month", "year"),
		field.Int("duration").Optional(),
		field.Int("session").Optional(),
	}

}

// Edges of the Price.
func (Price) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("plan", Plan.Type).Ref("prices").Unique(),
		edge.To("subscriptions", Subscription.Type),
	}
}

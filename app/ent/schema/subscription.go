package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Subscription struct {
	ent.Schema
}

// Fields of the Subscription.
func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(36).Unique().Immutable(),
		field.Time("start_date"),
		field.Time("end_date"),
		field.String("flutterwave_subscription_id").Optional(),
	}
}

// Edges of the Subscription.
func (Subscription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("subscriptions").Unique(),
		edge.From("plan", Plan.Type).Ref("subscriptions").Unique(),
		edge.From("price", Price.Type).Ref("subscriptions").Unique(),
	}
}

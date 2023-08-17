package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type TransactionHistory struct {
	ent.Schema
}

func (TransactionHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("transaction_type").Values("purchase", "sale", "return", "usage"),
		field.Int("quantity"),
		field.Time("date").Default(time.Now),
	}
}

func (TransactionHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("place_inventory", PlaceInventory.Type).Ref("transaction_histories").Unique(),
		edge.From("user", User.Type).Ref("transaction_histories").Unique(),
	}
}

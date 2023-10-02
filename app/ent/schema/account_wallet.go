package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type AccountWallet struct {
	ent.Schema
}

func (AccountWallet) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.Float("balance").Default(0),
		field.Float("total_deposited").Default(0),
		field.Float("total_withdrawn").Default(0),
		field.Float("total_earned").Default(0),
		field.Float("total_spent").Default(0),
		field.Float("total_refunded").Default(0),
		field.Float("total_fees").Default(0),
		field.Float("total_tax").Default(0),
		field.Float("total_discount").Default(0),
		field.Float("total_revenue").Default(0),
		field.Float("total_expenses").Default(0),
		field.Float("total_profit").Default(0),
		field.Float("total_loss").Default(0),
	}
}

func (AccountWallet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("wallet").Unique(),
		edge.From("business", Business.Type).Ref("wallet").Unique(),
	}
}

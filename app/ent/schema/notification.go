package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("title").MaxLen(255),
		field.String("message").MaxLen(255),
		field.String("link").MaxLen(255),
		field.Bool("is_read").Default(false),
		field.Int("type").Default(0),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.String("notifiable_type").MaxLen(255), // for example: "Post", "Comment", "Order", etc.
		field.String("notifiable_id").MaxLen(36),    // for example: "Post ID", "Comment ID", "Order ID", etc.
		field.String("triggered_by").MaxLen(36),
		field.String("triggered_to").MaxLen(36),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("notifications"),
		edge.From("business_account", Business.Type).Ref("notifications"),
	}
}

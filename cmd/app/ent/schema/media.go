package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Media holds the schema definition for the Media, entity.
type Media struct {
	ent.Schema
}

// Fields of the Media.
func (Media) Fields() []ent.Field {
	return []ent.Field{
		field.String("ID").Unique(),
		field.String("PostID").
			StructTag(`json:"post_id,omitempty"`),
		field.String("URL").
			StructTag(`json:"url,omitempty"`),
		field.String("MediaType").
			StructTag(`json:"media_type,omitempty"`).
			Comment("image, gif, or video"),
		field.Time("CreatedAt").
			Default(time.Now).
			Immutable().
			StructTag(`json:"created_at,omitempty"`),
		field.Time("UpdatedAt").
			Default(time.Now).
			UpdateDefault(time.Now).
			StructTag(`json:"updated_at,omitempty"`),
	}
}

// Edges of the Media.
func (Media) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("post", Post.Type).
			Unique().
			Required().
			Field("PostID"),
	}
}

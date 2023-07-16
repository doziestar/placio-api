package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Media holds the schema definition for the Media entity.
type Media struct {
	ent.Schema
}

// Fields of the Media.
func (Media) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("URL"),
		field.String("MediaType").Comment("image, gif, or video"),
		field.Time("CreatedAt").Default(time.Now).Immutable(),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
		field.Int("likeCount").
			Default(0).
			Comment("Count of likes for this media."),
		field.Int("dislikeCount").
			Default(0).
			Comment("Count of dislikes for this media."),
	}
}

// Edges of the Media.
func (Media) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", Post.Type).
			Ref("medias").
			Unique(),
		edge.From("review", Review.Type).
			Ref("medias").
			Unique(),
		edge.To("categories", Category.Type),
	}
}

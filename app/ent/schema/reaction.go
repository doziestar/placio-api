package schema

import (
	"entgo.io/ent"
)

type Reaction struct {
	ent.Schema
}

//func (Reaction) Fields() []ent.Field {
//	return []ent.Field{
//		field.String("id").
//			MaxLen(36).
//			Unique().
//			Immutable(),
//		field.Time("CreatedAt").Default(time.Now),
//		field.Time("UpdatedAt").UpdateDefault(time.Now),
//		field.Enum("Type").Values("like", "love", "wow", "sad", "angry"),
//	}
//}
//
//func (Reaction) Edges() []ent.Edge {
//	return []ent.Edge{
//		edge.From("user", User.Type).
//			Ref("reactions").
//			Unique(),
//		edge.To("post", Post.Type).Unique(),
//	}
//}

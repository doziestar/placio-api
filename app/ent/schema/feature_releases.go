package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

type FeatureRelease struct {
	ent.Schema
}

// Fields of the FeatureRelease.
func (FeatureRelease) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("feature_name"),
		field.String("description").Optional(),
		field.Enum("state").
			Values("testing", "staging", "live", "deprecated"),
		field.Time("release_date").
			Default(time.Now),
		field.JSON("eligibility_rules", map[string]interface{}{}).Optional(),
		field.String("documentation_link").Optional(),
		field.JSON("metadata", map[string]interface{}{}).Optional(),
	}
}

// Edges of the FeatureRelease.
func (FeatureRelease) Edges() []ent.Edge {
	return []ent.Edge{}
}

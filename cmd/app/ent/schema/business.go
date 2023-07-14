package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"fmt"
	gen "placio-app/ent"
	"placio-app/ent/hook"
	_ "placio-app/ent/runtime"
	"placio-app/utility"
)

// Business holds the schema definition for the Business entity.
type Business struct {
	ent.Schema
}

// Fields of the Business.
func (Business) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("name"),
		field.Text("description").Optional(),
		field.String("picture").Optional(),
		field.String("cover_image").Optional().Default("https://res.cloudinary.com/placio/image/upload/v1686842319/mjl8stmbn5xmfsm50vbg.jpg"),
		field.String("website").Optional(),
		field.String("location").Optional(),
		field.String("longitude").Optional(),
		field.JSON("map_coordinates", map[string]interface{}{}).Optional(),
		field.String("latitude").Optional(),
		field.String("email").Optional(),
		field.String("phone").Optional(),
		field.JSON("business_settings", map[string]interface{}{}).Optional(),
		field.String("url").Optional(),
		field.String("search_text").Optional(),
		field.Float("relevance_score").Optional(),
	}
}

// Edges of the Business.
func (Business) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("userBusinesses", UserBusiness.Type),
		edge.To("business_account_settings", AccountSettings.Type).
			Unique(),
		edge.To("posts", Post.Type),
		edge.To("followedUsers", BusinessFollowUser.Type),
		edge.To("followerUsers", UserFollowBusiness.Type),
		edge.To("followedBusinesses", BusinessFollowBusiness.Type),
		edge.To("followerBusinesses", BusinessFollowBusiness.Type),
		edge.To("places", Place.Type),
		edge.To("categories", Category.Type),
		edge.To("categoryAssignments", CategoryAssignment.Type),
		edge.To("events", Event.Type),
		edge.To("businessFollowEvents", BusinessFollowEvent.Type),
		edge.To("faqs", FAQ.Type),
	}
}

func (Business) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.BusinessFunc(func(ctx context.Context, m *gen.BusinessMutation) (ent.Value, error) {
					location, exist := m.Location()
					if exist {
						// Assuming the new location can be broken down to lat and long
						data, err := utility.GetCoordinates(location)
						if err != nil {
							return nil, fmt.Errorf("failed to get coordinates: %w", err)
						}
						latitude := fmt.Sprintf("%f", data.Features[0].Geometry.Coordinates[1])
						longitude := fmt.Sprintf("%f", data.Features[0].Geometry.Coordinates[0])

						m.SetLatitude(latitude)
						m.SetLongitude(longitude)

						mapCoordinates, err := utility.StructToMap(data.Features[0])
						if err != nil {
							return nil, fmt.Errorf("failed to convert struct to map: %w", err)
						}

						m.SetMapCoordinates(mapCoordinates)
					}

					return next.Mutate(ctx, m)
				})
			},
			// Add the hook only for Create operation.
			ent.OpCreate,
		),

		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.BusinessFunc(func(ctx context.Context, m *gen.BusinessMutation) (ent.Value, error) {
					location, exist := m.Location()
					oldLocation, _ := m.OldLocation(ctx)
					if exist && location != oldLocation {
						// Assuming the new location can be broken down to lat and long
						data, err := utility.GetCoordinates(location)
						if err != nil {
							return nil, fmt.Errorf("failed to get coordinates: %w", err)
						}
						latitude := fmt.Sprintf("%f", data.Features[0].Geometry.Coordinates[1])
						longitude := fmt.Sprintf("%f", data.Features[0].Geometry.Coordinates[0])

						m.SetLatitude(latitude)
						m.SetLongitude(longitude)

						mapCoordinates, err := utility.StructToMap(data.Features[0])
						if err != nil {
							return nil, fmt.Errorf("failed to convert struct to map: %w", err)
						}

						m.SetMapCoordinates(mapCoordinates)
					}

					return next.Mutate(ctx, m)
				})
			},
			// Add the hook only for Update operation.
			ent.OpUpdate,
		),
	}
}

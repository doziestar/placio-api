package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/getsentry/sentry-go"
	gen "placio-app/ent"
	"placio-app/ent/hook"
	"placio-pkg/errors"

	//"github.com/getsentry/sentry-go"
	//gen "placio-app/ent/ent"
	//"placio-app/ent/ent/hook"
	"time"
)

type EventOrganizer struct {
	ent.Schema
}

func (EventOrganizer) Fields() []ent.Field {
	return []ent.Field{
		field.String("organizerID").NotEmpty(),
		field.String("organizerType").NotEmpty().Validate(func(s string) error {
			if s != "user" && s != "business" {
				return errors.New("invalid organizer type")
			}
			return nil
		}),
	}
}

func (EventOrganizer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("event_organizers").Unique().Required(),
	}
}

type Event struct {
	ent.Schema
}

func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.String("name").Optional(),
		field.Enum("EventType").
			Values("event", "place", "business", "free", "paid").Optional(),
		field.String("status").Optional(),
		field.String("location").Optional(),
		field.String("url").Optional(),
		field.String("title").Optional(),
		field.String("time_zone").Optional(),
		field.Time("start_time").Optional(),
		field.Time("end_time").Optional(),
		field.String("start_date").Optional(),
		field.String("end_date").Optional(),
		field.Enum("frequency").Values("once", "daily", "weekly", "monthly", "yearly").Optional(),
		field.String("frequency_interval").Optional(),
		field.String("frequency_day_of_week").Optional(),
		field.String("frequency_day_of_month").Optional(),
		field.String("frequency_month_of_year").Optional(),
		field.Enum("venue_type").Values("online", "in_person", "hybrid").Optional(),
		field.String("venue_name").Optional(),
		field.String("venue_address").Optional(),
		field.String("venue_city").Optional(),
		field.String("venue_state").Optional(),
		field.String("venue_country").Optional(),
		field.String("venue_zip").Optional(),
		field.String("venue_lat").Optional(),
		field.String("venue_lon").Optional(),
		field.String("venue_url").Optional(),
		field.String("venue_phone").Optional(),
		field.String("venue_email").Optional(),
		field.JSON("tags", []string{}).Optional(),
		field.Text("description").Optional(),
		field.JSON("event_settings", map[string]interface{}{}).Optional(),
		field.String("cover_image").Optional().Default("https://res.cloudinary.com/placio/image/upload/v1686842319/mjl8stmbn5xmfsm50vbg.jpg"),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
		field.JSON("map_coordinates", map[string]interface{}{}).Optional(),
		field.String("longitude").Optional(),
		field.String("latitude").Optional(),
		field.String("search_text").Optional(),
		field.Float("relevance_score").Optional(),
		field.Int("follower_count").Default(0),
		field.Int("following_count").Default(0),
		field.Bool("is_Premium").Default(false),
		field.Bool("is_published").Default(false),
		field.Bool("is_Online").Default(false),
		field.Bool("is_Free").Default(false),
		field.Bool("is_Paid").Default(false),
		field.Bool("is_public").Default(false),
		field.Bool("is_Online_Only").Default(false),
		field.Bool("is_In_Person_Only").Default(false),
		field.Bool("is_Hybrid").Default(false),
		field.Bool("is_Online_And_In_Person").Default(false),
		field.Bool("is_Online_And_In_Person_Only").Default(false),
		field.Bool("is_Online_And_In_Person_Or_Hybrid").Default(false),
		field.Bool("likedByCurrentUser").Default(false),
		field.Bool("followedByCurrentUser").Default(false),
		field.Enum("registration_type").Values("none", "required", "optional", "closed").Optional(),
		field.String("registration_url").Optional(),
		field.Bool("is_physically_accessible").Default(false),
		field.String("accessibility_info").Optional(),
		field.Bool("is_virtually_accessible").Default(false),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tickets", Ticket.Type),
		edge.To("ticket_options", TicketOption.Type),
		edge.To("place", Place.Type),
		edge.To("event_categories", Category.Type),
		edge.To("event_category_assignments", CategoryAssignment.Type),
		edge.From("ownerUser", User.Type).
			Ref("ownedEvents").
			Unique(),
		edge.From("ownerBusiness", Business.Type).
			Ref("events").
			Unique(),
		edge.From("userFollowers", UserFollowEvent.Type).
			Ref("event"),
		edge.From("businessFollowers", BusinessFollowEvent.Type).
			Ref("event"),
		edge.From("faqs", FAQ.Type).Ref("event"),
		edge.To("ratings", Rating.Type),
		edge.To("additional_organizers", User.Type),
		edge.To("media", Media.Type),
		edge.To("event_comments", Comment.Type),
		edge.To("event_reviews", Review.Type),
		edge.To("performers", User.Type),
		edge.To("event_organizers", EventOrganizer.Type),
	}
}

func (Event) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.EventFunc(func(ctx context.Context, m *gen.EventMutation) (ent.Value, error) {
					oldLocation, _ := m.OldLocation(ctx)
					err := ProcessLocation(m, oldLocation)
					if err != nil {
						sentry.CaptureEvent(&sentry.Event{
							Message: "Failed to process location",
							Level:   sentry.LevelError,
							Extra: map[string]interface{}{
								"error": err,
								"ctx":   ctx,
							},
						})
					}
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate|ent.OpUpdate,
		),
	}
}

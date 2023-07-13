package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

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
		// TODO: convert tags to array
		field.String("tags").Optional(),
		field.Text("description").Optional(),
		field.JSON("event_settings", map[string]interface{}{}).Optional(),
		field.String("cover_image").Optional().Default("https://res.cloudinary.com/placio/image/upload/v1686842319/mjl8stmbn5xmfsm50vbg.jpg"),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
		field.String("search_text").Optional(),
		field.Float("relevance_score").Optional(),
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
	}
}

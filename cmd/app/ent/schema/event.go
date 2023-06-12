package schema

// Event holds the schema definition for the Event, entity.
//type Event struct {
//	ent.Schema
//}
//
//func (Event) Fields() []ent.Field {
//	return []ent.Field{
//		//field.String("EventID").Unique(),
//		field.String("Name").NotEmpty().Unique(),
//		field.Time("Date"),
//		field.Time("Time"),
//		field.Time("EndDate"),
//		field.Time("EndTime"),
//		field.String("Location"),
//		field.String("Address"),
//		field.String("City"),
//		field.String("State"),
//		field.String("Country"),
//		field.String("Description").NotEmpty(),
//		field.String("Category"),
//		field.Strings("Tags"),
//		field.String("ImageURL"),
//		field.String("Organizer"),
//		field.String("OrganizerEmail"),
//		field.String("OrganizerPhone"),
//		field.String("Website"),
//		field.String("TicketURL"),
//		field.String("PriceRange"),
//		field.Int("Capacity"),
//		field.Bool("IsFree"),
//		field.Bool("IsPublic"),
//		field.Bool("IsOnline"),
//		field.String("AccountID").NotEmpty(),
//		field.Time("CreatedAt").Default(time.Now),
//		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
//	}
//}
//
//// Edges of the Event.
//func (Event) Edges() []ent.Edge {
//	return []ent.Edge{
//		//edge.From("attendees", Attendee.Type).
//		//	Ref("event").
//		//	Unique(), // this indicates that each Attendee is connected to exactly one Event
//		edge.To("tickets", Ticket.Type).
//			Field("EventID"),
//		edge.To("ticket_options", TicketOption.Type).
//			Field("EventID"),
//		// Additional edges ...
//	}
//}

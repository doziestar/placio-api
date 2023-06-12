package schema

//
//// Attendee holds the schema definition for the Attendee entity.
//type Attendee struct {
//	ent.Schema
//}
//
//// Fields of the Attendee.
//func (Attendee) Fields() []ent.Field {
//	return []ent.Field{
//		field.String("AttendeeID").Unique(),
//		field.String("EventID"),
//		field.String("CustomUserID").Optional(),
//		field.String("TicketID"),
//		field.Bool("Attended"),
//		field.Time("CreatedAt").Default(time.Now),
//		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
//	}
//}
//
//// Edges of the Attendee.
//func (Attendee) Edges() []ent.Edge {
//	return []ent.Edge{
//		// Edges of the Attendee
//		edge.To("event", Event.Type).
//			Unique(),
//		//edge.To("user", User.Type).
//		//	Unique(),
//		edge.To("ticket", Ticket.Type).
//			Unique(),
//	}
//}

package reviews

import "placio-app/ent"

type RatingDTO struct {
	User     *ent.User
	Place    *ent.Place
	Event    *ent.Event
	Business *ent.Business
	Score    int
}

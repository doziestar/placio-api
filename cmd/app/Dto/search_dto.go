package Dto

import "placio-app/ent"

type SearchResponse struct {
	users *ent.User 
	places *ent.Place
	event *ent.Event
	business *ent.Business
}
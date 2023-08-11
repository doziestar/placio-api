package search

import "placio-app/ent"

type SearchResponse struct {
	users    *ent.User
	places   *ent.Place
	event    *ent.Event
	business *ent.Business
}

type SearchResponses struct {
	Users      []*ent.User     `json:"users,omitempty"`
	Businesses []*ent.Business `json:"businesses,omitempty"`
	Events     []*ent.Event    `json:"events,omitempty"`
	Places     []*ent.Place    `json:"places,omitempty"`
}

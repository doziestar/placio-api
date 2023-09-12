package notifications

import "placio-app/ent"

type Notification struct {
	ID               string            `json:"id"`
	Title            string            `json:"title"`
	Message          string            `json:"message"`
	Link             string            `json:"link"`
	IsRead           bool              `json:"is_read"`
	Type             int               `json:"type"`
	NotificationMeta *notificationMeta `json:"notification_meta"`
	CreatedAt        string            `json:"created_at"`
	UpdatedAt        string            `json:"updated_at"`
}

type notificationMeta struct {
	TriggeredBy trigger   `json:"triggered_by"`
	TriggeredTo triggerTo `json:"triggered_to"`
}

type trigger struct {
	user             ent.User
	business         ent.Business
	post             ent.Post
	comment          ent.Comment
	order            ent.Order
	place            ent.Place
	review           ent.Review
	userFollower     ent.UserFollowUser
	businessFollower ent.UserFollowBusiness
}

type triggerTo struct {
	user     ent.User
	business ent.Business
}

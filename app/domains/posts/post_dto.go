package posts

import (
	"placio-app/ent"
	"time"
)

type PostDto struct {
	ID         string     `json:"id"`
	Content    string     `json:"content"`
	UserID     string     `json:"user_id"`
	BusinessID string     `json:"business_id,omitempty"` // It can be empty if it does not exist
	CreatedAt  time.Time  `json:"created_at"`
	Medias     []MediaDto `json:"medias"`
	Privacy    string     `json:"privacy"`
}

type MediaDto struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type PostResponseDto struct {
	ID        string        `json:"id"`
	Content   string        `json:"content"`
	User      *ent.User     `json:"user"`
	Business  *ent.Business `json:"business,omitempty"` // It can be empty if it does not exist
	CreatedAt time.Time     `json:"created_at"`
	Medias    []MediaDto    `json:"medias"`
}

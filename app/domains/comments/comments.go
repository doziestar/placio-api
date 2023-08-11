package comments

import (
	"placio-app/ent"
	"time"
)

// CommentDto is a Data Transfer Object for comments
type CommentDto struct {
	Content string `json:"content" binding:"required"`
	//PostID  string `json:"postId" binding:"required"`
}

// CommentResponseDto is the response structure for comments
type CommentResponseDto struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	User      *ent.User `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

package comments

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"placio-app/ent"
	"placio-app/ent/comment"
	"time"
)

type CommentService interface {
	CreateComment(ctx context.Context, userID string, postID string, content string) (*ent.Comment, error)
	UpdateComment(ctx context.Context, userID string, commentID string, newContent string) (*ent.Comment, error)
	DeleteComment(ctx context.Context, userID string, commentID string) error
}

type CommentServiceImpl struct {
	client *ent.Client
}

func NewCommentService(client *ent.Client) CommentService {
	return &CommentServiceImpl{client: client}
}

func (cs *CommentServiceImpl) CreateComment(ctx context.Context, userID string, postID string, content string) (*ent.Comment, error) {
	log.Println("about to create a comment")

	user, err := cs.client.User.Get(ctx, userID)
	if ent.IsNotFound(err) {
		return nil, fmt.Errorf("user does not exist: %w", err)
	} else if err != nil {
		return nil, fmt.Errorf("failed retrieving user: %w", err)
	}

	post, err := cs.client.Post.Get(ctx, postID)
	if ent.IsNotFound(err) {
		return nil, fmt.Errorf("post does not exist: %w", err)
	} else if err != nil {
		return nil, fmt.Errorf("failed retrieving post: %w", err)
	}

	comment, err := cs.client.Comment.
		Create().
		SetID(uuid.New().String()).
		SetContent(content).
		SetUser(user).
		SetUpdatedAt(time.Now()).
		SetPost(post).
		Save(ctx)

	if err != nil {
		log.Println("Failed to add comment", err)
		return nil, fmt.Errorf("failed adding comment: %w", err)
	}

	return comment, nil
}

func (cs *CommentServiceImpl) UpdateComment(ctx context.Context, userID string, commentID string, newContent string) (*ent.Comment, error) {
	comment, err := cs.client.Comment.
		Query().
		Where(comment.ID(commentID)).
		WithUser().
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed retrieving comment: %w", err)
	}

	if comment.Edges.User.ID != userID {
		return nil, fmt.Errorf("unauthorized")
	}

	comment, err = cs.client.Comment.
		UpdateOneID(commentID).
		SetContent(newContent).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed updating comment: %w", err)
	}

	return comment, nil
}

func (cs *CommentServiceImpl) DeleteComment(ctx context.Context, userID string, commentID string) error {
	comment, err := cs.client.Comment.
		Query().
		Where(comment.ID(commentID)).
		WithUser().
		Only(ctx)

	if err != nil {
		return fmt.Errorf("failed retrieving comment: %w", err)
	}

	if comment.Edges.User.ID != userID {
		return fmt.Errorf("unauthorized")
	}

	err = cs.client.Comment.DeleteOneID(commentID).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed deleting comment: %w", err)
	}

	return nil
}

package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"placio-app/ent"
	"placio-app/ent/comment"
	"placio-app/ent/post"
	"placio-app/ent/user"
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
	// Check if user exists
	uExists, err := cs.client.User.Query().Where(user.ID(userID)).Exist(ctx)
	if err != nil || !uExists {
		return nil, fmt.Errorf("user does not exist: %w", err)
	}

	// Check if post exists
	pExists, err := cs.client.Post.Query().Where(post.ID(postID)).Exist(ctx)
	if err != nil || !pExists {
		return nil, fmt.Errorf("post does not exist: %w", err)
	}

	user, err := cs.client.User.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving user: %w", err)
	}

	post, err := cs.client.Post.Get(ctx, postID)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving post: %w", err)
	}

	comment, err := cs.client.Comment.
		Create().
		SetID(uuid.New().String()).
		SetContent(content).
		SetUser(user).
		SetPost(post).
		Save(ctx)

	if err != nil {
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

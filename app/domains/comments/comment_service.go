package comments

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"placio-app/domains/notifications"
	"placio-app/ent"
	"placio-app/ent/comment"
	"time"
)

type CommentService interface {
	CreateComment(ctx context.Context, userID string, postID string, content string) (*ent.Comment, error)
	UpdateComment(ctx context.Context, userID string, commentID string, newContent string) (*ent.Comment, error)
	DeleteComment(ctx context.Context, userID string, commentID string) error
	CreateReply(ctx context.Context, userID string, parentCommentID string, content string) (*ent.Comment, error)
}

type CommentServiceImpl struct {
	client        *ent.Client
	notifications notifications.INotification
}

func NewCommentService(client *ent.Client, notifications notifications.INotification) CommentService {
	return &CommentServiceImpl{client: client, notifications: notifications}
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

	// Create a new notification entity
	notification := &ent.Notification{
		Title:          fmt.Sprintf("%s commented on your post", user.Name),
		Message:        content,
		Type:           1,
		TriggeredBy:    user.ID,
		TriggeredTo:    comment.Edges.User.ID,
		NotifiableType: "comment",
		NotifiableID:   comment.ID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Edges: ent.NotificationEdges{
			User:    []*ent.User{user},       // Associate the notification with the user
			Comment: []*ent.Comment{comment}, // Associate the notification with the comment (reply)
		},
	}

	// Attempt to create the notification in the database
	_, err = cs.notifications.CreateNotification(ctx, notification)
	if err != nil {
		// Log the error if notification creation fails
		log.Printf("Failed to create notification: %v", err)
		// Decide how you want to handle the error - you can return it, log it, or ignore it
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

func (cs *CommentServiceImpl) CreateReply(ctx context.Context, userID string, parentCommentID string, content string) (*ent.Comment, error) {
	parentComment, err := cs.client.Comment.Get(ctx, parentCommentID)
	if ent.IsNotFound(err) {
		return nil, fmt.Errorf("parent comment does not exist: %w", err)
	} else if err != nil {
		return nil, fmt.Errorf("failed retrieving parent comment: %w", err)
	}

	user, err := cs.client.User.Get(ctx, userID)
	if ent.IsNotFound(err) {
		return nil, fmt.Errorf("user does not exist: %w", err)
	} else if err != nil {
		return nil, fmt.Errorf("failed retrieving user: %w", err)
	}

	reply, err := cs.client.Comment.
		Create().
		SetID(uuid.New().String()).
		SetContent(content).
		SetUser(user).
		SetUpdatedAt(time.Now()).
		SetParentComment(parentComment).
		Save(ctx)

	if err != nil {
		log.Println("Failed to add comment", err)
		return nil, fmt.Errorf("failed adding comment: %w", err)
	}

	if parentComment.Edges.User != nil && parentComment.Edges.User.ID != userID {
		notificationContent := fmt.Sprintf("%s replied to your comment", user.Name)
		notification := &ent.Notification{
			Title:          fmt.Sprintf("%s replied to your comment", user.Name),
			Message:        notificationContent,
			Type:           1,
			TriggeredBy:    user.ID,
			TriggeredTo:    parentComment.Edges.User.ID,
			NotifiableType: "comment",
			NotifiableID:   parentComment.ID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			Edges: ent.NotificationEdges{
				User: []*ent.User{
					user,
				},
				Comment: []*ent.Comment{
					reply,
				},
			},
		}
		_, err := cs.notifications.CreateNotification(ctx, notification)
		if err != nil {
			log.Printf("Failed to create notification: %v", err)
			// Decide how you want to handle the error - you can return it, log it, or ignore it
		}
	}

	return reply, nil
}

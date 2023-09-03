package posts

import (
	"context"
	"errors"
	"fmt"
	"log"
	"placio-app/ent"
	"placio-app/ent/comment"
	"placio-app/ent/post"
	"placio-app/ent/predicate"
	"placio-app/utility"
	"time"

	"github.com/getsentry/sentry-go"
)

type PostService interface {
	CreatePost(ctx context.Context, newPost *ent.Post, userID string, businessID *string) (*ent.Post, error)
	GetPost(ctx context.Context, postID string) (*ent.Post, error)
	UpdatePost(ctx context.Context, updatedPost *ent.Post) (*ent.Post, error)
	DeletePost(ctx context.Context, postID string) error
	ListPosts(ctx context.Context, page, pageSize int, sortBy []string, filters ...predicate.Post) ([]*ent.Post, error)
	GetPostFeeds(ctx context.Context) ([]*ent.Post, error)

	GetCommentsByPost(ctx context.Context, postID string) ([]*ent.Comment, error)
	GetComments(ctx context.Context, postID string) ([]*ent.Comment, error)
	AddMediaToPost(ctx context.Context, post *ent.Post, media *ent.Media) error

	//LikePost(postID string, userID string) error
	//UnlikePost(postID string, userID string) error
	//GetLikes(postID string, page, pageSize int) ([]*models.Like, error)
}

type PostServiceImpl struct {
	client *ent.Client
	cache  *utility.RedisClient
}

func NewPostService(client *ent.Client, cache *utility.RedisClient) PostService {
	return &PostServiceImpl{client: client, cache: cache}
}

func (ps *PostServiceImpl) GetPostFeeds(ctx context.Context) ([]*ent.Post, error) {
	//Get All Posts
	posts, err := ps.client.Post.
		Query().
		WithComments().
		WithLikes().
		WithUser().
		All(ctx)

	if err != nil {
		sentry.CaptureException(err)
		return nil, fmt.Errorf("failed getting posts: %w", err)
	}

	// Return the posts
	return posts, nil
}

func (ps *PostServiceImpl) CreatePost(ctx context.Context, newPost *ent.Post, userID string, businessID *string) (*ent.Post, error) {
	if newPost == nil {
		return nil, errors.New("post cannot be nil")
	}

	// Create builder
	postBuilder := ps.client.Post.
		Create().
		SetID(newPost.ID).
		SetContent(newPost.Content).SetUpdatedAt(time.Now())

	// Associate with user
	postBuilder = postBuilder.SetUserID(userID)
	// Associate with business, if business ID is provided
	if businessID != nil {
		postBuilder = postBuilder.SetBusinessAccountID(*businessID)
	}

	fmt.Println("saving post postBuilder", postBuilder)
	// Save post
	post, err := postBuilder.Save(ctx)
	log.Printf("post: %v", post)
	log.Printf("err: %v", err)
	if err != nil {
		fmt.Errorf("failed creating post: %w", err)
		return nil, fmt.Errorf("failed creating post: %w", err)
	}
	postToReturn, err := ps.GetPost(ctx, post.ID)
	if err != nil {
		return nil, fmt.Errorf("failed getting post: %w", err)
	}
	fmt.Println("saved post", post)

	return postToReturn, nil
}

func (ps *PostServiceImpl) AddMediaToPost(ctx context.Context, post *ent.Post, media *ent.Media) error {
	// Use the UpdateOneID method to find the post by its ID, then add the media using the AddMedias method, and finally call Save to apply the update.
	_, err := ps.client.Post.
		UpdateOneID(post.ID).
		AddMedias(media).
		Save(ctx)

	// Check if an error occurred and return it if so.
	if err != nil {
		return fmt.Errorf("failed adding media to post: %w", err)
	}

	// If no error occurred, return nil.
	return nil
}

func (ps *PostServiceImpl) GetPost(ctx context.Context, postID string) (*ent.Post, error) {
	fmt.Println("getting post", postID)

	post, err := ps.client.Post.
		Query().
		Where(post.ID(postID)).
		WithComments().
		WithLikes().
		WithUser().
		WithBusinessAccount().
		WithMedias().
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed getting post: %w", err)
	}

	return post, nil
}

func (ps *PostServiceImpl) UpdatePost(ctx context.Context, updatedPost *ent.Post) (*ent.Post, error) {
	if updatedPost == nil || updatedPost.Content == "" {
		return nil, errors.New("post or content cannot be nil or empty")
	}

	post, err := ps.client.Post.
		UpdateOneID(updatedPost.ID).
		SetContent(updatedPost.Content).
		Save(ctx)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed updating post: %w", err)
	}

	return post, nil
}

func (ps *PostServiceImpl) DeletePost(ctx context.Context, postID string) error {
	err := ps.client.Post.
		DeleteOneID(postID).
		Exec(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("post not found: %w", err)
		}
		return fmt.Errorf("failed deleting post: %w", err)
	}

	return nil
}

func (ps *PostServiceImpl) ListPosts(ctx context.Context, page, pageSize int, sortBy []string, filters ...predicate.Post) ([]*ent.Post, error) {
	offset := (page - 1) * pageSize

	posts, err := ps.client.Post.
		Query().
		Where(filters...).
		Order(ent.Asc(sortBy...)).
		Offset(offset).
		Limit(pageSize).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed listing posts: %w", err)
	}

	return posts, nil
}

func (ps *PostServiceImpl) GetComments(ctx context.Context, postID string) ([]*ent.Comment, error) {
	pExists, err := ps.client.Post.Query().Where(post.ID(postID)).Exist(ctx)
	if err != nil || !pExists {
		return nil, fmt.Errorf("post does not exist: %w", err)
	}

	comments, err := ps.client.Post.
		Query().
		Where(post.ID(postID)).
		QueryComments().
		WithUser().
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed getting comments: %w", err)
	}

	return comments, nil
}

func (ps *PostServiceImpl) GetCommentsByPost(ctx context.Context, postID string) ([]*ent.Comment, error) {
	// Check if post exists
	pExists, err := ps.client.Post.Query().Where(post.ID(postID)).Exist(ctx)
	if err != nil || !pExists {
		return nil, fmt.Errorf("post does not exist: %w", err)
	}

	// Query comments related to the post
	comments, err := ps.client.Comment.
		Query().
		Where(comment.HasPostWith(post.ID(postID))).
		WithUser().
		WithPost(func(pq *ent.PostQuery) {
			pq.WithComments().
				WithUser().
				WithMedias().
				WithLikes()
		}).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed getting comments: %w", err)
	}

	return comments, nil
}

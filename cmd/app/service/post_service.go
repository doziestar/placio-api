package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"placio-app/ent"
	"placio-app/ent/post"
	"placio-app/ent/predicate"
	"placio-app/utility"
	"time"
)

type PostService interface {
	CreatePost(ctx context.Context, newPost *ent.Post, userID string, businessID *string) (*ent.Post, error)
	GetPost(ctx context.Context, postID string) (*ent.Post, error)
	UpdatePost(ctx context.Context, updatedPost *ent.Post) (*ent.Post, error)
	DeletePost(ctx context.Context, postID string) error
	ListPosts(ctx context.Context, page, pageSize int, sortBy []string, filters ...predicate.Post) ([]*ent.Post, error)

	CreateComment(ctx context.Context, postID string, newComment *ent.Comment) (*ent.Comment, error)
	UpdateComment(ctx context.Context, updatedComment *ent.Comment) (*ent.Comment, error)
	DeleteComment(ctx context.Context, commentID string) error
	GetComments(ctx context.Context, postID string, page, pageSize int, sortBy []string, filters ...predicate.Comment) ([]*ent.Comment, error)

	//LikePost(postID string, userID string) error
	//UnlikePost(postID string, userID string) error
	//GetLikes(postID string, page, pageSize int) ([]*models.Like, error)
}

type PostServiceImpl struct {
	client   *ent.Client
	user     *ent.User
	post     *ent.Post
	comment  *ent.Comment
	like     *ent.Like
	business *ent.Business
	cache    *utility.RedisClient
}

func NewPostService(client *ent.Client, cache *utility.RedisClient) PostService {
	return &PostServiceImpl{client: client, cache: cache, user: &ent.User{}, post: &ent.Post{}, comment: &ent.Comment{}, like: &ent.Like{}, business: &ent.Business{}}
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
	fmt.Println("saved post", post)

	return post, nil
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
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed getting post: %w", err)
	}

	return post, nil
}

func (ps *PostServiceImpl) UpdatePost(ctx context.Context, updatedPost *ent.Post) (*ent.Post, error) {
	if updatedPost == nil {
		return nil, errors.New("post cannot be nil")
	}

	post, err := ps.client.Post.
		UpdateOneID(updatedPost.ID).
		SetContent(updatedPost.Content).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed updating post: %w", err)
	}

	return post, nil
}

func (ps *PostServiceImpl) DeletePost(ctx context.Context, postID string) error {
	err := ps.client.Post.
		DeleteOneID(postID).
		Exec(ctx)

	if err != nil {
		return fmt.Errorf("failed deleting post: %w", err)
	}

	return nil
}

func (ps *PostServiceImpl) CreateComment(ctx context.Context, postID string, newComment *ent.Comment) (*ent.Comment, error) {
	comment, err := ps.client.Comment.
		Create().
		SetContent(newComment.Content).
		SetUser(newComment.Edges.User).
		SetPostID(postID).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating comment: %w", err)
	}

	return comment, nil
}

func (ps *PostServiceImpl) UpdateComment(ctx context.Context, updatedComment *ent.Comment) (*ent.Comment, error) {
	if updatedComment == nil {
		return nil, errors.New("comment cannot be nil")
	}

	comment, err := ps.client.Comment.
		UpdateOneID(updatedComment.ID).
		SetContent(updatedComment.Content).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed updating comment: %w", err)
	}

	return comment, nil
}

func (ps *PostServiceImpl) DeleteComment(ctx context.Context, commentID string) error {
	err := ps.client.Comment.
		DeleteOneID(commentID).
		Exec(ctx)

	if err != nil {
		return fmt.Errorf("failed deleting comment: %w", err)
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

func (ps *PostServiceImpl) GetComments(ctx context.Context, postID string, page, pageSize int, sortBy []string, filters ...predicate.Comment) ([]*ent.Comment, error) {
	offset := (page - 1) * pageSize

	comments, err := ps.client.Post.
		Query().
		Where(post.IDEQ(postID)).
		QueryComments().
		Where(filters...).
		Order(ent.Asc(sortBy...)).
		Offset(offset).
		Limit(pageSize).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed getting comments: %w", err)
	}

	return comments, nil
}

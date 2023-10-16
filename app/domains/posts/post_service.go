package posts

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log"
	"placio-app/domains/media"
	"placio-app/ent"
	"placio-app/ent/comment"
	"placio-app/ent/post"
	"placio-app/ent/predicate"
	"placio-app/utility"
	"placio-pkg/kafka"
	"sync"
	"time"

	"github.com/getsentry/sentry-go"
)

type PostService interface {
	CreatePost(ctx context.Context, newPost *ent.Post, userID string, businessID string, media []MediaDto) (*ent.Post, error)
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
	client       *ent.Client
	cache        *utility.RedisClient
	mediaService media.MediaService
	producer     *kafka.Producer
}

func NewPostService(client *ent.Client, cache *utility.RedisClient, mediaService media.MediaService, producer *kafka.Producer) PostService {
	return &PostServiceImpl{client: client, cache: cache, mediaService: mediaService, producer: producer}
}

func (ps *PostServiceImpl) GetPostFeeds(ctx context.Context) ([]*ent.Post, error) {
	log.Println("Getting post feeds from DB...", ctx.Value("user"))
	md, ok := metadata.FromIncomingContext(ctx)

	var userId string

	if ok {
		userId = md.Get("user")[0]
	}

	userId = ctx.Value("user").(string)

	//Get All Posts
	posts, err := ps.client.Post.
		Query().
		WithComments(func(query *ent.CommentQuery) {
			query.WithUser()
			query.WithReplies(func(query *ent.CommentQuery) {
				query.WithUser()
			}).WithParentComment(func(query *ent.CommentQuery) {
				query.WithUser()
			})
		}).
		WithMedias().
		WithLikes(func(query *ent.LikeQuery) {
			query.WithUser()
		}).
		WithUser().
		All(ctx)

	var wg sync.WaitGroup

	for _, post := range posts {
		for _, like := range post.Edges.Likes {
			if like.Edges.User.ID == userId {
				post.LikedByMe = true
			}
		}
		wg.Add(1)
		go func(post *ent.Post) {
			defer wg.Done()
			post.LikeCount = len(post.Edges.Likes)
			post.CommentCount = len(post.Edges.Comments)
		}(post)
	}
	wg.Wait()

	if err != nil {
		sentry.CaptureException(err)
		return nil, fmt.Errorf("failed getting posts: %w", err)
	}

	// Return the posts
	return posts, nil
}

func (ps *PostServiceImpl) CreatePost(ctx context.Context, newPost *ent.Post, userID string, businessID string, medias []MediaDto) (*ent.Post, error) {
	if newPost == nil {
		return nil, errors.New("post cannot be nil")
	}

	// Create builder
	postBuilder := ps.client.Post.
		Create().
		SetID(newPost.ID).
		SetContent(newPost.Content).
		SetUpdatedAt(time.Now()).
		SetUserID(userID)

	// Associate with business, if business ID is provided
	if businessID != "" {
		postBuilder = postBuilder.SetBusinessAccountID(businessID)
	}

	// Save post
	post, err := postBuilder.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating post: %w", err)
	}
	log.Printf("post saved: %v", post)

	for _, mediaDto := range medias {
		createdMedia, err := ps.mediaService.CreateMedia(ctx, mediaDto.URL, mediaDto.Type)
		if err != nil {
			return nil, fmt.Errorf("failed creating media: %w", err)
		}

		err = ps.AddMediaToPost(ctx, newPost, createdMedia)
		if err != nil {
			return nil, fmt.Errorf("failed adding media to post: %w", err)
		}
	}

	postToReturn, err := ps.GetPost(ctx, post.ID)
	if err != nil {
		return nil, fmt.Errorf("failed getting post: %w", err)
	}

	// Publish post created event
	postBytes, err := json.Marshal(postToReturn)
	if err != nil {
		log.Println("error serializing post:", err)
		return postToReturn, nil
	}

	err = ps.producer.PublishMessage(ctx, []byte(postToReturn.ID), postBytes)
	if err != nil {
		log.Println("error sending post to Kafka:", err)
	} else {
		log.Println("published post:created event to Kafka")
	}

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
		WithComments(func(query *ent.CommentQuery) {
			query.WithUser()
			query.WithReplies(func(query *ent.CommentQuery) {
				query.WithUser()
			}).WithParentComment(func(query *ent.CommentQuery) {
				query.WithUser()
			})
		}).
		WithLikes(func(query *ent.LikeQuery) {
			query.WithUser()
		}).
		WithUser().
		WithBusinessAccount().
		WithMedias().
		Only(ctx)

	if user, ok := ctx.Value("user").(string); ok {
		userId := user

		for _, like := range post.Edges.Likes {
			if like.Edges.User.ID == userId {
				post.LikedByMe = true
			}
		}

		post.LikeCount = len(post.Edges.Likes)
		post.CommentCount = len(post.Edges.Comments)
	}

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

	postData, err := ps.GetPost(ctx, post.ID)
	if err != nil {
		return nil, fmt.Errorf("failed getting post: %w", err)
	}

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed updating post: %w", err)
	}

	return postData, nil
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
		WithReplies(func(query *ent.CommentQuery) {
			query.WithUser()
		}).WithParentComment(func(query *ent.CommentQuery) {
		query.WithUser()
	}).
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
		WithReplies(func(query *ent.CommentQuery) {
			query.WithUser()
		}).WithParentComment(func(query *ent.CommentQuery) {
		query.WithUser()
	}).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed getting comments: %w", err)
	}

	return comments, nil
}

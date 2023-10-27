package posts

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"log"
	"mime/multipart"
	"placio-app/domains/media"
	"placio-app/ent"
	"placio-app/ent/comment"
	"placio-app/ent/post"
	"placio-app/ent/predicate"
	"placio-app/utility"
	"placio-pkg/kafka"
	"time"

	"github.com/getsentry/sentry-go"
)

type PostService interface {
	CreatePost(ctx context.Context, newPost *ent.Post, userID string, businessID string, mediaFiles []*multipart.FileHeader) (*ent.Post, error)
	GetPost(ctx context.Context, postID string) (*ent.Post, error)
	Repost(ctx context.Context, originalPostID, content, userID, businessID string) (*ent.Post, error)
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
	} else {
		userId = ctx.Value("user").(string)
	}

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
		WithBusinessAccount().
		WithUser().
		WithOriginalPost(func(query *ent.PostQuery) {
			query.WithUser()
			query.WithMedias()
			query.WithLikes()
			query.WithComments(func(query *ent.CommentQuery) {
				query.WithUser()
				query.WithReplies(func(query *ent.CommentQuery) {
					query.WithUser()
				}).WithParentComment(func(query *ent.CommentQuery) {
					query.WithUser()
				})
			})
			query.WithBusinessAccount()
		}).
		All(ctx)

	for _, post := range posts {
		post.LikeCount = len(post.Edges.Likes)
		post.CommentCount = len(post.Edges.Comments)

		for _, like := range post.Edges.Likes {
			if like.Edges.User.ID == userId {
				post.LikedByMe = true
				break
			}
		}
	}

	if err != nil {
		sentry.CaptureException(err)
		return nil, fmt.Errorf("failed getting posts: %w", err)
	}

	// Return the posts
	return posts, nil
}

func (ps *PostServiceImpl) Repost(ctx context.Context, originalPostID, content, userID, businessID string) (*ent.Post, error) {
	if originalPostID == "" {
		return nil, errors.New("original post ID, content, and user ID must be provided")
	}

	tx, err := ps.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed starting a transaction: %w", err)
	}

	originalPost, err := tx.Post.
		Query().
		Where(post.IDEQ(originalPostID)).
		Only(ctx)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	// Increment repost count of original post
	_, err = tx.Post.
		UpdateOne(originalPost).
		AddRepostCount(1).
		Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("failed incrementing repost count: %w", err)
	}

	// Create a new post as a repost
	repostBuilder := tx.Post.
		Create().
		SetID(uuid.New().String()).
		SetContent(content).
		SetUpdatedAt(time.Now()).
		SetIsRepost(true).
		AddOriginalPost(originalPost)

	// Associate with business account if business ID is provided
	if businessID != "" {
		repostBuilder.SetBusinessAccountID(businessID)
	}

	if content == "" {
		repostBuilder.SetContent(originalPost.Content)
	}

	if userID == "" {
		repostBuilder.SetUserID(userID)
	}

	// Save post
	repost, err := repostBuilder.Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("failed creating repost: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed committing transaction: %w", err)
	}

	log.Printf("repost saved: %v", repost)
	return repost, nil
}

func (ps *PostServiceImpl) CreatePost(ctx context.Context, newPost *ent.Post, userID string, businessID string, mediaFiles []*multipart.FileHeader) (*ent.Post, error) {
	if newPost == nil {
		return nil, errors.New("post cannot be nil")
	}

	// Create builder
	postBuilder := ps.client.Post.
		Create().
		SetID(uuid.New().String()).
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

	if len(mediaFiles) > 0 {
		go func() {
			medias, err := ps.mediaService.UploadAndCreateMedia(ctx, mediaFiles)
			if err != nil {
				log.Printf("failed uploading and creating media: %v", err)
				return
			}

			_, err = ps.client.Post.
				UpdateOneID(post.ID).
				AddMedias(medias...).
				Save(ctx)
			if err != nil {
				log.Printf("failed adding media to post: %v", err)
				return
			}

		}()
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
		WithOriginalPost(func(query *ent.PostQuery) {
			query.WithUser()
			query.WithMedias()
			query.WithLikes()
			query.WithComments(func(query *ent.CommentQuery) {
				query.WithUser()
				query.WithReplies(func(query *ent.CommentQuery) {
					query.WithUser()
				}).WithParentComment(func(query *ent.CommentQuery) {
					query.WithUser()
				})
			})
			query.WithBusinessAccount()
		}).
		Only(ctx)

	if user, ok := ctx.Value("user").(string); ok {
		userId := user

		for _, like := range post.Edges.Likes {
			if like.Edges.User.ID == userId {
				post.LikedByMe = true
				break
			}
		}
	}

	post.LikeCount = len(post.Edges.Likes)
	post.CommentCount = len(post.Edges.Comments)

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

package service

import (
	"context"
	"github.com/google/uuid"
	"placio-app/ent"
	"placio-app/ent/like"
	"placio-app/ent/post"
	"placio-app/ent/user"
	"placio-app/utility"
)

type LikeService interface {
	LikePost(ctx context.Context, userID string, postID string) (*ent.Like, error)
	UnlikePost(ctx context.Context, likeID string) error
	GetUserLikes(ctx context.Context, userID string) ([]*ent.Like, error)
	GetPostLikes(ctx context.Context, postID string) ([]*ent.Like, error)
}

type LikeServiceImpl struct {
	client *ent.Client
	cache  *utility.RedisClient
}

func NewLikeService(client *ent.Client, cache *utility.RedisClient) *LikeServiceImpl {
	return &LikeServiceImpl{client: client, cache: cache}
}

func (s *LikeServiceImpl) LikePost(ctx context.Context, userID string, postID string) (*ent.Like, error) {
	user, err := s.client.User.
		Query().
		Where(user.ID(userID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	post, err := s.client.Post.
		Query().
		Where(post.ID(postID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return s.client.Like.
		Create().
		SetID(uuid.New().String()).
		SetUser(user).
		SetPost(post).
		Save(ctx)
}

func (s *LikeServiceImpl) UnlikePost(ctx context.Context, likeID string) error {
	_, err := s.client.Like.
		Delete().
		Where(like.ID(likeID)).
		Exec(ctx)
	if err != nil {
		return err
	}
	//if like == nil {
	//	return nil
	//}
	return nil
}

func (s *LikeServiceImpl) GetUserLikes(ctx context.Context, userID string) ([]*ent.Like, error) {
	return s.client.User.
		Query().
		Where(user.ID(userID)).
		QueryLikes().
		All(ctx)
}

func (s *LikeServiceImpl) GetPostLikes(ctx context.Context, postID string) ([]*ent.Like, error) {
	return s.client.Post.
		Query().
		Where(post.ID(postID)).
		QueryLikes().
		All(ctx)
}

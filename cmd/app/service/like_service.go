package service

import (
	"context"
	"github.com/google/uuid"
	"placio-app/ent"
	"placio-app/ent/like"
	"placio-app/ent/place"
	"placio-app/ent/post"
	"placio-app/ent/user"
	"placio-app/ent/userlikeplace"
	"placio-app/utility"
	"time"
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

type UserLikePlaceService interface {
	LikePlace(ctx context.Context, userID string, placeID string) (*ent.UserLikePlace, error)
	UnlikePlace(ctx context.Context, userLikePlaceID string) error
	GetUserLikedPlaces(ctx context.Context, userID string) ([]*ent.UserLikePlace, error)
	GetPlaceLikes(ctx context.Context, placeID string) ([]*ent.UserLikePlace, error)
	CheckIfUserLikesPlace(ctx context.Context, userID string, placeID string) (bool, error)
}

type UserLikePlaceServiceImpl struct {
	client *ent.Client
	cache  *utility.RedisClient
}

func NewUserLikePlaceService(client *ent.Client, cache *utility.RedisClient) *UserLikePlaceServiceImpl {
	return &UserLikePlaceServiceImpl{client: client, cache: cache}
}

func (s *UserLikePlaceServiceImpl) CheckIfUserLikesPlace(ctx context.Context, userID, placeID string) (bool, error) {
	count, err := s.client.UserLikePlace.
		Query().
		Where(userlikeplace.HasUserWith(user.ID(userID))).
		Where(userlikeplace.HasPlaceWith(place.ID(placeID))).
		Count(ctx)
	if err != nil {
		return false, err
	}

	return count > 0, nil
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

func (s *UserLikePlaceServiceImpl) LikePlace(ctx context.Context, userID string, placeID string) (*ent.UserLikePlace, error) {
	user, err := s.client.User.
		Query().
		Where(user.ID(userID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	place, err := s.client.Place.
		Query().
		Where(place.ID(placeID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	userLike, err := s.client.UserLikePlace.
		Create().
		SetID(uuid.New().String()).
		SetUser(user).
		SetPlace(place).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return userLike, nil
}

func (s *UserLikePlaceServiceImpl) UnlikePlace(ctx context.Context, userLikePlaceID string) error {
	_, err := s.client.UserLikePlace.
		Delete().
		Where(userlikeplace.ID(userLikePlaceID)).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserLikePlaceServiceImpl) GetUserLikedPlaces(ctx context.Context, userID string) ([]*ent.UserLikePlace, error) {
	return s.client.UserLikePlace.
		Query().
		Where(userlikeplace.HasUserWith(user.ID(userID))).
		All(ctx)
}

func (s *UserLikePlaceServiceImpl) GetPlaceLikes(ctx context.Context, placeID string) ([]*ent.UserLikePlace, error) {
	return s.client.UserLikePlace.
		Query().
		Where(userlikeplace.HasPlaceWith(place.ID(placeID))).
		All(ctx)
}

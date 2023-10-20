package like

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"log"
	"placio-app/domains/cache"
	"placio-app/domains/search"
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
	UnlikePlace(ctx context.Context, userId, placeID string) error
	GetUserLikedPlaces(ctx context.Context, userID string) ([]*ent.UserLikePlace, error)
	GetPlaceLikes(ctx context.Context, placeID string) ([]*ent.UserLikePlace, error)
	CheckIfUserLikesPlace(ctx context.Context, userID string, placeID string) (bool, error)
}

type UserLikePlaceServiceImpl struct {
	client        *ent.Client
	cache         *utility.RedisClient
	searchService search.SearchService
	cacheService  cache.ICacheService
}

func NewUserLikePlaceService(client *ent.Client, cache *utility.RedisClient, searchService search.SearchService, cacheService cache.ICacheService) *UserLikePlaceServiceImpl {
	return &UserLikePlaceServiceImpl{client: client, cache: cache, searchService: searchService, cacheService: cacheService}
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
	log.Println("LikePost called ", userID, postID)

	// Check if the user already liked the post
	existingLike, err := s.client.Like.
		Query().
		Where(like.HasUserWith(user.ID(userID))).
		Where(like.HasPostWith(post.ID(postID))).
		Only(ctx)

	if err == nil && existingLike != nil {
		// If a like exists, delete it to 'unlike' the post
		err = s.client.Like.DeleteOne(existingLike).Exec(ctx)
		if err != nil {
			log.Println("Error unliking post: ", err)
			return nil, err
		}
		return nil, nil // Returning nil since no like was created
	} else if !ent.IsNotFound(err) {
		// If it's an error other than 'not found', return the error
		log.Println("Error checking existing like: ", err)
		return nil, err
	}

	// If no existing like, create a new like
	likes, err := s.client.Like.
		Create().
		SetID(uuid.New().String()).
		SetUserID(userID).
		SetPostID(postID).
		SetLike(true).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		log.Println("Error creating like: ", err)
		return nil, err
	}

	return likes, nil
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

	// Check if the UserLikePlace relation already exists
	existingUserLikePlaces, err := s.client.UserLikePlace.
		Query().
		Where(userlikeplace.HasUserWith(user.ID(userID))).
		Where(userlikeplace.HasPlaceWith(place.ID(placeID))).
		All(ctx)
	if err != nil {
		log.Println("Failed to query UserLikePlace:", err)
		return nil, err
	}

	if len(existingUserLikePlaces) > 0 {
		log.Println("UserLikePlace already exists for user and place:", userID, placeID)
		return nil, errors.New("you already like this place")
	}

	userData, err := s.client.User.
		Query().
		Where(user.ID(userID)).
		Only(ctx)
	if err != nil {
		log.Println("Failed to query user:", userID)
		return nil, err
	}

	placeData, err := s.client.Place.
		Query().
		Where(place.ID(placeID)).
		Only(ctx)
	if err != nil {
		log.Println("Failed to query place:", placeID)
		return nil, err
	}

	userLike, err := s.client.UserLikePlace.
		Create().
		SetID(uuid.New().String()).
		SetUser(userData).
		SetPlace(placeData).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		log.Println("Failed to create UserLikePlace:", err)
		return nil, err
	}

	//go s.cacheService.AddPlaceToCacheAndSearchIndex(ctx, placeData)

	log.Println("Successfully created UserLikePlace for user and place:", userID, placeID)
	return userLike, nil
}

func (s *UserLikePlaceServiceImpl) UnlikePlace(ctx context.Context, userId, placeID string) error {
	userLikePlaces, err := s.client.UserLikePlace.
		Query().
		Where(userlikeplace.HasUserWith(user.ID(userId))).
		Where(userlikeplace.HasPlaceWith(place.ID(placeID))).
		All(ctx)
	if err != nil {
		return err
	}

	if len(userLikePlaces) == 0 {
		return errors.New("you don't like this place")
	}

	// Loop through all userLikePlaces and delete them
	for _, userLikePlace := range userLikePlaces {
		_, err = s.client.UserLikePlace.
			Delete().
			Where(userlikeplace.ID(userLikePlace.ID)).
			Exec(ctx)
		if err != nil {
			return err
		}
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

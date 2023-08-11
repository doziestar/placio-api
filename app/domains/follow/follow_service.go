package follow

import (
	"context"
	"github.com/google/uuid"
	"placio-app/ent"
	"placio-app/ent/business"
	"placio-app/ent/event"
	"placio-app/ent/place"
	"placio-app/ent/user"
	"placio-app/ent/userfollowbusiness"
	"placio-app/ent/userfollowevent"
	"placio-app/ent/userfollowplace"
	"placio-app/ent/userfollowuser"
	"time"
)

type IFollowService interface {
	FollowUserToBusiness(ctx context.Context, userID, businessID string) error
	UnfollowUserToBusiness(ctx context.Context, userID, businessID string) error
	GetFollowedBusinessesByUser(ctx context.Context, userID string) ([]*ent.Business, error)

	FollowUserToUser(ctx context.Context, followerID, followedID string) error
	UnfollowUserToUser(ctx context.Context, followerID, followedID string) error
	GetFollowedUsersByUser(ctx context.Context, userID string) ([]*ent.User, error)

	FollowUserToPlace(ctx context.Context, userID, placeID string) error
	UnfollowUserToPlace(ctx context.Context, userID, placeID string) error
	GetFollowedPlacesByUser(ctx context.Context, userID string) ([]*ent.Place, error)

	FollowUserToEvent(ctx context.Context, userID, eventID string) error
	UnfollowUserToEvent(ctx context.Context, userID, eventID string) error
	GetFollowedEventsByUser(ctx context.Context, userID string) ([]*ent.Event, error)

	CheckIfUserFollowsBusiness(ctx context.Context, userID, businessID string) (bool, error)
	CheckIfUserFollowsUser(ctx context.Context, followerID, followedID string) (bool, error)
	CheckIfUserFollowsPlace(ctx context.Context, userID, placeID string) (bool, error)
	CheckIfUserFollowsEvent(ctx context.Context, userID, eventID string) (bool, error)
}

type FollowService struct {
	client *ent.Client
}

func NewFollowService(client *ent.Client) *FollowService {
	return &FollowService{client: client}
}

// FollowUserToBusiness User-Business methods
func (s *FollowService) FollowUserToBusiness(ctx context.Context, userID, businessID string) error {
	_, err := s.client.UserFollowBusiness.
		Create().
		SetID(uuid.New().String()).
		SetUserID(userID).
		SetBusinessID(businessID).
		Save(ctx)

	return err
}

func (s *FollowService) UnfollowUserToBusiness(ctx context.Context, userID, businessID string) error {
	_, err := s.client.UserFollowBusiness.
		Delete().
		Where(userfollowbusiness.HasUserWith(user.ID(userID)), userfollowbusiness.HasBusinessWith(business.ID(businessID))).
		Exec(ctx)

	return err
}

func (s *FollowService) GetFollowedBusinessesByUser(ctx context.Context, userID string) ([]*ent.Business, error) {
	businesses, err := s.client.User.
		Query().
		Where(user.ID(userID)).
		QueryFollowedBusinesses().
		QueryBusiness().
		All(ctx)

	return businesses, err
}

// FollowUserToUser User-User methods
func (s *FollowService) FollowUserToUser(ctx context.Context, followerID, followedID string) error {
	_, err := s.client.UserFollowUser.
		Create().
		SetID(uuid.New().String()).
		SetFollowerID(followerID).
		SetFollowedID(followedID).
		Save(ctx)

	return err
}

func (s *FollowService) UnfollowUserToUser(ctx context.Context, followerID, followedID string) error {
	_, err := s.client.UserFollowUser.
		Delete().
		Where(userfollowuser.HasFollowerWith(user.ID(followerID)), userfollowuser.HasFollowedWith(user.ID(followedID))).
		Exec(ctx)

	return err
}

func (s *FollowService) GetFollowedUsersByUser(ctx context.Context, userID string) ([]*ent.User, error) {
	users, err := s.client.User.
		Query().
		Where(user.ID(userID)).
		QueryFollowedUsers().
		QueryFollowed().
		All(ctx)

	return users, err
}

// FollowUserToPlace User-Place methods
func (s *FollowService) FollowUserToPlace(ctx context.Context, userID, placeID string) error {

	_, err := s.client.UserFollowPlace.
		Create().
		SetID(uuid.New().String()).
		SetUserID(userID).
		SetPlaceID(placeID).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return err
	}

	placeToUpdate, err := s.client.Place.Get(ctx, placeID)
	if err != nil {
		return err
	}
	_, err = s.client.Place.
		UpdateOne(placeToUpdate).
		SetFollowerCount(placeToUpdate.FollowerCount + 1).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *FollowService) UnfollowUserToPlace(ctx context.Context, userID, placeID string) error {
	_, err := s.client.UserFollowPlace.
		Delete().
		Where(userfollowplace.HasUserWith(user.ID(userID)), userfollowplace.HasPlaceWith(place.ID(placeID))).
		Exec(ctx)

	return err
}

func (s *FollowService) GetFollowedPlacesByUser(ctx context.Context, userID string) ([]*ent.Place, error) {
	places, err := s.client.User.
		Query().
		Where(user.ID(userID)).
		QueryFollowedPlaces().
		QueryPlace().
		All(ctx)

	return places, err
}

// FollowUserToEvent User-Event methods
func (s *FollowService) FollowUserToEvent(ctx context.Context, userID, eventID string) error {
	_, err := s.client.UserFollowEvent.
		Create().
		SetUserID(userID).
		SetEventID(eventID).
		Save(ctx)

	return err
}

func (s *FollowService) UnfollowUserToEvent(ctx context.Context, userID, eventID string) error {
	_, err := s.client.UserFollowEvent.
		Delete().
		Where(userfollowevent.HasUserWith(user.ID(userID)), userfollowevent.HasEventWith(event.ID(eventID))).
		Exec(ctx)

	return err
}

func (s *FollowService) GetFollowedEventsByUser(ctx context.Context, userID string) ([]*ent.Event, error) {
	events, err := s.client.User.
		Query().
		Where(user.ID(userID)).
		QueryUserFollowEvents().
		QueryEvent().
		All(ctx)

	return events, err
}

func (s *FollowService) CheckIfUserFollowsBusiness(ctx context.Context, userID, businessID string) (bool, error) {
	exists, err := s.client.UserFollowBusiness.
		Query().
		Where(
			userfollowbusiness.HasUserWith(user.ID(userID)),
			userfollowbusiness.HasBusinessWith(business.ID(businessID)),
		).
		Exist(ctx)

	if err != nil {
		return false, err
	}
	return exists, nil
}

func (s *FollowService) CheckIfUserFollowsUser(ctx context.Context, followerID, followedID string) (bool, error) {
	exists, err := s.client.UserFollowUser.
		Query().
		Where(
			userfollowuser.HasFollowerWith(user.ID(followerID)),
			userfollowuser.HasFollowedWith(user.ID(followedID)),
		).
		Exist(ctx)

	if err != nil {
		return false, err
	}
	return exists, nil
}

func (s *FollowService) CheckIfUserFollowsPlace(ctx context.Context, userID, placeID string) (bool, error) {
	exists, err := s.client.UserFollowPlace.
		Query().
		Where(
			userfollowplace.HasUserWith(user.ID(userID)),
			userfollowplace.HasPlaceWith(place.ID(placeID)),
		).
		Exist(ctx)

	if err != nil {
		return false, err
	}
	return exists, nil
}

func (s *FollowService) CheckIfUserFollowsEvent(ctx context.Context, userID, eventID string) (bool, error) {
	exists, err := s.client.UserFollowEvent.
		Query().
		Where(
			userfollowevent.HasUserWith(user.ID(userID)),
			userfollowevent.HasEventWith(event.ID(eventID)),
		).
		Exist(ctx)

	if err != nil {
		return false, err
	}
	return exists, nil
}

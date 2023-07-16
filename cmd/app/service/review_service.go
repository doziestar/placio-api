package service

import (
	"context"
	"github.com/google/uuid"
	"placio-app/ent"
	"placio-app/ent/business"
	"placio-app/ent/event"
	"placio-app/ent/place"
	"placio-app/ent/review"
	"placio-app/ent/user"
	"placio-app/errors"
	"time"
)

// ReviewService represents the contract for your review related operations.
type ReviewService interface {
	RatePlace(placeID, userID string, score float64, content string) error
	RateEvent(eventID, userID string, score float64, content string) error
	RateBusiness(businessID, userID string, score float64, content string) error
	RemoveReview(reviewID, userID string) error
	GetReviewByID(reviewID string) (*ent.Review, error)
	GetReviewsByUserID(userID string) ([]*ent.Review, error)
	GetReviewsByPlaceID(placeID string) ([]*ent.Review, error)
	GetReviewsByEventID(eventID string) ([]*ent.Review, error)
	GetReviewsByBusinessID(businessID string) ([]*ent.Review, error)
	GetReviewsByScore(score float64) ([]*ent.Review, error)
	GetReviewsInDateRange(startDate, endDate time.Time) ([]*ent.Review, error)
	LikeReview(reviewID, userID string) error
	DislikeReview(reviewID, userID string) error
	UpdateReviewContent(reviewID, userID, newContent string) error
	AddMediaToReview(reviewID string, media *ent.Media) error
	// GetReviewsByLikeCount FlagReview(reviewID, userID string) error
	//AddResponseToReview(reviewID, userID, response string) error
	GetReviewsByLikeCount() ([]*ent.Review, error)
	GetReviewsByDislikeCount() ([]*ent.Review, error)
}

type ReviewServiceImpl struct {
	client *ent.Client
}

func NewReviewService(client *ent.Client) ReviewService {
	return &ReviewServiceImpl{client: client}
}

// RatePlace allows a user to rate and review a place.
func (rs *ReviewServiceImpl) RatePlace(placeID, userID string, score float64, content string) error {
	placeData, err := rs.client.Place.Get(context.Background(), placeID)
	if err != nil {
		return err
	}

	userData, _ := rs.client.User.Get(context.Background(), userID)

	// Check if user has already rated this place
	_, err = rs.client.Review.Query().
		Where(review.HasPlaceWith(place.HasUsersWith(user.ID(userID)))).
		Only(context.Background())
	if err == nil {
		return errors.New("user has already rated this place")
	}

	_, err = rs.client.Review.Create().
		SetID(uuid.New().String()).
		SetPlace(placeData).
		SetUser(userData).
		SetScore(score).
		SetContent(content).
		Save(context.Background())
	return err
}

// RemoveReview allows a user to remove a review.
func (rs *ReviewServiceImpl) RemoveReview(reviewID, userID string) error {
	review, err := rs.client.Review.Get(context.Background(), reviewID)
	if err != nil {
		return err
	}
	if review.Edges.User.ID != userID {
		return errors.New("user does not have permission to delete this review")
	}
	return rs.client.Review.DeleteOneID(reviewID).Exec(context.Background())
}

// GetReviewByID retrieves a review by its ID.
func (rs *ReviewServiceImpl) GetReviewByID(reviewID string) (*ent.Review, error) {
	return rs.client.Review.Get(context.Background(), reviewID)
}

// UpdateReviewContent allows a user to update the content of their review.
func (rs *ReviewServiceImpl) UpdateReviewContent(reviewID, userID, newContent string) error {
	review, err := rs.client.Review.Get(context.Background(), reviewID)
	if err != nil {
		return err
	}
	if review.Edges.User.ID != userID {
		return errors.New("user does not have permission to update this review")
	}
	_, err = rs.client.Review.UpdateOneID(reviewID).SetContent(newContent).Save(context.Background())
	return err
}

// AddMediaToReview allows a user to add media to their review.
func (rs *ReviewServiceImpl) AddMediaToReview(reviewID string, media *ent.Media) error {
	_, err := rs.client.Review.UpdateOneID(reviewID).AddMedias(media).Save(context.Background())
	return err
}

// GetReviewsByUserID retrieves all reviews by a user.
func (rs *ReviewServiceImpl) GetReviewsByUserID(userID string) ([]*ent.Review, error) {
	return rs.client.Review.Query().Where(review.HasUserWith(user.ID(userID))).All(context.Background())
}

// GetReviewsByPlaceID retrieves all reviews for a place.
func (rs *ReviewServiceImpl) GetReviewsByPlaceID(placeID string) ([]*ent.Review, error) {
	return rs.client.Review.Query().Where(review.HasPlaceWith(place.ID(placeID))).All(context.Background())
}

// GetReviewsByEventID retrieves all reviews for an event.
func (rs *ReviewServiceImpl) GetReviewsByEventID(eventID string) ([]*ent.Review, error) {
	return rs.client.Review.Query().Where(review.HasEventWith(event.ID(eventID))).All(context.Background())
}

// GetReviewsByBusinessID retrieves all reviews for a business.
func (rs *ReviewServiceImpl) GetReviewsByBusinessID(businessID string) ([]*ent.Review, error) {
	return rs.client.Review.Query().Where(review.HasBusinessWith(business.ID(businessID))).All(context.Background())
}

// GetReviewsByScore retrieves all reviews with a given score.
func (rs *ReviewServiceImpl) GetReviewsByScore(score float64) ([]*ent.Review, error) {
	return rs.client.Review.Query().Where(review.ScoreEQ(score)).All(context.Background())
}

// GetReviewsInDateRange retrieves all reviews in a given date range.
func (rs *ReviewServiceImpl) GetReviewsInDateRange(startDate, endDate time.Time) ([]*ent.Review, error) {
	return rs.client.Review.Query().Where(review.CreatedAtGTE(startDate), review.CreatedAtLTE(endDate)).All(context.Background())
}

// GetReviewsByLikeCount retrieves all reviews by like count.
func (rs *ReviewServiceImpl) GetReviewsByLikeCount() ([]*ent.Review, error) {
	return rs.client.Review.Query().Order(ent.Desc(review.FieldLikeCount)).All(context.Background())
}

// GetReviewsByDislikeCount retrieves all reviews by dislike count.
func (rs *ReviewServiceImpl) GetReviewsByDislikeCount() ([]*ent.Review, error) {
	return rs.client.Review.Query().Order(ent.Desc(review.FieldDislikeCount)).All(context.Background())
}

// LikeReview // LikeReview allows a user to like a review.
func (rs *ReviewServiceImpl) LikeReview(reviewID, userID string) error {
	_, err := rs.client.Review.UpdateOneID(reviewID).AddLikeCount(1).Save(context.Background())
	return err
}

// DislikeReview // DislikeReview allows a user to dislike a review.
func (rs *ReviewServiceImpl) DislikeReview(reviewID, userID string) error {
	_, err := rs.client.Review.UpdateOneID(reviewID).AddDislikeCount(1).Save(context.Background())
	// get a user and add the review to their disliked reviews
	//user, err := rs.client.User.Get(context.Background(), userID)
	//if err != nil {
	//	return err
	//}
	return err
}

// RateEvent allows a user to rate and review an event.
func (rs *ReviewServiceImpl) RateEvent(eventID, userID string, score float64, content string) error {
	event, err := rs.client.Event.Get(context.Background(), eventID)
	if err != nil {
		return err
	}

	user, err := rs.client.User.Get(context.Background(), userID)
	if err != nil {
		return err
	}

	_, err = rs.client.Review.Create().
		SetID(uuid.New().String()).
		SetScore(score).
		SetEventID(eventID).
		SetEvent(event).
		SetUserID(userID).
		SetUser(user).
		SetContent(content).
		Save(context.Background())
	return err
}

// RateBusiness allows a user to rate and review a business.
func (rs *ReviewServiceImpl) RateBusiness(businessID, userID string, score float64, content string) error {
	business, err := rs.client.Business.Get(context.Background(), businessID)
	if err != nil {
		return err
	}

	user, err := rs.client.User.Get(context.Background(), userID)
	if err != nil {
		return err
	}

	_, err = rs.client.Review.Create().
		SetID(uuid.New().String()).
		SetScore(score).
		SetBusinessID(businessID).
		SetBusiness(business).
		SetUserID(userID).
		SetUser(user).
		SetContent(content).
		Save(context.Background())
	return err
}

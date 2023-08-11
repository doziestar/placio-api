package reviews

import (
	"context"
	"placio-app/ent"
	"placio-app/ent/rating"
)

type RatingService interface {
	CreateRating(ctx context.Context, ratingDTO *RatingDTO) (*ent.Rating, error)
	UpdateRating(ctx context.Context, ratingID string, newRating int) (*ent.Rating, error)
	GetRating(ctx context.Context, ratingID string) (*ent.Rating, error)
	ListRatings(ctx context.Context) ([]*ent.Rating, error)
	DeleteRating(ctx context.Context, ratingID string) error
}

type RatingServiceImpl struct {
	client *ent.Client
}

func NewRatingService(client *ent.Client) RatingService {
	return &RatingServiceImpl{client: client}
}

func (rs *RatingServiceImpl) CreateRating(ctx context.Context, r *RatingDTO) (*ent.Rating, error) {
	rating, err := rs.client.Rating.
		Create().
		SetUser(r.User).
		SetPlace(r.Place).
		SetEvent(r.Event).
		SetBusiness(r.Business).
		SetScore(r.Score).
		Save(ctx)
	return rating, err
}

func (rs *RatingServiceImpl) UpdateRating(ctx context.Context, ratingID string, score int) (*ent.Rating, error) {
	rating, err := rs.client.Rating.
		UpdateOneID(ratingID).
		SetScore(score).
		Save(ctx)
	return rating, err
}

func (rs *RatingServiceImpl) GetRating(ctx context.Context, ratingID string) (*ent.Rating, error) {
	rating, err := rs.client.Rating.
		Query().
		Where(rating.IDEQ(ratingID)).
		Only(ctx)
	return rating, err
}

func (rs *RatingServiceImpl) ListRatings(ctx context.Context) ([]*ent.Rating, error) {
	ratings, err := rs.client.Rating.
		Query().
		All(ctx)
	return ratings, err
}

func (rs *RatingServiceImpl) DeleteRating(ctx context.Context, ratingID string) error {
	err := rs.client.Rating.
		DeleteOneID(ratingID).
		Exec(ctx)
	return err
}

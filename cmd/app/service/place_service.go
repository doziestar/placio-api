package service

import (
	"context"
	"placio-app/ent"
)

type PlaceService interface {
	GetPlace(ctx context.Context, placeID string) (*ent.Place, error)
	CreatePlace(ctx context.Context, placeData map[string]interface{}) (*ent.Place, error)
	UpdatePlace(ctx context.Context, placeID string, placeData map[string]interface{}) (*ent.Place, error)
	DeletePlace(ctx context.Context, placeID string) error
}

type PlaceServiceImpl struct {
	client *ent.Client
}

func NewPlaceService(client *ent.Client) *PlaceServiceImpl {
	return &PlaceServiceImpl{client: client}
}

func (s *PlaceServiceImpl) GetPlace(ctx context.Context, placeID string) (*ent.Place, error) {
	return s.client.Place.Get(ctx, placeID)
}

func (s *PlaceServiceImpl) CreatePlace(ctx context.Context, placeData map[string]interface{}) (*ent.Place, error) {
	// Here you would parse placeData and use it to create a new Place.
	// This is just a basic example, you may need to handle additional fields and validation.
	return s.client.Place.
		Create().
		SetName(placeData["name"].(string)).
		SetDescription(placeData["description"].(string)).
		Save(ctx)
}

func (s *PlaceServiceImpl) UpdatePlace(ctx context.Context, placeID string, placeData map[string]interface{}) (*ent.Place, error) {
	// Similar to CreatePlace, you'd parse placeData and use it to update the Place.
	// This is a basic example, you'll need to handle additional fields and validation.
	return s.client.Place.
		UpdateOneID(placeID).
		SetName(placeData["name"].(string)).
		SetDescription(placeData["description"].(string)).
		Save(ctx)
}

func (s *PlaceServiceImpl) DeletePlace(ctx context.Context, placeID string) error {
	return s.client.Place.
		DeleteOneID(placeID).
		Exec(ctx)
}

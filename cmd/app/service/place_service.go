package service

import (
	"context"
	"github.com/google/uuid"
	"placio-app/Dto"
	"placio-app/ent"
)

type PlaceService interface {
	GetPlace(ctx context.Context, placeID string) (*ent.Place, error)
	CreatePlace(ctx context.Context, placeData Dto.CreatePlaceDTO) (*ent.Place, error)
	UpdatePlace(ctx context.Context, placeID string, placeData Dto.UpdatePlaceDTO) (*ent.Place, error)
	DeletePlace(ctx context.Context, placeID string) error
}

type PlaceServiceImpl struct {
	client        *ent.Client
	searchService SearchService
}

func NewPlaceService(client *ent.Client, searchService SearchService) *PlaceServiceImpl {
	return &PlaceServiceImpl{client: client, searchService: searchService}
}

func (s *PlaceServiceImpl) GetPlace(ctx context.Context, placeID string) (*ent.Place, error) {
	return s.client.Place.Get(ctx, placeID)
}

func (s *PlaceServiceImpl) CreatePlace(ctx context.Context, placeData Dto.CreatePlaceDTO) (*ent.Place, error) {
	// Here you would parse placeData and use it to create a new Place.
	// This is just a basic example, you may need to handle additional fields and validation.
	place, err := s.client.Place.
		Create().
		SetName(placeData.Name).
		SetID(uuid.New().String()).
		SetDescription(*placeData.Description).
		SetPicture(*placeData.Picture).
		SetCoverImage(*placeData.CoverImage).
		SetWebsite(*placeData.Website).
		SetLocation(placeData.Location).
		SetEmail(*placeData.Email).
		SetPhone(*placeData.Phone).
		SetAvailability(*placeData.Availability).
		SetImages(*placeData.Images).
		SetFeatures(*placeData.Features).
		SetMapCoordinates(placeData.MapCoordinates).
		SetAdditionalInfo(*placeData.AdditionalInfo).
		SetType(placeData.Type).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// Add the new place to the search index.
	if err := s.searchService.CreateOrUpdatePlace(ctx, place); err != nil {
		return nil, err
	}

	return place, nil
}

func (s *PlaceServiceImpl) UpdatePlace(ctx context.Context, placeID string, placeData Dto.UpdatePlaceDTO) (*ent.Place, error) {
	// Similar to CreatePlace, you'd parse placeData and use it to update the Place.
	// This is a basic example, you'll need to handle additional fields and validation.
	place, err := s.client.Place.
		UpdateOneID(placeID).
		SetDescription(*placeData.Description).
		SetPicture(*placeData.Picture).
		SetCoverImage(*placeData.CoverImage).
		SetWebsite(*placeData.Website).
		SetEmail(*placeData.Email).
		SetPhone(*placeData.Phone).
		SetAvailability(*placeData.Availability).
		SetImages(*placeData.Images).
		SetFeatures(*placeData.Features).
		SetAdditionalInfo(*placeData.AdditionalInfo).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// Update the place in the search index.
	if err := s.searchService.CreateOrUpdatePlace(ctx, place); err != nil {
		return nil, err
	}

	return place, nil
}

func (s *PlaceServiceImpl) DeletePlace(ctx context.Context, placeID string) error {
	return s.client.Place.
		DeleteOneID(placeID).
		Exec(ctx)
}

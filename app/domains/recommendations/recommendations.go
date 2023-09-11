package recommendations

import (
	"context"
	"placio-app/domains/places"
	"placio-app/domains/users"
	"placio-app/ent"
)

type IRecommendations interface {
	GetPlacesRecommendations(ctx context.Context) ([]*ent.Place, error)
	GetRestaurantsRecommendations(ctx context.Context) ([]*ent.Place, error)
	GetHotelsRecommendations(ctx context.Context) ([]*ent.Place, error)
	GetInventoryRecommendations(ctx context.Context) ([]*ent.PlaceInventory, error)
	GetUsersRecommendations(ctx context.Context) ([]*ent.User, error)
}

type RecommendationService struct {
	client        *ent.Client
	userService   users.UserService
	placesService places.PlaceService
}

func NewRecommendations(client *ent.Client, userService users.UserService, placesService places.PlaceService) *RecommendationService {
	return &RecommendationService{
		client:        client,
		userService:   userService,
		placesService: placesService,
	}
}

func (r *RecommendationService) GetPlacesRecommendations(ctx context.Context) ([]*ent.Place, error) {
	return nil, nil
}

func (r *RecommendationService) GetRestaurantsRecommendations(ctx context.Context) ([]*ent.Place, error) {
	return nil, nil
}

func (r *RecommendationService) GetHotelsRecommendations(ctx context.Context) ([]*ent.Place, error) {
	return nil, nil
}

func (r *RecommendationService) GetInventoryRecommendations(ctx context.Context) ([]*ent.PlaceInventory, error) {
	return nil, nil
}

func (r *RecommendationService) GetUsersRecommendations(ctx context.Context) ([]*ent.User, error) {
	return nil, nil
}

package recommendations

import (
	"context"
	"log"
	"math/rand"
	"placio-app/domains/places"
	"placio-app/domains/users"
	"placio-app/ent"
	"time"
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
	user := ctx.Value("user").(string)
	userData, err := r.userService.GetUser(ctx, user)
	if err != nil {
		return nil, err
	}
	log.Println("recommendations for user", userData.Username)
	places, err := r.client.Place.
		Query().
		WithMedias().
		All(ctx)
	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(places), func(i, j int) {
		places[i], places[j] = places[j], places[i]
	})

	// Limit to 6 random places.
	if len(places) > 6 {
		places = places[:6]
	}

	return places, nil
}

func (r *RecommendationService) GetRestaurantsRecommendations(ctx context.Context) ([]*ent.Place, error) {
	//user := ctx.Value("user").(string)
	////userData, err := r.userService.GetUser(ctx, user)
	//if err != nil {
	//	return nil, err
	//}

	placesData, err := r.client.Place.
		Query().
		//Where(place.HasCategoriesWith(categories.Place).HasNameWith("Restaurant")).
		WithMedias().
		All(ctx)
	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(placesData), func(i, j int) {
		placesData[i], placesData[j] = placesData[j], placesData[i]
	})

	// Limit to 6 random places.
	if len(placesData) > 6 {
		placesData = placesData[:6]
	}

	return placesData, nil
}

func (r *RecommendationService) GetHotelsRecommendations(ctx context.Context) ([]*ent.Place, error) {
	//user := ctx.Value("user").(string)
	////userData, err := r.userService.GetUser(ctx, user)
	//if err != nil {
	//	return nil, err
	//}

	placesData, err := r.client.Place.
		Query().
		//Where(place.HasCategoriesWith(categories.Place).HasNameWith("Hotel")).
		WithMedias().
		All(ctx)
	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(placesData), func(i, j int) {
		placesData[i], placesData[j] = placesData[j], placesData[i]
	})

	// Limit to 6 random places.
	if len(placesData) > 6 {
		placesData = placesData[:6]
	}

	return placesData, nil
}

func (r *RecommendationService) GetInventoryRecommendations(ctx context.Context) ([]*ent.PlaceInventory, error) {
	placeInv, err := r.client.PlaceInventory.
		Query().
		WithPlace().
		WithMedia().
		All(ctx)

	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(placeInv), func(i, j int) {
		placeInv[i], placeInv[j] = placeInv[j], placeInv[i]
	})

	// Limit to 6 random places.
	if len(placeInv) > 6 {
		placeInv = placeInv[:6]
	}

	return placeInv, nil
}

func (r *RecommendationService) GetUsersRecommendations(ctx context.Context) ([]*ent.User, error) {
	users, err := r.client.User.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(users), func(i, j int) {
		users[i], users[j] = users[j], users[i]
	})

	//// Limit to 6 random places.
	//if len(users) > 6 {
	//	users = users[:6]
	//}

	return users, nil
}

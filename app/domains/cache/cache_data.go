package cache

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"log"
	"placio-app/domains/search"
	"placio-app/ent"
	"placio-app/ent/place"
	"placio-app/utility"
)

type ICacheService interface {
	AddPlaceToCacheAndSearchIndex(ctx context.Context, placeData *ent.Place, other ...string) error
}

type CacheService struct {
	client        *ent.Client
	cache         utility.RedisClient
	searchService search.SearchService
}

func NewCacheService(client *ent.Client, cache utility.RedisClient, searchService search.SearchService) *CacheService {
	return &CacheService{
		client:        client,
		cache:         cache,
		searchService: searchService,
	}
}

func (s *CacheService) AddPlaceToCacheAndSearchIndex(ctx context.Context, placeData *ent.Place, other ...string) error {
	log.Println("adding place to cache")
	fullPlace, err := s.client.Place.
		Query().
		Where(place.IDEQ(placeData.ID)).
		WithBusiness().
		WithUsers().
		WithFaqs().
		WithCategories().
		WithRooms().
		WithMedias().
		WithRatings().
		WithAmenities().
		WithReviews().
		WithReviews().
		Only(ctx)

	if err != nil {
		sentry.CaptureException(err)
		return err
	}
	// Add the new place to the search index.
	go func() {
		if err := s.searchService.CreateOrUpdatePlace(ctx, fullPlace); err != nil {
			sentry.CaptureException(err)
		}
	}()

	// add the new place to cache
	go func() {
		cacheKey := fmt.Sprintf("place:%s", fullPlace.ID)
		if err := s.cache.SetCache(ctx, cacheKey, fullPlace); err != nil {
			sentry.CaptureException(err)
			return
		}
	}()

	return nil
}

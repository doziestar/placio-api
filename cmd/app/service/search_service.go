package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"log"
	"os"
	"placio-app/ent"
)

type SearchService interface {
	SearchUsers(ctx context.Context, searchText string) (string, error)
	CreateOrUpdateUser(ctx context.Context, user *ent.User) error
	DeleteUser(ctx context.Context, userID string) error

	SearchPlaces(ctx context.Context, searchText string) (string, error)
	CreateOrUpdatePlace(ctx context.Context, place *ent.Place) error
	DeletePlace(ctx context.Context, placeID string) error

	SearchEvents(ctx context.Context, searchText string) (string, error)
	CreateOrUpdateEvent(ctx context.Context, event *ent.Event) error
	DeleteEvent(ctx context.Context, eventID string) error

	SearchBusinesses(ctx context.Context, searchText string) (string, error)
	CreateOrUpdateBusiness(ctx context.Context, business *ent.Business) error
	DeleteBusiness(ctx context.Context, businessID string) error
}

type SearchServiceImpl struct {
	client *search.Client
}

// NewSearchService Initiate a new search service with an Elasticsearch client
func NewSearchService() (SearchService, error) {
	log.Println("Creating algolia client")
	client := search.NewClient(os.Getenv("ALGOLIA_APP_ID"), os.Getenv("ALGOLIA_API_KEY"))
	if client == nil {
		return nil, errors.New("error creating algolia client")
	}
	log.Println("Created algolia client")
	return &SearchServiceImpl{client: client}, nil
}

// CreateOrUpdateUser updates an existing user or creates a new one.
func (s *SearchServiceImpl) CreateOrUpdateUser(ctx context.Context, user *ent.User) error {
	// index name is "users"
	index := s.client.InitIndex("users")
	_, err := index.SaveObject(user)
	if err != nil {
		log.Fatalf("Error when saving user: %s", err)
		return err
	}
	return nil
}

// DeleteUser deletes a user from the Elasticsearch index.
func (s *SearchServiceImpl) DeleteUser(ctx context.Context, userID string) error {
	index := s.client.InitIndex("users")
	_, err := index.DeleteObject(userID)
	if err != nil {
		log.Fatalf("Error when deleting user: %s", err)
		return err
	}
	return nil
}

// SearchUsers performs a search in the User index.
func (s *SearchServiceImpl) SearchUsers(ctx context.Context, searchText string) (string, error) {
	// index name is "users"
	index := s.client.InitIndex("users")
	res, err := index.Search(searchText, nil)
	if err != nil {
		log.Println("Error when searching users: %s", err)
		return "", err
	}

	bytes, err := json.Marshal(res)
	if err != nil {
		log.Printf("Error when converting search result to JSON: %v", err)
		return "", err
	}

	return string(bytes), nil

}

// SearchPlaces performs a search in the Place index.
func (s *SearchServiceImpl) SearchPlaces(ctx context.Context, searchText string) (string, error) {
	// index name is "places"
	index := s.client.InitIndex("places")
	res, err := index.Search(searchText, nil)
	if err != nil {
		log.Println("Error when searching places: %s", err)
		return "", err
	}

	bytes, err := json.Marshal(res)
	if err != nil {
		log.Printf("Error when converting search result to JSON: %v", err)
		return "", err
	}

	return string(bytes), nil

}

// SearchEvents performs a search in the Event index.
func (s *SearchServiceImpl) SearchEvents(ctx context.Context, searchText string) (string, error) {
	// index name is "events"
	index := s.client.InitIndex("events")
	res, err := index.Search(searchText, nil)
	if err != nil {
		log.Println("Error when searching events: %s", err)
		return "", err
	}

	bytes, err := json.Marshal(res)
	if err != nil {
		log.Printf("Error when converting search result to JSON: %v", err)
		return "", err
	}

	return string(bytes), nil

}

// SearchBusinesses performs a search in the Business index.
func (s *SearchServiceImpl) SearchBusinesses(ctx context.Context, searchText string) (string, error) {
	// index name is "businesses"
	index := s.client.InitIndex("businesses")
	res, err := index.Search(searchText, nil)
	if err != nil {
		log.Println("Error when searching businesses: %s", err)
		return "", err
	}

	bytes, err := json.Marshal(res)
	if err != nil {
		log.Println("Error when converting search result to JSON: %v", err)
		return "", err
	}

	return string(bytes), nil

}

// CreateOrUpdatePlace updates an existing place or creates a new one.
func (s *SearchServiceImpl) CreateOrUpdatePlace(ctx context.Context, place *ent.Place) error {
	// index name is "places"
	index := s.client.InitIndex("places")
	_, err := index.SaveObject(place)
	if err != nil {
		log.Fatalf("Error when saving place: %s", err)
		return err
	}

	return nil
}

// DeletePlace deletes a place from the Elasticsearch index.
func (s *SearchServiceImpl) DeletePlace(ctx context.Context, placeID string) error {
	index := s.client.InitIndex("places")
	_, err := index.DeleteObject(placeID)
	if err != nil {
		log.Fatalf("Error when deleting place: %s", err)
		return err
	}
	return nil
}

// CreateOrUpdateEvent updates an existing event or creates a new one.
func (s *SearchServiceImpl) CreateOrUpdateEvent(ctx context.Context, event *ent.Event) error {
	// index name is "events"
	index := s.client.InitIndex("events")
	_, err := index.SaveObject(event)
	if err != nil {
		log.Fatalf("Error when saving event: %s", err)
		return err
	}

	return nil
}

// DeleteEvent deletes an event from the Elasticsearch index.
func (s *SearchServiceImpl) DeleteEvent(ctx context.Context, eventID string) error {
	index := s.client.InitIndex("events")
	_, err := index.DeleteObject(eventID)
	if err != nil {
		log.Fatalf("Error when deleting event: %s", err)
		return err
	}
	return nil
}

// CreateOrUpdateBusiness updates an existing business or creates a new one.
func (s *SearchServiceImpl) CreateOrUpdateBusiness(ctx context.Context, business *ent.Business) error {
	// index name is "businesses"
	index := s.client.InitIndex("businesses")
	_, err := index.SaveObject(business)
	if err != nil {
		log.Fatalf("Error when saving business: %s", err)
		return err
	}
	return nil
}

// DeleteBusiness deletes a business from the Elasticsearch index.
func (s *SearchServiceImpl) DeleteBusiness(ctx context.Context, businessID string) error {
	// index name is "businesses"
	index := s.client.InitIndex("businesses")
	_, err := index.DeleteObject(businessID)
	if err != nil {
		log.Fatalf("Error when deleting business: %s", err)
		return err
	}
	return nil
}

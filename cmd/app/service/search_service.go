package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"placio-app/ent"

	"github.com/elastic/go-elasticsearch/v8"
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
	client *elasticsearch.Client
}

// NewSearchService Initiate a new search service with an Elasticsearch client
func NewSearchService() (SearchService, error) {
	log.Println("=========connecting to elastic search=====")
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	if err != nil {
		return nil, err
	}

	log.Println("======connection complete======")

	return &SearchServiceImpl{client: es}, nil
}

// CreateOrUpdateUser updates an existing user or creates a new one.
func (s *SearchServiceImpl) CreateOrUpdateUser(ctx context.Context, user *ent.User) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(user); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := s.client.Index(
		"users", // Index name.
		&buf,    // Data.
		s.client.Index.WithContext(ctx),
		s.client.Index.WithDocumentID(user.ID),
		s.client.Index.WithRefresh("true"),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%s", res.Status(), user.ID)
	} else {
		//log.Printf("[%s] %s; version=%d", res.Status(), res.Result, res.Version)
	}

	return nil
}

// DeleteUser deletes a user from the Elasticsearch index.
func (s *SearchServiceImpl) DeleteUser(ctx context.Context, userID string) error {
	res, err := s.client.Delete(
		"users", // Index name.
		userID,  // Document ID.
		s.client.Delete.WithContext(ctx),
		s.client.Delete.WithRefresh("true"),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error deleting document ID=%s", res.Status(), userID)
	} else {
		//	log.Printf("[%s] %s; version=%d", res.Status(), res.Result, res.Version)
	}

	return nil
}

// SearchUsers performs a search in the User index.
func (s *SearchServiceImpl) SearchUsers(ctx context.Context, searchText string) (string, error) {
	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  searchText,
				"fields": []string{"name", "username", "location", "bio", "website"},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := s.client.Search(
		s.client.Search.WithContext(ctx),
		s.client.Search.WithIndex("users"),
		s.client.Search.WithBody(&buf),
		s.client.Search.WithTrackTotalHits(true),
		s.client.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return "", err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
		return "", errors.New("elasticsearch search error")
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Process the response.
	out, _ := json.MarshalIndent(r, "", "  ")
	return string(out), nil
}

// SearchPlaces performs a search in the Place index.
func (s *SearchServiceImpl) SearchPlaces(ctx context.Context, searchText string) (string, error) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  searchText,
				"fields": []string{"name", "type", "description", "location"},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	res, err := s.client.Search(
		s.client.Search.WithContext(ctx),
		s.client.Search.WithIndex("places"),
		s.client.Search.WithBody(&buf),
		s.client.Search.WithTrackTotalHits(true),
		s.client.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return "", err
	}
	defer res.Body.Close()

	// Other error handling and response processing code as before...
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
		return "", errors.New("elasticsearch search error")
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Process the response.
	out, _ := json.MarshalIndent(r, "", "  ")
	return string(out), nil

}

// SearchEvents performs a search in the Event index.
func (s *SearchServiceImpl) SearchEvents(ctx context.Context, searchText string) (string, error) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  searchText,
				"fields": []string{"name"},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	res, err := s.client.Search(
		s.client.Search.WithContext(ctx),
		s.client.Search.WithIndex("events"),
		s.client.Search.WithBody(&buf),
		s.client.Search.WithTrackTotalHits(true),
		s.client.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return "", err
	}
	defer res.Body.Close()

	// Other error handling and response processing code as before...
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
		return "", errors.New("elasticsearch search error")
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Process the response.
	out, _ := json.MarshalIndent(r, "", "  ")
	return string(out), nil

}

// SearchBusinesses performs a search in the Business index.
func (s *SearchServiceImpl) SearchBusinesses(ctx context.Context, searchText string) (string, error) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  searchText,
				"fields": []string{"name", "description", "location", "website"},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	res, err := s.client.Search(
		s.client.Search.WithContext(ctx),
		s.client.Search.WithIndex("businesses"),
		s.client.Search.WithBody(&buf),
		s.client.Search.WithTrackTotalHits(true),
		s.client.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return "", err
	}
	defer res.Body.Close()

	// Other error handling and response processing code as before...
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
		return "", errors.New("elasticsearch search error")
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Process the response.
	out, _ := json.MarshalIndent(r, "", "  ")
	return string(out), nil

}

// CreateOrUpdatePlace updates an existing place or creates a new one.
func (s *SearchServiceImpl) CreateOrUpdatePlace(ctx context.Context, place *ent.Place) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(place); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := s.client.Index(
		"places",
		&buf,
		s.client.Index.WithContext(ctx),
		s.client.Index.WithDocumentID(place.ID),
		s.client.Index.WithRefresh("true"),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%s", res.Status(), place.ID)
	}

	return nil
}

// DeletePlace deletes a place from the Elasticsearch index.
func (s *SearchServiceImpl) DeletePlace(ctx context.Context, placeID string) error {
	res, err := s.client.Delete(
		"places",
		placeID,
		s.client.Delete.WithContext(ctx),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error deleting document ID=%s", res.Status(), placeID)
	}

	return nil
}

// CreateOrUpdateEvent updates an existing event or creates a new one.
func (s *SearchServiceImpl) CreateOrUpdateEvent(ctx context.Context, event *ent.Event) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(event); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := s.client.Index(
		"events",
		&buf,
		s.client.Index.WithContext(ctx),
		s.client.Index.WithDocumentID(event.ID),
		s.client.Index.WithRefresh("true"),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%s", res.Status(), event.ID)
	}

	return nil
}

// DeleteEvent deletes an event from the Elasticsearch index.
func (s *SearchServiceImpl) DeleteEvent(ctx context.Context, eventID string) error {
	res, err := s.client.Delete(
		"events",
		eventID,
		s.client.Delete.WithContext(ctx),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error deleting document ID=%s", res.Status(), eventID)
	}

	return nil
}

// CreateOrUpdateBusiness updates an existing business or creates a new one.
func (s *SearchServiceImpl) CreateOrUpdateBusiness(ctx context.Context, business *ent.Business) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(business); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := s.client.Index(
		"businesses",
		&buf,
		s.client.Index.WithContext(ctx),
		s.client.Index.WithDocumentID(business.ID),
		s.client.Index.WithRefresh("true"),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%s", res.Status(), business.ID)
	}

	return nil
}

// DeleteBusiness deletes a business from the Elasticsearch index.
func (s *SearchServiceImpl) DeleteBusiness(ctx context.Context, businessID string) error {
	res, err := s.client.Delete(
		"businesses",
		businessID,
		s.client.Delete.WithContext(ctx),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error deleting document ID=%s", res.Status(), businessID)
	}

	return nil
}

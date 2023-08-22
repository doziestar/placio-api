package feature_releases

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"placio-app/ent"
	"placio-app/ent/featurerelease"
	"placio-app/utility"
	"time"
)

type FeatureReleaseService interface {
	GetFeature(ctx context.Context, featureID string) (*ent.FeatureRelease, error)
	CreateFeature(ctx context.Context, featureData FeatureReleaseDTO) (*ent.FeatureRelease, error)
	UpdateFeature(ctx context.Context, featureID string, featureData FeatureReleaseDTO) (*ent.FeatureRelease, error)
	DeleteFeature(ctx context.Context, featureID string) error
	ListFeatures(ctx context.Context, lastId string, limit int) ([]*ent.FeatureRelease, string, error)
	SetFeatureState(ctx context.Context, featureID string, newState string) error
}

type FeatureReleaseServiceImpl struct {
	client *ent.Client
	cache  utility.RedisClient
}

func NewFeatureReleaseService(client *ent.Client, cache utility.RedisClient) *FeatureReleaseServiceImpl {
	return &FeatureReleaseServiceImpl{client: client, cache: cache}
}

func (s *FeatureReleaseServiceImpl) GetFeature(ctx context.Context, featureID string) (*ent.FeatureRelease, error) {
	cacheKey := fmt.Sprintf("feature:%s", featureID)
	data, err := s.cache.GetCache(ctx, cacheKey)
	if err == nil && data != nil {
		var feature ent.FeatureRelease
		if err := json.Unmarshal(data, &feature); err == nil {
			return &feature, nil
		}
	}

	// If cache miss or unmarshal error, fetch from the database
	featureData, err := s.client.FeatureRelease.Get(ctx, featureID)
	if err != nil {
		return nil, err
	}

	// Update cache
	if err := s.cache.SetCache(ctx, cacheKey, featureData); err != nil {
		log.Println("Failed to update cache:", err)
	}

	return featureData, nil
}

func (s *FeatureReleaseServiceImpl) CreateFeature(ctx context.Context, featureData FeatureReleaseDTO) (*ent.FeatureRelease, error) {
	// convert release date to time.Time
	releaseDate, err := time.Parse(time.RFC3339, featureData.ReleaseDate)
	if err != nil {
		return nil, errors.New("invalid release date")
	}

	createdFeature, err := s.client.FeatureRelease.Create().
		SetFeatureID(uuid.New().String()).
		SetFeatureName(featureData.FeatureName).
		SetDescription(featureData.Description).
		SetState(featurerelease.State(featureData.State)).
		SetEligibilityRules(featureData.EligibilityRules).
		SetReleaseDate(releaseDate).
		SetDocumentationLink(featureData.DocumentationLink).
		SetMetadata(featureData.Metadata).
		Save(ctx)
	return createdFeature, err
}

func (s *FeatureReleaseServiceImpl) UpdateFeature(ctx context.Context, featureID string, featureData FeatureReleaseDTO) (*ent.FeatureRelease, error) {
	// convert release date to time.Time
	releaseDate, err := time.Parse(time.RFC3339, featureData.ReleaseDate)
	if err != nil {
		return nil, errors.New("invalid release date")
	}
	upd := s.client.FeatureRelease.UpdateOneID(featureID).
		SetFeatureName(featureData.FeatureName).
		SetDescription(featureData.Description).
		SetState(featurerelease.State(featureData.State)).
		SetEligibilityRules(featureData.EligibilityRules).
		SetReleaseDate(releaseDate).
		SetDocumentationLink(featureData.DocumentationLink).
		SetMetadata(featureData.Metadata)

	updatedFeature, err := upd.Save(ctx)

	// If updated successfully, remove the outdated cache
	if err == nil {
		cacheKey := fmt.Sprintf("feature:%s", featureID)
		if delErr := s.cache.DeleteCache(ctx, cacheKey); delErr != nil {
			log.Println("Failed to delete outdated cache:", delErr)
		}
	}
	return updatedFeature, err
}

func (s *FeatureReleaseServiceImpl) DeleteFeature(ctx context.Context, featureID string) error {
	err := s.client.FeatureRelease.DeleteOneID(featureID).Exec(ctx)

	// If deleted successfully, remove the cache
	if err == nil {
		cacheKey := fmt.Sprintf("feature:%s", featureID)
		if delErr := s.cache.DeleteCache(ctx, cacheKey); delErr != nil {
			log.Println("Failed to delete cache:", delErr)
		}
	}
	return err
}

func (s *FeatureReleaseServiceImpl) ListFeatures(ctx context.Context, lastId string, limit int) ([]*ent.FeatureRelease, string, error) {
	if limit == 0 {
		limit = 10
	}

	// Construct the base query
	query := s.client.FeatureRelease.
		Query().
		Order(ent.Asc(featurerelease.FieldID)).
		Limit(limit + 1)

	// If lastId is provided, fetch records after it
	if lastId != "" {
		query = query.Where(featurerelease.IDGT(lastId))
	}

	features, err := query.All(ctx)
	if err != nil {
		return nil, "", err
	}

	// Determine the nextId for pagination
	var nextId string
	if len(features) == limit+1 {
		nextId = features[len(features)-1].ID
		features = features[:limit]
	}

	return features, nextId, nil
}

func (s *FeatureReleaseServiceImpl) SetFeatureState(ctx context.Context, featureID string, newState string) error {
	validStates := map[string]bool{
		"testing":    true,
		"staging":    true,
		"live":       true,
		"deprecated": true,
		"archived":   true,
	}
	if !validStates[newState] {
		return errors.New("invalid state provided")
	}

	upd := s.client.FeatureRelease.UpdateOneID(featureID).SetState(featurerelease.State(newState))
	_, err := upd.Save(ctx)

	// If state updated successfully, remove the outdated cache
	if err == nil {
		cacheKey := fmt.Sprintf("feature:%s", featureID)
		if delErr := s.cache.DeleteCache(ctx, cacheKey); delErr != nil {
			log.Println("Failed to delete outdated cache:", delErr)
		}
	}
	return err
}

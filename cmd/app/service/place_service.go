package service

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"log"
	"placio-app/Dto"
	"placio-app/ent"
	"placio-app/ent/business"
	"placio-app/ent/place"
	"placio-app/errors"
	"placio-app/utility"
	"sync"
)

type PlaceFilter struct {
	IDs      []string
	Name     []string
	Type     []string
	Country  []string
	City     []string
	State    []string
	Tags     []string
	Features []string
	Email    []string
	Phone    []string
	Website  []string
}

type PlaceService interface {
	GetPlace(ctx context.Context, placeID string) (*ent.Place, error)
	CreatePlace(ctx context.Context, placeData Dto.CreatePlaceDTO) (*ent.Place, error)
	UpdatePlace(ctx context.Context, placeID string, placeData Dto.UpdatePlaceDTO) (*ent.Place, error)
	DeletePlace(ctx context.Context, placeID string) error
	GetPlacesAssociatedWithBusinessAccount(c context.Context, businessId string) ([]*ent.Place, error)
	GetPlaces(ctx context.Context, filter *PlaceFilter, lastId string, limit int) ([]*ent.Place, string, error)
	AddAmenitiesToPlace(ctx context.Context, placeID string, amenityIDs []string) error
	GetAllPlaces(ctx context.Context, nextPageToken string, limit int) ([]*ent.Place, string, error)
}

type PlaceServiceImpl struct {
	client        *ent.Client
	searchService SearchService
	userLikes     UserLikePlaceService
	followService FollowService
	cache         utility.RedisClient
}

func NewPlaceService(client *ent.Client, searchService SearchService, userLikes UserLikePlaceService, followService FollowService, cache utility.RedisClient) *PlaceServiceImpl {
	return &PlaceServiceImpl{client: client, searchService: searchService, userLikes: userLikes, followService: followService}
}

func (s *PlaceServiceImpl) GetPlacesAssociatedWithBusinessAccount(c context.Context, businessId string) ([]*ent.Place, error) {
	places, err := s.client.Place.
		Query().
		Where(place.HasBusinessWith(business.ID(businessId))).
		All(c)
	if err != nil {
		return nil, err
	}

	return places, errors.LogAndReturnError(err)
}
func (s *PlaceServiceImpl) GetPlace(ctx context.Context, placeID string) (*ent.Place, error) {
	placeData, err := s.client.Place.
		Query().
		Where(place.ID(placeID)).
		WithUsers().
		WithBusiness().
		WithCategories().
		WithCategoryAssignments().
		WithEvents().
		WithAmenities().
		WithMenus().
		WithFaqs().
		First(ctx)
	if err != nil {
		return nil, errors.LogAndReturnError(err)
	}

	userID, ok := ctx.Value("user").(string)
	if !ok {
		return placeData, nil
	}
	if userID != "" {
		if err := s.checkUserInteraction(ctx, userID, placeID, placeData); err != nil {
			return nil, errors.LogAndReturnError(err)
		}
	}

	go s.addPlaceToCacheAndSearchIndex(ctx, placeData)

	return placeData, nil
}

func (s *PlaceServiceImpl) GetAllPlaces(ctx context.Context, lastId string, limit int) ([]*ent.Place, string, error) {
	if limit == 0 {
		limit = 10
	}
	query := s.client.Place.
		Query().
		WithUsers().
		WithBusiness().
		WithCategories().
		WithEvents().
		WithAmenities().
		WithReviews().
		WithMenus().
		WithFaqs().
		Limit(limit + 1) // We retrieve one extra record to determine if there are more pages

	if lastId != "" {
		// If lastId is provided, we fetch records after it
		query = query.Where(place.IDGT(lastId))
	}

	places, err := query.All(ctx)
	if err != nil {
		return nil, "", err
	}

	var nextId string
	if len(places) == limit+1 {
		// We have an extra record, that means there is a next page.
		// We take the ID of the last item as the cursor for the next page.
		// Also, we remove the extra item from the list that we return.
		nextId = places[len(places)-1].ID
		places = places[:limit]
	}

	userID, ok := ctx.Value("user").(string)
	if !ok {
		return places, nextId, nil
	}
	var wg sync.WaitGroup
	var firstErr error
	var lock sync.Mutex
	log.Println("checking if the user likes and follows the place", userID)
	if userID != "" {
		for _, place := range places {
			wg.Add(1)
			go func(place *ent.Place) {
				defer wg.Done()
				err := s.checkUserInteraction(ctx, userID, place.ID, place)
				if err != nil && firstErr == nil { // only capture the first error
					lock.Lock()
					if firstErr == nil {
						firstErr = err
					}
					lock.Unlock()
				}
			}(place)
		}
		wg.Wait()
		if firstErr != nil {
			return nil, "", firstErr
		}
	}

	//go func() {
	//	cacheKey := fmt.Sprintf("places:%s:%s", nextId, strconv.Itoa(limit))
	//	err := s.cache.SetCache(ctx, cacheKey, places)
	//	if err != nil {
	//		errors.LogAndReturnError(err)
	//	}
	//}()

	return places, nextId, nil
}

func (s *PlaceServiceImpl) checkUserInteraction(ctx context.Context, userID, placeID string, place *ent.Place) error {
	var wg sync.WaitGroup
	var userLikesPlace, userFollowsPlace bool
	var errUserLikesPlace, errUserFollowsPlace error

	wg.Add(2)
	go func() {
		defer wg.Done()
		userLikesPlace, errUserLikesPlace = s.userLikes.CheckIfUserLikesPlace(ctx, userID, placeID)
	}()

	go func() {
		defer wg.Done()
		userFollowsPlace, errUserFollowsPlace = s.followService.CheckIfUserFollowsPlace(ctx, userID, placeID)
	}()

	wg.Wait()

	if errUserLikesPlace != nil {
		return errUserLikesPlace
	}
	if errUserFollowsPlace != nil {
		return errUserFollowsPlace
	}

	place.FollowedByCurrentUser = userFollowsPlace
	place.LikedByCurrentUser = userLikesPlace

	return nil
}

func (s *PlaceServiceImpl) CreatePlace(ctx context.Context, placeData Dto.CreatePlaceDTO) (*ent.Place, error) {
	// get the user from the context
	user := ctx.Value("user").(string)
	// get user from database
	userEnt, err := s.client.User.Get(ctx, user)
	if err != nil {
		return nil, err
	}

	// get business from database
	business, err := s.client.Business.Get(ctx, placeData.BusinessID)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	var amenities []*ent.Amenity
	var categories []*ent.Category

	// Fetch amenities, loop through the IDs and get the amenity from the database. do it using goroutines
	if len(placeData.AmenityIDs) != 0 {
		wg.Add(len(placeData.AmenityIDs))
		amenities = make([]*ent.Amenity, len(placeData.AmenityIDs))
		for i, amenityID := range placeData.AmenityIDs {
			go func(i int, amenityID string) {
				defer wg.Done()
				amenities[i], err = s.client.Amenity.Get(ctx, amenityID)
				if err != nil {
					return
				}
			}(i, amenityID)
			wg.Wait()
		}
	}

	// Fetch categories, loop through the IDs and get the category from the database. do it using goroutines
	if len(placeData.Categories) != 0 {
		wg.Add(len(placeData.Categories))
		categories = make([]*ent.Category, len(placeData.Categories))
		for i, categoryID := range placeData.Categories {
			go func(i int, categoryID string) {
				defer wg.Done()
				categories[i], err = s.client.Category.Get(ctx, categoryID)
				if err != nil {
					return
				}
			}(i, categoryID)
			wg.Wait()
		}
	}

	// Begin a transaction to ensure all operations occur atomically
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, errors.LogAndReturnError(err)
	}

	// Create the Place without associating Categories directly
	place, err := tx.Place.
		Create().
		SetID(uuid.New().String()).
		SetName(placeData.Name).
		SetDescription(placeData.Description).
		SetPicture(placeData.Picture).
		SetCoverImage(placeData.CoverImage).
		SetWebsite(placeData.Website).
		SetLocation(placeData.Location).
		SetCountry(placeData.Country).
		SetCity(placeData.City).
		SetState(placeData.State).
		SetEmail(placeData.Email).
		SetPhone(placeData.Phone).
		SetAvailability(placeData.Availability).
		SetImages(placeData.Images).
		SetFeatures(placeData.Features).
		SetMapCoordinates(placeData.MapCoordinates).
		SetAdditionalInfo(placeData.AdditionalInfo).
		SetType(placeData.Type).
		SetBusiness(business).
		SetBusinessID(business.ID).
		AddUsers(userEnt).
		AddAmenities(amenities...).
		Save(ctx)
	if err != nil {
		_ = tx.Rollback() // rollback the transaction in case of error
		return nil, errors.LogAndReturnError(err)
	}

	// Now loop over each Category and create a CategoryAssignment for each
	for _, category := range categories {
		_, err := tx.CategoryAssignment.
			Create().
			SetEntityID(place.ID).
			SetEntityType("place").
			SetCategoryID(category.ID).
			Save(ctx)
		if err != nil {
			_ = tx.Rollback() // rollback the transaction in case of error
			return nil, errors.LogAndReturnError(err)
		}
	}

	// Now that all operations have succeeded, commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, errors.LogAndReturnError(err)
	}

	// Add the new place to the search index and cache
	go s.addPlaceToCacheAndSearchIndex(ctx, place)

	return place, nil
}

func (s *PlaceServiceImpl) addPlaceToCacheAndSearchIndex(ctx context.Context, place interface{}, other ...string) error {
	// Add the new place to the search index.
	go func() {
		if err := s.searchService.CreateOrUpdatePlace(ctx, place.(*ent.Place)); err != nil {
			errors.LogAndReturnError(err)
		}
	}()

	// add the new place to cache
	go func() {
		// check if other is not empty
		if len(other) != 0 {
			cacheKey := fmt.Sprintf("place:%s:%s", place.(ent.Place).ID, other[0])
			if err := s.cache.SetCache(ctx, cacheKey, place); err != nil {
				errors.LogAndReturnError(err)
			}
			return
		}
		cacheKey := fmt.Sprintf("place:%s", place.(ent.Place).ID)
		if err := s.cache.SetCache(ctx, cacheKey, place); err != nil {
			errors.LogAndReturnError(err)
		}
	}()

	return nil
}

func (s *PlaceServiceImpl) AddAmenitiesToPlace(ctx context.Context, placeID string, amenityIDs []string) error {
	// Fetch place
	place, err := s.client.Place.Get(ctx, placeID)
	if err != nil {
		return err
	}

	// Fetch amenities
	wg := sync.WaitGroup{}
	wg.Add(len(amenityIDs))
	amenities := make([]*ent.Amenity, len(amenityIDs))
	for i, amenityID := range amenityIDs {
		go func(i int, amenityID string) {
			defer wg.Done()
			amenities[i], err = s.client.Amenity.Get(ctx, amenityID)
			if err != nil {
				errors.LogAndReturnError(err)
				return
			}
		}(i, amenityID)
		wg.Wait()

	}

	// Update place with new amenities
	placeData, err := place.Update().AddAmenities(amenities...).Save(ctx)
	if err != nil {
		errors.LogAndReturnError(err)
	}

	// Add the new place to the search index and cache
	go s.addPlaceToCacheAndSearchIndex(ctx, placeData)

	return nil
}

func (s *PlaceServiceImpl) UpdatePlace(ctx context.Context, placeID string, placeData Dto.UpdatePlaceDTO) (*ent.Place, error) {

	place, err := s.client.Place.
		UpdateOneID(placeID).
		SetDescription(placeData.Description).
		SetPicture(placeData.Picture).
		SetCoverImage(placeData.CoverImage).
		SetWebsite(placeData.Website).
		SetEmail(placeData.Email).
		SetPhone(placeData.Phone).
		SetAvailability(placeData.Availability).
		SetImages(placeData.Images).
		SetFeatures(placeData.Features).
		SetAdditionalInfo(placeData.AdditionalInfo).
		SetLocation(placeData.Location).
		SetCountry(placeData.Country).
		SetCity(placeData.City).
		SetState(placeData.State).
		Save(ctx)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	// Update the place in the search index.
	go s.addPlaceToCacheAndSearchIndex(ctx, place)

	return place, nil
}

func (s *PlaceServiceImpl) DeletePlace(ctx context.Context, placeID string) error {
	return s.client.Place.
		DeleteOneID(placeID).
		Exec(ctx)
}

//func (s *PlaceServiceImpl) GetPlaces(ctx context.Context, filter *PlaceFilter, page int, pageSize int) ([]*ent.Place, error) {
//	query := s.client.Place.
//		Query().
//		WithBusiness().
//		WithUsers()
//
//	// Apply filters
//	if len(filter.IDs) > 0 {
//		query = query.Where(place.IDIn(filter.IDs...))
//	}
//	if len(filter.Name) > 0 {
//		query = query.Where(place.NameIn(filter.Name...))
//	}
//	if len(filter.Type) > 0 {
//		query = query.Where(place.TypeIn(filter.Type...))
//	}
//	if len(filter.Country) > 0 {
//		query = query.Where(place.CountryIn(filter.Country...))
//	}
//	if len(filter.City) > 0 {
//		query = query.Where(place.CityIn(filter.City...))
//	}
//	if len(filter.State) > 0 {
//		query = query.Where(place.StateIn(filter.State...))
//	}
//	if len(filter.Tags) > 0 {
//		// To use Tags, Features, etc you may need additional handling as these fields are stored as JSON
//		// This is just a sample, it might not work as is
//		query = query.Where(place.HasTagsWith(filter.Tags...))
//	}
//	if len(filter.Features) > 0 {
//		// Similarly to Tags
//		query = query.Where(place.HasFeaturesWith(filter.Features...))
//	}
//
//	// Apply pagination
//	query = query.Offset((page - 1) * pageSize).Limit(pageSize)
//
//	// Execute query
//	places, err := query.All(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	return places, nil
//}

func (s *PlaceServiceImpl) GetPlaces(ctx context.Context, filter *PlaceFilter, lastId string, limit int) ([]*ent.Place, string, error) {
	if limit == 0 {
		limit = 10
	}
	query := s.client.Place.
		Query().
		WithBusiness().
		WithUsers().
		Limit(limit + 1)

	// Apply filters
	if len(filter.IDs) > 0 {
		query = query.Where(place.IDIn(filter.IDs...))
	}
	if len(filter.Name) > 0 {
		query = query.Where(place.NameIn(filter.Name...))
	}
	if len(filter.Type) > 0 {
		query = query.Where(place.TypeIn(filter.Type...))
	}
	if len(filter.Country) > 0 {
		query = query.Where(place.CountryIn(filter.Country...))
	}
	if len(filter.City) > 0 {
		query = query.Where(place.CityIn(filter.City...))
	}
	if len(filter.State) > 0 {
		query = query.Where(place.StateIn(filter.State...))
	}
	// ...

	if lastId != "" {
		// If lastId is provided, we fetch records after it
		query = query.Where(place.IDGT(lastId))
	}

	places, err := query.All(ctx)
	if err != nil {
		return nil, "", err
	}

	var nextId string
	if len(places) == limit+1 {
		nextId = places[len(places)-1].ID
		places = places[:limit]
	}

	userID, ok := ctx.Value("user").(string)
	if !ok {
		return places, nextId, nil
	}
	var wg sync.WaitGroup
	var firstErr error
	var lock sync.Mutex
	log.Println("checking if the user likes and follows the place", userID)
	if userID != "" {
		for _, place := range places {
			wg.Add(1)
			go func(place *ent.Place) {
				defer wg.Done()
				err := s.checkUserInteraction(ctx, userID, place.ID, place)
				if err != nil && firstErr == nil { // only capture the first error
					lock.Lock()
					if firstErr == nil {
						firstErr = err
					}
					lock.Unlock()
				}
			}(place)
		}
		wg.Wait()
		if firstErr != nil {
			return nil, "", firstErr
		}
	}

	// Filter places by tags and features at the application level
	filteredPlaces := []*ent.Place{}
	for _, p := range places {
		// Check if place tags include all tags from the filter
		if len(filter.Tags) > 0 && !containsAll(p.Tags, filter.Tags) {
			continue
		}
		// Check if place features include all features from the filter
		if len(filter.Features) > 0 && !containsAll(p.Features, filter.Features) {
			continue
		}
		filteredPlaces = append(filteredPlaces, p)
	}

	places = filteredPlaces

	return places, nextId, nil
}

// containsAll checks if all elements of subset are in the set
func containsAll(set []string, subset []string) bool {
	setMap := make(map[string]bool)
	for _, s := range set {
		setMap[s] = true
	}
	for _, s := range subset {
		if !setMap[s] {
			return false
		}
	}
	return true
}

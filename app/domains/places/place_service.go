package places

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"log"
	"mime/multipart"
	"placio-app/domains/amenities"
	"placio-app/domains/follow"
	"placio-app/domains/like"
	"placio-app/domains/media"
	"placio-app/domains/search"
	"placio-app/ent"
	"placio-app/ent/amenity"
	"placio-app/ent/business"
	"placio-app/ent/place"
	"placio-app/utility"
	"placio-pkg/errors"
	"sync"
)

type PlaceFilter struct {
	Name     string
	Type     string
	Country  string
	City     string
	State    string
	Tags     string
	Features string
	Website  string
}

type PlaceService interface {
	GetPlace(ctx context.Context, placeID string) (*ent.Place, error)
	CreatePlace(ctx context.Context, placeData CreatePlaceDTO) (*ent.Place, error)
	UpdatePlace(ctx context.Context, placeID string, placeData UpdatePlaceDTO) (*ent.Place, error)
	AddMediaToPlace(ctx context.Context, placeID string, files []*multipart.FileHeader) error
	RemoveMediaToPlace(ctx context.Context, placeID string, mediaID []string) error
	DeletePlace(ctx context.Context, placeID string) error
	GetPlacesAssociatedWithBusinessAccount(c context.Context, businessId string) ([]*ent.Place, error)
	GetPlaces(ctx context.Context, filter *PlaceFilter, lastId string, limit int) ([]*ent.Place, string, error)
	AddAmenitiesToPlace(ctx context.Context, placeID string, amenities []amenities.CreateAmenityInput) error
	GetAllPlaces(ctx context.Context, nextPageToken string, limit int) ([]*ent.Place, string, error)
	RemoveAmenitiesFromPlace(ctx context.Context, placeID string, amenityIDs []string) error
	RemoveMediaFromPlace(ctx context.Context, placeID string, mediaID string) error
}

type PlaceServiceImpl struct {
	client        *ent.Client
	searchService search.SearchService
	userLikes     like.UserLikePlaceService
	followService follow.FollowService
	cache         utility.RedisClient
	mediaService  media.MediaService
}

func NewPlaceService(client *ent.Client, searchService search.SearchService, userLikes like.UserLikePlaceService, followService follow.FollowService, cache utility.RedisClient, mediaService media.MediaService) *PlaceServiceImpl {
	return &PlaceServiceImpl{client: client, searchService: searchService, userLikes: userLikes, followService: followService, cache: cache, mediaService: mediaService}
}

func (s *PlaceServiceImpl) GetPlacesAssociatedWithBusinessAccount(c context.Context, businessId string) ([]*ent.Place, error) {
	places, err := s.client.Place.
		Query().
		Where(place.HasBusinessWith(business.ID(businessId))).
		All(c)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	return places, nil
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
		sentry.CaptureException(err)
		return nil, err
	}

	userID, ok := ctx.Value("user").(string)
	if !ok {
		return placeData, nil
	}
	if userID != "" {
		if err := s.checkUserInteraction(ctx, userID, placeID, placeData); err != nil {
			sentry.CaptureException(err)
			return nil, err
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
			sentry.CaptureException(firstErr)
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

func (s *PlaceServiceImpl) CreatePlace(ctx context.Context, placeData CreatePlaceDTO) (*ent.Place, error) {
	// get the user from the context
	user := ctx.Value("user").(string)
	// get user from db
	userEnt, err := s.client.User.Get(ctx, user)
	if err != nil {
		return nil, err
	}

	// get business from db
	business, err := s.client.Business.Get(ctx, placeData.BusinessID)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	var wg sync.WaitGroup
	var amenities []*ent.Amenity
	var categories []*ent.Category

	// Fetch amenities, loop through the IDs and get the amenity from the db. do it using goroutines
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

	// Fetch categories, loop through the IDs and get the category from the db. do it using goroutines
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
		sentry.CaptureException(err)
		return nil, err
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
		sentry.CaptureException(err)
		return nil, err
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
			_ = tx.Rollback()
			sentry.CaptureException(err) // rollback the transaction in case of error
			return nil, err
		}
	}

	// Now that all operations have succeeded, commit the transaction
	if err := tx.Commit(); err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	// Add the new place to the search index and cache
	go s.addPlaceToCacheAndSearchIndex(ctx, place)

	return place, nil
}

func (s *PlaceServiceImpl) addPlaceToCacheAndSearchIndex(ctx context.Context, placeData *ent.Place, other ...string) error {
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

func (s *PlaceServiceImpl) RemoveAmenitiesFromPlace(ctx context.Context, placeID string, amenityIDs []string) error {
	placeData, err := s.client.Place.
		UpdateOneID(placeID).
		RemoveAmenityIDs(amenityIDs...).
		Save(ctx)

	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	go s.addPlaceToCacheAndSearchIndex(ctx, placeData)

	return nil
}

func (s *PlaceServiceImpl) AddMediaToPlace(ctx context.Context, placeID string, files []*multipart.FileHeader) error {

	// Fetch place
	place, err := s.client.Place.Get(ctx, placeID)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	// Upload files to cloudinary
	uploadedFiles, err := s.mediaService.UploadAndCreateMedia(ctx, files)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	log.Println("media uploaded", uploadedFiles)
	_, err = s.client.Place.UpdateOneID(placeID).AddMedias(uploadedFiles...).Save(ctx)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}
	log.Println("data saved successfully")

	// Add the updated place to the search index and cache
	go s.addPlaceToCacheAndSearchIndex(ctx, place)

	return nil
}

func (s *PlaceServiceImpl) RemoveMediaFromPlace(ctx context.Context, placeID string, mediaID string) error {
	// Fetch the place
	placeData, err := s.client.Place.Get(ctx, placeID)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	// Remove the media from the place
	_, err = s.client.Place.UpdateOneID(placeID).RemoveMediaIDs(mediaID).Save(ctx)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	go s.addPlaceToCacheAndSearchIndex(ctx, placeData)

	// Ideally, we'd also want to delete the media from wherever it's stored,
	// but for now, we're just removing the association.

	return nil
}

func (s *PlaceServiceImpl) AddAmenitiesToPlace(ctx context.Context, placeID string, amenities []amenities.CreateAmenityInput) error {
	// Step 1: Fetch the existing amenities for the place.
	existingAmenities, err := s.client.Place.
		Query().
		Where(place.ID(placeID)).
		QueryAmenities().
		All(ctx)

	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	existingAmenityNames := make(map[string]struct{})
	for _, amenity := range existingAmenities {
		existingAmenityNames[amenity.Name] = struct{}{}
	}

	// Step 2: Initialize a slice to hold amenity references.
	var amenityList []*ent.Amenity

	// Step 3: Iterate over each amenity in the input.
	for _, amenityInput := range amenities {
		// If the amenity already exists for the place, skip.
		if _, exists := existingAmenityNames[amenityInput.Name]; exists {
			continue
		}

		// Check if the amenity with the provided name exists globally.
		amenity, err := s.client.Amenity.
			Query().
			Where(amenity.NameEQ(amenityInput.Name)).
			Only(ctx)

		if err != nil && !ent.IsNotFound(err) {
			sentry.CaptureException(err)
			return err
		}

		if amenity != nil {
			amenityList = append(amenityList, amenity)
		} else {
			// Create the amenity.
			newAmenity, err := s.client.Amenity.
				Create().
				SetName(amenityInput.Name).
				SetIcon(amenityInput.Icon).
				Save(ctx)

			if err != nil {
				sentry.CaptureException(err)
				return err
			}

			amenityList = append(amenityList, newAmenity)
		}
	}

	// Step 4: Associate the new amenities with the place.
	if len(amenityList) > 0 {
		placeData, err := s.client.Place.
			UpdateOneID(placeID).
			AddAmenities(amenityList...).
			Save(ctx)

		if err != nil {
			sentry.CaptureException(err)
			return err
		}

		go s.addPlaceToCacheAndSearchIndex(ctx, placeData)
	}

	return nil
}

func (s *PlaceServiceImpl) RemoveMediaToPlace(ctx context.Context, placeID string, mediaIDs []string) error {
	// Fetch place
	place, err := s.client.Place.Get(ctx, placeID)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	// Create a map to quickly lookup media IDs
	mediaLookup := make(map[string]bool)
	for _, mID := range mediaIDs {
		mediaLookup[mID] = true
	}

	// Check if the mediaIDs are associated with the place
	var associatedMediaIDs []string
	for _, associatedMedia := range place.Edges.Medias {
		if mediaLookup[associatedMedia.ID] {
			associatedMediaIDs = append(associatedMediaIDs, associatedMedia.ID)
		}
	}

	if len(associatedMediaIDs) == 0 {
		// None of the provided media IDs are associated with the place
		return errors.IDMissing
	}

	// Remove media from the place
	_, err = s.client.Place.UpdateOneID(placeID).RemoveMediaIDs(associatedMediaIDs...).Save(ctx)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	// TODO: Check if the media is no longer associated with any other places.
	// If not, then remove the media from Cloudinary or your media hosting service.

	// Update the place in cache and search index
	go s.addPlaceToCacheAndSearchIndex(ctx, place)

	return nil
}

//func (s *PlaceServiceImpl) AddAmenitiesToPlace(ctx context.Context, placeID string, amenityIDs []string) error {
//	// Fetch place
//	place, err := s.client.Place.Get(ctx, placeID)
//	if err != nil {
//		return err
//	}
//
//	// Fetch amenities
//	wg := sync.WaitGroup{}
//	wg.Add(len(amenityIDs))
//	amenities := make([]*ent.Amenity, len(amenityIDs))
//	for i, amenityID := range amenityIDs {
//		go func(i int, amenityID string) {
//			defer wg.Done()
//			amenities[i], err = s.client.Amenity.Get(ctx, amenityID)
//			if err != nil {
//				sentry.CaptureException(err)
//				return
//			}
//		}(i, amenityID)
//		wg.Wait()
//
//	}
//
//	// Update place with new amenities
//	placeData, err := place.Update().AddAmenities(amenities...).Save(ctx)
//	if err != nil {
//		sentry.CaptureException(err)
//		return err
//	}
//
//	// Add the new place to the search index and cache
//	go s.addPlaceToCacheAndSearchIndex(ctx, placeData)
//
//	return nil
//}

func (s *PlaceServiceImpl) UpdatePlace(ctx context.Context, placeID string, placeData UpdatePlaceDTO) (*ent.Place, error) {

	place, err := s.client.Place.
		UpdateOneID(placeID).
		SetDescription(placeData.Description).
		SetName(placeData.Name).
		SetPicture(placeData.Picture).
		SetCoverImage(placeData.CoverImage).
		SetWebsite(placeData.Website).
		SetEmail(placeData.Email).
		SetPhone(placeData.Phone).
		SetPlaceSettings(placeData.PlaceSettings).
		SetOpeningHours(placeData.OpeningHours).
		SetSocialMedia(placeData.SocialMedia).
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

func (s *PlaceServiceImpl) GetPlaces(ctx context.Context, filter *PlaceFilter, lastId string, limit int) ([]*ent.Place, string, error) {
	if limit == 0 {
		limit = 10
	}
	query := s.client.Place.
		Query().
		Order(ent.Asc(place.FieldID)).
		WithBusiness().
		WithUsers().
		Limit(limit + 1)

	// Apply filters
	if filter.Name != "" {
		query = query.Where(place.Or(place.NameEQ(filter.Name), place.NameContains(filter.Name), place.Name(filter.Name)))
	}

	if filter.Type != "" {
		query = query.Where(place.Or(place.TypeEQ(filter.Type), place.TypeContains(filter.Type), place.Type(filter.Type)))
	}

	if filter.Country != "" {
		query = query.Where(place.Or(place.CountryEQ(filter.Country), place.CountryContains(filter.Country), place.Country(filter.Country)))
	}

	if filter.City != "" {
		query = query.Where(place.Or(place.CityEQ(filter.City), place.CityContains(filter.City), place.City(filter.City)))
	}

	if filter.State != "" {
		query = query.Where(place.Or(place.StateEQ(filter.State), place.StateContains(filter.State), place.State(filter.State)))
	}

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
	var filteredPlaces []*ent.Place
	for _, p := range places {
		// Check if place tags include the tag from the filter
		if filter.Tags != "" && !contains(p.Tags, filter.Tags) {
			continue
		}
		// Check if place features include the feature from the filter
		if filter.Features != "" && !contains(p.Features, filter.Features) {
			continue
		}
		filteredPlaces = append(filteredPlaces, p)
	}

	places = filteredPlaces

	return places, nextId, nil
}

// contains checks if the set contains the element
func contains(set []string, element string) bool {
	for _, s := range set {
		if s == element {
			return true
		}
	}
	return false
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

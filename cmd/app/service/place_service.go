package service

import (
	"context"
	"github.com/google/uuid"
	"placio-app/Dto"
	"placio-app/ent"
	"placio-app/ent/business"
	"placio-app/ent/place"
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
	GetPlaces(ctx context.Context, filter *PlaceFilter, page int, pageSize int) ([]*ent.Place, error)
}

type PlaceServiceImpl struct {
	client        *ent.Client
	searchService SearchService
}

func NewPlaceService(client *ent.Client, searchService SearchService) *PlaceServiceImpl {
	return &PlaceServiceImpl{client: client, searchService: searchService}
}

func (s *PlaceServiceImpl) GetPlacesAssociatedWithBusinessAccount(c context.Context, businessId string) ([]*ent.Place, error) {
	places, err := s.client.Place.
		Query().
		Where(place.HasBusinessWith(business.ID(businessId))).
		All(c)
	if err != nil {
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
		First(ctx)
	if err != nil {
		return nil, err
	}

	return placeData, nil
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

	place, err := s.client.Place.
		Create().
		SetID(uuid.New().String()).
		SetName(placeData.Name).
		SetID(uuid.New().String()).
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

func (s *PlaceServiceImpl) GetPlaces(ctx context.Context, filter *PlaceFilter, page int, pageSize int) ([]*ent.Place, error) {
	query := s.client.Place.
		Query().
		WithBusiness().
		WithUsers()

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

	// Execute query
	places, err := query.All(ctx)
	if err != nil {
		return nil, err
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

	// Apply pagination
	// Validate page and pageSize values
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	// Apply pagination after filtering by tags and features
	start := (page - 1) * pageSize
	if start >= len(filteredPlaces) {
		return []*ent.Place{}, nil // Return an empty slice if page is beyond the total number of pages
	}
	end := start + pageSize
	if end > len(filteredPlaces) {
		end = len(filteredPlaces)
	}
	paginatedPlaces := filteredPlaces[start:end]

	return paginatedPlaces, nil
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

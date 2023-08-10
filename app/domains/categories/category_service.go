package categories

import (
	"context"
	"github.com/google/uuid"
	"log"
	"mime/multipart"
	media2 "placio-app/domains/media"
	"placio-app/ent"
	"placio-app/ent/category"
	"placio-app/ent/categoryassignment"
	"placio-app/ent/media"
	"placio-app/ent/menu"
	"placio-app/ent/post"
)

type CategoryService interface {
	GetUsersByCategory(ctx context.Context, name string) ([]*ent.User, error)
	CreateCategory(ctx context.Context, icon string, name string, image []*multipart.FileHeader) (*ent.Category, error)
	UpdateCategory(ctx context.Context, id string, name string, image string) (*ent.Category, error)
	DeleteCategory(ctx context.Context, id string) error
	GetAllCategories(ctx context.Context) ([]*ent.Category, error)
	GetCategory(ctx context.Context, id string) (*ent.Category, error)
	GetPostsByCategory(ctx context.Context, name string) ([]*ent.Post, error)
	GetMediasByCategory(ctx context.Context, name string) ([]*ent.Media, error)
	GetMenusByCategory(ctx context.Context, name string) ([]*ent.Menu, error)
	GetEntitiesByCategory(ctx context.Context, name, lastId string, limit int) (*CategorySearchResult, error)
	AssignBusinessToCategory(ctx context.Context, businessID string, categoryID string) (*ent.CategoryAssignment, error)
	AssignPlaceToCategory(ctx context.Context, placeID string, categoryID string) (*ent.CategoryAssignment, error)
	AssignUserToCategory(ctx context.Context, userID string, categoryID string) (*ent.CategoryAssignment, error)
	GetPlacesByCategory(ctx context.Context, categoryID, lastId string, limit int) ([]*ent.Place, string, error)
	CreateCategoryAssignment(ctx context.Context, entityID string, entityType string, categoryID string) (*ent.CategoryAssignment, error)
	GetBusinessesByCategory(ctx context.Context, categoryID, lastId string, limit int) ([]*ent.Business, string, error)
}

type EntityType int

const (
	Place EntityType = iota
	Business
	Event
)

type EntityQueryFunc func(*ent.Query) *ent.Query

type CategoryServiceImpl struct {
	client       *ent.Client
	mediaService media2.MediaService
}

func NewCategoryService(client *ent.Client, mediaService media2.MediaService) CategoryService {
	return &CategoryServiceImpl{client: client, mediaService: mediaService}
}

type CategorySearchResult struct {
	Users         []*ent.User
	Posts         []*ent.Post
	Medias        []*ent.Media
	Menus         []*ent.Menu
	Places        []*ent.Place
	Businesses    []*ent.Business
	NextPageToken string "json:nextPageToken"
}

func (cs *CategoryServiceImpl) CreateCategory(ctx context.Context, icon string, name string, image []*multipart.FileHeader) (*ent.Category, error) {
	// Create the category without the image first
	category, err := cs.client.Category.
		Create().
		SetID(uuid.New().String()).
		SetIcon(icon).
		SetName(name).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	if len(image) > 0 {
		go func() {
			media, err := cs.mediaService.UploadAndCreateMedia(ctx, image)
			if err != nil {
				log.Println("error uploading image", err.Error())
				return
			}

			_, err = cs.client.Category.
				UpdateOneID(category.ID).
				SetImage(media[0].URL).
				Save(ctx)

			if err != nil {
				log.Println("error adding image to category", err.Error())
				return
			}
		}()
	}

	return category, nil
}

func (cs *CategoryServiceImpl) UpdateCategory(ctx context.Context, id string, name string, image string) (*ent.Category, error) {
	return cs.client.Category.
		UpdateOneID(id).
		SetName(name).
		SetImage(image).
		Save(ctx)
}

func (cs *CategoryServiceImpl) DeleteCategory(ctx context.Context, id string) error {
	return cs.client.Category.
		DeleteOneID(id).
		Exec(ctx)
}

func (cs *CategoryServiceImpl) GetAllCategories(ctx context.Context) ([]*ent.Category, error) {
	return cs.client.Category.
		Query().
		All(ctx)
}

func (cs *CategoryServiceImpl) GetCategory(ctx context.Context, id string) (*ent.Category, error) {
	return cs.client.Category.
		Get(ctx, id)
}

func (cs *CategoryServiceImpl) SearchByCategory(ctx context.Context, name string) {
	cs.client.Category.
		Query().
		Where(category.Name(name)).
		All(ctx)
}

func (cs *CategoryServiceImpl) GetPostsByCategory(ctx context.Context, name string) ([]*ent.Post, error) {
	return cs.client.Post.
		Query().
		Where(post.HasCategoriesWith(category.Name(name))).
		All(ctx)
}

func (cs *CategoryServiceImpl) GetMediasByCategory(ctx context.Context, name string) ([]*ent.Media, error) {
	return cs.client.Media.
		Query().
		Where(media.HasCategoriesWith(category.Name(name))).
		All(ctx)
}

func (cs *CategoryServiceImpl) GetMenusByCategory(ctx context.Context, name string) ([]*ent.Menu, error) {
	return cs.client.Menu.
		Query().
		Where(menu.HasCategoriesWith(category.Name(name))).
		All(ctx)
}

func (cs *CategoryServiceImpl) CreateCategoryAssignment(ctx context.Context, entityID string, entityType string, categoryID string) (*ent.CategoryAssignment, error) {
	return cs.client.CategoryAssignment.
		Create().
		SetEntityID(entityID).
		SetEntityType(entityType).
		SetCategoryID(categoryID).
		Save(ctx)
}

func (cs *CategoryServiceImpl) AssignBusinessToCategory(ctx context.Context, businessID string, categoryID string) (*ent.CategoryAssignment, error) {
	return cs.CreateCategoryAssignment(ctx, businessID, "Business", categoryID)
}

func (cs *CategoryServiceImpl) AssignPlaceToCategory(ctx context.Context, placeID string, categoryID string) (*ent.CategoryAssignment, error) {
	return cs.CreateCategoryAssignment(ctx, placeID, "Place", categoryID)
}

func (cs *CategoryServiceImpl) AssignUserToCategory(ctx context.Context, userID string, categoryID string) (*ent.CategoryAssignment, error) {
	return cs.CreateCategoryAssignment(ctx, userID, "User", categoryID)
}

func (cs *CategoryServiceImpl) GetEntitiesByCategory(ctx context.Context, name, lastId string, limit int) (*CategorySearchResult, error) {
	//users, nextPageToken, err := cs.GetUsersByCategory(ctx, name, lastId, limit)
	//if err != nil {
	//	return nil, err
	//}

	//posts, nextPageToken, err := cs.GetPostsByCategory(ctx, name, lastId, limit)
	//if err != nil {
	//	return nil, err
	//}
	//
	//medias, nextPageToken, err := cs.GetMediasByCategory(ctx, name, lastId, limit)
	//if err != nil {
	//	return nil, err
	//}
	//
	//menus, nextPageToken, err := cs.GetMenusByCategory(ctx, name, lastId, limit)
	//if err != nil {
	//	return nil, err
	//}

	places, nextPageToken, err := cs.GetPlacesByCategory(ctx, name, lastId, limit)
	if err != nil {
		return nil, err
	}

	businesses, nextPageToken, err := cs.GetBusinessesByCategory(ctx, name, lastId, limit)
	if err != nil {
		return nil, err
	}

	result := &CategorySearchResult{
		//Users:         users,
		//Posts:         posts,
		//Medias:        medias,
		//Menus:         menus,
		Places:        places,
		Businesses:    businesses,
		NextPageToken: nextPageToken,
	}

	return result, nil
}

func (cs *CategoryServiceImpl) GetBusinessesByCategory(ctx context.Context, categoryID, lastId string, limit int) ([]*ent.Business, string, error) {
	query := cs.client.CategoryAssignment.
		Query().
		Where(categoryassignment.CategoryID(categoryID), categoryassignment.EntityType("Business")).
		Limit(limit + 1) // Retrieve one extra record to determine if there are more pages

	if lastId != "" {
		query = query.Where(categoryassignment.IDGT(lastId)) // Fetch records after the lastId
	}

	assignments, err := query.All(ctx)
	if err != nil {
		return nil, "", err
	}

	var businesses []*ent.Business
	for _, assignment := range assignments {
		business, err := cs.client.Business.Get(ctx, assignment.EntityID)
		if err != nil {
			return nil, "", err
		}
		businesses = append(businesses, business)
	}

	var nextId string
	if len(businesses) == limit+1 {
		nextId = businesses[len(businesses)-1].ID
		businesses = businesses[:limit]
	}

	return businesses, nextId, nil
}

func (cs *CategoryServiceImpl) GetPlacesByCategory(ctx context.Context, categoryID, lastId string, limit int) ([]*ent.Place, string, error) {
	query := cs.client.CategoryAssignment.
		Query().
		Where(categoryassignment.CategoryID(categoryID), categoryassignment.EntityType("Place")).
		Limit(limit + 1) // Retrieve one extra record to determine if there are more pages

	if lastId != "" {
		query = query.Where(categoryassignment.IDGT(lastId)) // Fetch records after the lastId
	}

	assignments, err := query.All(ctx)
	if err != nil {
		return nil, "", err
	}

	var places []*ent.Place
	for _, assignment := range assignments {
		place, err := cs.client.Place.Get(ctx, assignment.EntityID)
		if err != nil {
			return nil, "", err
		}
		places = append(places, place)
	}

	var nextId string
	if len(places) == limit+1 {
		nextId = places[len(places)-1].ID
		places = places[:limit]
	}

	return places, nextId, nil
}

func (cs *CategoryServiceImpl) GetUsersByCategory(ctx context.Context, categoryID string) ([]*ent.User, error) {
	assignments, err := cs.client.CategoryAssignment.
		Query().
		Where(categoryassignment.CategoryID(categoryID), categoryassignment.EntityType("User")).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var users []*ent.User
	for _, assignment := range assignments {
		user, err := cs.client.User.Get(ctx, assignment.EntityID)
		if err != nil {
			return nil, err // Or you might want to continue and skip this one.
		}
		users = append(users, user)
	}

	return users, nil
}

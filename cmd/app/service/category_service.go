package service

import (
	"context"
	"placio-app/ent"
	"placio-app/ent/business"
	"placio-app/ent/category"
	"placio-app/ent/media"
	"placio-app/ent/menu"
	"placio-app/ent/place"
	"placio-app/ent/post"
	"placio-app/ent/user"
)

type CategoryService interface {
	GetUsersByCategory(ctx context.Context, name string) ([]*ent.User, error)
	CreateCategory(ctx context.Context, id string, name string, image string) (*ent.Category, error)
	UpdateCategory(ctx context.Context, id string, name string, image string) (*ent.Category, error)
	DeleteCategory(ctx context.Context, id string) error
	GetCategory(ctx context.Context, id string) (*ent.Category, error)
	GetPostsByCategory(ctx context.Context, name string) ([]*ent.Post, error)
	GetMediasByCategory(ctx context.Context, name string) ([]*ent.Media, error)
	GetMenusByCategory(ctx context.Context, name string) ([]*ent.Menu, error)
	GetPlacesByCategory(ctx context.Context, name string) ([]*ent.Place, error)
	GetBusinessesByCategory(ctx context.Context, name string) ([]*ent.Business, error)
	GetEntitiesByCategory(ctx context.Context, name string) (*CategorySearchResult, error)
}

type CategoryServiceImpl struct {
	client *ent.Client
}

func NewCategoryService(client *ent.Client) CategoryService {
	return &CategoryServiceImpl{client: client}
}

type CategorySearchResult struct {
	Users      []*ent.User
	Posts      []*ent.Post
	Medias     []*ent.Media
	Menus      []*ent.Menu
	Places     []*ent.Place
	Businesses []*ent.Business
}

func (cs *CategoryServiceImpl) CreateCategory(ctx context.Context, id string, name string, image string) (*ent.Category, error) {
	return cs.client.Category.
		Create().
		SetID(id).
		SetName(name).
		SetImage(image).
		Save(ctx)
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

func (cs *CategoryServiceImpl) GetUsersByCategory(ctx context.Context, name string) ([]*ent.User, error) {
	return cs.client.User.
		Query().
		Where(user.HasCategoriesWith(category.Name(name))).
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

func (cs *CategoryServiceImpl) GetPlacesByCategory(ctx context.Context, name string) ([]*ent.Place, error) {
	return cs.client.Place.
		Query().
		Where(place.HasCategoriesWith(category.Name(name))).
		All(ctx)
}

func (cs *CategoryServiceImpl) GetBusinessesByCategory(ctx context.Context, name string) ([]*ent.Business, error) {
	return cs.client.Business.
		Query().
		Where(business.HasCategoriesWith(category.Name(name))).
		All(ctx)
}

func (cs *CategoryServiceImpl) GetEntitiesByCategory(ctx context.Context, name string) (*CategorySearchResult, error) {
	users, err := cs.GetUsersByCategory(ctx, name)
	if err != nil {
		return nil, err
	}

	posts, err := cs.GetPostsByCategory(ctx, name)
	if err != nil {
		return nil, err
	}

	medias, err := cs.GetMediasByCategory(ctx, name)
	if err != nil {
		return nil, err
	}

	menus, err := cs.GetMenusByCategory(ctx, name)
	if err != nil {
		return nil, err
	}

	places, err := cs.GetPlacesByCategory(ctx, name)
	if err != nil {
		return nil, err
	}

	businesses, err := cs.GetBusinessesByCategory(ctx, name)
	if err != nil {
		return nil, err
	}

	result := &CategorySearchResult{
		Users:      users,
		Posts:      posts,
		Medias:     medias,
		Menus:      menus,
		Places:     places,
		Businesses: businesses,
	}

	return result, nil
}

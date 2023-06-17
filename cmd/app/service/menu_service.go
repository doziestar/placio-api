package service

import (
	"context"
	"placio-app/ent"
)

type MenuService interface {
	GetMenu(ctx context.Context, menuID string) (*ent.Menu, error)
	CreateMenu(ctx context.Context, menuData map[string]interface{}) (*ent.Menu, error)
	UpdateMenu(ctx context.Context, menuID string, menuData map[string]interface{}) (*ent.Menu, error)
	DeleteMenu(ctx context.Context, menuID string) error
}

type MenuServiceImpl struct {
	client *ent.Client
}

func NewMenuService(client *ent.Client) *MenuServiceImpl {
	return &MenuServiceImpl{client: client}
}

func (s *MenuServiceImpl) GetMenu(ctx context.Context, menuID string) (*ent.Menu, error) {
	return s.client.Menu.Get(ctx, menuID)
}

func (s *MenuServiceImpl) CreateMenu(ctx context.Context, menuData map[string]interface{}) (*ent.Menu, error) {
	return s.client.Menu.
		Create().
		//SetName(menuData["name"].(string)).
		Save(ctx)
}

func (s *MenuServiceImpl) UpdateMenu(ctx context.Context, menuID string, menuData map[string]interface{}) (*ent.Menu, error) {
	return s.client.Menu.
		UpdateOneID(menuID).
		//SetName(menuData["name"].(string)).
		Save(ctx)
}

func (s *MenuServiceImpl) DeleteMenu(ctx context.Context, menuID string) error {
	return s.client.Menu.
		DeleteOneID(menuID).
		Exec(ctx)
}

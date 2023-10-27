package smartMenu

import (
	"context"
	"github.com/google/uuid"
	"placio-app/ent"
	"placio-app/ent/menu"
	"placio-app/ent/place"
	"placio-app/ent/placetable"
	"time"
)

type SmartMenuService struct {
	client *ent.Client
}

type ISmartMenu interface {
	CreateMenu(context.Context, string, *ent.Menu) (*ent.Menu, error)
	GetMenus(context.Context, string) ([]*ent.Menu, error)
	GetMenuByID(context.Context, string) (*ent.Menu, error)
	UpdateMenu(context.Context, string, *ent.Menu) (*ent.Menu, error)
	DeleteMenu(context.Context, string) error
	RestoreMenu(context.Context, string) (*ent.Menu, error)

	CreateTable(context.Context, string, *ent.PlaceTable) (*ent.PlaceTable, error)
	GetTables(context.Context, string) ([]*ent.PlaceTable, error)
	GetTableByID(context.Context, string) (*ent.PlaceTable, error)
	UpdateTable(context.Context, string, *ent.PlaceTable) (*ent.PlaceTable, error)
	DeleteTable(context.Context, string) error
	RestoreTable(context.Context, string) (*ent.PlaceTable, error)
	RegenerateQRCode(context.Context, string) (*ent.PlaceTable, error)

	CreateOrder(context.Context, string, string, *ent.Order) (*ent.Order, error)
	GetOrdersForATable(context.Context, string) ([]*ent.Order, error)
	GetOrders(context.Context, string) ([]*ent.Order, error)
	GetOrderByID(context.Context, string) (*ent.Order, error)
	UpdateOrder(context.Context, string, *ent.Order) (*ent.Order, error)
	DeleteOrder(context.Context, string) error
	RestoreOrder(context.Context, string) (*ent.Order, error)
}

func NewSmartMenuService(client *ent.Client) *SmartMenuService {
	return &SmartMenuService{client: client}
}

func (s *SmartMenuService) CreateMenu(ctx context.Context, placeId string, menu *ent.Menu) (*ent.Menu, error) {
	return s.client.Menu.Create().
		SetName(menu.Name).
		SetPlaceID(placeId).
		Save(ctx)
}

func (s *SmartMenuService) GetMenus(ctx context.Context, placeId string) ([]*ent.Menu, error) {
	return s.client.Place.
		Query().
		Where(place.ID(placeId)).
		QueryMenus().
		All(ctx)
}

func (s *SmartMenuService) GetMenuByID(ctx context.Context, menuId string) (*ent.Menu, error) {
	return s.client.Menu.
		Query().
		Where(menu.ID(menuId)).
		Only(ctx)
}

func (s *SmartMenuService) UpdateMenu(ctx context.Context, menuId string, menu *ent.Menu) (*ent.Menu, error) {
	return s.client.Menu.
		UpdateOneID(menuId).
		SetName(menu.Name).
		Save(ctx)
}

func (s *SmartMenuService) DeleteMenu(ctx context.Context, menuId string) error {
	return s.client.Menu.
		UpdateOneID(menuId).
		SetDeletedAt(time.Now().String()).
		SetIsDeleted(true).
		Exec(ctx)
}

func (s *SmartMenuService) RestoreMenu(ctx context.Context, menuId string) (*ent.Menu, error) {
	return s.client.Menu.
		UpdateOneID(menuId).
		ClearDeletedAt().
		SetIsDeleted(false).
		Save(ctx)
}

func (s *SmartMenuService) CreateTable(ctx context.Context, placeId string, table *ent.PlaceTable) (*ent.PlaceTable, error) {
	return s.client.PlaceTable.
		Create().
		SetNumber(table.Number).
		SetPlaceID(placeId).
		Save(ctx)
}

func (s *SmartMenuService) GetTables(ctx context.Context, placeId string) ([]*ent.PlaceTable, error) {
	return s.client.Place.
		Query().
		Where(place.ID(placeId)).
		QueryTables().
		All(ctx)
}

func (s *SmartMenuService) GetTableByID(ctx context.Context, tableId string) (*ent.PlaceTable, error) {
	return s.client.PlaceTable.
		Query().
		Where(placetable.ID(tableId)).
		Only(ctx)
}

func (s *SmartMenuService) UpdateTable(ctx context.Context, tableId string, table *ent.PlaceTable) (*ent.PlaceTable, error) {
	return s.client.PlaceTable.
		UpdateOneID(tableId).
		SetNumber(table.Number).
		Save(ctx)
}

func (s *SmartMenuService) DeleteTable(ctx context.Context, tableId string) error {
	return s.client.PlaceTable.
		UpdateOneID(tableId).
		SetDeletedAt(time.Now().String()).
		SetIsDeleted(true).
		Exec(ctx)
}

func (s *SmartMenuService) RestoreTable(ctx context.Context, tableId string) (*ent.PlaceTable, error) {
	return s.client.PlaceTable.
		UpdateOneID(tableId).
		ClearDeletedAt().
		SetIsDeleted(false).
		Save(ctx)
}

func (s *SmartMenuService) RegenerateQRCode(ctx context.Context, tableId string) (*ent.PlaceTable, error) {
	//return s.client.PlaceTable.
	//	UpdateOneID(tableId).
	//	SetQRCode("").
	//	Save(ctx)
	return nil, nil
}

func (s *SmartMenuService) CreateOrder(ctx context.Context, tableId string, placeId string, order *ent.Order) (*ent.Order, error) {
	return s.client.Order.
		Create().
		SetID(uuid.New().String()).
		Save(ctx)
}

func (s *SmartMenuService) GetOrdersForATable(ctx context.Context, s2 string) ([]*ent.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SmartMenuService) GetOrders(ctx context.Context, s2 string) ([]*ent.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SmartMenuService) GetOrderByID(ctx context.Context, s2 string) (*ent.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SmartMenuService) UpdateOrder(ctx context.Context, s2 string, order *ent.Order) (*ent.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SmartMenuService) DeleteOrder(ctx context.Context, s2 string) error {
	//TODO implement me
	panic("implement me")
}

func (s *SmartMenuService) RestoreOrder(ctx context.Context, s2 string) (*ent.Order, error) {
	//TODO implement me
	panic("implement me")
}

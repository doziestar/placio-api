package smartMenu

import (
	"context"
	"log"
	"mime/multipart"
	"placio-app/domains/media"
	"placio-app/ent"
	"placio-app/ent/category"
	"placio-app/ent/menu"
	"placio-app/ent/place"
	"placio-app/ent/placetable"
	"strings"
	"time"

	"github.com/google/uuid"
)

type SmartMenuService struct {
	client *ent.Client
	mediaService media.MediaService
}

type ISmartMenu interface {
	CreateMenu(context.Context, string, *ent.Menu, string, []*multipart.FileHeader)(*ent.Menu, error)
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

func NewSmartMenuService(client *ent.Client, mediaService media.MediaService) *SmartMenuService {
	return &SmartMenuService{client: client, mediaService: mediaService}
}

func (s *SmartMenuService) CreateMenu(ctx context.Context, placeId string, menuDto *ent.Menu, cat string, medias []*multipart.FileHeader) (*ent.Menu, error) {
	lowercaseCategory := strings.ToLower(cat)

	// Ensure the category exists in the database, creating it if necessary
	catData, err := s.ensureCategoryExists(ctx, lowercaseCategory)
	if err != nil {
		return nil, err
	}

	// Check if the menu already exists in the database for this place
	existingMenu, err := s.client.Menu.
		Query().
		Where(menu.Name(strings.ToLower(menuDto.Name))).
		Where(menu.HasPlaceWith(place.ID(placeId))).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		log.Println("Error checking if menu with name", menuDto.Name, "exists for place", placeId, ":", err)
		return nil, err
	}
	if existingMenu != nil {
		log.Println("Menu with name", menuDto.Name, "already exists for place", placeId)
		return existingMenu, nil
	}

	log.Println("Creating menu with name", menuDto.Name)

	// Create and return the new menu
	newMenu, err := s.client.Menu.
		Create().
		SetID(uuid.New().String()).
		SetName(menuDto.Name).
		SetDescription(menuDto.Description).
		SetPreparationTime(menuDto.PreparationTime).
		SetOptions(menuDto.Options).
		SetPrice(menuDto.Price).
		SetIsAvailable(menuDto.IsAvailable).
		AddCategories(catData).
		AddPlaceIDs(placeId).
		Save(ctx)


	if err != nil {
		log.Println("Error creating menu with name", menuDto.Name, ":", err)
		return nil, err
	}

	// Create the media for the menu
	if len(medias) > 0 {
		go func(menuID string, mediaFiles []*multipart.FileHeader) {
			ctx := context.Background() // Use a background context for the asynchronous operation
			log.Println("Uploading media for menu with ID", menuID)
			
			media, err := s.mediaService.UploadAndCreateMedia(ctx, mediaFiles)
			if err != nil {
				log.Println("Error uploading media for menu with ID", menuID, ":", err)
				return
			}
			
			_, err = s.client.Menu.
				UpdateOneID(menuID).
				AddMedia(media...).
				Save(ctx)
			if err != nil {
				log.Println("Error adding media to menu with ID", menuID, ":", err)
				return
			}
			
			log.Println("Media uploaded and added to menu with ID", menuID)
		}(newMenu.ID, medias)
	}

	return newMenu, nil
}

func (s *SmartMenuService) ensureCategoryExists(ctx context.Context, categoryName string) (*ent.Category, error) {
	catData, err := s.client.Category.
		Query().
		Where(category.Name(categoryName)).
		Only(ctx)
	if ent.IsNotFound(err) {
		return s.client.Category.
			Create().
			SetName(categoryName).
			Save(ctx)
	}
	return catData, err
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
		WithMedia().
		WithPlace(func(q *ent.PlaceQuery) {
			q.WithBusiness()
		}).
		WithCategories(func(q *ent.CategoryQuery) {
		}).
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

package inventory

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"placio-app/domains/business"
	"placio-app/domains/cache"
	"placio-app/domains/media"
	"placio-app/domains/places"
	"placio-app/ent"
	"placio-app/ent/inventoryattribute"
	"placio-app/ent/inventorytype"
	"placio-app/ent/place"
	"placio-app/ent/placeinventory"
	"placio-app/ent/placeinventoryattribute"
	"placio-pkg/errors"
	"time"
)

type InventoryService interface {
	CreateInventoryType(ctx context.Context, data *inventoryTypeData) (*ent.InventoryType, error)
	UpdateInventoryType(ctx context.Context, inventoryTypeID string, data *inventoryTypeData) (*ent.InventoryType, error)
	DeleteInventoryType(ctx context.Context, inventoryTypeID string) error
	ListInventoryTypes(ctx context.Context, limit, offset int) ([]*ent.InventoryType, error)
	CreateInventoryAttribute(ctx context.Context, data *inventoryAttributeData) (*ent.InventoryAttribute, error)
	UpdateInventoryAttribute(ctx context.Context, inventoryAttributeID string, data *inventoryAttributeData) (*ent.InventoryAttribute, error)
	DeleteInventoryAttribute(ctx context.Context, inventoryAttributeID string) error
	ListInventoryAttributes(ctx context.Context, limit, offset int) ([]*ent.InventoryAttribute, error)
	SearchInventoryAttributes(ctx context.Context, query string, limit, offset int) ([]*ent.InventoryAttribute, error)
	CreatePlaceInventory(ctx context.Context, placeID string, data *placeInventoryData) (*ent.PlaceInventory, error)
	GetPlaceInventory(ctx context.Context, placeInventoryID string) (*ent.PlaceInventory, error)
	UpdatePlaceInventory(ctx context.Context, placeInventoryID string, data *placeInventoryData) (*ent.PlaceInventory, error)
	DeletePlaceInventory(ctx context.Context, placeInventoryID string) error
	ListPlaceInventories(ctx context.Context, placeID string, limit, offset int) ([]*ent.PlaceInventory, error)
	CreatePlaceInventoryAttribute(ctx context.Context, placeInventoryID string, data *placeInventoryAttributeData) (*ent.PlaceInventoryAttribute, error)
	GetPlaceInventoryAttribute(ctx context.Context, placeInventoryAttributeID string) (*ent.PlaceInventoryAttribute, error)
	UpdatePlaceInventoryAttribute(ctx context.Context, placeInventoryAttributeID string, data *placeInventoryAttributeData) (*ent.PlaceInventoryAttribute, error)
	DeletePlaceInventoryAttribute(ctx context.Context, placeInventoryAttributeID string) error
	//ListPlaceInventoryAttributes(ctx context.Context, placeInventoryID string, limit, offset int) ([]*ent.PlaceInventoryAttribute, error)
	//AddInventoryToPlace(ctx context.Context, placeID string, data *PlaceInventoryData) (*ent.PlaceInventory, error)
	//AddAttributeToPlaceInventory(ctx context.Context, placeInventoryID string, data *PlaceInventoryAttributeData) (*ent.PlaceInventoryAttribute, error)
}

type InventoryServiceImpl struct {
	client          *ent.Client
	cache           *cache.CacheService
	mediaService    *media.MediaService
	placeService    places.PlaceService
	businessService *business.BusinessAccountService
}

func NewInventoryService(client *ent.Client, cache *cache.CacheService, mediaService *media.MediaService, placeService places.PlaceService, businessService *business.BusinessAccountService) *InventoryServiceImpl {
	return &InventoryServiceImpl{client: client, cache: cache, mediaService: mediaService, placeService: placeService, businessService: businessService}
}

func (s *InventoryServiceImpl) CreateInventoryType(ctx context.Context, data *inventoryTypeData) (*ent.InventoryType, error) {

	if !IsValidIndustryType(data.IndustryType) {
		return nil, errors.ErrUnprocessable
	}

	return s.client.InventoryType.Create().
		SetName(data.Name).
		SetDescription(data.Description).
		SetMeasurementUnit(data.MeasurementUnit).
		SetIndustryType(inventorytype.IndustryType(data.IndustryType)).
		Save(ctx)
}

func (s *InventoryServiceImpl) UpdateInventoryType(ctx context.Context, inventoryTypeID string, data *inventoryTypeData) (*ent.InventoryType, error) {

	if !IsValidIndustryType(data.IndustryType) {
		return nil, errors.ErrUnprocessable
	}

	return s.client.InventoryType.UpdateOneID(inventoryTypeID).
		SetName(data.Name).
		SetDescription(data.Description).
		SetMeasurementUnit(data.MeasurementUnit).
		SetIndustryType(inventorytype.IndustryType(data.IndustryType)).
		Save(ctx)
}

func (s *InventoryServiceImpl) DeleteInventoryType(ctx context.Context, inventoryTypeID string) error {
	return s.client.InventoryType.DeleteOneID(inventoryTypeID).Exec(ctx)
}

func (s *InventoryServiceImpl) ListInventoryTypes(ctx context.Context, limit, offset int) ([]*ent.InventoryType, error) {
	if limit == 0 {
		limit = 10
	}
	return s.client.InventoryType.Query().Limit(limit).Offset(offset).All(ctx)
}

// CreateInventoryAttribute creates a new inventory attribute
func (s *InventoryServiceImpl) CreateInventoryAttribute(ctx context.Context, data *inventoryAttributeData) (*ent.InventoryAttribute, error) {
	inventoryType, err := s.client.InventoryType.Get(ctx, data.InventoryTypeID)
	if err != nil {
		sentry.CaptureException(err)
		return nil, errors.ErrNotFound
	}

	attribute, err := s.client.InventoryAttribute.Create().
		SetName(data.Name).
		SetIsMandatory(data.IsMandatory).
		SetInventoryType(inventoryType).
		Save(ctx)

	if err != nil {
		sentry.CaptureException(err)
		return nil, errors.ErrUnprocessable
	}

	return attribute, nil
}

// UpdateInventoryAttribute updates an existing inventory attribute
func (s *InventoryServiceImpl) UpdateInventoryAttribute(ctx context.Context, inventoryAttributeID string, data *inventoryAttributeData) (*ent.InventoryAttribute, error) {
	attribute, err := s.client.InventoryAttribute.UpdateOneID(inventoryAttributeID).
		SetName(data.Name).
		SetIsMandatory(data.IsMandatory).
		Save(ctx)

	if err != nil {
		sentry.CaptureException(err)
		return nil, errors.ErrUnprocessable
	}

	return attribute, nil
}

// DeleteInventoryAttribute deletes an existing inventory attribute
func (s *InventoryServiceImpl) DeleteInventoryAttribute(ctx context.Context, inventoryAttributeID string) error {
	return s.client.InventoryAttribute.DeleteOneID(inventoryAttributeID).Exec(ctx)
}

// ListInventoryAttributes lists all inventory attributes
func (s *InventoryServiceImpl) ListInventoryAttributes(ctx context.Context, limit, offset int) ([]*ent.InventoryAttribute, error) {
	if limit == 0 {
		limit = 10
	}
	return s.client.InventoryAttribute.Query().Limit(limit).Offset(offset).All(ctx)
}

// SearchInventoryAttributes searches inventory attributes by name
func (s *InventoryServiceImpl) SearchInventoryAttributes(ctx context.Context, query string, limit, offset int) ([]*ent.InventoryAttribute, error) {
	if limit == 0 {
		limit = 10
	}
	return s.client.InventoryAttribute.Query().Where(inventoryattribute.NameContains(query)).Limit(limit).Offset(offset).All(ctx)
}

// CreatePlaceInventory creates a new place inventory
func (s *InventoryServiceImpl) CreatePlaceInventory(ctx context.Context, placeID string, data *placeInventoryData) (*ent.PlaceInventory, error) {
	place, err := s.placeService.GetPlace(ctx, placeID)
	if err != nil {
		return nil, err
	}

	inventoryType, err := s.client.InventoryType.Get(ctx, data.InventoryTypeID)
	if err != nil {
		return nil, err
	}

	placeInventory, err := s.client.PlaceInventory.Create().
		SetName(data.Name).
		SetStockQuantity(data.Quantity).
		SetLastUpdated(time.Now()).
		SetSku(uuid.New().String()).
		SetColor(data.Color).
		SetBrand(data.Brand).
		SetSize(data.Size).
		SetBusiness(place.Edges.Business).
		SetInventoryType(inventoryType).
		SetPlace(place).
		Save(ctx)

	if err != nil {
		sentry.CaptureException(err)
		return nil, errors.ErrUnprocessable
	}

	return placeInventory, nil
}

// UpdatePlaceInventory updates an existing place inventory
func (s *InventoryServiceImpl) UpdatePlaceInventory(ctx context.Context, placeInventoryID string, data *placeInventoryData) (*ent.PlaceInventory, error) {
	placeInventory, err := s.client.PlaceInventory.UpdateOneID(placeInventoryID).
		SetName(data.Name).
		SetStockQuantity(data.Quantity).
		SetLastUpdated(time.Now()).
		SetSku(uuid.New().String()).
		SetColor(data.Color).
		SetBrand(data.Brand).
		SetSize(data.Size).
		Save(ctx)

	if err != nil {
		sentry.CaptureException(err)
		return nil, errors.ErrUnprocessable
	}

	return placeInventory, nil
}

// DeletePlaceInventory deletes an existing place inventory
func (s *InventoryServiceImpl) DeletePlaceInventory(ctx context.Context, placeInventoryID string) error {
	return s.client.PlaceInventory.DeleteOneID(placeInventoryID).Exec(ctx)
}

// ListPlaceInventories lists all place inventories
func (s *InventoryServiceImpl) ListPlaceInventories(ctx context.Context, placeID string, limit, offset int) ([]*ent.PlaceInventory, error) {
	if limit == 0 {
		limit = 10
	}
	return s.client.PlaceInventory.Query().Where(placeinventory.HasPlaceWith(place.ID(placeID))).Limit(limit).Offset(offset).All(ctx)
}

// SearchPlaceInventories searches place inventories by name
func (s *InventoryServiceImpl) SearchPlaceInventories(ctx context.Context, query string, limit, offset int) ([]*ent.PlaceInventory, error) {
	if limit == 0 {
		limit = 10
	}
	return s.client.PlaceInventory.Query().Where(placeinventory.NameContains(query)).Limit(limit).Offset(offset).All(ctx)
}

// CreatePlaceInventoryAttribute creates a new place inventory attribute
func (s *InventoryServiceImpl) CreatePlaceInventoryAttribute(ctx context.Context, placeInventoryID string, data *placeInventoryAttributeData) (*ent.PlaceInventoryAttribute, error) {
	placeInventory, err := s.client.PlaceInventory.Get(ctx, placeInventoryID)
	if err != nil {
		return nil, err
	}

	inventoryAttribute, err := s.client.InventoryAttribute.Get(ctx, data.InventoryAttributeID)
	if err != nil {
		return nil, err
	}

	placeInventoryAttribute, err := s.client.PlaceInventoryAttribute.Create().
		SetValue(data.Value).
		SetAttributeType(inventoryAttribute).
		SetInventory(placeInventory).
		Save(ctx)

	if err != nil {
		sentry.CaptureException(err)
		return nil, errors.ErrUnprocessable
	}

	return placeInventoryAttribute, nil
}

func (s *InventoryServiceImpl) GetPlaceInventory(ctx context.Context, placeInventoryID string) (*ent.PlaceInventory, error) {
	return s.client.PlaceInventory.Get(ctx, placeInventoryID)
}

// UpdatePlaceInventoryAttribute updates an existing place inventory attribute
func (s *InventoryServiceImpl) UpdatePlaceInventoryAttribute(ctx context.Context, placeInventoryAttributeID string, data *placeInventoryAttributeData) (*ent.PlaceInventoryAttribute, error) {
	placeInventoryAttribute, err := s.client.PlaceInventoryAttribute.UpdateOneID(placeInventoryAttributeID).
		SetValue(data.Value).
		Save(ctx)

	if err != nil {
		sentry.CaptureException(err)
		return nil, errors.ErrUnprocessable
	}

	return placeInventoryAttribute, nil
}

// DeletePlaceInventoryAttribute deletes an existing place inventory attribute
func (s *InventoryServiceImpl) DeletePlaceInventoryAttribute(ctx context.Context, placeInventoryAttributeID string) error {
	return s.client.PlaceInventoryAttribute.DeleteOneID(placeInventoryAttributeID).Exec(ctx)
}

// ListPlaceInventoryAttributes lists all place inventory attributes
func (s *InventoryServiceImpl) ListPlaceInventoryAttributes(ctx context.Context, limit, offset int) ([]*ent.PlaceInventoryAttribute, error) {
	if limit == 0 {
		limit = 10
	}
	return s.client.PlaceInventoryAttribute.Query().Limit(limit).Offset(offset).All(ctx)
}

// SearchPlaceInventoryAttributes searches place inventory attributes by name
func (s *InventoryServiceImpl) SearchPlaceInventoryAttributes(ctx context.Context, query string, limit, offset int) ([]*ent.PlaceInventoryAttribute, error) {
	if limit == 0 {
		limit = 10
	}
	return s.client.PlaceInventoryAttribute.Query().Where(placeinventoryattribute.ValueContains(query)).Limit(limit).Offset(offset).All(ctx)
}

func (s *InventoryServiceImpl) GetPlaceInventoryAttribute(ctx context.Context, placeInventoryAttributeID string) (*ent.PlaceInventoryAttribute, error) {
	return s.client.PlaceInventoryAttribute.Get(ctx, placeInventoryAttributeID)
}

//// CreateInventoryTransaction creates a new inventory transaction
//func (s *InventoryServiceImpl) CreateInventoryTransaction(ctx context.Context, data *InventoryTransactionData) (*ent.InventoryTransaction, error) {
//	placeInventory, err := s.client.PlaceInventory.Get(ctx, data.PlaceInventoryID)
//	if err != nil {
//		return nil, err
//	}
//
//	inventoryTransaction, err := s.client.InventoryTransaction.Create().
//		SetQuantity(data.Quantity).
//		SetType(data.Type).
//		SetPlaceInventory(placeInventory).
//		Save(ctx)
//
//	if err != nil {
//		sentry.CaptureException(err)
//		return nil, errors.ErrUnprocessable
//	}
//
//	return inventoryTransaction, nil
//}

// AddInventoryToPlace adds inventory to a place
//func (s *InventoryServiceImpl) AddInventoryToPlace(ctx context.Context, placeID, inventoryId string) (*ent.PlaceInventory, error) {
//	place, err := s.client.Place.Get(ctx, placeID)
//	if err != nil {
//		return nil, err
//	}
//
//	inventory, err := s.client.Inventory.Get(ctx, inventoryId)
//
//	return placeInventory, nil
//}

package inventory

//type InventoryService interface {
//	CreateInventory(ctx context.Context, input CreateInventoryDTO) (*ent.PlaceInventory, error)
//	FetchInventoriesByPlace(ctx context.Context, placeID string) ([]*ent.PlaceInventory, error)
//	EditInventory(ctx context.Context, inventoryID string, input EditInventoryDTO) (*ent.PlaceInventory, error)
//	DeleteInventory(ctx context.Context, inventoryID string) error
//}
//
//type InventoryServiceImpl struct {
//	client       *ent.Client
//	mediaService media.MediaService // Assuming you have a media service for media operations
//}
//
//func NewInventoryService(client *ent.Client, mediaService media.MediaService) *InventoryServiceImpl {
//	return &InventoryServiceImpl{
//		client:       client,
//		mediaService: mediaService,
//	}
//}
//
//// CreateInventory creates an inventory including media, custom fields, and categorizations.
//func (s *InventoryServiceImpl) CreateInventory(ctx context.Context, input CreateInventoryDTO) (*ent.Inventory, error) {
//	// Create media
//	media, err := s.mediaService.CreateMedia(ctx, input.Media)
//	if err != nil {
//		return nil, err
//	}
//
//	// Create inventory with associated media and other data
//	inventory, err := s.client.Inventory.
//		Create().
//		SetMedia(media).
//		// Set other fields...
//		Save(ctx)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return inventory, nil
//}
//
//// FetchInventoriesByPlace fetches all inventory items attached to a place.
//func (s *InventoryServiceImpl) FetchInventoriesByPlace(ctx context.Context, placeID string) ([]*ent.Inventory, error) {
//	inventories, err := s.client.Inventory.
//		Query().
//		Where(inventory.HasPlaceWith(place.ID(placeID))).
//		All(ctx)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return inventories, nil
//}
//
//// EditInventory edits a single inventory.
//func (s *InventoryServiceImpl) EditInventory(ctx context.Context, inventoryID string, input EditInventoryDTO) (*ent.Inventory, error) {
//	// Edit inventory. Handle media, custom fields, and categorizations similar to creation.
//	inventory, err := s.client.Inventory.
//		UpdateOneID(inventoryID).
//		// Set updated fields...
//		Save(ctx)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return inventory, nil
//}
//
//// DeleteInventory deletes a single inventory.
//func (s *InventoryServiceImpl) DeleteInventory(ctx context.Context, inventoryID string) error {
//	err := s.client.Inventory.
//		DeleteOneID(inventoryID).
//		Exec(ctx)
//
//	return err
//}

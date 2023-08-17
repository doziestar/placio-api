package inventory

//type InventoryController struct {
//	inventoryService InventoryService
//}
//
//func NewInventoryController(inventoryService InventoryService) *InventoryController {
//	return &InventoryController{
//		inventoryService: inventoryService,
//	}
//}
//
//func (c *InventoryController) RegisterRoutes(router *gin.RouterGroup) {
//	inventoryRouter := router.Group("/inventories")
//	{
//		inventoryRouter.POST("/", c.createInventory)
//		inventoryRouter.GET("/place/:placeID", c.getInventoriesByPlace)
//		inventoryRouter.PATCH("/:id", c.editInventory)
//		inventoryRouter.DELETE("/:id", c.deleteInventory)
//	}
//}
//
//// createInventory handles the endpoint to create an inventory.
//func (c *InventoryController) createInventory(ctx *gin.Context) {
//	var input CreateInventoryDTO
//	if err := ctx.ShouldBindJSON(&input); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	inventory, err := c.inventoryService.CreateInventory(ctx, input)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, inventory)
//}
//
//// getInventoriesByPlace handles the endpoint to fetch all inventory items attached to a place.
//func (c *InventoryController) getInventoriesByPlace(ctx *gin.Context) {
//	placeID := ctx.Param("placeID")
//
//	inventories, err := c.inventoryService.FetchInventoriesByPlace(ctx, placeID)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, inventories)
//}
//
//// editInventory handles the endpoint to edit a single inventory.
//func (c *InventoryController) editInventory(ctx *gin.Context) {
//	var input EditInventoryDTO
//	if err := ctx.ShouldBindJSON(&input); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	inventoryID := ctx.Param("id")
//	inventory, err := c.inventoryService.EditInventory(ctx, inventoryID, input)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, inventory)
//}
//
//// deleteInventory handles the endpoint to delete a single inventory.
//func (c *InventoryController) deleteInventory(ctx *gin.Context) {
//	inventoryID := ctx.Param("id")
//
//	err := c.inventoryService.DeleteInventory(ctx, inventoryID)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{"message": "Inventory deleted successfully"})
//}

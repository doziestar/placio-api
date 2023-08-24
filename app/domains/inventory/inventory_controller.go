package inventory

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"placio-app/utility"
	"placio-pkg/middleware"
	"strconv"
)

type InventoryController struct {
	inventoryService InventoryService
	cache            utility.RedisClient
}

func NewInventoryController(inventoryService InventoryService, cache utility.RedisClient) *InventoryController {
	return &InventoryController{inventoryService: inventoryService, cache: cache}
}

func (ic *InventoryController) RegisterRoutes(router *gin.RouterGroup) {
	inventoryRouter := router.Group("/inventory")
	{
		inventoryRouter.POST("/types", middleware.ErrorMiddleware(ic.createInventoryType))
		inventoryRouter.PUT("/types/:id", middleware.ErrorMiddleware(ic.updateInventoryType))
		inventoryRouter.DELETE("/types/:id", middleware.ErrorMiddleware(ic.deleteInventoryType))
		inventoryRouter.GET("/types", middleware.ErrorMiddleware(ic.listInventoryTypes))

		inventoryRouter.POST("/attributes", middleware.ErrorMiddleware(ic.createInventoryAttribute))
		inventoryRouter.PUT("/attributes/:id", middleware.ErrorMiddleware(ic.updateInventoryAttribute))
		inventoryRouter.DELETE("/attributes/:id", middleware.ErrorMiddleware(ic.deleteInventoryAttribute))
		inventoryRouter.GET("/attributes", middleware.ErrorMiddleware(ic.listInventoryAttributes))
		inventoryRouter.GET("/attributes/search", middleware.ErrorMiddleware(ic.searchInventoryAttributes))

		// Group for "places"
		placesRouter := inventoryRouter.Group("/places")
		{
			placesRouter.POST("/:placeID", middleware.ErrorMiddleware(ic.createPlaceInventory))
			placesRouter.GET("/:placeID", middleware.ErrorMiddleware(ic.getPlaceInventory))
			placesRouter.PUT("/:placeID", middleware.ErrorMiddleware(ic.updatePlaceInventory))
			placesRouter.DELETE("/:placeID", middleware.ErrorMiddleware(ic.deletePlaceInventory))
			placesRouter.GET("", middleware.ErrorMiddleware(ic.listPlaceInventories))

			// Nested group for "attributes" within "places"
			attributesRouter := placesRouter.Group("/:placeID/attributes")
			{
				attributesRouter.POST("", middleware.ErrorMiddleware(ic.createPlaceInventoryAttribute))
				attributesRouter.PUT("/:id", middleware.ErrorMiddleware(ic.updatePlaceInventoryAttribute))
				attributesRouter.DELETE("/:id", middleware.ErrorMiddleware(ic.deletePlaceInventoryAttribute))
			}
		}
	}
}

func (ic *InventoryController) createInventoryType(c *gin.Context) error {
	var data inventoryTypeData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	result, err := ic.inventoryService.CreateInventoryType(c, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(result, "success", "Inventory Type created successfully", ""))
	return nil
}

func (ic *InventoryController) updateInventoryType(c *gin.Context) error {
	id := c.Param("id")

	var data inventoryTypeData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	result, err := ic.inventoryService.UpdateInventoryType(c, id, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(result, "success", "Inventory Type updated successfully", ""))
	return nil
}

func (ic *InventoryController) deleteInventoryType(c *gin.Context) error {
	id := c.Param("id")

	err := ic.inventoryService.DeleteInventoryType(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
	return nil
}

func (ic *InventoryController) listInventoryTypes(c *gin.Context) error {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	result, err := ic.inventoryService.ListInventoryTypes(c, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(result, "success", "Inventory Type updated successfully", ""))
	return nil
}

func (ic *InventoryController) createInventoryAttribute(c *gin.Context) error {
	var data inventoryAttributeData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	result, err := ic.inventoryService.CreateInventoryAttribute(c, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(result, "success", "Inventory Attribute updated successfully", ""))
	return nil
}

func (ic *InventoryController) updateInventoryAttribute(c *gin.Context) error {
	id := c.Param("id")

	var data inventoryAttributeData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	result, err := ic.inventoryService.UpdateInventoryAttribute(c, id, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(result, "success", "Inventory Type updated successfully", ""))
	return nil
}

func (ic *InventoryController) deleteInventoryAttribute(c *gin.Context) error {
	id := c.Param("id")

	err := ic.inventoryService.DeleteInventoryAttribute(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
	return nil
}

func (ic *InventoryController) listInventoryAttributes(c *gin.Context) error {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	result, err := ic.inventoryService.ListInventoryAttributes(c, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(result, "success", "Inventory retrieved", ""))
	return err
}

func (ic *InventoryController) searchInventoryAttributes(c *gin.Context) error {
	query := c.DefaultQuery("query", "")

	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	result, err := ic.inventoryService.SearchInventoryAttributes(c, query, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(result, "success"))
	return nil
}

func (ic *InventoryController) createPlaceInventory(c *gin.Context) error {
	placeID := c.Param("placeID")

	var data placeInventoryData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	result, err := ic.inventoryService.CreatePlaceInventory(c, placeID, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	c.JSON(http.StatusOK, result)
	return nil
}

func (ic *InventoryController) getPlaceInventory(c *gin.Context) error {
	id := c.Param("id")

	result, err := ic.inventoryService.GetPlaceInventory(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	c.JSON(http.StatusOK, result)
	return nil
}

func (ic *InventoryController) updatePlaceInventory(c *gin.Context) error {
	id := c.Param("id")

	var data placeInventoryData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	result, err := ic.inventoryService.UpdatePlaceInventory(c, id, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	c.JSON(http.StatusOK, result)
	return nil
}

func (ic *InventoryController) deletePlaceInventory(c *gin.Context) error {
	id := c.Param("id")

	err := ic.inventoryService.DeletePlaceInventory(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
	return nil
}

func (ic *InventoryController) listPlaceInventories(c *gin.Context) error {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	id := c.Param("id")

	result, err := ic.inventoryService.ListPlaceInventories(c, id, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	c.JSON(http.StatusOK, result)
	return nil
}

func (ic *InventoryController) createPlaceInventoryAttribute(c *gin.Context) error {
	placeInventoryID := c.Param("placeInventoryID")

	var data placeInventoryAttributeData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	result, err := ic.inventoryService.CreatePlaceInventoryAttribute(c, placeInventoryID, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	c.JSON(http.StatusOK, result)
	return nil
}

func (ic *InventoryController) updatePlaceInventoryAttribute(c *gin.Context) error {
	id := c.Param("id")

	var data placeInventoryAttributeData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	result, err := ic.inventoryService.UpdatePlaceInventoryAttribute(c, id, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	c.JSON(http.StatusOK, result)
	return nil
}

func (ic *InventoryController) deletePlaceInventoryAttribute(c *gin.Context) error {
	id := c.Param("id")

	err := ic.inventoryService.DeletePlaceInventoryAttribute(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
	return nil
}

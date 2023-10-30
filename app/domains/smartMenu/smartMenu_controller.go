package smartMenu

import (
	"fmt"
	"log"
	"net/http"
	"placio-app/ent"
	"placio-app/utility"
	"placio-pkg/errors"
	"placio-pkg/middleware"

	"github.com/gin-gonic/gin"
)

type SmartMenuController struct {
	smartMenuService ISmartMenu
}

func NewSmartMenuController(smartMenuService ISmartMenu) *SmartMenuController {
	return &SmartMenuController{smartMenuService: smartMenuService}
}

func (c *SmartMenuController) RegisterRoutes(router *gin.RouterGroup) {
	const (
		placeIDParam = "placeId"
		menuIDParam  = "menuId"
		tableIDParam = "tableId"
		orderIDParam = "orderId"
	)

	menuRouter := router.Group("/menus")
	{
		menuRouter.POST(fmt.Sprintf("/place/:%s", placeIDParam), middleware.ErrorMiddleware(c.createMenu))
		menuRouter.GET(fmt.Sprintf("/place/:%s", placeIDParam), middleware.ErrorMiddleware(c.getMenus))
		menuRouter.GET(fmt.Sprintf("/:%s", menuIDParam), middleware.ErrorMiddleware(c.getMenuByID))
		menuRouter.PUT(fmt.Sprintf("/:%s", menuIDParam), middleware.ErrorMiddleware(c.updateMenu))
		menuRouter.DELETE(fmt.Sprintf("/:%s", menuIDParam), middleware.ErrorMiddleware(c.deleteMenu))
		menuRouter.PATCH(fmt.Sprintf("/:%s/restore", menuIDParam), middleware.ErrorMiddleware(c.restoreMenu))
	}

	tableRouter := router.Group("/tables")
	{
		tableRouter.POST(fmt.Sprintf("/place/:%s", placeIDParam), middleware.ErrorMiddleware(c.createTable))
		tableRouter.GET(fmt.Sprintf("/place/:%s", placeIDParam), middleware.ErrorMiddleware(c.getTables))
		tableRouter.GET(fmt.Sprintf("/:%s", tableIDParam), middleware.ErrorMiddleware(c.getTableByID))
		tableRouter.PUT(fmt.Sprintf("/:%s", tableIDParam), middleware.ErrorMiddleware(c.updateTable))
		tableRouter.DELETE(fmt.Sprintf("/:%s", tableIDParam), middleware.ErrorMiddleware(c.deleteTable))
		tableRouter.PATCH(fmt.Sprintf("/:%s/restore", tableIDParam), middleware.ErrorMiddleware(c.restoreTable))
		tableRouter.POST(fmt.Sprintf("/:%s/regenerate-qr", tableIDParam), middleware.ErrorMiddleware(c.regenerateQRCode))
	}

	orderRouter := router.Group("/orders")
	{
		orderRouter.POST("/business/:businessId", middleware.ErrorMiddleware(c.createOrder))
		orderRouter.GET("/", middleware.ErrorMiddleware(c.getOrders))
		orderRouter.GET(fmt.Sprintf("/:%s", orderIDParam), middleware.ErrorMiddleware(c.getOrderByID))
		orderRouter.PUT(fmt.Sprintf("/:%s", orderIDParam), middleware.ErrorMiddleware(c.updateOrder))
		orderRouter.DELETE(fmt.Sprintf("/:%s", orderIDParam), middleware.ErrorMiddleware(c.deleteOrder))
		orderRouter.PATCH(fmt.Sprintf("/:%s/restore", orderIDParam), middleware.ErrorMiddleware(c.restoreOrder))
	}
}

// CreateMenu creates a new menu.
// @Summary Create a new menu
// @Description Create a new menu for the authenticated user
// @Tags Menu
// @Accept json
// @Produce json
// @Param placeId path string true "Place ID"
// @Param menu body Menu true "Menu"
// @Success 200 {object} MenuResponseDto "Successfully created a new menu"
// @Failure 400 {object} ErrorDTO "Bad Request"
// @Router /menus/{placeId} [post]
func (c *SmartMenuController) createMenu(ctx *gin.Context) error {
	log.Println("createMenu")
	placeId := ctx.Param("placeId")
	var menu ent.Menu
	
	form, err := ctx.MultipartForm()
	if err != nil {
		log.Println("Error parsing form:", err)
		return nil
	}

	// It's a good practice to check if the form values exist before accessing them
	if name, exists := form.Value["name"]; exists {
		menu.Name = name[0]
	}
	if description, exists := form.Value["description"]; exists {
		menu.Description = description[0]
	}
	if price, exists := form.Value["price"]; exists {
		menu.Price = price[0]
	}
	if preparationTime, exists := form.Value["preparationTime"]; exists {
		menu.PreparationTime = preparationTime[0]
	}
	if isAvailable, exists := form.Value["isAvailable"]; exists {
		menu.IsAvailable = isAvailable[0] == "true"
	}
	if options, exists := form.Value["options"]; exists {
		menu.Options = options[0]
	}

	log.Println("menu", menu)
	if category, exists := form.Value["category"]; exists && len(category) > 0 {
		log.Println("category", category[0])
		medias := form.File["medias"]
		log.Println("menu", menu, "category", category[0], "medias", medias)

		createdMenu, err := c.smartMenuService.CreateMenu(ctx.Request.Context(), placeId, &menu, category[0], medias)
		if err != nil {
			log.Println("Error creating menu:", err)
			return nil
		}

		ctx.JSON(http.StatusOK, utility.ProcessResponse(createdMenu))
		return nil
	}

	return errors.ErrUnprocessable

}


// GetMenus returns a list of menus.
// @Summary Get menus
// @Description Get menus for the authenticated user
// @Tags Menu
// @Accept json
// @Produce json
// @Param placeId path string true "Place ID"
// @Success 200 {object} []MenuResponseDto "Successfully retrieved menus"
// @Failure 400 {object} ErrorDTO "Bad Request"
// @Router /menus/{placeId} [get]
func (c *SmartMenuController) getMenus(ctx *gin.Context) error {
	placeId := ctx.Param("placeId")
	menus, err := c.smartMenuService.GetMenus(ctx, placeId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(menus))
	return nil
}

// GetMenuByID returns a menu by ID.
// @Summary Get menu by ID
// @Description Get menu by ID for the authenticated user
// @Tags Menu
// @Accept json
// @Produce json
// @Param menuId path string true "Menu ID"
// @Success 200 {object} MenuResponseDto "Successfully retrieved menu"
// @Failure 400 {object} ErrorDTO "Bad Request"
// @Router /menus/{menuId} [get]
func (c *SmartMenuController) getMenuByID(ctx *gin.Context) error {
	menuId := ctx.Param("menuId")
	menu, err := c.smartMenuService.GetMenuByID(ctx, menuId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(menu))
	return nil
}

// UpdateMenu updates a menu.
// @Summary Update menu
// @Description Update menu for the authenticated user
// @Tags Menu
// @Accept json
// @Produce json
// @Param menuId path string true "Menu ID"
// @Param menu body Menu true "Menu"
// @Success 200 {object} MenuResponseDto "Successfully updated menu"
// @Failure 400 {object} ErrorDTO "Bad Request"
// @Router /menus/{menuId} [put]
func (c *SmartMenuController) updateMenu(ctx *gin.Context) error {
	menuId := ctx.Param("menuId")
	var menu *ent.Menu
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		return err
	}
	updatedMenu, err := c.smartMenuService.UpdateMenu(ctx, menuId, menu)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(updatedMenu))
	return nil
}

// DeleteMenu deletes a menu.
// @Summary Delete menu
// @Description Delete menu for the authenticated user
// @Tags Menu
// @Accept json
// @Produce json
// @Param menuId path string true "Menu ID"
// @Success 200 {object} MenuResponseDto "Successfully deleted menu"
// @Failure 400 {object} ErrorDTO "Bad Request"
// @Router /menus/{menuId} [delete]
func (c *SmartMenuController) deleteMenu(ctx *gin.Context) error {
	menuId := ctx.Param("menuId")
	err := c.smartMenuService.DeleteMenu(ctx, menuId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(nil))
	return nil
}

// RestoreMenu restores a menu.
// @Summary Restore menu
// @Description Restore menu for the authenticated user
// @Tags Menu
// @Accept json
// @Produce json
// @Param menuId path string true "Menu ID"
// @Success 200 {object} MenuResponseDto "Successfully restored menu"
// @Failure 400 {object} ErrorDTO "Bad Request"
// @Router /menus/{menuId}/restore [patch]
func (c *SmartMenuController) restoreMenu(ctx *gin.Context) error {
	menuId := ctx.Param("menuId")
	restoredMenu, err := c.smartMenuService.RestoreMenu(ctx, menuId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(restoredMenu))
	return nil
}

// CreateTable creates a new table.
// @Summary Create a new table
// @Description Create a new table for the authenticated user
// @Tags Table
// @Accept json
// @Produce json
// @Param placeId path string true "Place ID"
// @Param table body Table true "Table"
// @Success 200 {object} TableResponseDto "Successfully created a new table"
// @Failure 400 {object} ErrorDTO "Bad Request"
// @Router /tables/{placeId} [post]
func (c *SmartMenuController) createTable(ctx *gin.Context) error {
	placeId := ctx.Param("placeId")
	var table *ent.PlaceTable
	if err := ctx.ShouldBindJSON(&table); err != nil {
		return err
	}

	createdTable, err := c.smartMenuService.CreateTable(ctx, placeId, table)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(createdTable))
	return nil
}

// GetTables returns a list of tables.
// @Summary Get tables
// @Description Get tables for the authenticated user
// @Tags Table
// @Accept json
// @Produce json
// @Param placeId path string true "Place ID"
// @Success 200 {object} []TableResponseDto "Successfully retrieved tables"
// @Failure 400 {object} ErrorDTO "Bad Request"
// @Router /tables/{placeId} [get]
func (c *SmartMenuController) getTables(ctx *gin.Context) error {
	placeId := ctx.Param("placeId")
	tables, err := c.smartMenuService.GetTables(ctx, placeId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(tables))
	return nil
}

// GetTableByID returns a table by ID.
// @Summary Get table by ID
// @Description Get table by ID for the authenticated user
// @Tags Table
// @Accept json
// @Produce json
// @Param tableId path string true "Table ID"
// @Success 200 {object} TableResponseDto "Successfully retrieved table"
// @Failure 400 {object} ErrorDTO "Bad Request"
// @Router /tables/{tableId} [get]
func (c *SmartMenuController) getTableByID(ctx *gin.Context) error {
	tableId := ctx.Param("tableId")
	table, err := c.smartMenuService.GetTableByID(ctx, tableId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(table))
	return nil
}

// UpdateTable updates a table.
// @Summary Update table
// @Description Update table for the authenticated user
// @Tags Table
// @Accept json
// @Produce json
// @Param tableId path string true "Table ID"
// @Param table body Table true "Table"
// @Success 200 {object} TableResponseDto "Successfully updated table"
// @Failure 400 {object} ErrorDTO "Bad Request"
// @Router /tables/{tableId} [put]
func (c *SmartMenuController) updateTable(ctx *gin.Context) error {
	table := ctx.Param("tableId")
	var tableBody *ent.PlaceTable
	if err := ctx.ShouldBindJSON(&tableBody); err != nil {
		return err
	}

	updatedTable, err := c.smartMenuService.UpdateTable(ctx, table, tableBody)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(updatedTable))
	return nil
}

// DeleteTable deletes a table.
// @Summary Delete table
// @Description Delete table for the authenticated user
// @Tags Table
// @Accept json
// @Produce json
// @Param tableId path string true "Table ID"
// @Success 200 {object} TableResponseDto "Successfully deleted table"
// @Failure 400 {object} ErrorDTO "Bad Request"
// @Router /tables/{tableId} [delete]
func (c *SmartMenuController) deleteTable(ctx *gin.Context) error {
	tableId := ctx.Param("tableId")
	err := c.smartMenuService.DeleteTable(ctx, tableId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(nil))
	return nil
}

// RestoreTable restores a table.
// @Summary Restore table
// @Description Restore table for the authenticated user
// @Tags Table
// @Accept json
// @Produce json
// @Param tableId path string true "Table ID"
// @Success 200 {object} TableResponseDto "Successfully restored table"
// @Failure 400 {object} ErrorDTO "Bad Request"
// @Router /tables/{tableId}/restore [patch]
func (c *SmartMenuController) restoreTable(ctx *gin.Context) error {
	tableId := ctx.Param("tableId")
	restoredTable, err := c.smartMenuService.RestoreTable(ctx, tableId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(restoredTable))
	return nil
}

// RegenerateQRCode regenerates a QR code.
// @Summary Regenerate QR code
// @Description Regenerate QR code for the authenticated user
// @Tags Table
// @Accept json
// @Produce json
// @Param tableId path string true "Table ID"
// @Success 200 {object} TableResponseDto "Successfully regenerated QR code"
// @Failure 400 {object} ErrorDTO "Bad Request"
// @Router /tables/{tableId}/regenerate-qr [post]
func (c *SmartMenuController) regenerateQRCode(ctx *gin.Context) error {
	tableId := ctx.Param("tableId")
	qrcode, err := c.smartMenuService.RegenerateQRCode(ctx, tableId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(qrcode))
	return nil
}

// CreateOrder creates a new order.
// @Summary Create a new order
// @Description Create a new order for the authenticated user
// @Tags Order
// @Accept json
// @Produce json
// @Param businessId path string true "Business ID"
// @Param order body Order true "Order"
// @Success 200 {object} OrderResponseDto "Successfully created a new order"
// @Failure 400 {object} ErrorDTO "Bad Request"
// @Router /orders/{businessId} [post]
func (c *SmartMenuController) createOrder(ctx *gin.Context) error {
	return nil
}

// GetOrders returns a list of orders.
// @Summary Get orders
// @Description Get orders for the authenticated user
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {object} []OrderResponseDto "Successfully retrieved orders"
// @Failure 400 {object} ErrorDTO "Bad Request"
// @Router /orders [get]
func (c *SmartMenuController) getOrders(ctx *gin.Context) error {
	return nil
}

// GetOrderByID returns an order by ID.
// @Summary Get order by ID
// @Description Get order by ID for the authenticated user
// @Tags Order
// @Accept json
// @Produce json
// @Param orderId path string true "Order ID"
// @Success 200 {object} OrderResponseDto "Successfully retrieved order"
// @Failure 400 {object} ErrorDTO "Bad Request"
// @Router /orders/{orderId} [get]
func (c *SmartMenuController) getOrderByID(ctx *gin.Context) error {
	return nil
}

// UpdateOrder updates an order.
// @Summary Update order
// @Description Update order for the authenticated user
// @Tags Order
// @Accept json
// @Produce json
// @Param orderId path string true "Order ID"
// @Param order body Order true "Order"
// @Success 200 {object} OrderResponseDto "Successfully updated order"
// @Failure 400 {object} ErrorDTO "Bad Request"
// @Router /orders/{orderId} [put]
func (c *SmartMenuController) updateOrder(ctx *gin.Context) error {
	return nil
}

// DeleteOrder deletes an order.
// @Summary Delete order
// @Description Delete order for the authenticated user
// @Tags Order
// @Accept json
// @Produce json
// @Param orderId path string true "Order ID"
// @Success 200 {object} OrderResponseDto "Successfully deleted order"
// @Failure 400 {object} ErrorDTO "Bad Request"
// @Router /orders/{orderId} [delete]
func (c *SmartMenuController) deleteOrder(ctx *gin.Context) error {
	return nil
}

// RestoreOrder restores an order.
// @Summary Restore order
// @Description Restore order for the authenticated user
// @Tags Order
// @Accept json
// @Produce json
// @Param orderId path string true "Order ID"
// @Success 200 {object} OrderResponseDto "Successfully restored order"
// @Failure 400 {object} ErrorDTO "Bad Request"
// @Router /orders/{orderId}/restore [patch]
func (c *SmartMenuController) restoreOrder(ctx *gin.Context) error {
	return nil
}

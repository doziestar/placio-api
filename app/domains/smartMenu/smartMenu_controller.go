package smartMenu

import (
	"github.com/gin-gonic/gin"
	"placio-pkg/middleware"
)

type SmartMenuController struct {
	smartMenuService *SmartMenuService
}

func NewSmartMenuController(smartMenuService *SmartMenuService) *SmartMenuController {
	return &SmartMenuController{smartMenuService: smartMenuService}
}

func (c *SmartMenuController) RegisterRoutes(router *gin.Engine) {
	menuRouter := router.Group("/api/menus")
	tableRouter := router.Group("/api/tables")
	orderRouter := router.Group("/api/orders")

	{
		menuRouter.POST("/:placeId", middleware.ErrorMiddleware(c.createMenu))
		menuRouter.GET("/:placeId", middleware.ErrorMiddleware(c.getMenus))
		menuRouter.GET("/:menuId", middleware.ErrorMiddleware(c.getMenuByID))
		menuRouter.PUT("/:menuId", middleware.ErrorMiddleware(c.updateMenu))
		menuRouter.DELETE("/:menuId", middleware.ErrorMiddleware(c.deleteMenu))
		menuRouter.PATCH("/:menuId/restore", middleware.ErrorMiddleware(c.restoreMenu))
	}

	{
		tableRouter.POST("/:placeId", middleware.ErrorMiddleware(c.createTable))
		tableRouter.GET("/:placeId", middleware.ErrorMiddleware(c.getTables))
		tableRouter.GET("/:tableId", middleware.ErrorMiddleware(c.getTableByID))
		tableRouter.PUT("/:tableId", middleware.ErrorMiddleware(c.updateTable))
		tableRouter.DELETE("/:tableId", middleware.ErrorMiddleware(c.deleteTable))
		tableRouter.PATCH("/:tableId/restore", middleware.ErrorMiddleware(c.restoreTable))
		tableRouter.POST("/:tableId/regenerate-qr", middleware.ErrorMiddleware(c.regenerateQRCode))
	}

	{
		orderRouter.POST("/:businessId", middleware.ErrorMiddleware(c.createOrder))
		orderRouter.GET("/", middleware.ErrorMiddleware(c.getOrders))
		orderRouter.GET("/:orderId", middleware.ErrorMiddleware(c.getOrderByID))
		orderRouter.PUT("/:orderId", middleware.ErrorMiddleware(c.updateOrder))
		orderRouter.DELETE("/:orderId", middleware.ErrorMiddleware(c.deleteOrder))
		orderRouter.PATCH("/:orderId/restore", middleware.ErrorMiddleware(c.restoreOrder))
	}
}

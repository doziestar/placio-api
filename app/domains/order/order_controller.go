package order

import (
	"log"
	"net/http"
	"placio-app/ent"
	"placio-pkg/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService OrderServices
}

func NewOrderController(service OrderServices) OrderController {
	return OrderController{orderService: service}
}

func (oc *OrderController) RegisterRoutes(route *gin.RouterGroup) {
	orderRoute := route.Group("/orders")
	{
		orderRoute.POST("/", middleware.ErrorMiddleware(oc.createOrder))
		orderRoute.PUT("/:orderID", middleware.ErrorMiddleware(oc.updateOrder))
		orderRoute.DELETE("/:orderID", middleware.ErrorMiddleware(oc.deleteOrder))
		orderRoute.GET("/:orderID", middleware.ErrorMiddleware(oc.getOrder))
		orderRoute.GET("/place/:placeID", middleware.ErrorMiddleware(oc.getOrders))
		orderRoute.GET("/user/:userID", middleware.ErrorMiddleware(oc.getOrdersByUserID))
		orderRoute.GET("/table/:tableID", middleware.ErrorMiddleware(oc.getOrdersByTableID))
		orderRoute.GET("/table/:tableID/status/:status", middleware.ErrorMiddleware(oc.getOrdersByTableIDAndStatus))
		orderRoute.GET("/user/:userID/status/:status", middleware.ErrorMiddleware(oc.getOrdersByUserIDAndStatus))
	}
}

func (oc *OrderController) createOrder(ctx *gin.Context) error {
	tableID := ctx.DefaultQuery("tableID", "")
	userID := ctx.MustGet("user").(string)

	var orderItems OrderWithItemsDTO
	if err := ctx.ShouldBindJSON(&orderItems); err != nil {
		log.Println(err)
		return err
	}

	newOrder, err := oc.orderService.CreateOrder(ctx, tableID, userID, orderItems)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, newOrder)
	return nil
}

func (oc *OrderController) updateOrder(ctx *gin.Context) error {
	orderID := ctx.Param("orderID")
	var updateDto ent.Order
	if err := ctx.ShouldBindJSON(&updateDto); err != nil {
		return err
	}

	// Similar to createOrder, adjust based on your request structure
	var newMenuItems OrderWithItemsDTO
	if err := ctx.ShouldBindJSON(&newMenuItems); err != nil {
		return err
	}

	updatedOrder, err := oc.orderService.UpdateOrder(ctx, orderID, newMenuItems)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	ctx.JSON(http.StatusOK, updatedOrder)
	return nil
}

func (oc *OrderController) deleteOrder(ctx *gin.Context) error {
	orderID := ctx.Param("orderID")
	err := oc.orderService.DeleteOrder(ctx, orderID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
	return nil
}

func (oc *OrderController) getOrder(ctx *gin.Context) error {
	orderID := ctx.Param("orderID")
	order, err := oc.orderService.GetOrder(ctx, orderID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	ctx.JSON(http.StatusOK, order)
	return nil
}

func (oc *OrderController) getOrders(ctx *gin.Context) error {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	placeId := ctx.Param("placeID")

	orders, err := oc.orderService.GetOrders(ctx, placeId, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	ctx.JSON(http.StatusOK, orders)
	return nil
}

func (oc *OrderController) getOrdersByUserID(ctx *gin.Context) error {
	userID := ctx.Param("userID")
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	orders, err := oc.orderService.GetOrdersByUserID(ctx, userID, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	ctx.JSON(http.StatusOK, orders)
	return nil
}

func (oc *OrderController) getOrdersByTableID(ctx *gin.Context) error {
	tableID := ctx.Param("tableID")
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	orders, err := oc.orderService.GetOrdersByTableID(ctx, tableID, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	ctx.JSON(http.StatusOK, orders)
	return nil
}

func (oc *OrderController) getOrdersByTableIDAndStatus(ctx *gin.Context) error {
	tableID := ctx.Param("tableID")
	status := ctx.Param("status")
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	orders, err := oc.orderService.GetOrdersByTableIDAndStatus(ctx, tableID, status, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	ctx.JSON(http.StatusOK, orders)
	return nil
}

func (oc *OrderController) getOrdersByUserIDAndStatus(ctx *gin.Context) error {
	userID := ctx.Param("userID")
	status := ctx.Param("status")
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	orders, err := oc.orderService.GetOrdersByUserIDAndStatus(ctx, userID, status, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	ctx.JSON(http.StatusOK, orders)
	return nil
}

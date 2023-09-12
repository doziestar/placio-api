package notifications

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"placio-app/utility"
	"placio-pkg/middleware"
	"strconv"
)

type NotificationController struct {
	notificationService INotification
}

func NewNotificationController(notificationService INotification) *NotificationController {
	return &NotificationController{
		notificationService: notificationService,
	}
}

func (n *NotificationController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/notification/:notificationID", middleware.ErrorMiddleware(n.GetNotification))
	router.GET("/notifications", middleware.ErrorMiddleware(n.GetNotifications))
	router.GET("/notifications/unread", middleware.ErrorMiddleware(n.GetUnreadNotifications))
	router.POST("/notification/", middleware.ErrorMiddleware(n.CreateNotification))
	router.DELETE("/notification/:notificationID", middleware.ErrorMiddleware(n.DeleteNotification))
}

func (n *NotificationController) GetNotification(c *gin.Context) error {
	// get the notification id
	notificationID := c.Param("notificationID")
	// get the notification
	notification, err := n.notificationService.GetNotification(c, notificationID)
	if err != nil {
		return err
	}
	// return the notification
	c.JSON(http.StatusOK, utility.ProcessResponse(notification))
	return nil
}

func (n *NotificationController) GetNotifications(c *gin.Context) error {
	// get the user id
	userID := c.Query("userID")
	// get the limit
	limit := c.Query("limit")
	// get the offset
	offset := c.Query("offset")

	limits, err := strconv.Atoi(limit)
	if err != nil {
		return err
	}

	offsets, err := strconv.Atoi(offset)
	if err != nil {
		return err
	}

	// get the notifications
	notifications, err := n.notificationService.GetNotifications(c, userID, limits, offsets)
	if err != nil {
		return err
	}
	// return the notifications
	c.JSON(http.StatusOK, utility.ProcessResponse(notifications))
	return nil
}

func (n *NotificationController) GetUnreadNotifications(c *gin.Context) error {
	// get the user id
	userID := c.Query("userID")
	// get the limit
	limit := c.Query("limit")
	// get the offset
	offset := c.Query("offset")

	limits, err := strconv.Atoi(limit)
	if err != nil {
		return err
	}

	offsets, err := strconv.Atoi(offset)
	if err != nil {
		return err
	}

	// get the notifications
	notifications, err := n.notificationService.GetUnreadNotifications(c, userID, limits, offsets)
	if err != nil {
		return err
	}
	// return the notifications
	c.JSON(http.StatusOK, utility.ProcessResponse(notifications))
	return nil
}

func (n *NotificationController) CreateNotification(c *gin.Context) error {
	var notification *Notification
	// bind the notification
	err := c.BindJSON(&notification)
	// get the notification
	notificationData, err := n.notificationService.CreateNotification(c, notification)
	if err != nil {
		return err
	}
	// return the notification
	c.JSON(http.StatusOK, utility.ProcessResponse(notificationData))
	return nil
}

func (n *NotificationController) DeleteNotification(c *gin.Context) error {
	// get the notification id
	notificationID := c.Param("notificationID")
	// delete the notification
	err := n.notificationService.DeleteNotification(c, notificationID)
	if err != nil {
		return err
	}
	// return the notification
	c.JSON(http.StatusOK, utility.ProcessResponse("notification deleted successfully"))
	return nil
}

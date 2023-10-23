package notifications

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"placio-app/utility"
	"placio-pkg/errors"
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
	router.DELETE("/notification/:notificationID", middleware.ErrorMiddleware(n.DeleteNotification))
	router.GET("/notifications/unread/count", middleware.ErrorMiddleware(n.GetUnreadNotificationsCount))
}

func (n *NotificationController) GetNotification(c *gin.Context) error {
	notificationID := c.Param("notificationID")
	notification, err := n.notificationService.GetNotification(c.Request.Context(), notificationID)
	if err != nil {
		return errors.ErrNotFound
	}
	c.JSON(http.StatusOK, utility.ProcessResponse(notification))
	return nil
}

func (n *NotificationController) GetUnreadNotificationsCount(c *gin.Context) error {
	userID := c.GetString("user")
	count, err := n.notificationService.GetUnreadNotificationsCount(c.Request.Context(), userID)
	if err != nil {
		return errors.New("Failed to retrieve notifications count")
	}
	c.JSON(http.StatusOK, utility.ProcessResponse(count))
	return nil
}

func (n *NotificationController) GetNotifications(c *gin.Context) error {
	userID := c.GetString("user")
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		return errors.ErrInvalidLimitValue
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		return errors.ErrInvalidOffset
	}

	notifications, err := n.notificationService.GetNotifications(c.Request.Context(), userID, limit, offset)
	if err != nil {
		return errors.New("Failed to retrieve notifications")
	}
	c.JSON(http.StatusOK, utility.ProcessResponse(notifications))
	return nil
}

func (n *NotificationController) GetUnreadNotifications(c *gin.Context) error {
	userID := c.GetString("user")
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		return errors.ErrInvalidLimitValue
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		return errors.ErrInvalidOffset
	}

	notifications, err := n.notificationService.GetUnreadNotifications(c.Request.Context(), userID, limit, offset)
	if err != nil {
		return errors.New("Failed to retrieve notifications")
	}
	c.JSON(http.StatusOK, utility.ProcessResponse(notifications))
	return nil
}

func (n *NotificationController) DeleteNotification(c *gin.Context) error {
	notificationID := c.Param("notificationID")
	err := n.notificationService.DeleteNotification(c.Request.Context(), notificationID)
	if err != nil {
		return errors.New("Failed to delete notification")
	}
	c.JSON(http.StatusOK, utility.ProcessResponse("Notification deleted successfully"))
	return nil
}

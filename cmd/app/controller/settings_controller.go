package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2/middleware/session"
	"net/http"
	"placio-app/errors"
	"placio-app/models"
	"placio-app/service"
	"placio-app/utility"
)

type SettingsController struct {
	store service.ISettingsService
}

func NewSettingsController(store service.ISettingsService) *SettingsController {
	return &SettingsController{store: store}
}

func (c *SettingsController) RegisterRoutes(app *gin.RouterGroup, session *session.Store) {
	settingsGroup := app.Group("/settings")

	settingsGroup.GET("/general", utility.Use(c.getGeneralSettings))
	settingsGroup.PUT("/general", utility.Use(c.updateUserSettings))

	settingsGroup.GET("/notifications", utility.Use(c.getNotificationsSettings))
	settingsGroup.PUT("/notifications", utility.Use(c.updateNotificationsSettings))

	//settingsGroup.GET("/account", c.getAccountSettings)
	settingsGroup.PUT("/account", utility.Use(c.updateAccountSettings))

	settingsGroup.GET("/content", utility.Use(c.getContentSettings))
	settingsGroup.PUT("/content", utility.Use(c.updateContentSettings))
}

// GetGeneralSettings godoc
// @Summary GET general settings
// @Description GET general settings
// @Tags Settings
// @Accept  json
// @Produce  json
// @Success 200 {object} models.GeneralSettings
// @Failure 500 {object} models.ErrorResponse
// @Router /settings/general [get]
func (c *SettingsController) getGeneralSettings(ctx *gin.Context) error {
	userID, ok := ctx.Get("UserID")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch general settings",
		})
		return errors.ErrInvalid
	}

	settings, err := c.store.GetGeneralSettings((userID).(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch general settings",
		})
		return errors.ErrInvalid
	}

	ctx.JSON(http.StatusOK, settings)
	return nil
}

// UpdateUserSettings godoc
// @Summary Update general settings
// @Description Update general settings
// @Tags Settings
// @Accept  json
// @Produce  json
// @Param body body models.GeneralSettings true "General settings"
// @Success 200 {object} models.GeneralSettings
// @Failure 500 {object} models.ErrorResponse
// @Router /settings/general [put]
func (c *SettingsController) updateUserSettings(ctx *gin.Context) error {
	userID, ok := ctx.Get("UserID")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update general settings",
		})
		return errors.ErrInvalid
	}

	var settings models.GeneralSettings
	if err := ctx.BindJSON(&settings); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse request body",
		})
		return err
	}

	settings.UserID = (userID).(string)
	if err := c.store.UpdateUserSettings((userID).(string), &settings); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update general settings",
		})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "General settings updated successfully",
	})
	return nil
}

// GetNotificationsSettings godoc
// @Summary GET notifications settings
// @Description GET notifications settings
// @Tags Settings
// @Accept  json
// @Produce  json
// @Success 200 {object} models.NotificationsSettings
// @Failure 500 {object} models.ErrorResponse
// @Router /settings/notifications [get]
func (c *SettingsController) getNotificationsSettings(ctx *gin.Context) error {
	userID, ok := ctx.Get("UserID")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch notifications settings",
		})
		return errors.ErrInvalid
	}

	settings, err := c.store.GetNotificationsSettings((userID).(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch notifications settings",
		})
		return err
	}

	ctx.JSON(http.StatusOK, settings)
	return nil
}

// UpdateNotificationsSettings godoc
// @Summary Update notifications settings
// @Description Update notifications settings
// @Tags Settings
// @Accept  json
// @Produce  json
// @Param body body models.NotificationsSettings true "Notifications settings"
// @Success 200 {object} models.NotificationsSettings
// @Failure 500 {object} models.ErrorResponse
// @Router /settings/notifications [put]
func (c *SettingsController) updateNotificationsSettings(ctx *gin.Context) error {
	userID, ok := ctx.Get("UserID")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update notifications settings",
		})
		return errors.ErrInvalid
	}

	var settings models.NotificationsSettings
	if err := ctx.BindJSON(&settings); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse request body",
		})
		return err
	}

	if err := c.store.UpdateNotificationsSettings((userID).(string), &settings); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update notifications settings",
		})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Notifications settings updated successfully",
	})
	return nil
}

// GetAccountSettings godoc
// @Summary GET account settings
// @Description GET account settings
// @Tags Settings
// @Accept  json
// @Produce  json
// @Success 200 {object} models.AccountSettings
// @Failure 500 {object} models.ErrorResponse
// @Router /settings/account [get]
//func (c *SettingsController) getAccountSettings(ctx *fiber.Ctx) error {
//	userID := ctx.Locals("UserID").(string)
//
//	settings, err := c.store.GetAccountSettings(userID)
//	if err != nil {
//		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"error": "Failed to fetch account settings",
//		})
//	}
//
//	return ctx.JSON(settings)
//}

// UpdateAccountSettings godoc
// @Summary Update account settings
// @Description Update account settings
// @Tags Settings
// @Accept  json
// @Produce  json
// @Param body body models.AccountSettings true "Account settings"
// @Success 200 {object} models.AccountSettings
// @Failure 500 {object} models.ErrorResponse
// @Router /settings/account [put]
func (c *SettingsController) updateAccountSettings(ctx *gin.Context) error {
	userID, ok := ctx.Get("UserID")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update account settings",
		})
		return errors.ErrInvalid
	}

	var settings models.AccountSettings
	if err := ctx.BindJSON(&settings); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse request body",
		})
		return err
	}

	if err := c.store.UpdateAccountSettings((userID).(string), &settings); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update account settings",
		})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Account settings updated successfully",
	})
	return nil
}

// GetContentSettings godoc
// @Summary GET content settings
// @Description GET content settings
// @Tags Settings
// @Accept  json
// @Security BearerAuth
// @Produce  json
// @Success 200 {object} models.ContentSettings
// @Failure 500 {object} models.ErrorResponse
// @Router /settings/content [get]
func (c *SettingsController) getContentSettings(ctx *gin.Context) error {
	userID, ok := ctx.Get("UserID")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch content settings",
		})
		return errors.ErrInvalid
	}

	settings, err := c.store.GetContentSettings((userID).(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch content settings",
		})
	}

	ctx.JSON(http.StatusOK, settings)
	return nil
}

// UpdateContentSettings godoc
// @Summary Update content settings
// @Description Update content settings
// @Tags Settings
// @Accept  json
// @Produce  json
// @Param body body models.ContentSettings true "Content settings"
// @Success 200 {object} models.ContentSettings
// @Failure 500 {object} models.ErrorResponse
// @Router /settings/content [put]
func (c *SettingsController) updateContentSettings(ctx *gin.Context) error {
	userID, ok := ctx.Get("UserID")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update content settings",
		})
		return errors.ErrInvalid
	}
	var settings models.ContentSettings
	if err := ctx.BindJSON(&settings); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse request body",
		})
		return err
	}

	if err := c.store.UpdateContentSettings((userID).(string), &settings); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update content settings",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Content settings updated successfully",
	})
	return nil
}

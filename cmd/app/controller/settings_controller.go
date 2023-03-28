package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"placio-app/models"
)

type SettingsController struct {
	store models.ISettingsService
}

func NewSettingsController(store models.ISettingsService) *SettingsController {
	return &SettingsController{store: store}
}

func (c *SettingsController) RegisterRoutes(app fiber.Router, session *session.Store) {
	settingsGroup := app.Group("/settings")

	settingsGroup.Get("/general", c.getGeneralSettings)
	settingsGroup.Put("/general", c.updateUserSettings)

	settingsGroup.Get("/notifications", c.getNotificationsSettings)
	settingsGroup.Put("/notifications", c.updateNotificationsSettings)

	settingsGroup.Get("/account", c.getAccountSettings)
	settingsGroup.Put("/account", c.updateAccountSettings)

	settingsGroup.Get("/content", c.getContentSettings)
	settingsGroup.Put("/content", c.updateContentSettings)
}

// GetGeneralSettings godoc
// @Summary Get general settings
// @Description Get general settings
// @Tags Settings
// @Accept  json
// @Produce  json
// @Success 200 {object} models.GeneralSettings
// @Failure 500 {object} models.ErrorResponse
// @Router /settings/general [get]
func (c *SettingsController) getGeneralSettings(ctx *fiber.Ctx) error {
	userID := ctx.Locals("UserID").(string)

	settings, err := c.store.GetGeneralSettings(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch general settings",
		})
	}

	return ctx.JSON(settings)
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
func (c *SettingsController) updateUserSettings(ctx *fiber.Ctx) error {
	userID := ctx.Locals("UserID").(string)

	var settings models.GeneralSettings
	if err := ctx.BodyParser(&settings); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	settings.UserID = userID
	if err := c.store.UpdateUserSettings(userID, &settings); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update general settings",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "General settings updated successfully",
	})
}

// GetNotificationsSettings godoc
// @Summary Get notifications settings
// @Description Get notifications settings
// @Tags Settings
// @Accept  json
// @Produce  json
// @Success 200 {object} models.NotificationsSettings
// @Failure 500 {object} models.ErrorResponse
// @Router /settings/notifications [get]
func (c *SettingsController) getNotificationsSettings(ctx *fiber.Ctx) error {
	userID := ctx.Locals("UserID").(string)

	settings, err := c.store.GetNotificationsSettings(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch notifications settings",
		})
	}

	return ctx.JSON(settings)
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
func (c *SettingsController) updateNotificationsSettings(ctx *fiber.Ctx) error {
	userID := ctx.Locals("UserID").(string)

	var settings models.NotificationsSettings
	if err := ctx.BodyParser(&settings); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	if err := c.store.UpdateNotificationsSettings(userID, &settings); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update notifications settings",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Notifications settings updated successfully",
	})
}

// GetAccountSettings godoc
// @Summary Get account settings
// @Description Get account settings
// @Tags Settings
// @Accept  json
// @Produce  json
// @Success 200 {object} models.AccountSettings
// @Failure 500 {object} models.ErrorResponse
// @Router /settings/account [get]
func (c *SettingsController) getAccountSettings(ctx *fiber.Ctx) error {
	userID := ctx.Locals("UserID").(string)

	settings, err := c.store.GetAccountSettings(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch account settings",
		})
	}

	return ctx.JSON(settings)
}

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
func (c *SettingsController) updateAccountSettings(ctx *fiber.Ctx) error {
	userID := ctx.Locals("UserID").(string)

	var settings models.AccountSettings
	if err := ctx.BodyParser(&settings); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	if err := c.store.UpdateAccountSettings(userID, &settings); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update account settings",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Account settings updated successfully",
	})
}

// GetContentSettings godoc
// @Summary Get content settings
// @Description Get content settings
// @Tags Settings
// @Accept  json
// @Security BearerAuth
// @Produce  json
// @Success 200 {object} models.ContentSettings
// @Failure 500 {object} models.ErrorResponse
// @Router /settings/content [get]
func (c *SettingsController) getContentSettings(ctx *fiber.Ctx) error {
	userID := ctx.Locals("UserID").(string)

	settings, err := c.store.GetContentSettings(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch content settings",
		})
	}

	return ctx.JSON(settings)
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
func (c *SettingsController) updateContentSettings(ctx *fiber.Ctx) error {
	userID := ctx.Locals("UserID").(string)
	var settings models.ContentSettings
	if err := ctx.BodyParser(&settings); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	if err := c.store.UpdateContentSettings(userID, &settings); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update content settings",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Content settings updated successfully",
	})
}

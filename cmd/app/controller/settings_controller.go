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

func (c *SettingsController) RegisterRoutes(app *fiber.App, session *session.Store) {
	settingsGroup := app.Group("/settings")

	settingsGroup.Get("/general", c.GetGeneralSettings)
	settingsGroup.Put("/general", c.UpdateUserSettings)

	settingsGroup.Get("/notifications", c.GetNotificationsSettings)
	settingsGroup.Put("/notifications", c.UpdateNotificationsSettings)

	settingsGroup.Get("/account", c.GetAccountSettings)
	settingsGroup.Put("/account", c.UpdateAccountSettings)

	settingsGroup.Get("/content", c.GetContentSettings)
	settingsGroup.Put("/content", c.UpdateContentSettings)
}

func (c *SettingsController) GetGeneralSettings(ctx *fiber.Ctx) error {
	userID := ctx.Locals("UserID").(string)

	settings, err := c.store.GetGeneralSettings(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch general settings",
		})
	}

	return ctx.JSON(settings)
}

func (c *SettingsController) UpdateUserSettings(ctx *fiber.Ctx) error {
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

func (c *SettingsController) GetNotificationsSettings(ctx *fiber.Ctx) error {
	userID := ctx.Locals("UserID").(string)

	settings, err := c.store.GetNotificationsSettings(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch notifications settings",
		})
	}

	return ctx.JSON(settings)
}

func (c *SettingsController) UpdateNotificationsSettings(ctx *fiber.Ctx) error {
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

func (c *SettingsController) GetAccountSettings(ctx *fiber.Ctx) error {
	userID := ctx.Locals("UserID").(string)

	settings, err := c.store.GetAccountSettings(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch account settings",
		})
	}

	return ctx.JSON(settings)
}

func (c *SettingsController) UpdateAccountSettings(ctx *fiber.Ctx) error {
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

func (c *SettingsController) GetContentSettings(ctx *fiber.Ctx) error {
	userID := ctx.Locals("UserID").(string)

	settings, err := c.store.GetContentSettings(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch content settings",
		})
	}

	return ctx.JSON(settings)
}

func (c *SettingsController) UpdateContentSettings(ctx *fiber.Ctx) error {
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

package api

import "github.com/gofiber/fiber/v2"

func InitializeRoutes(app *fiber.App) {
	HealthCheckRoutes(app)
}

package api

import (
	"placio-app/controller"

	"github.com/gofiber/fiber/v2"
)

func HealthCheckRoutes(app *fiber.App) {
	apiRouter := app.Group("/api")
	v1 := apiRouter.Group("/v1")
	{
		v1.Get("/health", controller.HealthCheck)

	}
}

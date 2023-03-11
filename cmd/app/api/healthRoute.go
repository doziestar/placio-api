package api

import (
	"github.com/gofiber/fiber/v2"
	"placio-app/controller"
)

func HealthCheckRoutes(app *fiber.App) {
	apiRouter := app.Group("/api")
	v1 := apiRouter.Group("/v1")
	{
		v1.Get("/health", controller.HealthCheck)
	}
}

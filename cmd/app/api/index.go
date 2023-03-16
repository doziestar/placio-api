package api

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "placio-api/docs/app"
)

func InitializeRoutes(app *fiber.App) {
	app.Get("/docs/*", swagger.HandlerDefault)
	HealthCheckRoutes(app)

}

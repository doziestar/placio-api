package api

import (
	"github.com/gofiber/fiber/v2"
	"placio-app/controller"
)

func AuthRoutes(app *fiber.App) {
	// Auth Routes
	routerGroup := app.Group("/api")
	v1 := routerGroup.Group("/v1/auth")
	{
		v1.Post("/login", controller.Login)
		v1.Post("/signup", controller.SignUp)
	}
}

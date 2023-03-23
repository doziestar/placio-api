package api

import (
	"placio-app/controller"
	"placio-app/utility"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	// Auth Routes
	routerGroup := app.Group("/api")
	v1 := routerGroup.Group("/v1/auth")
	{
		v1.Post("/auth/signin", utility.Use(controller.Signin))
		v1.Post("/signup", utility.Use(controller.SignUp))
		v1.Post("/logout", utility.Use(controller.LogOut))
		v1.Get("/refresh", utility.Use(controller.RefreshToken))
		v1.Post("/verify", utility.Use(controller.ChangePassword))
		v1.Post("/verify", utility.Use(controller.VerifyEmail))
		v1.Post("/reset", utility.Use(controller.ResetPassword))
		v1.Post("/verify", utility.Use(controller.VerifyPhone))
	}
}

package api

import (
	"placio-app/controller"
	"placio-app/middleware"
	"placio-app/utility"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	// Auth Routes
	routerGroup := app.Group("/api/v1")
	api := routerGroup.Group("/auth")
	{
		api.Post("/", utility.Use(controller.Signin))
		api.Post("/logout", middleware.Verify("user"), utility.Use(controller.LogOut))
		api.Get("/refresh", utility.Use(controller.RefreshToken))
		api.Post("/verify/password", middleware.Verify("user"), utility.Use(controller.ChangePassword))
		api.Post("/verify/email", middleware.Verify("user"), utility.Use(controller.VerifyEmail))
		api.Post("/reset", utility.Use(controller.ResetPassword))
		api.Post("/verify/phone", middleware.Verify("user"), utility.Use(controller.VerifyPhone))
		// api.post(
		// 	"/api/auth/otp",
		// 	use(authController.signin.otp))

		// api.post(
		// 	"/api/auth/magic",
		// 	use(authController.magic))

		// api.post(
		// 	"/api/auth/magic/verify",
		// 	use(authController.magic.verify))

		// api.post(
		// 	"/api/auth/password/reset/request",
		// 	limiter(throttle.password_reset),
		// 	use(userController.password.reset.request))

		// api.post(
		// 	"/api/auth/password/reset",
		// 	limiter(throttle.password_reset),
		// 	use(userController.password.reset))

		// api.post("/api/auth/switch", utility.Verify("user", ""), use(authController.switch))

		// api.post("/api/auth/impersonate", use(authController.impersonate))

		// api.delete("/api/auth", auth.verify("user"), use(authController.signout))
	}
}

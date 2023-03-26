package api

import (
	"placio-app/controller"
	"placio-app/middleware"
	"placio-app/utility"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(api fiber.Router) {
	{
		api.Post("/", utility.Use(controller.Signin))
		api.Post("/sign-out", middleware.Verify("user"), utility.Use(controller.SignOut))
		api.Get("/refresh", utility.Use(controller.RefreshToken))
		api.Post("/verify/password", middleware.Verify("user"), utility.Use(controller.ChangePassword))
		api.Post("/verify/email", middleware.Verify("user"), utility.Use(controller.VerifyEmail))
		api.Post("/reset", utility.Use(controller.ResetPassword))
		api.Post("/verify/phone", middleware.Verify("user"), utility.Use(controller.VerifyPhone))
		api.Get("status", middleware.Verify("user"), utility.Use(controller.GetAuthStatus))
		api.Post(
			"/otp",
			utility.Use(controller.GetOTP))

		api.Post(
			"/magic",
			utility.Use(controller.GetMagicLink))

		api.Post(
			"/magic/verify",
			utility.Use(controller.VerifyMagicLink))

		api.Post(
			"/password/reset/request",
			//limiter(throttle.password_reset),
			utility.Use(controller.RequestPasswordReset))

		api.Post("/switch-account", middleware.Verify("user"), utility.Use(controller.SwitchAccount))

		api.Post("/impersonate", utility.Use(controller.ImpersonateUser))
	}
}

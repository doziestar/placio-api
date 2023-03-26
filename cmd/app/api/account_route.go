package api

import (
	"placio-app/middleware"
	"placio-app/utility"

	"placio-app/controller"

	"github.com/gofiber/fiber/v2"
)

func AccountRoutes(api fiber.Router) {
	api.Post("/", utility.Use(controller.CreateAccount))

	api.Post("/plan", middleware.Verify("owner"), utility.Use(controller.Plan))

	api.Patch("/plan", middleware.Verify("owner"), utility.Use(controller.UpdatePlan))

	api.Get(
		"/",
		middleware.Verify("owner"),
		utility.Use(controller.GetAccounts))

	api.Get(
		"/card",
		middleware.Verify("owner"),
		utility.Use(controller.GetAccount))

	api.Patch(
		"/card",
		middleware.Verify("owner"),
		utility.Use(controller.UpdateInvoice))

	api.Get(
		"/invoice",
		middleware.Verify("owner"),
		utility.Use(controller.GetInvoice))

	api.Get(
		"/plans",
		middleware.Verify("public"),
		utility.Use(controller.GetPlans))

	// api.Get(
	// 	"/api/account/utility.Users",
	// 	middleware.Verify("admin", "account.read"),
	// 	utility.Use(controller.utility.Users))

	api.Get(
		"/subscription",
		middleware.Verify("utility.User"),
		utility.Use(controller.GetSubscription))

	api.Post(
		"/upgrade",
		middleware.Verify("owner"),
		utility.Use(controller.UpgradePlan))

	api.Delete(
		"/",
		middleware.Verify("owner"),
		utility.Use(controller.CancelSubscription))

	api.Delete(
		"/:id",
		middleware.Verify("owner"),
		utility.Use(controller.DeleteAccount))
}

package api

import (
	"placio-app/utility"

	"placio-app/controller"

	"github.com/gofiber/fiber/v2"
)

func AccountRoutes(app *fiber.App) {
	api := app.Group("/api/v1/accounts")
	{
		api.Post("/", utility.Use(controller.CreateAccount))

		// api.Post("/api/account/plan", middleware.Verify("owner", ""), utility.Use(controller.plan))

		// api.Patch("/api/account/plan", middleware.Verify("owner", "billing.update"), utility.Use(controller.plan.update))

		// api.Get(
		// 	"/api/account",
		// 	middleware.Verify("owner", "account.read"),
		// 	utility.Use(controller.get))

		// api.Get(
		// 	"/api/account/card",
		// 	middleware.Verify("owner", "billing.read"),
		// 	utility.Use(controller.card))

		// api.Patch(
		// 	"/api/account/card",
		// 	middleware.Verify("owner", "billing.update"),
		// 	utility.Use(controller.card.update))

		// api.Get(
		// 	"/api/account/invoice",
		// 	middleware.Verify("owner", "billing.read"),
		// 	utility.Use(controller.invoice))

		// api.Get(
		// 	"/api/account/plans",
		// 	middleware.Verify("public"),
		// 	utility.Use(controller.plans))

		// api.Get(
		// 	"/api/account/utility.Users",
		// 	middleware.Verify("admin", "account.read"),
		// 	utility.Use(controller.utility.Users))

		// api.Get(
		// 	"/api/account/subscription",
		// 	middleware.Verify("utility.User", "billing.read"),
		// 	utility.Use(controller.subscription))

		// api.Post(
		// 	"/api/account/upgrade",
		// 	middleware.Verify("owner", "billing.update"),
		// 	utility.Use(controller.upgrade))

		// api.Delete(
		// 	"/api/account",
		// 	middleware.Verify("owner", "account.delete"),
		// 	utility.Use(controller.close))

		// api.Delete(
		// 	"/api/account/:id",
		// 	middleware.Verify("owner", "account.delete"),
		// 	utility.Use(controller.close))
	}
}

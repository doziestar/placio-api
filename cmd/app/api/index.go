package api

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "placio-api/docs/app"
	"placio-app/controller"
	"placio-app/database"
	"placio-app/models"
	"placio-app/service"
)

func InitializeRoutes(app *fiber.App) {
	app.Get("/docs/*", swagger.HandlerDefault)
	routerGroupV1 := app.Group("/api/v1")

	healthApi := routerGroupV1.Group("/health")
	{
		HealthCheckRoutes(healthApi)
	}
	authApi := routerGroupV1.Group("/auth")
	{
		AuthRoutes(authApi)
	}

	// user
	userService := service.NewUserService(database.DB, &models.User{})
	userController := controller.NewUserController(userService)
	userController.RegisterRoutes(routerGroupV1)

	// settings
	store := service.NewSettingsService(database.DB)
	settingsController := controller.NewSettingsController(store)
	settingsController.RegisterRoutes(routerGroupV1, nil)

	// account
	accountService := service.NewAccountService(database.DB)
	accountController := controller.NewAccountController(accountService)
	accountController.RegisterRoutes(app)

}

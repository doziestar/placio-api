package api

import (
	"fmt"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "placio-api/docs/app"
	"placio-app/controller"
	"placio-app/database"
	"placio-app/models"
	"placio-app/service"
	"placio-app/utility"
)

func InitializeRoutes(app *fiber.App) {
	fmt.Println("Initializing routes...")
	app.Get("/docs/*", swagger.HandlerDefault)
	routerGroupV1 := app.Group("/api/v1")

	healthApi := routerGroupV1.Group("/health")
	{
		HealthCheckRoutes(healthApi)
	}

	// instances
	var user models.User
	var account models.Account

	// utility
	newUtils := utility.NewUtility()

	// auth
	authService := service.NewAuthService(database.DB, &models.User{})
	authController := controller.NewAuthController(authService, newUtils)
	authController.RegisterRoutes(routerGroupV1)

	// user
	userService := service.NewUserService(database.DB, &user, &account)
	userController := controller.NewUserController(userService)
	userController.RegisterRoutes(routerGroupV1)

	// settings
	store := service.NewSettingsService(database.DB)
	settingsController := controller.NewSettingsController(store)
	settingsController.RegisterRoutes(routerGroupV1, nil)

	// account
	accountService := service.NewAccountService(database.DB, account, user)
	accountController := controller.NewAccountController(accountService, newUtils)
	accountController.RegisterRoutes(routerGroupV1)

}

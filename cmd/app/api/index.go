package api

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "placio-api/docs/app"
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
	userApi := routerGroupV1.Group("/users")
	{
		UserRoutes(userApi)
	}
	accountApi := routerGroupV1.Group("/accounts")
	{
		AccountRoutes(accountApi)
	}
	//utilityApi := routerGroupV1.Group("/utility")
	//{
	//	UtilityRoutes(utilityApi)
	//}

}

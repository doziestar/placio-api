package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"placio-app/api"
	"placio-app/database"
	_ "placio-app/ent/runtime"
	"placio-app/start"
)

// @title Placio Application Api
// @version 0.01
// @description This is the documentation for the Placio Application Api
// @termsOfService https://placio.io/terms
// @privacyPolicy https://placio.io/privacy-policy
// @contact.name Placio Ltd
// @contact.url https://placio.io
// @contact.email support@placio.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host https://api.palnight.com
// @Authorization Bearer <token>
// @BasePath /api/v1
// @schemes http https
func main() {
	// initialize gin app
	app := gin.Default()

	// initialize middleware
	start.Middleware(app)

	app.Use(start.PrometheusMiddleware())
	app.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// initialize database
	client := database.EntClient(context.Background())

	// initialize routes
	api.InitializeRoutes(app, client)

	// set port
	start.Initialize(app)

}

package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"placio-app/cmd"
	"placio-app/db"
	_ "placio-app/ent/runtime"
	"time"
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
	app := gin.New()

	//AllowOrigins:     []string{"http://localhost:3000", "https://placio.io", "https://www.placio.io", "https://control.placio.io"},
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length,Content-Type,Authorization,X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	app.Use(func(c *gin.Context) {
		c.Header("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Next()
	})

	//c.Header("X-Frame-Options", "SAMEORIGIN")
	// apply gin middleware
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	app.Use(cmd.PrometheusMiddleware())
	app.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// initialize middleware
	cmd.Middleware(app)

	// initialize db
	client := db.EntClient(context.Background())

	// initialize routes
	cmd.InitializeRoutes(app, client)

	// set port
	cmd.Initialize(app)

}

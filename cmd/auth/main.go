package main

import (
	"os"
	"placio-pkg/start"

	"github.com/gofiber/fiber/v2"
)

// @title Placio Application Api
// @version 0.01
// @description This is the documentation for the Placio Application Api
// @termsOfService https://placio.io/terms

// @contact.name Darc Technologies
// @contact.url https://placio.io
// @contact.email support@placio.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:7070
// @BasePath /
// @schemes http
func main() {
	// get port from env
	port := os.Getenv("PORT")

	// initialize fiber app
	app := fiber.New()

	// set port
	start.Initialize(port, app)

}

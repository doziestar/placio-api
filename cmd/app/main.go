package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"os"
	"placio-app/api"
	"placio-app/database"
	"placio-app/start"
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
	// if port is not set, set it to 3000

	// initialize fiber app
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// initialize routes
	api.InitializeRoutes(app)
	// initialize database
	//env, _ := config.LoadConfig("./config")
	_, err := database.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		return
	}
	// set port
	start.Initialize(port, app)

}

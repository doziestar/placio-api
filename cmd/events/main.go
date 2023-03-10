package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"placio-pkg/start"
)

func main() {
	// get port from env
	port := os.Getenv("PORT")

	// initialize fiber app
	app := fiber.New()

	// set port
	start.Initialize(port, app)

}

package main

import (
	"os"
	"placio-pkg/start"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// get port from env
	port := os.Getenv("PORT")

	// initialize fiber app
	app := fiber.New()

	// set port
	start.Initialize(port, app)

}

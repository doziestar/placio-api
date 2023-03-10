package main

import (
	"context"
	"log"
	"os"
	"placio-pkg/start"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// get port from env
	port := os.Getenv("PORT")
	// if port is not set, set it to 3000

	ctx := context.Background()
	log.Println(ctx, "Starting app on port: "+port)

	// initialize fiber app
	app := fiber.New()

	log.Println(ctx, "App started on port: "+port)
	// set port
	start.Initialize(port, app)

}

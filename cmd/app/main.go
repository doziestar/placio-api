package main

import (
	"os"
	"placio-pkg/pkg/start"

	"github.com/gofiber/fiber"
)


func main(){
	// get port from env
	port := os.Getenv("PORT")

	// initialize fiber app
	app := fiber.New()

	// set port
	start.Initialize(port, app)

}
package start

import (
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"

	//"github.com/gofiber/secure/v2"
	"time"
	//"github.com/gofiber/fiber/v2/middleware/recover"
	//"github.com/gofiber/fiber/v2/middleware/timeout"\
	// import sqlite3 driver\
)

func Initialize(PORT string, app *fiber.App) {

	Middleware(app)

	err := sentry.Init(sentry.ClientOptions{
		Dsn: "",
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
		ServerName:       "placio-api",
		Release:          "1.0.0",
		Environment:      "development",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)
	err = app.Listen(":" + PORT)
	if err != nil {
		return
	}
}

package start

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

func Initialize(PORT string, app *gin.Engine) {
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
	err = app.Run(":" + PORT)
	if err != nil {
		return
	}
}

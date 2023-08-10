package cmd

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/getsentry/sentry-go"
)

func Initialize(app *gin.Engine) {

	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate:   1.0,
		EnableTracing:      true,
		AttachStacktrace:   true,
		ProfilesSampleRate: 1.0,
		TracesSampler: sentry.TracesSampler(func(ctx sentry.SamplingContext) float64 {
			// As an example, this custom sampler does not send some
			// transactions to Sentry based on their name.
			hub := sentry.GetHubFromContext(context.Background())
			hub.Scope().SetTag("transaction", ctx.Parent.Name)
			if hub == nil {
				return 0.0
			}
			return 1.0
		}),
		ServerName:  "placio-api",
		Release:     "1.0.0",
		Environment: os.Getenv("ENVIRONMENT"),
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	srv := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: app,
	}

	go func() {
		// service connections
		log.Println("Listening on port " + os.Getenv("PORT"))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sentry.CaptureEvent(&sentry.Event{
				Level:   sentry.LevelError,
				Message: "Server Listen",
				Extra: map[string]interface{}{
					"error": err,
				},
			})

			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		sentry.CaptureEvent(&sentry.Event{
			Level:   sentry.LevelError,
			Message: "Server Shutdown",
			Extra: map[string]interface{}{
				"error": err,
			},
		})
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)
}

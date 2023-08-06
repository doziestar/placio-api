package start

import (
	"context"
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
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
		Dsn:                os.Getenv("SENTRY_DSN"),
		TracesSampleRate:   1.0,
		EnableTracing:      true,
		AttachStacktrace:   true,
		ProfilesSampleRate: 1.0,
		TracesSampler: sentry.TracesSampler(func(ctx sentry.SamplingContext) float64 {
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

	tlsConfig := &tls.Config{
		// Cause the server to use HTTP/2.
		NextProtos: []string{
			http2.NextProtoTLS,
		},
	}

	srv := &http.Server{
		Addr:      ":" + os.Getenv("PORT"),
		Handler:   app,
		TLSConfig: tlsConfig,
	}

	go func() {
		// service connections
		log.Println("Listening on port " + os.Getenv("PORT"))
		if err := srv.ListenAndServeTLS("/app/cert/server.pem", "/app/cert/server.key"); err != nil && err != http.ErrServerClosed {
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

	quit := make(chan os.Signal)
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

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")

	defer sentry.Flush(2 * time.Second)
}

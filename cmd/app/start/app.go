package start

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Initialize(PORT string, app *gin.Engine) {

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "",
		TracesSampleRate: 1.0,
		ServerName:       "placio-api",
		Release:          "1.0.0",
		Environment:      "development",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	srv := &http.Server{
		Addr:    ":" + PORT,
		Handler: app,
	}

	// Enable HTTP/2
	http2.ConfigureServer(srv, &http2.Server{})

	go func() {
		if err := srv.ListenAndServeTLS("server.crt", "server.key"); err != nil && err != http.ErrServerClosed {
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
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
	defer sentry.Flush(2 * time.Second)
}

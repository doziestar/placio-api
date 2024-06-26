package cmd

import (
	"context"
	"fmt"
	"github.com/axiaoxin-com/logging"
	"github.com/gin-contrib/requestid"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"os"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
)

func Middleware(app *gin.Engine) {
	// Rate limiting
	lmt := tollbooth.NewLimiter(20, nil)
	app.Use(tollbooth_gin.LimitHandler(lmt))

	// Request ID
	app.Use(requestid.New())

	// Session middleware
	store := cookie.NewStore([]byte("secret"))
	app.Use(sessions.Sessions("mysession", store))

	gin.SetMode(os.Getenv("GIN_MODE"))
	// you can custom the config or use logging.GinLogger() by default config
	conf := logging.GinLoggerConfig{
		Formatter: func(c context.Context, m logging.GinLogDetails) string {
			return fmt.Sprintf("%s use %s request %s at %v, handler %s use %f seconds to respond it with %d",
				m.ClientIP, m.Method, m.RequestURI, m.ReqTime, m.HandlerName, m.Latency, m.StatusCode)
		},
		SkipPaths:     []string{},
		EnableDetails: false,
		TraceIDFunc:   func(context.Context) string { return "fake-uuid" },
	}
	app.Use(logging.GinLoggerWithConfig(conf))

	app.Delims("{{", "}}")

}

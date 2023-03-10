package start

import (
	"time"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cookie"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/redirect"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/secure"
	"github.com/gofiber/fiber/v2/middleware/session"
	// "github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/storage/sqlite3"
	"github.com/gofiber/utils"
)

func Initialize(PORT string, app *fiber.App) {
	app.Use(recover.New())
	// app.Use(timeout.New(
	// 	timeout.Config{
	// 		Timeout: 5 * time.Second,
	// 	},
	// ))
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        20,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendFile("./toofast.html")
		},
		// Storage: myCustomStorage{},
	}))
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))
	app.Use(pprof.New())
	app.Use(etag.New())
	app.Use(compress.New())
	app.Use(secure.New())
	app.Use(helmet.New())
	app.Use(filesystem.New())
	app.Use(session.New())
	storage := sqlite3.New()
	store := session.New(session.Config{
		Storage: storage,
	})
	app.Use(store)
	app.Use(cookie.New())
	app.Use(csrf.New(csrf.Config{
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     "csrf_",
		CookieSameSite: "Strict",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
		Extractor:      func(c *fiber.Ctx) (string, error) { return c.FormValue("csrf"), nil },
	}))
	app.Use(redirect.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))
	app.Listen(":" + PORT)
}

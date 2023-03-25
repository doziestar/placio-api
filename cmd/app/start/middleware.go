package start

import (
	"os"
	"time"

	"github.com/gofiber/contrib/fibersentry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/helmet/v2"
	// "gorm.io/gorm/logger"
	// "gorm.io/gorm/logger"
)

func Middleware(app *fiber.App) {
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
		Format:       "${pid} ${locals:requestid} ${status} - ${method} ${path}​ ${ip} ${latency}\n",
		TimeFormat:   "02-Jan-2006",
		TimeZone:     "WAT",
		TimeInterval: 0,
		Output:       os.Stdout,
	}))
	fibersentry.New(fibersentry.Config{
		Repanic:         true,
		WaitForDelivery: true,
	})
	app.Use(pprof.New())
	app.Use(etag.New())
	app.Use(compress.New())
	//app.Use(csrf.New(csrf.Config{
	//	KeyLookup:      "header:X-Csrf-Token", // string in the form of '<source>:<key>' that is used to extract token from the request
	//	CookieName:     "my_csrf_",            // name of the session cookie
	//	CookieSameSite: "Strict",              // indicates if CSRF cookie is requested by SameSite
	//	Expiration:     3 * time.Hour,         // expiration is the duration before CSRF token will expire
	//	KeyGenerator:   utils.UUID,            // creates a new CSRF token
	//}))
	//app.Use(secure.New())
	app.Use(helmet.New())
	//app.Use(filesystem.New())
	//app.Use(session.New())
	//storage := sqlite3.New()
	//store := session.New(session.Config{
	//	Storage: storage,
	//})
	//app.Use(store)
	//app.Use(cookie.New())
	app.Use(csrf.New(csrf.Config{
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     "csrf_",
		CookieSameSite: "Strict",
		Expiration:     3 * time.Hour,
		KeyGenerator:   utils.UUID,
		Extractor:      func(c *fiber.Ctx) (string, error) { return c.FormValue("csrf"), nil },
	}))
	//app.Use(redirect.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-AccountType, Accept",
	}))
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))
}

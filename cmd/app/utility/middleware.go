package utility

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Scope string `json:"scope"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx *gin.Context) error {
	return nil
}

// EnsureValidToken is a middleware that will check the validity of our JWT.
func EnsureValidToken() gin.HandlerFunc {
	issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer URL: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE")},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to set up the JWT validator: %v", err)
	}

	errorHandler := func(c *gin.Context, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to validate JWT.",
		})
		c.Abort()
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(convertErrorHandler(errorHandler)),
	)

	return func(c *gin.Context) {
		err := middleware.CheckJWT(c.Writer, c.Request)
		if err != nil {
			c.Abort()
		}
	}
}

// convertErrorHandler adapts the Gin context-based error handler to match jwtmiddleware.ErrorHandler signature.
func convertErrorHandler(handler func(*gin.Context, error)) jwtmiddleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request, err string) {
		ctx := r.Context()
		c, _ := gin.ContextFromContext(ctx)
		handler(c, error(err))
	}
}

// HasScope checks whether our claims have a specific scope.
func (c CustomClaims) HasScope(expectedScope string) bool {
	result := strings.Split(c.Scope, " ")
	for i := range result {
		if result[i] == expectedScope {
			return true
		}
	}

	return false
}

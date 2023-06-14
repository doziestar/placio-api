package middleware

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"placio-app/ent"
	"placio-app/ent/user"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

type CustomClaims struct {
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Nickname      string `json:"nickname"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
	UpdatedAt     string `json:"updated_at"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Scope         string `json:"scope"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// EnsureValidToken is a middleware that will check the validity of our JWT.
func EnsureValidToken(client *ent.Client) gin.HandlerFunc {
	issuerURL, err := url.Parse(os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE"), "KpDGogGXqWeuGQfZ4Wu30neiHS79hGiU"},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is missing"})
			c.Abort()
			return
		}

		// remove the Bearer prefix
		tokenString = tokenString[7:]
		tokenInterface, err := jwtValidator.ValidateToken(context.Background(), tokenString)
		if err != nil {
			log.Printf("Encountered error while validating JWT: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to validate JWT."})
			c.Abort()
			return
		}

		if validatedClaims, ok := tokenInterface.(*validator.ValidatedClaims); ok {
			// Now you can access the fields of validatedClaims
			//fmt.Println(validatedClaims.RegisteredClaims.Issuer)
			//fmt.Println(validatedClaims.RegisteredClaims.Subject)
			//fmt.Println(validatedClaims.RegisteredClaims.Audience)
			//
			//fmt.Println(validatedClaims.CustomClaims.(*CustomClaims).GivenName)
			//fmt.Println(validatedClaims.CustomClaims.(*CustomClaims).FamilyName)
			//fmt.Println(validatedClaims.CustomClaims.(*CustomClaims).Nickname)
			//fmt.Println(validatedClaims.CustomClaims.(*CustomClaims).Name)
			//fmt.Println(validatedClaims.CustomClaims.(*CustomClaims).Picture)
			//fmt.Println(validatedClaims.CustomClaims.(*CustomClaims).Locale)
			//fmt.Println(validatedClaims.CustomClaims.(*CustomClaims).UpdatedAt)
			//fmt.Println(validatedClaims.CustomClaims.(*CustomClaims).Email)

			//split := strings.Split(validatedClaims.RegisteredClaims.Subject, "|")
			//// split[0] will have the provider and split[1] will have the ID
			//id := split[1]
			// make a request to the user service to get the user details
			user, err := client.User.Query().Where(user.Auth0IDEQ(validatedClaims.RegisteredClaims.Subject)).Only(context.Background())
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
				c.Abort()
				return
			}
			c.Set("user", user.ID)
			c.Set("auth0_id", validatedClaims.RegisteredClaims.Subject)
			//c.Set("token", tokenString)
			c.Next()
		} else {
			// handle error, the assertion failed
		}
	}
}

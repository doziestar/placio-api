package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type JWTClaims struct {
	AccountID  uint   `json:"accountId"`
	UserID     uint   `json:"userId"`
	Permission string `json:"permission"`
	Provider   string `json:"provider"`
	jwt.StandardClaims
}

func GenerateToken(data map[string]interface{}, secret string, duration int) (string, error) {
	claims := JWTClaims{
		AccountID:  uint(data["accountId"].(float64)),
		UserID:     uint(data["userId"].(float64)),
		Permission: data["permission"].(string),
		Provider:   data["provider"].(string),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Minute * time.Duration(duration))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func VerifyToken(tokenString string, secret string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

func Verify(permission string, scope string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.Get("Authorization")

		if header == "" {
			if permission == "public" {
				return c.Next()
			} else {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "No authorization header provided",
				})
			}
		}

		parts := strings.Split(header, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unrecognised header type",
			})
		}

		token := parts[1]

		claims, err := jwt.ParseWithClaims(token, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_SECRET")), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token",
			})
		}

		if claims.Valid && claims.Claims.(*JwtCustomClaims).AccountId != "" &&
			claims.Claims.(*JwtCustomClaims).UserId != "" &&
			claims.Claims.(*JwtCustomClaims).Permission != "" &&
			claims.Claims.(*JwtCustomClaims).Provider != "" {

			if permission == "public" || permissions[claims.Claims.(*JwtCustomClaims).Permission][permission] {
				c.Locals("account", claims.Claims.(*JwtCustomClaims).AccountId)
				c.Locals("user", claims.Claims.(*JwtCustomClaims).UserId)
				c.Locals("permission", claims.Claims.(*JwtCustomClaims).Permission)
				c.Locals("provider", claims.Claims.(*JwtCustomClaims).Provider)
				return c.Next()
			} else {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "User doesn't have permission",
				})
			}

		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid auth token",
			})
		}
	}
}

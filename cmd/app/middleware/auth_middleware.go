package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"log"
	"os"
	"placio-app/database"
	"placio-app/models"
	"placio-pkg/logger"
	"strings"
	"time"
)

type JWTClaims struct {
	AccountID  uint   `json:"accountId"`
	UserID     uint   `json:"userId"`
	Permission string `json:"permission"`
	Provider   string `json:"provider"`
	jwt.StandardClaims
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

func Verify(permission string) fiber.Handler {
	logger.Info(context.Background(), fmt.Sprintf("permission: %s", permission))
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
				"message": "Unrecognized header type",
			})
		}
		var err error
		token := parts[1]
		//token, err = hash.DecryptString(token, "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
		//logger.Info(context.Background(), fmt.Sprintf("token: %s", token))

		claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token",
			})
		}
		customClaims, ok := claims.Claims.(jwt.MapClaims)

		logger.Info(context.Background(), fmt.Sprintf("customClaims: %v", customClaims))
		logger.Info(context.Background(), fmt.Sprintf("ok: %v", ok))
		logger.Info(context.Background(), fmt.Sprintf("claims.Valid: %v", claims.Valid))
		if !ok || !claims.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid auth token",
			})
		}

		// Define your permissions map here
		permissions := map[string]map[string]bool{
			"admin": {
				"admin":   true,
				"manager": true,
				"user":    true,
			},
			"manager": {
				"admin":   false,
				"manager": true,
				"user":    true,
			},
			"user": {
				"admin":   false,
				"manager": false,
				"user":    true,
			},
		}

		var user *models.User
		//var account *models.Account

		logger.Info(context.Background(), fmt.Sprintf("customClaims[\"sub\"]: %v", customClaims["sub"].(string)))
		userAccount, err := user.GetUserById(customClaims["sub"].(string), database.DB)
		//logger.Info(context.Background(), fmt.Sprintf("userAccount: %v", userAccount))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid auth token",
			})
		}

		// check token expiration time
		if customClaims["exp"].(float64) < float64(time.Now().Unix()) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Token expired",
			})
		}

		if permission == "user" || permissions[userAccount.Permission][permission] {
			//c.Locals("account", userAccount.Accounts[0].ID)
			//logger.Info(context.Background(), fmt.Sprintf("userAccount.ID: %s", userAccount.ID))
			//logger.Info(context.Background(), fmt.Sprintf("userAccount.Permission: %s", userAccount.Permission))
			log.Println("user permissions", permissions[userAccount.Permission][permission])
			c.Locals("user", userAccount.ID)
			c.Locals("account", func() string {
				if userAccount.CurrentActiveAccount != "" {
					return userAccount.CurrentActiveAccount
				} else {
					return userAccount.DefaultAccount
				}
			}())
			c.Locals("tokenID", customClaims["jti"].(string))
			c.Locals("permission", userAccount.Permission)
			c.Locals("provider", "app")
			return c.Next()
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "User doesn't have permission",
			})
		}
	}
}

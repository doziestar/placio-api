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
)

type JWTClaims struct {
	AccountID  uint   `json:"accountId"`
	UserID     uint   `json:"userId"`
	Permission string `json:"permission"`
	Provider   string `json:"provider"`
	jwt.StandardClaims
}

//
//func GenerateToken(data map[string]interface{}, secret string, duration int) (string, error) {
//	claims := JWTClaims{
//		AccountID:  uint(data["accountId"].(float64)),
//		UserID:     uint(data["userId"].(float64)),
//		Permission: data["permission"].(string),
//		Provider:   data["provider"].(string),
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: jwt.At(time.Now().Add(time.Minute * time.Duration(duration))),
//		},
//	}
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	return token.SignedString([]byte(secret))
//}

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

//func Verify(permission string, scope string) fiber.Handler {
//	return func(c *fiber.Ctx) error {
//		header := c.Get("Authorization")
//
//		if header == "" {
//			if permission == "public" {
//				return c.Next()
//			} else {
//				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
//					"message": "No authorization header provided",
//				})
//			}
//		}
//
//		parts := strings.Split(header, " ")
//
//		if len(parts) != 2 || parts[0] != "Bearer" {
//			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
//				"message": "Unrecognised header type",
//			})
//		}
//
//		token := parts[1]
//
//		claims, err := jwt.ParseWithClaims(token, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
//			return []byte(os.Getenv("TOKEN_SECRET")), nil
//		})
//
//		if err != nil {
//			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
//				"message": "Invalid token",
//			})
//		}
//
//		if claims.Valid && claims.Claims.(*JwtCustomClaims).AccountId != "" &&
//			claims.Claims.(*JwtCustomClaims).UserId != "" &&
//			claims.Claims.(*JwtCustomClaims).Permission != "" &&
//			claims.Claims.(*JwtCustomClaims).Provider != "" {
//
//			if permission == "public" || permissions[claims.Claims.(*JwtCustomClaims).Permission][permission] {
//				c.Locals("account", claims.Claims.(*JwtCustomClaims).AccountId)
//				c.Locals("user", claims.Claims.(*JwtCustomClaims).UserId)
//				c.Locals("permission", claims.Claims.(*JwtCustomClaims).Permission)
//				c.Locals("provider", claims.Claims.(*JwtCustomClaims).Provider)
//				return c.Next()
//			} else {
//				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
//					"message": "User doesn't have permission",
//				})
//			}
//
//		} else {
//			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
//				"message": "Invalid auth token",
//			})
//		}
//	}
//}

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

		token := parts[1]
		logger.Info(context.Background(), fmt.Sprintf("token: %s", token))

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

		userAccount, err := user.GetUserById(customClaims["sub"].(string), database.DB)
		logger.Info(context.Background(), fmt.Sprintf("userAccount: %v", userAccount))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid auth token",
			})
		}

		if permission == "user" || permissions[userAccount.Permission][permission] {
			//c.Locals("account", userAccount.Accounts[0].ID)
			//logger.Info(context.Background(), fmt.Sprintf("userAccount.ID: %s", userAccount.ID))
			//logger.Info(context.Background(), fmt.Sprintf("userAccount.Permission: %s", userAccount.Permission))
			log.Println("user permissions", permissions[userAccount.Permission][permission])
			c.Locals("user", userAccount.ID)
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

package middleware

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"placio-app/Dto"
	"placio-app/database"
	"placio-app/models"
	"placio-pkg/logger"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

func Verify(permission string) gin.HandlerFunc {
	//logger.Info(context.Background(), fmt.Sprintf("permission: %s", permission))
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		if header == "" {
			if permission == "public" {
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "No authorization header provided",
				})
				c.Abort()
			}
			return
		}

		parts := strings.Split(header, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unrecognized header type",
			})
			c.Abort()
			return
		}
		var err error
		token := parts[1]

		claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
			c.Abort()
			return
		}
		customClaims, ok := claims.Claims.(jwt.MapClaims)

		logger.Info(context.Background(), fmt.Sprintf("customClaims: %v", customClaims))
		logger.Info(context.Background(), fmt.Sprintf("ok: %v", ok))
		logger.Info(context.Background(), fmt.Sprintf("claims.Valid: %v", claims.Valid))
		if !ok || !claims.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid auth token",
			})
			c.Abort()
			return
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

		logger.Info(context.Background(), fmt.Sprintf("customClaims[\"sub\"]: %v", customClaims["sub"].(string)))
		userAccount, err := user.GetUserById(customClaims["sub"].(string), database.DB)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid auth token",
			})
			c.Abort()
			return
		}

		// check token expiration time
		if customClaims["exp"].(float64) < float64(time.Now().Unix()) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token expired",
			})
			c.Abort()
			return
		}

		if permission == "user" || permissions[userAccount.Permission][permission] {
			log.Println("user permissions", permissions[userAccount.Permission][permission])
			c.Set("user", userAccount.ID)
			c.Set("account", userAccount.ActiveAccountID)
			c.Set("tokenID", customClaims["jti"].(string))
			c.Set("permission", userAccount.Permission)
			c.Set("provider", "app")
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "User doesn't have permission",
			})
			c.Abort()

		}
	}
}

func AuthorizeUser(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")

		if authorizationHeader == "" {
			if permission == "public" {
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "No authorization header provided",
				})
				c.Abort()
			}
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")

		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unrecognized authorization header type",
			})
			c.Abort()
			return
		}

		token := headerParts[1]

		loginData := Status(token, c)

		c.Set("user", loginData.User.ID)
		c.Set("email", loginData.User.Email)
		c.Next()
	}
}

func Status(token string, c *gin.Context) Dto.LoginData {
	url := fmt.Sprintf("http://localhost:3004/api/v1/auth/authorize?token=%s&type=%s", token, "Bearer")
	log.Println("url", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("err", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		c.Abort()
		return Dto.LoginData{}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid response body",
		})
		c.Abort()
		return Dto.LoginData{}
	}

	loginData, err := Dto.UnmarshalLoginData(body)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		c.Abort()
		return Dto.LoginData{}
	}

	return loginData
}


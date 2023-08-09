package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"placio-app/database"
	"placio-app/domains/search"
	"placio-app/domains/users"
	"placio-app/utility"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestUserController(t *testing.T) {
	client := database.EntClient(context.Background())
	redisClient := utility.NewRedisClient("redis://default:a3677c1a7b84402eb34efd55ad3cf059@golden-colt-33790.upstash.io:33790", 0, utility.CacheDuration)
	_ = redisClient.ConnectRedis()
	searchService, _ := search.NewSearchService()

	userService := users.NewUserService(client, redisClient, searchService)
	userController := users.NewUserController(userService)

	gin.SetMode(gin.TestMode)

	t.Run("GetUser", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)

		c.Request = req
		userController.GetUser(c)

		require.Equal(t, http.StatusOK, recorder.Code)
		// further assertions based on what the response is supposed to look like
	})

	t.Run("UpdateUser", func(t *testing.T) {
		user := map[string]interface{}{
			"name": "New Name",
		}
		b, err := json.Marshal(user)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPatch, "/users", bytes.NewBuffer(b))
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)

		c.Request = req
		userController.UpdateUser(c)

		require.Equal(t, http.StatusOK, recorder.Code)
		// further assertions based on what the response is supposed to look like
	})
}

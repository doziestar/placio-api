package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	appErrors "placio-pkg/errors"
	"testing"
)

func TestUse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		mockFn         func(*gin.Context) error
		expectedStatus int
		expectedError  string
	}{
		{
			name: "Handle ErrInvalid",
			mockFn: func(c *gin.Context) error {
				return appErrors.Wrap(appErrors.ErrInvalid)
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  appErrors.ErrInvalid.Error(),
		},
		{
			name: "Handle ErrUnauthorized",
			mockFn: func(c *gin.Context) error {
				return appErrors.Wrap(appErrors.ErrUnauthorized)
			},
			expectedStatus: http.StatusUnauthorized,
			expectedError:  appErrors.ErrUnauthorized.Error(),
		},
		// ... Continue similarly for all other errors in the map and additional errors
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			r.GET("/", Use(tt.mockFn))

			w := performRequest(r, "GET", "/")

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedError)
		})
	}
}

func performRequest(r *gin.Engine, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

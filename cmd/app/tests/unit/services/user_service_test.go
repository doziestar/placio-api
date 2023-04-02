package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
)

type MockAccountService struct {
	mock.Mock
}

func (m *MockAccountService) SwitchUserAccount(ctx *fiber.Ctx) (*fiber.Map, error) {
	args := m.Called(ctx)
	return args.Get(0).(*fiber.Map), args.Error(1)
}

func TestSwitchUserAccount(t *testing.T) {
	t.Run("switch account success", func(t *testing.T) {
		app := fiber.New()

		mockAccountService := new(MockAccountService)
		mockAccountService.On("SwitchUserAccount", mock.Anything).Return(&fiber.Map{
			"message": "Account switched successfully",
			"data":    nil,
		}, nil)

		// Assuming the route for your endpoint is "/switch-account"
		app.Post("/switch-account", func(c *fiber.Ctx) error {
			res, err := mockAccountService.SwitchUserAccount(c)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Internal Server Error",
				})
			}
			return c.Status(fiber.StatusOK).JSON(res)
		})

		req := httptest.NewRequest("POST", "/switch-account", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		mockAccountService.AssertExpectations(t)
	})
}

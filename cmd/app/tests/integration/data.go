package integration

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"net/http"
	"placio-app/Dto"
	"placio-app/models"
)

var (
	user     models.User
	userData = Dto.SignUpDto{
		Name:            "Test User",
		Email:           "test2@example.com",
		Password:        "test_password",
		ConfirmPassword: "test_password",
		AccountType:     "user",
	}

	app = fiber.New(fiber.Config{})
	ctx = app.AcquireCtx(&fasthttp.RequestCtx{})

	baseUrl = "http://localhost:7070/api/v1"
	client  = &http.Client{}
)

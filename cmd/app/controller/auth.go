package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	Dto "placio-app/Dto"
	"placio-app/database"
	"placio-app/models"
	"placio-app/service"
	"placio-pkg/logger"
)

// Login godoc
// @Summary Login and get token
// @Description Login to the server
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} Dto.UserResponseDto
// @Failure 400 {object} map[string]interface{}
// @QueryParam email query string true "email"
// @QueryParam password query string true "password"
// @Body 200 {object} Dto.LoginDto
// @Router /api/v1/auth/login [get]
func Login(c *fiber.Ctx) error {
	// get data from request body
	data := new(Dto.LoginDto)
	if err := c.BodyParser(data); err != nil {
		return c.JSON(fiber.Map{"status": "error", "message": "invalid data", "data": err})
	}

	// validate data
	email, password := data.IsValid()

	// call service
	auth := service.NewAuth(&database.Database{}, &models.User{
		Email:    email.(string),
		Password: password.(string),
	})

	user, err := auth.LogIn(*data)
	if err != nil {
		return c.JSON(fiber.Map{"status": "error", "message": "invalid data", "data": err})
	}

	return c.JSON(fiber.Map{"status": "ok", "data": user})

}

// SignUp godoc
// @Summary Sign up
// @Description Sign up to the server
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} Dto.UserResponseDto
// @Failure 400 {object} map[string]interface{}
// @Param user body Dto.SignUpDto true "user"
// @Body 200 {object} Dto.SignUpDto
// @Router /api/v1/auth/signup [post]
func SignUp(c *fiber.Ctx) error {
	var data Dto.SignUpDto
	logger.Info(context.Background(), string(c.Body()))
	if err := c.BodyParser(&data); err != nil {
		logger.Info(context.Background(), data.Email)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "invalid data", "data": err})
	}

	// validate data
	userData, err := data.IsValid()
	if err != nil {
		logger.Info(context.Background(), err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "invalid data", "data": err})
	}
	logger.Info(context.Background(), userData.Email)

	// call service
	auth := service.NewAuth(&database.Database{}, &models.User{
		Email:    userData.Email,
		Password: userData.Password,
		Name:     userData.Name,
		Role:     userData.Role,
	})

	user, err := auth.SignUp(data)
	if err != nil {
		return c.JSON(fiber.Map{"status": "error", "message": "invalid data", "data": err})
	}

	return c.JSON(fiber.Map{"status": "ok", "data": user})
}

// LogOut godoc
// @Summary Log out
// @Description Log out from the server
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/auth/logout [get]
func LogOut(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

// RefreshToken godoc
// @Summary Refresh token
// @Description Refresh token
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/auth/refresh [get]
func RefreshToken(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

// ResetPassword godoc
// @Summary Reset password
// @Description Reset password
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/auth/reset [get]
func ResetPassword(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

// ChangePassword godoc
// @Summary Change password
// @Description Change password
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/auth/change [get]
func ChangePassword(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

// VerifyEmail godoc
// @Summary Verify email
// @Description Verify email
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/auth/verify [get]
func VerifyEmail(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

// VerifyPhone godoc
// @Summary Verify phone
// @Description Verify phone
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/auth/verify [get]
func VerifyPhone(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

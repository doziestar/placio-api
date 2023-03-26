package controller

import "github.com/gofiber/fiber/v2"

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/health [get]
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

package controller

import (
	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/health [get]
func HealthCheck(c *gin.Context) error {
	c.JSON(200, gin.H{
		"status": "ok",
	})
	return nil
}

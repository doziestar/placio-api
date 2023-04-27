package api

import (
	"github.com/gin-gonic/gin"
	"placio-app/controller"
	"placio-app/utility"
)

func HealthCheckRoutes(api *gin.RouterGroup) {
	api.GET("/", utility.Use(controller.HealthCheck))

}

package routes

import (
	pingControllers "api/internal/ping/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	pingController := pingControllers.NewPingController()
	router.GET("/ping", pingController.Index)
}

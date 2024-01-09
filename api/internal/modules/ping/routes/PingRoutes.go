package routes

import (
	PingController "api/internal/modules/ping/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	pingController := PingController.New()
	router.GET("/ping", pingController.Index)
}

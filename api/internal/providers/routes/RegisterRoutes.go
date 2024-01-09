package routes

import (
	pingRoutes "api/internal/modules/ping/routes"
	todoRoutes "api/internal/modules/todo/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	pingRoutes.RegisterRoutes(router)
	todoRoutes.Register(router)
}

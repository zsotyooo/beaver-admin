package routes

import (
	pingRoutes "api/app/modules/ping/routes"
	todoRoutes "api/app/modules/todo/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	pingRoutes.RegisterRoutes(router)
	todoRoutes.Register(router)
}

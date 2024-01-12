package routes

import (
	authRoutes "api/internal/modules/auth/routes" // Import the missing package
	pingRoutes "api/internal/modules/ping/routes"
	todoRoutes "api/internal/modules/todo/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	pingRoutes.RegisterRoutes(router)
	todoRoutes.Register(router)
	authRoutes.Register(router) // Register the routes
}

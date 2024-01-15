package routesprovider

import (
	authRoutes "api/internal/auth/routes" // Import the missing package
	pingRoutes "api/internal/ping/routes"
	swaggerRoutes "api/internal/swagger/routes"
	todoRoutes "api/internal/todo/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	pingRoutes.RegisterRoutes(router)
	todoRoutes.Register(router)
	authRoutes.Register(router)
	swaggerRoutes.Register(router)
}

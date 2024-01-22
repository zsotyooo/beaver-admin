package routesprovider

import (
	authRoutes "api/internal/auth/routes" // Import the missing package
	pingRoutes "api/internal/ping/routes"
	swaggerRoutes "api/internal/swagger/routes"
	todoRoutes "api/internal/todo/routes"
	userRoutes "api/internal/user/routes"
	"api/pkg/routing"
)

func RegisterRoutes() {
	router := routing.GetRouter()
	pingRoutes.RegisterRoutes(router)
	authRoutes.RegisterRoutes(router)
	userRoutes.RegisterRoutes(router)
	todoRoutes.RegisterRoutes(router)
	swaggerRoutes.RegisterRoutes(router)
}

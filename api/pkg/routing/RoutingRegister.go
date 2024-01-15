package routing

import (
	routesProvider "api/internal/providers/routes"
)

func Register() {
	routesProvider.RegisterRoutes(GetRouter())
}

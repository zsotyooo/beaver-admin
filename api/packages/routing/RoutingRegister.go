package routing

import (
	"api/internal/providers/routes"
)

func Register() {
	routes.RegisterRoutes(GetRouter())
}

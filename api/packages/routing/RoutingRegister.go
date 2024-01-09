package routing

import (
	"api/app/providers/routes"
)

func Register() {
	routes.RegisterRoutes(GetRouter())
}

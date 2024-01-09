package bootstrap

import (
	"api/packages/config"
	"api/packages/database"
	"api/packages/routing"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	routing.Register()

	routing.Serve()
}

package bootstrap

import (
	"api/packages/config"
	"api/packages/cors"
	"api/packages/database"
	"api/packages/logger"
	"api/packages/routing"
)

func Serve() {
	logger.Init()

	config.Set()

	database.Connect()

	routing.Init()

	cors.Init()

	routing.Register()

	routing.Serve()
}

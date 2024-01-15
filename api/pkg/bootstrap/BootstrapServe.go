package bootstrap

import (
	"api/pkg/config"
	"api/pkg/cors"
	"api/pkg/database"
	"api/pkg/logger"
	"api/pkg/routing"
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

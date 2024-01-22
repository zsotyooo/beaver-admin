package bootstrap

import (
	routesProvider "api/internal/providers/routes"
	sessionProvider "api/internal/providers/session"
	"api/pkg/config"
	"api/pkg/cors"
	database "api/pkg/database/postgresql"
	redis "api/pkg/database/redis"
	"api/pkg/logger"
	"api/pkg/routing"
	"api/pkg/session"
)

func Serve() {
	logger.Init()

	config.Set()

	database.PostgresqlConnect()

	redis.RedisConnect()

	routing.Init()

	session.InitSession()

	cors.Init()

	sessionProvider.RegisterSessions()
	routesProvider.RegisterRoutes()

	routing.Serve()
}

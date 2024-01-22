package bootstrap

import (
	"api/internal/database/migration"
	"api/pkg/config"
	database "api/pkg/database/postgresql"
	"api/pkg/logger"
)

func Migrate() {
	logger.Init()

	config.Set()

	database.PostgresqlConnect()

	migration.Migrate()
}

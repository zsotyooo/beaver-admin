package bootstrap

import (
	"api/internal/modules/database/migration"
	"api/packages/config"
	"api/packages/database"
	"api/packages/logger"
)

func Migrate() {
	logger.Init()

	config.Set()

	database.Connect()

	migration.Migrate()
}

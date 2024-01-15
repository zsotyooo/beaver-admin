package bootstrap

import (
	"api/internal/database/migration"
	"api/pkg/config"
	"api/pkg/database"
	"api/pkg/logger"
)

func Migrate() {
	logger.Init()

	config.Set()

	database.Connect()

	migration.Migrate()
}

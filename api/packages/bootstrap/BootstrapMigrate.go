package bootstrap

import (
	"api/internal/modules/database/migration"
	"api/packages/config"
	"api/packages/database"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()
}

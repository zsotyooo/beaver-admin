package migration

import (
	holidayModels "api/internal/modules/holiday/models"
	todoModels "api/internal/modules/todo/models"
	userModels "api/internal/modules/user/models"
	"api/packages/database"
	"fmt"
	"log"
)

func Migrate() {
	db := database.Connection()

	err := db.AutoMigrate(&userModels.User{}, &todoModels.Todo{}, &holidayModels.Holiday{})

	if err != nil {
		log.Fatal("Can't migrate!")
		return
	}

	fmt.Println("Migration done.")
}

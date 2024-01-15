package migration

import (
	holidayModels "api/internal/holiday/models"
	todoModels "api/internal/todo/models"
	userModels "api/internal/user/models"
	"api/pkg/database"
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

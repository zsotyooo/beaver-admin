package migration

import (
	todoModels "api/app/modules/todo/models"
	userModels "api/app/modules/user/models"
	"api/packages/database"
	"fmt"
	"log"
)

func Migrate() {
	db := database.Connection()

	err := db.AutoMigrate(&userModels.User{}, &todoModels.Todo{})

	if err != nil {
		log.Fatal("Can't migrate!")
		return
	}

	fmt.Println("Migration done.")
}

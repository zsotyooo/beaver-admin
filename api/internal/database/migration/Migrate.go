package migration

import (
	"api/internal/todo"
	"api/internal/user"
	database "api/pkg/database/postgresql"
	"fmt"
	"log"
)

func Migrate() {
	db := database.DBConnection()

	err := db.AutoMigrate(&user.User{}, &todo.Todo{})

	if err != nil {
		log.Fatal("Can't migrate!")
		return
	}

	fmt.Println("Migration done.")
}

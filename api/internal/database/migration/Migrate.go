package migration

import (
	"api/internal/todo"
	"api/internal/user"
	"api/pkg/database"
	"fmt"
	"log"
)

func Migrate() {
	db := database.Connection()

	err := db.AutoMigrate(&user.User{}, &todo.Todo{})

	if err != nil {
		log.Fatal("Can't migrate!")
		return
	}

	fmt.Println("Migration done.")
}

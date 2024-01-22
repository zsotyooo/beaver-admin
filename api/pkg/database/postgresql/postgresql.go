package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func PostgresqlConnect() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			os.Getenv("DATABASE_USERNAME"),
			os.Getenv("DATABASE_PASSWORD"),
			os.Getenv("DATABASE_HOST"),
			os.Getenv("DATABASE_DBNAME"),
		),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to database")
		return
	}

	DB = db
}

func DBConnection() *gorm.DB {
	return DB
}

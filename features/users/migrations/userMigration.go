package main

import (
	"log"

	"github.com/Bits-Fusion/the_application_backend/config"
	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/features/users/entities"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	userMigrate(db)
}

func userMigrate(db database.Database) {
	log.Println("[info] creating users table")
	if err := db.GetDb().Migrator().CreateTable(&entities.User{}); err != nil {
		log.Println(err)
	}
	log.Println("[info] finished creating user table")
}

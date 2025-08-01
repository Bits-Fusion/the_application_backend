package main

import (
	"log"

	"github.com/Bits-Fusion/the_application_backend/config"
	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/features/leads/entities"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)

	leadsMigrate(db)
}

func leadsMigrate(db database.Database) {
	log.Println("[info] creating Leads table")
	// _ = db.GetDb().Migrator().DropTable(&entities.Lead{})
	if err := db.GetDb().AutoMigrate(&entities.Lead{}); err != nil {
		log.Println(err)
	}
	log.Println("[info] finished creating Leads table")
}

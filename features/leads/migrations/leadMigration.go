package main

import (
	"log"

	"github.com/Bits-Fusion/the_application_backend/config"
	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/features/leads/entities"
	// "gorm.io/gorm"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	// ensurePriorityEnum(db.GetDb())
	// ensureStatusEnum(db.GetDb())
	leadsMigrate(db)
}

func leadsMigrate(db database.Database) {
	log.Println("[info] creating Leads table")
	// _ = db.GetDb().Migrator().DropTable(&entities.Task{})
	if err := db.GetDb().AutoMigrate(&entities.Lead{}); err != nil {
		log.Println(err)
	}
	log.Println("[info] finished creating Leads table")
}

//
// func ensurePriorityEnum(db *gorm.DB) {
// 	createEnumSQL := `
// 	DO $$ BEGIN
// 		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'priority_enum') THEN
// 			CREATE TYPE priority_enum AS ENUM ('high', 'mid', 'low');
// 		END IF;
// 	END$$;
// 	`
//
// 	if err := db.Exec(createEnumSQL).Error; err != nil {
// 		log.Fatalf("failed to create enum type: %v", err)
// 	}
// }
//
// func ensureStatusEnum(db *gorm.DB) {
// 	createEnumSQL := `
// 	DO $$ BEGIN
// 		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_enum') THEN
// 			CREATE TYPE status_enum AS ENUM ('complete', 'inprogress', 'canceled');
// 		END IF;
// 	END$$;
// 	`
//
// 	if err := db.Exec(createEnumSQL).Error; err != nil {
// 		log.Fatalf("failed to create enum type: %v", err)
// 	}
// }

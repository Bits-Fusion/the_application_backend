package main

import (
	"log"

	"github.com/Bits-Fusion/the_application_backend/config"
	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/features/users/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	ensureRoleEnum(db.GetDb())
	userMigrate(db)
}

func userMigrate(db database.Database) {
	log.Println("[info] creating users table")
	// _ = db.GetDb().Migrator().DropTable(&entities.User{})
	if err := db.GetDb().AutoMigrate(&entities.User{}); err != nil {
		log.Println(err)
	}
	log.Println("[info] finished creating user table")
}

func ensureRoleEnum(db *gorm.DB) {
	createEnumSQL := `
	DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_enum') THEN
			CREATE TYPE role_enum AS ENUM ('admin', 'user', 'editor');
		END IF;
	END$$;
	`

	if err := db.Exec(createEnumSQL).Error; err != nil {
		log.Fatalf("failed to create enum type: %v", err)
	}
}

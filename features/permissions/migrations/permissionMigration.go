package main

import (
	"log"

	"github.com/Bits-Fusion/the_application_backend/config"
	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/features/permissions/entities"
	"github.com/Bits-Fusion/the_application_backend/features/permissions/repositories"
	"gorm.io/gorm"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	userMigrate(db)
	initiatePermissions(db)
}

func userMigrate(db database.Database) {
	log.Println("[info] creating Permission table")
	permissionActionEnum(db.GetDb())
	// 	_ = db.GetDb().Migrator().DropTable(&entities.Permission{})
	if err := db.GetDb().AutoMigrate(&entities.Permission{}); err != nil {
		log.Println(err)
	}
	log.Println("[info] finished creating Permission table")
}

func initiatePermissions(db database.Database) {
	log.Println("[info] initating permissions")
	permissionRepo := repositories.NewPermissionRepository(db)

	resources := []string{"lead", "user", "task"}
	actions := []string{"view", "create", "update", "delete"}

	for _, action := range actions {
		for _, res := range resources {
			permissionRepo.CreatePermission(action, res)
		}
	}
}

func permissionActionEnum(db *gorm.DB) {
	createEnumSQL := `
	DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'permission_action_enum') THEN
			CREATE TYPE permission_action_enum AS ENUM ('create', 'update', 'view', 'delete');
		END IF;
	END$$;
	`

	if err := db.Exec(createEnumSQL).Error; err != nil {
		log.Fatalf("failed to create enum type: %v", err)
	}
}

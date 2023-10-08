package migrate

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"github.com/Joshua-FeatureFlag/backend/models"
)

func RunMigration(db *gorm.DB) {

	// Automatically migrate your schema, to keep it up to date.
	if err := db.AutoMigrate(&models.Organization{}, &models.User{}); err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}

	fmt.Println("Migration completed successfully!")
}

package config

import (
	"fmt"
	"log"
	"os"
	"squadify-app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	fmt.Println("Connecting to the database...")

	dbURL := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"),
	)
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	DB.AutoMigrate(&models.User{}, &models.UserProfile{})
	// if err := migrations.MigrateUserProfiles(DB); err != nil {
	// 	log.Fatalf("Migration failed: %v", err)
	// }
	// TODO: Add other models here
	// DB.AutoMigrate(&models.Contact{})

	fmt.Println("Database connection established.")
}

package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rsp84/ride-with-me-backend/graph/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	dbURL := os.Getenv("DB_CONNECTION_STRING")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.AutoMigrate(models.Models...)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db, nil
}

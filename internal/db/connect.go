package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// InitDB initializes and returns the database connection
func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("events.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// AutoMigrate will create or update the tables based on the model
	err = db.AutoMigrate(&Event{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return db
}

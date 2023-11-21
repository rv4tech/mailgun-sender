package database

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Fallback filename for database.
const defaultDatabaseName = "mailgun_sender.db"

// Database initializer.
// Takes name string as an argument.
// Default value is mailgun_sender.db.
func InitDataBase(databaseName string) *gorm.DB {
	var database *gorm.DB
	// In case empty string was passed, use default hardcoded value.
	if databaseName == "" {
		databaseName = defaultDatabaseName
	}
	database, err := gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connet to database %s\nError: %s\n", databaseName, err)
	}
	// Applying schemas does not require any kind of checks.
	ApplySchemas(database)
	fmt.Printf("Successfully connected to database %s\n", databaseName)
	return database
}

// Apply schemas from models.go file.
// Schemas structs are hardcoded inside of a function.
// Schemas are Campaign, Translation, SendStat
// GORM will not apply schemas of they already filled in the database.
func ApplySchemas(db *gorm.DB) {
	db.AutoMigrate(&Campaign{})
	db.AutoMigrate(&Translation{})
	db.AutoMigrate(&SendStat{})
}

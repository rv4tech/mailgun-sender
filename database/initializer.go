package database

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	database, err := gorm.Open(sqlite.Open(databaseName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Could not connet to database %s\nError: %s\n", databaseName, err)
	}
	// Applying schemas does not require any kind of checks.
	ApplySchemas(database)
	fmt.Printf("Successfully connected to database %s\n\n", databaseName)
	return database
}

// Apply schemas from models.go file.
// Schemas structs are hardcoded inside of a function.
// Schemas are Campaign, Translation, SendStat
// GORM will not apply schemas if they already injected.
func ApplySchemas(db *gorm.DB) {
	db.AutoMigrate(&Campaign{})
	db.AutoMigrate(&Translation{})
	db.AutoMigrate(&SendStat{})
}

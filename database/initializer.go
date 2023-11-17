package database

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DBName = "rv4.db"

func InitDataBase(dbName string, withSchemas bool) (*gorm.DB, error) {
	var database *gorm.DB
	if dbName == "" {
		dbName = DBName
	}
	database, err := gorm.Open(sqlite.Open(DBName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connet to database %s\nError: %s", DBName, err)
	} else {
		fmt.Printf("Successfully connected to database %s", dbName)
		if withSchemas {
			applySchemas(database)
		}
	}
	return database, nil
}

func applySchemas(db *gorm.DB) {
	db.AutoMigrate(&Campaigns{})
	db.AutoMigrate(&Translations{})
	db.AutoMigrate(&SendStats{})
	fmt.Println("Applied schemas")
}

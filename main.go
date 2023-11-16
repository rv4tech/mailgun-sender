package main

import (
	"fmt"
	"rv4-request/database"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DBName = "rv4.db"

func main() {
	db, err := gorm.Open(sqlite.Open(DBName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&database.Campaigns{})

	// Create
	sum := 1
	for sum < 4000 {
		db.Create(&database.Campaigns{Name: fmt.Sprintf("testRecord %v", sum), MgTemplate: "testTemplate"})
		sum += 1
	}
	var campaigns []database.Campaigns

	result := db.Find(&campaigns)
	fmt.Printf("Data:\n %v", result)
}

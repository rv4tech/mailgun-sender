package database

import (
	"log"

	"gorm.io/gorm"
)

// Check if passed gorm query returns error.
func IsQueryValid(query *gorm.DB) bool {
	return query.Error == nil
}

func RaiseDataBaseError(message string) {
	log.Fatal(message)
}

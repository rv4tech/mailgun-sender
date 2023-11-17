package database

import (
	"log"

	"gorm.io/gorm"
)

// Check if passed gorm query returns error.
func CheckForQueryError(query *gorm.DB) {
	if query.Error != nil {
		log.Printf("Error occured [%s]", query.Error)
	}
}

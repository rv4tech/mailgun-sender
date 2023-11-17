package main

import (
	"fmt"
	"rv4-request/database"
)

func main() {
	// Initialize database. Exits if error occurs.
	db, _ := database.InitDataBase("", true)

	entry := database.Translations{
		CampID:  3,
		Lang:    "jp",
		From:    "meme",
		Subject: "check this meme",
	}
	// fmt.Println("Creating entry", entry)
	// database.InsertOneTranslationsRow(db, &entry)

	fmt.Println("Deleting entry", entry)
	database.HardDeleteTranslationsRowByCondition(db, "camp_id", "3")
}

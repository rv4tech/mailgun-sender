package main

import (
	"fmt"
	"rv4-request/database"
)

func main() {
	// Initialize database. Exits if error occurs.
	db, _ := database.InitDataBase("", true)

	entry := database.SendStats{
		Ts:         0,
		CampaignID: 1,
		Lang:       "en",
		Email:      "kekus.com",
		ExtID:      "2",
		Success:    true,
		ErrorMsg:   "",
	}
	fmt.Println("Creating entry", entry)
	database.InsertOneSendStatsRow(db, &entry)
}

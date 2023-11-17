package main

import (
	"rv4-request/arguments"
	"rv4-request/database"
)

func init() {
	arguments.ParseArguments()
}

func main() {
	// Initialize database. Exits if error occurs.
	db, _ := database.InitDataBase("", true)

	entry := database.SendStats{
		Ts:         0,
		CampaignID: 1,
		Lang:       "en",
		Email:      "kekus@pekus.com",
		ExtID:      "2",
		Success:    true,
		ErrorMsg:   "",
	}
	database.InsertOneSendStatsRow(db, &entry)
}

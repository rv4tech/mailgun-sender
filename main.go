package main

import (
	"rv4-request/arguments"
)

func init() {
	arguments.ParseArguments()
}

func main() {
	// Initialize database. Exits if error occurs.
	// db, _ := database.InitDataBase("", true)
	// s := database.SendStats{
	// 	CampaignID: 1,
	// 	Lang:       "en",
	// 	Email:      "kekus@pekus.com",
	// 	ExtID:      "2",
	// 	Success:    true,
	// 	ErrorMsg:   "",
	// }
	// s.InsertOne(db)
}

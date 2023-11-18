package main

import (
	"rv4-request/database"
	"rv4-request/io"
	"strconv"
)

func ApplyFixtures() {
	campaigns := io.ReadCsvFile("campaigns.csv")
	translations := io.ReadCsvFile("translations.csv")

	applyCampaigns(campaigns)
	applyTranslations(translations)
}

func applyCampaigns(data [][]string) {
	db, _ := database.InitDataBase("test.db", true)
	for index, row := range data {
		if index != 0 {
			campaign := database.Campaigns{
				Name:        row[0],
				MgTemplate:  row[1],
				DefaultLang: row[2],
			}
			campaign.InsertOne(db)
		}
	}
}

func applyTranslations(data [][]string) {
	db, _ := database.InitDataBase("test.db", true)
	for index, row := range data {
		if index != 0 {
			campID, _ := strconv.Atoi(row[0])
			translation := database.Translations{
				CampID:  campID,
				Lang:    row[1],
				From:    row[2],
				Subject: row[3],
			}
			translation.InsertOne(db)
		}
	}
}

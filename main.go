package main

import (
	"fmt"
	cmdarguments "rv4-request/cmd_arguments"
	"rv4-request/database"
	"rv4-request/io"
)

type fileData struct {
	client      string
	clientEmail string
	language    string
	externalID  string
}

type sendData struct {
	template string
	language string
	from     string
	subject  string
	to       string
}

func main() {
	// Initialize database. Exits if error occurs.
	db, _ := database.InitDataBase("")

	// Get parsed arguments as strings.
	mailListFileName, campaignName := cmdarguments.ParseArguments()
	fmt.Printf("filename: %s, campaign name: %s\n", *mailListFileName, *campaignName)

	// Read csv file.
	data := io.ReadCsvFile(*mailListFileName)

	// Encapsulate data from csv file into slice of structs.
	var preparationData []fileData
	for _, row := range data {
		toSend := fileData{
			client:      row[0],
			clientEmail: row[1],
			language:    row[2],
			externalID:  row[3],
		}
		preparationData = append(preparationData, toSend)
	}

	// Get db entity from passed campaign name.
	campaign := database.GetCampaignByName(db, *campaignName)

	// Get db entity of campaign related translations.
	translations := database.GetTranslationsByCampaignID(db, int(campaign.ID))
	for _, t := range translations {
		fmt.Printf("\nID IS: %v\nLANG IS: %v", t.CampID, t.Lang)
	}
}

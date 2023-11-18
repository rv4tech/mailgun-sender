package main

import (
	"flag"
	"fmt"
	"log"
	"rv4-request/database"
	"rv4-request/io"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type dataToSend struct {
	client      string
	clientEmail string
	language    string
	externalID  string
}

func init() {
	// ApplyFixtures()
}

func parseArguments() (*string, *string) {
	var maillistPath, campaignName string
	// Flag -ml to read from csv file with a given name. Default value is "maillist.csv"
	flag.StringVar(&maillistPath, "ml", "", "Name of the file to read from. No default value.")

	// Flag -camp to choose campaign with a given name. No default value.
	flag.StringVar(&campaignName, "camp", "", "Name of the campaign. No default value.")
	flag.Parse()

	// Assuming we need both arguments to not be empty.
	condition := campaignName != "" && maillistPath != ""
	if condition {
		data := io.ReadCsvFile(maillistPath)
		fmt.Println(data)
	} else {
		// Stdout -h in case of error or wrong arguments.
		flag.Usage()
	}
	return &maillistPath, &campaignName
}

func main() {
	// Initialize database. Exits if error occurs.
	db, err := gorm.Open(sqlite.Open("rv4.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connet to database %s\nError: %s", "rv4.db", err)
	}
	mailListFileName, campaignName := parseArguments()

	data := io.ReadCsvFile(*mailListFileName)

	var preparationData []dataToSend
	for _, row := range data {
		toSend := dataToSend{
			client:      row[0],
			clientEmail: row[1],
			language:    row[2],
			externalID:  row[3],
		}
		preparationData = append(preparationData, toSend)
	}

	var currentCampaign *database.Campaigns
	currentCampaign = database.Campaigns.GetOne(db, "name", campaignName)
	fmt.Printf("qe: %v+", &currentCampaign)
}

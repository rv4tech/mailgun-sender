package main

import (
	"fmt"
	"log"
	"os"
	cmdarguments "rv4-request/cmd_arguments"
	"rv4-request/database"
	"rv4-request/io"
	"rv4-request/sender"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	// Get dotenv variables from .env.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var domain string = os.Getenv("DOMAIN")
	var apiKey string = os.Getenv("APIKEY")
	// Initialize database from separate function.
	var db *gorm.DB = database.InitDataBase("")

	// Get cmd arguments from -ml and -camp commands.
	var fileName, campaignName string = cmdarguments.ParseArguments()

	// Get initial database data based on passed cmd arguments.
	campaign := database.GetCampaignByName(db, campaignName)
	translations := database.GetTranslationsByCampaignID(db, campaign.ID)

	// Get file data based on filename passed as cmd argument.
	var rawData [][]string = io.ReadCsvFile(fileName)
	var fileData []io.FileData

	// Encapsulate data from csv file for accessability.
	for _, row := range rawData {
		fileData = append(fileData, io.FileData{
			Client:      strings.Trim(row[0], " "),
			ClientEmail: strings.Trim(row[1], " "),
			Language:    strings.Trim(row[2], " "),
			ExternalID:  strings.Trim(row[3], " "),
		})
	}

	// Initialize slice that will containt data about message sent.
	var batchSendStatData []*database.SendStat
	for _, client := range fileData {
		for _, translation := range translations {
			if client.Language == translation.Lang {
				var tags []string
				tags = append(tags, campaign.MgTemplate)
			}

			if translation.Lang != "" {
				tags = append(tags, translation.Lang)
			} else {
				tags = append(tags, campaign.DefaultLang)
			}
			paramInstance := sender.MailGunParams{
				From:    translation.From,
				Subject: translation.Subject,
				Text:    "test",
				Tags:    tags,
				To:      clients,
			}
			responseMessage, responseID, responseError := sender.SendMailGunMessageV3(domain, apiKey, &paramInstance)
			stat := &database.SendStat{
				CampaignID: campaign.ID,
				Lang:       translation.Lang,
				ExtID:      "testID",
				Email:      client,
			}
			if responseError != nil {
				stat.ErrorMsg = fmt.Sprintf("%s", responseError)
				stat.Success = false
			} else {
				stat.ErrorMsg = ""
				stat.Success = true
			}
			batchSendStatData = append(batchSendStatData, stat)
			fmt.Printf("Response message: %s\n", responseMessage)
			fmt.Printf("Response id: %s\n\n", responseID)
		}
	}

	// Create entries as batch.
	db.Create(&batchSendStatData)
}

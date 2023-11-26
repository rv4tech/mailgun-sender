package main

import (
	"fmt"
	"log"
	"os"
	cmdarguments "rv4-request/cmd_arguments"
	"rv4-request/database"
	"rv4-request/io"
	"rv4-request/sender"

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
	campaign, err := database.GetCampaignByName(db, campaignName)
	if err != nil {
		log.Fatal(err)
	}

	var fileData [][]string = io.ReadCsvFile(fileName)
	var clients []*Client = CreateClientsSlice(fileData)

	for _, client := range clients {
		// Get related translation.
		translation, err := database.GetTranslationByCampaignIDAndLanguage(db, campaign.ID, client.Language)
		if err != nil {
			log.Fatal(err)
		}
		// Initialize empty slice with tags.
		var tags []string
		// Initialize empty string with lanuage.
		var language string
		// Add tag as per campaign table.
		tags = append(tags, campaign.MgTemplate)
		// We need to check if translations table related to the campaign contains lanuage string.
		// If not, we substitute it with default language from campaigns table.
		if translation.Lang == "" {
			language = campaign.DefaultLang
		} else {
			language = client.Language
		}
		tags = append(tags, language)
		// Instanciate parameters to fill mailgun request.
		paramInstance := sender.MailGunParams{
			From:    translation.From,
			Subject: translation.Subject,
			Text:    "test",
			Tags:    tags,
			To:      client.Email,
		}
		responseMessage, responseID, responseError := sender.SendMailGunMessageV3(domain, apiKey, &paramInstance)
		// Prepare data to fill send stats table.
		stat := database.SendStat{
			CampaignID: campaign.ID,
			Lang:       language,
			ExtID:      client.ExternalID,
			Email:      client.Email,
		}
		if stat.Exists(db) {
			stat.ErrorMsg = fmt.Sprintf("already sent %v", stat.ID)
		} else {
			if responseError != nil {
				stat.ErrorMsg = fmt.Sprintf("%s", responseError)
				stat.Success = false
			} else {
				stat.ErrorMsg = ""
				stat.Success = true
			}
		}
		db.Create(&stat)

		// Prints for testing purposes.
		fmt.Printf("Sending message to <%s %s>\n", client.Name, client.Email)
		fmt.Println("PARAMS:")
		fmt.Printf("From: %s\n", paramInstance.From)
		fmt.Printf("To: %s\n", paramInstance.To)
		fmt.Printf("Subject: %s\n", paramInstance.Subject)
		fmt.Printf("Text: %s\n", paramInstance.Text)
		fmt.Printf("Error: %s\n\n", stat.ErrorMsg)
		fmt.Println("MAILGUN RESPONSE")
		fmt.Printf("Response message: %s\n", responseMessage)
		fmt.Printf("Response id: %s\n", responseID)
		fmt.Printf("Response error: %s\n\n", responseError)
	}
}

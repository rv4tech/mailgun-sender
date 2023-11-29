package main

import (
	"fmt"
	"log"
	"os"
	cmd "rv4-request/cmd_arguments"
	"rv4-request/database"
	"rv4-request/io"
	"rv4-request/sender"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func init() {
	// Get dotenv variables from .env file.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	var db *gorm.DB = database.InitDataBase("test.db")
	var domain string = os.Getenv("DOMAIN")
	var apiKey string = os.Getenv("APIKEY")
	// Get cmd arguments from -ml and -camp commands.
	var fileName, campaignName string = cmd.ParseArguments()
	// Get initial database data based on passed cmd arguments.
	var campaign database.Campaign
	if err := db.Where("name = ?", campaignName).First(&campaign).Error; err != nil {
		log.Fatalln("Error querying for campaign", err)
	}

	// Structured slice of clients read from csv file.
	clients := CreateClientsSlice(io.ReadCsvFile(fileName))

	for _, client := range clients {
		// Payload for MailGun message sending function.
		var payload sender.MailGunPayload
		// Struct to enscapsulate data for 'send_stats' table.
		var stats database.SendStat

		// Check if row with requested 'camp_id' and 'lang' exists.
		// If not - we use campaign's default language.
		translation, translationDoesNotExist := database.TranslationDoesNotExist(db, campaign.ID, client.Language)
		if translationDoesNotExist {
			// Dereferrence new db entity for 'translation' struct.
			db.Where("camp_id = ? AND lang = ?", campaign.ID, campaign.DefaultLanguage).First(&translation)
		}

		// Create payload to fill MailGun send message function.
		payload.From = translation.From
		payload.Subject = translation.Subject
		payload.Text = fmt.Sprintf("This message is supposed to be in <%s> language", translation.Language)
		payload.To = client.Email
		payload.Tags = append(payload.Tags, translation.Language, campaign.MailgunTemplate)
		payload.TemplateVersion = translation.Language

		// Create some of statistics variables.
		// We do not initialize 'ErrorMessage' and 'Success' until after we send message.
		stats.CampaignID = campaign.ID
		stats.Email = client.Email
		stats.ExternalID = client.ExternalID
		stats.Language = translation.Language

		// Check if message was sent already by three parameters.
		// First is 'camp_id'.
		// Second is 'email'.
		// Third is 'success'. Must be true (1).
		// Even if 'camp_id' and 'email' already exists, but success was false (0), we send the message.
		statExists, statID := database.StatExists(db, &stats)
		if statExists {
			WriteLog("ERROR", fmt.Sprintf("Message for <%s> was not sent. Entry with this statistic already exists with id %v", client.Email, statID))
			stats.ErrorMessage = fmt.Sprintf("already sent <%v>", statID)
			stats.Success = false
			db.Create(&stats)
			continue
		}

		// MailGun response enteties.
		// responseMessage contains basic info about request. If message was sent successfully, it return "Queued. Thank you."
		// responseID contains MailGun message ID.
		// responseError contains error message if any occurs.
		responseMessage, responseID, responseError := sender.SendMailGunMessageV4(domain, apiKey, &payload)
		if responseError != nil {
			WriteLog("ERROR", fmt.Sprintf("Message for <%s> was not sent. MailGun responded with error: %s", client.Email, responseError))
			stats.ErrorMessage = fmt.Sprint(responseError)
			stats.Success = false
		} else {
			WriteLog("INFO", fmt.Sprintf("Message for <%s> was successfully sent. MG message: %s MG id: %s", client.Email, responseMessage, responseID))
			stats.ErrorMessage = ""
			stats.Success = true
		}
		db.Create(&stats)
	}
}

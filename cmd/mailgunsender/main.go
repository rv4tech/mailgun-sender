package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var fileName, campaignName string

func init() {
	// Get dotenv variables from .env file.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Flag -ml to read from csv file with a given name. Default value is "maillist.csv"
	flag.StringVar(&fileName, "ml", "", "Name of the file to read from. No default value.")

	// Flag -camp to choose campaign with a given name. No default value.
	flag.StringVar(&campaignName, "camp", "", "Name of the campaign. No default value.")

	// Parse arguments.
	flag.Parse()

	// Assuming we need both arguments to not be empty.
	if campaignName != "" && fileName != "" {
		ReadCsvFile(fileName)
	} else {
		// Stdout -h in case of error or wrong arguments.
		flag.Usage()
	}
}
func main() {
	var db *gorm.DB = InitDataBase("test.db")
	var domain string = os.Getenv("MG_DOMAIN")
	var apiKey string = os.Getenv("MG_API_KEY")
	// Get initial database data based on passed cmd arguments.
	var campaign Campaign
	if err := db.Where("name = ?", campaignName).First(&campaign).Error; err != nil {
		log.Fatalln("Error querying for campaign", err)
	}

	// Structured slice of clients read from csv file.
	clients := CreateClientsSlice(ReadCsvFile(fileName))

	for num, client := range clients {
		// Payload for MailGun message sending function.
		var payload MailGunPayload
		// Struct to enscapsulate data for 'send_stats' table.
		var stats SendStat
		// Check if row with requested 'camp_id' and 'lang' exists.
		// If not - we use campaign's default language.
		translation, translationDoesNotExist := TranslationDoesNotExist(db, campaign.ID, client.Language)
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
		payload.TemplateName = campaign.MailgunTemplate

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
		statExists, statID := StatExists(db, &stats)
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
		responseMessage, responseID, responseError := SendMailGunMessageV4(domain, apiKey, &payload)
		if responseError != nil {
			WriteLog("ERROR", fmt.Sprintf("Message for <%s> was not sent. MailGun responded with error: %s", client.Email, responseError))
			stats.ErrorMessage = fmt.Sprint(responseError)
			stats.Success = false
		} else {
			WriteLog("INFO", fmt.Sprintf("Message for <%s>:%s was successfully sent with %s lang. MG message: %s MG id: %s", client.Email, client.Language, translation.Language, responseMessage, responseID))
			stats.ErrorMessage = ""
			stats.Success = true
		}
		db.Create(&stats)

		if num > 0 && num%5 == 0 {
			sleepIntvl := 20*time.Second + time.Duration(rand.Intn(10))*time.Second
			WriteLog("INFO", fmt.Sprintf("Will sleep for %v before next send ...", sleepIntvl))
			time.Sleep(sleepIntvl)
		}
	}
}

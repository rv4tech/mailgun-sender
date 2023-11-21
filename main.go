package main

import (
	"log"
	"os"
	"rv4-request/sender"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var domain string = os.Getenv("DOMAIN")
	var apiKey string = os.Getenv("APIKEY")

	sender.SendMailGunMessageV3(domain, apiKey, sender.MailGunParams{})
}

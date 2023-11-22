package fixtures

import (
	"fmt"
	"math/rand"
	"rv4-request/database"

	"gorm.io/gorm"
)

// Fill Campaign table with dummy data.
func DummyCampaignData(db *gorm.DB, quantity uint8) {
	sum := 1
	for sum < int(quantity) {
		c := database.Campaign{
			Name:       fmt.Sprintf("Campaign %v", sum),
			MgTemplate: fmt.Sprintf("Template %v", sum),
		}
		db.Create(&c)
		sum += 1
	}
}

// Fill Translation table with dummy data.
func DummyTranslationData(db *gorm.DB, quantity uint8, domain string) {
	languageCodes := [20]string{
		"es", "en", "ru", "pt", "jp", "ab", "bi", "fr", "de", "ha",
		"hz", "kv", "la", "ko", "ms", "gv", "li", "pi", "fa", "sv",
	}
	subject := [3]string{"why hello there", "itsame", "test subject"}

	sum := 1
	for sum < int(quantity) {
		c := database.Translation{
			CampID:  rand.Intn(int(quantity)),
			Lang:    languageCodes[rand.Intn(cap(languageCodes))],
			From:    fmt.Sprintf("test@%s", domain),
			Subject: subject[rand.Intn(cap(subject))],
		}
		db.Create(&c)
		sum += 1
	}
}

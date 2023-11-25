package database

import (
	"log"

	"gorm.io/gorm"
)

// Get single row from Campaign table.
func GetCampaignByName(db *gorm.DB, campaignName string) (*Campaign, error) {
	var campaign *Campaign
	query := db.Where("name = ?", campaignName).First(&campaign)
	if query.Error != nil {
		log.Fatal(query.Error)
	}
	return campaign, query.Error
}

// Get related translation from Translation table.
func GetTranslationByCampaignIDAndLanguage(db *gorm.DB, id uint, clientLanguage string) (*Translation, error) {
	var translation *Translation
	query := db.Where("camp_id = ?", id).Where("lang = ?", clientLanguage).Find(&translation)
	if query.Error != nil {
		log.Fatal(query.Error)
	}
	return translation, query.Error
}

// Batch creation of send stat entity.
func CreateBatchSendStatRecord(db *gorm.DB, params []*SendStat) ([]*SendStat, error) {
	query := db.Create(&params)
	if query.Error != nil {
		log.Fatal(query.Error)
	}
	return params
}

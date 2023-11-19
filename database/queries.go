package database

import (
	"gorm.io/gorm"
)

// Get single row from Campaign table.
func GetCampaignByName(db *gorm.DB, campaignName string) Campaign {
	var campaign Campaign
	db.Where("name = ?", campaignName).First(&campaign)
	return campaign
}

// Get related translations from Translation table.
func GetTranslationsByCampaignID(db *gorm.DB, id int) []Translation {
	var translations []Translation
	db.Where("camp_id = ?", id).Find(&translations)
	return translations
}

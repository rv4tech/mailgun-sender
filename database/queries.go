package database

import (
	"gorm.io/gorm"
)

// Get single row from Campaign table.
func GetCampaignByName(db *gorm.DB, campaignName string) (*Campaign, error) {
	var campaign *Campaign
	query := db.
		Where("name = ?", campaignName).
		First(&campaign)
	return campaign, query.Error
}

// Get related translation from Translation table.
func GetTranslationByCampaignIDAndLanguage(db *gorm.DB, campID uint, clientLanguage string) (*Translation, error) {
	var translation *Translation
	query := db.
		Where("camp_id = ?", campID).
		Where("lang = ?", clientLanguage).
		Find(&translation)
	return translation, query.Error
}

// Batch creation of send stat entity.
func CreateBatchSendStatRecord(db *gorm.DB, params []*SendStat) error {
	query := db.Create(&params)
	return query.Error
}

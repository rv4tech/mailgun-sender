package database

import (
	"gorm.io/gorm"
)

// Get a single row from 'campaigns' table.
func GetCampaignByName(db *gorm.DB, campaignName string) (*Campaign, error) {
	var campaign *Campaign
	query := db.Where("name = ?", campaignName).First(&campaign)
	return campaign, query.Error
}

// Get related translation from 'translations' table.
func GetTranslationByCampaignIDAndLanguage(db *gorm.DB, campID uint, clientLanguage string) (*Translation, error) {
	var translation *Translation
	query := db.Where("camp_id = ?", campID).Where("lang = ?", clientLanguage).First(&translation)
	return translation, query.Error
}

// Get a single row from 'send_stats' table.
func GetSendStatEntry(db *gorm.DB, campaignID uint, email string) (*SendStat, error) {
	var stat *SendStat
	query := db.Where("camp_id = ?", campaignID).Where("email = ?", email).First(&stat)
	return stat, query.Error
}

// Checks if stat exists. Returns a boolean and an id of said stat.
func StatExists(db *gorm.DB, stat *SendStat) (bool, uint) {
	var entry SendStat
	query := db.Where("camp_id = ? AND email = ? AND success = 1", stat.CampaignID, stat.Email, stat.Success).First(&entry)
	return query.RowsAffected > 0, entry.ID
}

// Checks if translation exists. Returns entry itself and a boolean.
func TranslationDoesNotExist(db *gorm.DB, campaignID uint, language string) (*Translation, bool) {
	var entry Translation
	query := db.Where("camp_id = ? AND lang = ?", campaignID, language).First(&entry)
	return &entry, query.RowsAffected == 0
}

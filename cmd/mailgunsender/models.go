package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Struct representation for Campaign table. Gorm alias: campaigns.
type Campaign struct {
	ID              uint   `gorm:"primaryKey;autoIncrement"`
	Name            string `gorm:"unique;not null"`
	MailgunTemplate string `gorm:"column:mg_template"`
	DefaultLanguage string `gorm:"column:default_lang;default:'en'"`
}

// Struct representation for Translation table. Gorm alias: translations.
type Translation struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	CampaignID uint   `gorm:"column:camp_id;uniqueIndex:campaign_id_name_unique"`
	Language   string `gorm:"column:lang;uniqueIndex:campaign_id_name_unique;default:'en'"`
	From       string
	Subject    string
}

// Struct representation for SendStat table. Gorm alias: send_stats.
type SendStat struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	TimeStamp    int64  `gorm:"autoCreateTime;column:ts"` // Might want to change to uint64?
	CampaignID   uint   `gorm:"column:camp_id"`
	Language     string `gorm:"column:lang"`
	Email        string
	Name         string
	ExternalID   string `gorm:"column:ext_id"`
	Success      bool   `gorm:"default:0"`
	ErrorMessage string `gorm:"column:error_msg"`
}

const defaultDatabaseName = "mailgun_sender.db"

// Database initializer.
// Takes name string as an argument.
// Default value is mailgun_sender.db.
func InitDataBase(databaseName string) *gorm.DB {
	var database *gorm.DB
	// In case empty string was passed, use default hardcoded value.
	if databaseName == "" {
		databaseName = defaultDatabaseName
	}
	database, err := gorm.Open(sqlite.Open(databaseName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Could not conneсt to database %s\nError: %s\n", databaseName, err)
	}
	// Applying schemas does not require any kind of checks.
	ApplySchemas(database)
	fmt.Printf("Successfully connected to database %s\n\n", databaseName)
	return database
}

// Apply schemas from models.go file.
// Schemas structs are hardcoded inside of a function.
// Schemas are Campaign, Translation, SendStat
// GORM will not apply schemas if they already injected.
func ApplySchemas(db *gorm.DB) {
	var allModels = []any{Campaign{}, Translation{}, SendStat{}}

	for _, i := range allModels {
		log.Printf("interface type: %T, interface value: %+v", i, i)
		db.AutoMigrate(i)
	}
}

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

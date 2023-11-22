package database

import (
	"gorm.io/gorm"
)

// Struct representation for Campaign table. Gorm alias: campaigns.
type Campaign struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	MgTemplate  string
	DefaultLang string `gorm:"default:'en'"`
}

// Struct representation for Translation table. Gorm alias: translations.
type Translation struct {
	gorm.Model
	CampID  int    `gorm:"uniqueIndex:campaign_id_name_unique"`
	Lang    string `gorm:"uniqueIndex:campaign_id_name_unique"`
	From    string
	Subject string
}

// Struct representation for SendStat table. Gorm alias: send_stats.
type SendStat struct {
	gorm.Model
	Ts         int64 `gorm:"autoCreateTime"`
	CampaignID uint
	Lang       string
	Email      string
	ExtID      string
	Success    bool
	ErrorMsg   string
}

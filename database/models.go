package database

import (
	"fmt"

	"gorm.io/gorm"
)

type DBModel interface {
	printModel()
}

func (c *Campaigns) printModel() {
	fmt.Println(c)
}

func (c *Translations) printModel() {
	fmt.Println(c)
}

func (c *SendStats) printModel() {
	fmt.Println(c)
}

// Struct representation for Campaigns table. Gorm alias: campaigns.
type Campaigns struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	MgTemplate  string
	DefaultLang string `gorm:"default:'en'"`
}

// Struct representation for Translations table. Gorm alias: translations.
type Translations struct {
	gorm.Model
	CampID  int    `gorm:"uniqueIndex:campaign_id_name_unique"`
	Lang    string `gorm:"uniqueIndex:campaign_id_name_unique"`
	From    string
	Subject string
}

// Struct representation for SendStats table. Gorm alias: send_stats.
type SendStats struct {
	gorm.Model
	Ts         int64 `gorm:"autoCreateTime"`
	CampaignID int
	Lang       string
	Email      string
	ExtID      string
	Success    bool
	ErrorMsg   string
}

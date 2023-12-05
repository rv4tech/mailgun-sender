package database

// Struct representation for Campaign table. Gorm alias: campaigns.
type Campaign struct {
	ID              uint   `gorm:"primaryKey"`
	Name            string `gorm:"unique;not null"`
	MailgunTemplate string `gorm:"column:mg_template"`
	DefaultLanguage string `gorm:"column:default_lang;default:'en'"`
}

// Struct representation for Translation table. Gorm alias: translations.
type Translation struct {
	ID         uint   `gorm:"primaryKey"`
	CampaignID uint   `gorm:"column:camp_id;uniqueIndex:campaign_id_name_unique"`
	Language   string `gorm:"column:lang;uniqueIndex:campaign_id_name_unique;default:'en'"`
	From       string
	Subject    string
}

// Struct representation for SendStat table. Gorm alias: send_stats.
type SendStat struct {
	ID           uint   `gorm:"primaryKey"`
	TimeStamp    int64  `gorm:"autoCreateTime;column:ts"` // Might want to change to uint64?
	CampaignID   uint   `gorm:"column:camp_id"`
	Language     string `gorm:"column:lang"`
	Email        string
	ExternalID   string `gorm:"column:ext_id"`
	Success      bool   `gorm:"default:0"`
	ErrorMessage string `gorm:"column:error_msg"`
}

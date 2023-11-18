package requests

import (
	"rv4-request/database"

	"gorm.io/gorm"
)

func ProcessOneEntity(campaignName string, db *gorm.DB) *database.Campaigns {
	var campaign database.Campaigns
	var translations database.Translations
	var sendStats database.SendStats
	selectedCampaign := campaign.GetOne(db, "name", campaignName)
	getLanguage := translations.GetOne(db, "camp_id", &selectedCampaign.ID)
	return &campaign
}

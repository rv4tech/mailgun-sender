package database

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Insert one row into Translations table.
func InsertOneTranslationsRow(database *gorm.DB, data *Translations) int64 {
	query := database.Create(&data)
	CheckForQueryError(query)
	return query.RowsAffected
}

// Insert many rows into Translations table.
func InsertManyTranslationsRows(database *gorm.DB, data []*Translations) int64 {
	query := database.Create(data)
	CheckForQueryError(query)
	return query.RowsAffected
}

// Delete one row from Translations table.
func DeleteTranslationsRow(database *gorm.DB, data *Translations) int64 {
	query := database.Delete(&data)
	CheckForQueryError(query)
	return query.RowsAffected
}

// Hard(!) delete of Translations table row by condition.
func HardDeleteTranslationsRowByCondition(database *gorm.DB, field string, condition string) *gorm.DB {
	query := database.
		Unscoped().
		Clauses(clause.Returning{}).
		Where(fmt.Sprintf("%s = ?", field), condition).
		Delete(&Campaigns{})
	return query
}

// Insert one row into Campaigns table.
func InsertOneCampaignsRow(database *gorm.DB, data *Campaigns) int64 {
	query := database.Create(&data)
	CheckForQueryError(query)
	return query.RowsAffected
}

// Insert many rows into Campaigns table.
func InsertManyCampaignsRows(database *gorm.DB, data []*Campaigns) int64 {
	query := database.Create(data)
	CheckForQueryError(query)
	return query.RowsAffected
}

// Delete one row from Campaigns table.
func DeleteCampaignsRow(database *gorm.DB, data *Campaigns) int64 {
	query := database.Delete(&data)
	CheckForQueryError(query)
	return query.RowsAffected
}

// Hard(!) delete of Campaigns table row by condition.
func HardDeleteCampaignsRowByCondition(database *gorm.DB, field string, condition string) *gorm.DB {
	query := database.
		Unscoped().
		Clauses(clause.Returning{}).
		Where(fmt.Sprintf("%s = ?", field), condition).
		Delete(&Campaigns{})
	return query
}

// Insert one row into SendStats table.
func InsertOneSendStatsRow(database *gorm.DB, data *SendStats) int64 {
	query := database.Create(&data)
	CheckForQueryError(query)
	return query.RowsAffected
}

// Insert many rows into SendStats table.
func InsertManySendStatsRows(database *gorm.DB, data []*SendStats) int64 {
	query := database.Create(data)
	CheckForQueryError(query)
	return query.RowsAffected
}

// Delete one row from SendStats table.
func DeleteSendStatsRow(database *gorm.DB, data *SendStats) int64 {
	query := database.Delete(&data)
	CheckForQueryError(query)
	return query.RowsAffected
}

// Hard(!) delete of SendStats table row by condition.
func HardDeleteSendStatsRowByCondition(database *gorm.DB, field string, condition string) *gorm.DB {
	query := database.
		Unscoped().
		Clauses(clause.Returning{}).
		Where(fmt.Sprintf("%s = ?", field), condition).
		Delete(&SendStats{})
	return query
}

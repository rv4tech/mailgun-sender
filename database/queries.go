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
		Delete(&Translations{})
	return query
}

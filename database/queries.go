package database

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// type query interface {
// 	InsertOne()
// 	DeleteOne()
// 	HardDelete()
// }

// Insert one row into Translations table.
func (t *Translations) InsertOne(database *gorm.DB) int64 {
	query := database.Create(&t)
	CheckForQueryError(query)
	return query.RowsAffected
}

// Delete one row from Translations table.
func (t *Translations) DeleteOne(database *gorm.DB) int64 {
	query := database.Delete(&t)
	CheckForQueryError(query)
	return query.RowsAffected
}

// Hard(!) delete of Translations table row by condition.
func (t *Translations) HardDelete(database *gorm.DB, field string, condition string) *gorm.DB {
	query := database.
		Unscoped().
		Clauses(clause.Returning{}).
		Where(fmt.Sprintf("%s = ?", field), condition).
		Delete(&t)
	return query
}

// Insert one row into Campaigns table.
func (c *Campaigns) InsertOne(database *gorm.DB) int64 {
	query := database.Create(&c)
	CheckForQueryError(query)
	return query.RowsAffected
}

// Delete one row from Campaigns table.
func (c *Campaigns) DeleteOne(database *gorm.DB) int64 {
	query := database.Delete(&c)
	CheckForQueryError(query)
	return query.RowsAffected
}

// Hard(!) delete of Campaigns table row by condition.
func (c *Campaigns) HardDelete(database *gorm.DB, field string, condition string) *gorm.DB {
	query := database.
		Unscoped().
		Clauses(clause.Returning{}).
		Where(fmt.Sprintf("%s = ?", field), condition).
		Delete(&c)
	return query
}

// Insert one row into SendStats table.
func (s *SendStats) InsertOne(database *gorm.DB) int64 {
	query := database.Create(&s)
	CheckForQueryError(query)
	return query.RowsAffected
}

// Delete one row from SendStats table.
func (s *SendStats) DeleteOne(database *gorm.DB) int64 {
	query := database.Delete(&s)
	CheckForQueryError(query)
	return query.RowsAffected
}

// Hard(!) delete of SendStats table row by condition.
func (s *SendStats) HardDelete(database *gorm.DB, field string, condition string) *gorm.DB {
	query := database.
		Unscoped().
		Clauses(clause.Returning{}).
		Where(fmt.Sprintf("%s = ?", field), condition).
		Delete(&s)
	return query
}

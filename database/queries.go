package database

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

// type query interface {
// 	InsertOne()
// 	DeleteOne()
// 	HardDelete()
// }

// Get single row from Translations table.
func (t *Translations) GetOne(database *gorm.DB, field string, condition string) *Translations {
	query := database.
		Where(fmt.Sprintf("%s = ?", field), condition).
		First(&t)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Fatalf("Could not find entity with condition: %v, field: %v", condition, field)
	}
	return t
}

// Insert one row into Translations table.
func (t *Translations) InsertOne(database *gorm.DB) *Translations {
	query := database.Create(&t)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Fatalf("Could not find entity in Translations table")
	}
	return t
}

// Delete one row from Translations table.
func (t *Translations) DeleteOne(database *gorm.DB) *Translations {
	query := database.Delete(&t)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Fatalf("Could not find entity in Translations table")
	}
	return t
}

// Hard(!) delete of Translations table row by condition.
func (t *Translations) HardDelete(database *gorm.DB, field string, condition string) *Translations {
	query := database.
		Where(fmt.Sprintf("%s = ?", field), condition).
		Delete(&t)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Fatalf("Could not find entity with condition: %v, field: %v", condition, field)
	}
	return t
}

// Get single row from campaigns table.
func (c *Campaigns) GetOne(database *gorm.DB, field string, condition string) *gorm.DB {
	query := database.
		Where(fmt.Sprintf("%s = ?", field), condition).
		First(&c)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Fatalf("Could not find entity with condition: %v, field: %v", condition, field)
	}
	return query
}

// Insert one row into Campaigns table.
func (c *Campaigns) InsertOne(database *gorm.DB) *Campaigns {
	query := database.Create(&c)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Fatalf("Could not find entity in Campaigns table")
	}
	return c
}

// Delete one row from Campaigns table.
func (c *Campaigns) DeleteOne(database *gorm.DB) *Campaigns {
	query := database.Delete(&c)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Fatal("Could not find entity in Campaigns table")
	}
	return c
}

// Hard(!) delete of Campaigns table row by condition.
func (c *Campaigns) HardDelete(database *gorm.DB, field string, condition string) *Campaigns {
	query := database.
		Unscoped().
		Where(fmt.Sprintf("%s = ?", field), condition).
		Delete(&c)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Fatalf("Could not find entity with condition: %v, field: %v", condition, field)
	}
	return c
}

// Get single row from SendStats table.
func (s *SendStats) GetOne(database *gorm.DB, field string, condition string) *SendStats {
	query := database.
		Where(fmt.Sprintf("%s = ?", field), condition).
		First(&s)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Fatalf("Could not find entity with condition: %v, field: %v", condition, field)
	}
	return s
}

// Insert one row into SendStats table.
func (s *SendStats) InsertOne(database *gorm.DB) *SendStats {
	query := database.Create(&s)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Fatal("Could not find entity in Campaigns table")
	}
	return s
}

// Delete one row from SendStats table.
func (s *SendStats) DeleteOne(database *gorm.DB) *SendStats {
	query := database.Delete(&s)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Fatal("Could not find entity in Campaigns table")
	}
	return s
}

// Hard(!) delete of SendStats table row by condition.
func (s *SendStats) HardDelete(database *gorm.DB, field string, condition string) *SendStats {
	query := database.
		Unscoped().
		Where(fmt.Sprintf("%s = ?", field), condition).
		Delete(&s)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Fatalf("Could not find entity with condition: %v, field: %v", condition, field)
	}
	return s
}

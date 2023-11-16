package database

import "gorm.io/gorm"

type Campaigns struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	MgTemplate  string
	DefaultLang string `gorm:"default:'en'"`
}

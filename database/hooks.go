package database

import (
	"fmt"
	"log"
	"net/mail"

	"gorm.io/gorm"
)

// BeforeCreate hook to check if `StatsSend: Email` field is valid.
func (s *SendStat) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = mail.ParseAddress(s.Email)
	if err != nil {
		log.Fatalf(fmt.Sprintf("\nEmail is not valid.\nError: %s\nEmail: %s", err, s.Email))
	}
	return
}

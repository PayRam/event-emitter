package models

import (
	"gorm.io/gorm"
)

type EEEvent struct {
	gorm.Model // Embedded GORM model. Pointer not needed here.
	EventName  string
	ProfileID  string
	Attribute  string // JSON data as string
}

type Migration struct {
	ID       string
	Migrate  func(db *gorm.DB) error
	Rollback func(db *gorm.DB) error
}

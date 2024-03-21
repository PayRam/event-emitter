package model

import "gorm.io/gorm"

type Migration struct {
	ID       string
	Migrate  func(db *gorm.DB) error
	Rollback func(db *gorm.DB) error
}

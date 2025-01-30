package migration

import (
	"github.com/PayRam/event-emitter/internal/model"
	"github.com/PayRam/event-emitter/service/param"
	"gorm.io/gorm"
)

var AddValidUntil = model.Migration{
	ID: "202501301059-ee1234",
	Migrate: func(db *gorm.DB) error {
		return db.AutoMigrate(&param.EEEvent{})
	},
	Rollback: func(db *gorm.DB) error {
		return db.Migrator().DropColumn(&param.EEEvent{}, "valid_until")
	},
}

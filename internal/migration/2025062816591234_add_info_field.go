package migration

import (
	"github.com/PayRam/event-emitter/internal/model"
	"github.com/PayRam/event-emitter/service/param"
	"gorm.io/gorm"
)

var AddInfo = model.Migration{
	ID: "202506281659-ee1234",
	Migrate: func(db *gorm.DB) error {
		return db.AutoMigrate(&param.EEEvent{})
	},
	Rollback: func(db *gorm.DB) error {
		return db.Migrator().DropColumn(&param.EEEvent{}, "info")
	},
}

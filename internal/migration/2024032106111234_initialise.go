package migration

import (
	"github.com/PayRam/event-emitter/internal/model"
	"github.com/PayRam/event-emitter/service/param"
	"gorm.io/gorm"
)

var Initialise = model.Migration{
	ID: "202403210611-ee1234",
	Migrate: func(db *gorm.DB) error {
		return db.AutoMigrate(&param.EEEvent{})
	},
	Rollback: func(db *gorm.DB) error {
		return db.Migrator().DropTable(&param.EEEvent{})
	},
}

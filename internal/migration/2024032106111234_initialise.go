package migration

import (
	db2 "github.com/PayRam/event-emitter/internal/models"
	"gorm.io/gorm"
)

var Initialise = db2.Migration{
	ID: "2024032106111234",
	Migrate: func(db *gorm.DB) error {
		return db.AutoMigrate(&db2.EEEvent{})
	},
	Rollback: func(db *gorm.DB) error {
		return db.Migrator().DropTable(&db2.EEEvent{})
	},
}

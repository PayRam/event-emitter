package service

import (
	db2 "github.com/PayRam/event-emitter/internal/db"
	"github.com/PayRam/event-emitter/internal/serviceimpl"
	"github.com/PayRam/event-emitter/service/param"
	"gorm.io/gorm"
)

func NewEventServiceWithDB(db *gorm.DB) param.EventService {
	return serviceimpl.NewEventServiceWithDB(db2.Migrate(db))
}

func NewEventService(dbPath string) param.EventService {
	return serviceimpl.NewEventServiceWithDB(db2.InitDB(dbPath))
}

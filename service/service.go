package service

import (
	"github.com/PayRam/event-emitter/internal/db"
	"github.com/PayRam/event-emitter/internal/serviceimpl"
	"github.com/PayRam/event-emitter/service/param"
	"gorm.io/gorm"
)

func NewEventServiceWithDB(db *gorm.DB) param.EventService {
	// This assumes you have adjusted the visibility of serviceimpl or provided a way to access it from here.
	return serviceimpl.NewEventServiceWithDB(db)
}

func NewEventService(dbPath string) param.EventService {
	db := db.InitDB(dbPath)
	return serviceimpl.NewEventServiceWithDB(db)
}

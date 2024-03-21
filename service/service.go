package service

import (
	"github.com/PayRam/event-emitter/internal/db"
	_interface "github.com/PayRam/event-emitter/internal/interface"
	"github.com/PayRam/event-emitter/internal/serviceimpl"
	"gorm.io/gorm"
)

func NewEventServiceWithDB(db *gorm.DB) _interface.EventService {
	// This assumes you have adjusted the visibility of serviceimpl or provided a way to access it from here.
	return serviceimpl.NewEventServiceWithDB(db)
}

func NewEventService(dbPath string) _interface.EventService {
	db := db.InitDB(dbPath)
	return serviceimpl.NewEventServiceWithDB(db)
}

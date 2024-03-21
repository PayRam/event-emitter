package service

import (
	"github.com/PayRam/event-emitter/internal/db"
	"github.com/PayRam/event-emitter/internal/models"
	event_emitter "github.com/PayRam/event-emitter/pkg/event-emitter"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

func NewEventServiceWithDB(db *gorm.DB) event_emitter.EventService {
	return &service{db: db}
}

func NewEventService(dbPath string) event_emitter.EventService {
	db := db.InitDB(dbPath)
	return &service{db: db}
}

// CreateEvent adds a new event to the database.
func (s *service) CreateEvent(event models.EEEvent) error {
	result := s.db.Create(&event)
	return result.Error
}

// QueryEvents retrieves events based on the provided QuerySpec.
func (s *service) QueryEvents(query event_emitter.QuerySpec) ([]models.EEEvent, error) {
	tx := s.db.Model(&models.EEEvent{})

	if query.ID != nil {
		tx = tx.Where("id = ?", *query.ID)
	}
	if query.EventName != nil {
		tx = tx.Where("event_name = ?", *query.EventName)
	}
	if query.ProfileID != nil {
		tx = tx.Where("profile_id = ?", *query.ProfileID)
	}
	if query.CreatedAt != nil {
		tx = tx.Where("created_at = ?", *query.CreatedAt)
	}

	// Example of handling JSON attribute query; adjust based on actual needs
	if len(query.Attributes) > 0 {
		for key, value := range query.Attributes {
			// Here you need to construct the correct SQL for JSON querying depending on your schema and requirements
			// This is a simplistic example; actual implementation may vary
			tx = tx.Where("json_extract(attribute, ?) = ?", "$."+key, value)
		}
	}

	var events []models.EEEvent
	result := tx.Find(&events)
	return events, result.Error
}

package event_emitter

import (
	"github.com/PayRam/event-emitter/internal/models"
	"time"
)

type EventService interface {
	CreateEvent(event models.EEEvent) error
	QueryEvents(query QuerySpec) ([]models.EEEvent, error)
}

type QuerySpec struct {
	ID         *uint // Use pointers to distinguish between zero-value and non-provided
	EventName  *string
	ProfileID  *string
	CreatedAt  *time.Time
	Attributes map[string]interface{} // For querying JSON attributes dynamically
}

package event_emitter

import (
	"github.com/PayRam/event-emitter/internal/db"
	"time"
)

type EventService interface {
	CreateEvent(event db.Event) error
	QueryEvents(query QuerySpec) ([]db.Event, error)
}

type QuerySpec struct {
	ID         *uint // Use pointers to distinguish between zero-value and non-provided
	EventName  *string
	ProfileID  *string
	CreatedAt  *time.Time
	Attributes map[string]interface{} // For querying JSON attributes dynamically
}

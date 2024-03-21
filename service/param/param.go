package param

import (
	"gorm.io/gorm"
	"time"
)

type EventService interface {
	CreateEvent(event EEEvent) error
	QueryEvents(query QuerySpec) ([]EEEvent, error)
}

type EEEvent struct {
	gorm.Model // Embedded GORM model. Pointer not needed here.
	EventName  string
	ProfileID  string
	Attribute  string // JSON data as string
}

type QuerySpec struct {
	ID         *uint // Use pointers to distinguish between zero-value and non-provided
	EventName  *string
	ProfileID  *string
	CreatedAt  *time.Time
	Attributes map[string]interface{} // For querying JSON attributes dynamically
}

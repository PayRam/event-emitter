package param

import (
	"gorm.io/gorm"
	"time"
)

type EventService interface {
	CreateEvent(eventName string, jsonData string, profileID *string) (*EEEvent, error)
	CreateTimedEvent(eventName string, jsonData string, profileID *string, validUntil time.Time) (*EEEvent, error)
	CreateSimpleEvent(eventName string, jsonData string) (*EEEvent, error)
	QueryEvents(query QueryBuilder) ([]EEEvent, error)
}

type EEEvent struct {
	gorm.Model            // Embedded GORM model. Pointer not needed here.
	EventName  string     `gorm:"type:varchar(200);not null;index"`
	ProfileID  *string    `gorm:"type:varchar(200);index"`
	Attribute  string     `gorm:"type:text"`
	ValidUntil *time.Time `gorm:"index"` // Nullable ValidUntil field
}

type JoinClause struct {
	Exclude bool
	Clause  string
}

type QueryBuilder struct {
	ID                     []uint // Use pointers to distinguish between zero-value and non-provided
	EventNames             []string
	ProfileIDs             []string
	CreatedAtBefore        *time.Time
	CreatedAtAfter         *time.Time
	CreatedAtRelativeStart *time.Duration         // Duration relative to now for the start of the time range
	CreatedAtRelativeEnd   *time.Duration         // Duration relative to now for the end of the time range
	Attributes             map[string]interface{} // For querying JSON attributes dynamically
	JoinWhereClause        map[string]JoinClause  // For querying JSON attributes dynamically
	SubQueryBuilder        *QueryBuilder
}

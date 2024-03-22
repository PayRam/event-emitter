package param

import (
	"gorm.io/gorm"
	"time"
)

type EventService interface {
	CreateEvent(eventName string, profileId string, jsonData string) error
	CreateGenericEvent(eventName string, jsonData string) error
	QueryEvents(query QueryBuilder) ([]EEEvent, error)
}

type EEEvent struct {
	gorm.Model         // Embedded GORM model. Pointer not needed here.
	EventName  string  `gorm:"type:varchar(200);not null;index"`
	ProfileID  *string `gorm:"type:varchar(200);index"`
	Attribute  string  `gorm:"type:text"`
}

type JoinClause struct {
	Exclude bool
	Clause  string
}

type QueryBuilder struct {
	ID                []uint // Use pointers to distinguish between zero-value and non-provided
	EventName         []string
	ProfileID         []string
	CreatedAtBefore   *time.Time
	CreatedAtAfter    *time.Time
	Attributes        map[string]interface{} // For querying JSON attributes dynamically
	JoinWhereClause   map[string]JoinClause  // For querying JSON attributes dynamically
	QueryBuilderParam *QueryBuilder
}

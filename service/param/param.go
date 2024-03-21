package param

import "time"

type QuerySpec struct {
	ID         *uint // Use pointers to distinguish between zero-value and non-provided
	EventName  *string
	ProfileID  *string
	CreatedAt  *time.Time
	Attributes map[string]interface{} // For querying JSON attributes dynamically
}

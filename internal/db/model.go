package db

import "time"

type (
	Event struct {
		ID        uint `gorm:"primaryKey"`
		EventName string
		ProfileID string
		CreatedAt time.Time
		Attribute string // JSON data as string
	}
)

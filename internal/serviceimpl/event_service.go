package serviceimpl

import (
	"errors"
	"fmt"
	"github.com/PayRam/event-emitter/internal/db"
	"github.com/PayRam/event-emitter/service/param"
	"gorm.io/gorm"
	"strings"
	"time"
)

type service struct {
	db *gorm.DB
}

func NewEventServiceWithDB(db *gorm.DB) param.EventService {
	return &service{db: db}
}

func NewEventService(dbPath string) param.EventService {
	db := db.InitDB(dbPath)
	return &service{db: db}
}

// CreateEvent adds a new event to the database with optional ValidUntil.
func (s *service) CreateEvent(eventName string, jsonData string, profileID *string) (*param.EEEvent, error) {

	event := &param.EEEvent{
		EventName: eventName,
		ProfileID: profileID,
		Attribute: jsonData,
	}

	result := s.db.Create(event)
	if result.Error != nil {
		return nil, result.Error
	}

	return event, nil
}

// CreateEventWithInfo adds a new event to the database with additional info.
func (s *service) CreateEventWithInfo(eventName string, jsonData string, info string, profileID *string) (*param.EEEvent, error) {
	info = strings.TrimSpace(info)
	if info == "" {
		return nil, errors.New("info cannot be empty")
	}

	event := &param.EEEvent{
		EventName: eventName,
		ProfileID: profileID,
		Attribute: jsonData,
		Info:      &info,
	}

	result := s.db.Create(event)
	if result.Error != nil {
		return nil, result.Error
	}

	return event, nil
}

// CreateTimedEvent adds a new event to the database with a ValidUntil field.
func (s *service) CreateTimedEvent(eventName string, jsonData string, profileID *string, validUntil time.Time) (*param.EEEvent, error) {
	// Ensure ValidUntil is not in the past
	if validUntil.Before(time.Now()) {
		return nil, errors.New("validUntil cannot be set to a past time")
	}

	event := &param.EEEvent{
		EventName:  eventName,
		ProfileID:  profileID,
		Attribute:  jsonData,
		ValidUntil: &validUntil,
	}

	result := s.db.Create(event)
	if result.Error != nil {
		return nil, result.Error
	}

	return event, nil
}

// CreateTimedEventWithInfo adds a new event to the database with additional info and a ValidUntil field.
func (s *service) CreateTimedEventWithInfo(eventName string, jsonData string, info string, profileID *string, validUntil time.Time) (*param.EEEvent, error) {
	info = strings.TrimSpace(info)
	if info == "" {
		return nil, errors.New("info cannot be empty")
	}
	// Ensure ValidUntil is not in the past
	if validUntil.Before(time.Now()) {
		return nil, errors.New("validUntil cannot be set to a past time")
	}

	event := &param.EEEvent{
		EventName:  eventName,
		ProfileID:  profileID,
		Attribute:  jsonData,
		ValidUntil: &validUntil,
		Info:       &info,
	}

	result := s.db.Create(event)
	if result.Error != nil {
		return nil, result.Error
	}

	return event, nil
}

// CreateSimpleEvent adds a new event to the database which does not have a profile ID.
func (s *service) CreateSimpleEvent(eventName string, jsonData string) (*param.EEEvent, error) {
	event := &param.EEEvent{
		EventName: eventName,
		Attribute: jsonData,
	}

	result := s.db.Create(event)
	if result.Error != nil {
		return nil, result.Error
	}

	return event, nil
}

// CreateSimpleEventWithInfo adds a new event to the database with additional info and no profile ID.
func (s *service) CreateSimpleEventWithInfo(eventName string, jsonData string, info string) (*param.EEEvent, error) {
	info = strings.TrimSpace(info)
	if info == "" {
		return nil, errors.New("info cannot be empty")
	}
	event := &param.EEEvent{
		EventName: eventName,
		Attribute: jsonData,
		Info:      &info,
	}

	result := s.db.Create(event)
	if result.Error != nil {
		return nil, result.Error
	}

	return event, nil
}

func (s *service) QueryEvents(query param.QueryBuilder) ([]param.EEEvent, error) {
	db, err := s.queryEventsRecurse(query)
	if err != nil {
		return nil, err
	}

	var events []param.EEEvent
	if err := db.Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}

func (s *service) queryEventsRecurse(queryBuilder param.QueryBuilder) (*gorm.DB, error) {
	var errRec error
	subQuery := s.db.Model(&param.EEEvent{}) // Initialize subQuery at each recursion level

	// Recurse if there's a nested QueryBuilder
	if queryBuilder.SubQueryBuilder != nil {
		var nestedSubQuery *gorm.DB
		nestedSubQuery, errRec = s.queryEventsRecurse(*queryBuilder.SubQueryBuilder)
		if errRec != nil {
			return nil, errRec
		}

		for key, value := range queryBuilder.JoinWhereClause {
			if value.Exclude {
				subQuery = subQuery.Not(key+" IN (?)", nestedSubQuery.Select(value.Clause))
			} else {
				subQuery = subQuery.Or(key+" IN (?)", nestedSubQuery.Select(value.Clause))
			}
		}
	}

	if len(queryBuilder.EventNames) > 0 {
		subQuery = subQuery.Where("event_name IN ?", queryBuilder.EventNames)
	}

	if len(queryBuilder.ProfileIDs) > 0 {
		subQuery = subQuery.Where("profile_id IN ?", queryBuilder.ProfileIDs)
	}

	if queryBuilder.CreatedAtBefore != nil {
		subQuery = subQuery.Where("created_at < ?", queryBuilder.CreatedAtBefore)
	}
	if queryBuilder.CreatedAtAfter != nil {
		subQuery = subQuery.Where("created_at > ?", queryBuilder.CreatedAtAfter)
	}

	now := time.Now()
	// Apply dynamic time range based on relative start and end durations
	if queryBuilder.CreatedAtRelativeStart != nil {
		start := now.Add(*queryBuilder.CreatedAtRelativeStart)
		subQuery = subQuery.Where("created_at >= ?", start)
	}
	if queryBuilder.CreatedAtRelativeEnd != nil {
		end := now.Add(*queryBuilder.CreatedAtRelativeEnd)
		subQuery = subQuery.Where("created_at <= ?", end)
	}

	for key, value := range queryBuilder.Attributes {
		// 🟢 PostgreSQL JSONB syntax
		jsonQuery := fmt.Sprintf("attribute::jsonb ->> '%s' = ?", key)
		subQuery = subQuery.Where(jsonQuery, value)
	}

	subQuery = subQuery.Where("valid_until IS NULL OR valid_until >= ?", now)

	subQuery = subQuery.Order("created_at ASC")

	return subQuery, nil
}

package serviceimpl

import (
	"fmt"
	"github.com/PayRam/event-emitter/internal/db"
	"github.com/PayRam/event-emitter/service/param"
	"gorm.io/gorm"
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

// CreateEvent adds a new event to the database.
func (s *service) CreateEvent(eventName string, profileId string, jsonData string) error {
	result := s.db.Create(&param.EEEvent{
		EventName: eventName,
		ProfileID: &profileId,
		Attribute: jsonData,
	})
	return result.Error
}

// CreateGenericEvent adds a new event to the database which does not have profile id
func (s *service) CreateGenericEvent(eventName string, jsonData string) error {
	result := s.db.Create(&param.EEEvent{
		EventName: eventName,
		Attribute: jsonData,
	})
	return result.Error
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
	if queryBuilder.QueryBuilderParam != nil {
		var nestedSubQuery *gorm.DB
		nestedSubQuery, errRec = s.queryEventsRecurse(*queryBuilder.QueryBuilderParam)
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

	if len(queryBuilder.EventName) > 0 {
		subQuery = subQuery.Where("event_name IN ?", queryBuilder.EventName)
	}

	if len(queryBuilder.ProfileID) > 0 {
		subQuery = subQuery.Where("profile_id IN ?", queryBuilder.ProfileID)
	}

	if queryBuilder.CreatedAtBefore != nil {
		subQuery = subQuery.Where("created_at < ?", queryBuilder.CreatedAtBefore)
	}
	if queryBuilder.CreatedAtAfter != nil {
		subQuery = subQuery.Where("created_at > ?", queryBuilder.CreatedAtAfter)
	}

	for key, value := range queryBuilder.Attributes {
		jsonQuery := fmt.Sprintf("json_extract(attribute, '$.%s') = ?", key)
		subQuery = subQuery.Where(jsonQuery, value)
	}

	return subQuery, nil
}

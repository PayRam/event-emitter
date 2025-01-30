package serviceimpl_test

import (
	service3 "github.com/PayRam/event-emitter/service"
	"github.com/PayRam/event-emitter/service/param"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

var (
	db           *gorm.DB
	eventService param.EventService
)

func TestMain(m *testing.M) {
	// Initialize shared test database
	var err error
	db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	//db, err = gorm.Open(sqlite.Open("/Users/sameer/Documents/test1.db"), &gorm.Config{})
	if err != nil {
		panic("failed to initialize test database")
	}

	eventService = service3.NewEventServiceWithDB(db)

	m.Run()
}

func createSimpleEvent(t *testing.T, eventName, jsonData string) *param.EEEvent {

	event, err := eventService.CreateSimpleEvent(eventName, jsonData)
	assert.NoError(t, err)
	assert.NotNil(t, event)
	assert.Greater(t, event.ID, uint(0), "Failed to create simple event")
	assert.Equal(t, event.EventName, eventName)
	assert.Equal(t, event.Attribute, jsonData)
	assert.Nil(t, event.ValidUntil)
	assert.Nil(t, event.ProfileID)

	return event
}

func createEvent(t *testing.T, eventName, jsonData string, profileID string) *param.EEEvent {
	event, err := eventService.CreateEvent(eventName, jsonData, &profileID)
	assert.NoError(t, err)
	assert.NotNil(t, event)
	assert.Greater(t, event.ID, uint(0), "Failed to create event")
	assert.Equal(t, event.EventName, eventName)
	assert.Equal(t, *event.ProfileID, profileID)
	assert.Equal(t, event.Attribute, jsonData)
	assert.Nil(t, event.ValidUntil)
	return event
}

func createTimedEvent(t *testing.T, eventName, jsonData string, profileID string, validUntil time.Time) *param.EEEvent {
	event, err := eventService.CreateTimedEvent(eventName, jsonData, &profileID, validUntil)
	assert.NoError(t, err)
	assert.NotNil(t, event)
	assert.Greater(t, event.ID, uint(0), "Failed to create timed event")
	assert.Equal(t, event.EventName, eventName)
	assert.Equal(t, *event.ProfileID, profileID)
	assert.Equal(t, event.Attribute, jsonData)
	assert.Equal(t, *event.ValidUntil, validUntil)
	return event
}

func queryDepositReceivedEvents(t *testing.T, expectedIds []uint) {
	subQuery := param.QueryBuilder{
		EventNames: []string{"deposit-received-email-sent"},
	}

	eNames := []string{"deposit-received"}

	joinWhereClause := make(map[string]param.JoinClause)
	joinClause := param.JoinClause{
		Clause:  "json_extract(attribute, '$.refId')",
		Exclude: true,
	}
	joinWhereClause["json_extract(attribute, '$.refId')"] = joinClause

	builder := param.QueryBuilder{
		EventNames:      eNames,
		JoinWhereClause: joinWhereClause,
		SubQueryBuilder: &subQuery,
	}

	queryEvents, err := eventService.QueryEvents(builder)
	if err != nil {
		return
	}
	assert.Equal(t, len(expectedIds), len(queryEvents))

	for i, event := range queryEvents {
		assert.Equal(t, expectedIds[i], event.ID)
	}
}

func queryTimedEvent(t *testing.T, expectedIds []uint) {
	subQuery := param.QueryBuilder{
		EventNames: []string{"timed-deposit-received-email-sent"},
	}

	eNames := []string{"timed-deposit-received"}

	joinWhereClause := make(map[string]param.JoinClause)
	joinClause := param.JoinClause{
		Clause:  "json_extract(attribute, '$.refId')",
		Exclude: true,
	}
	joinWhereClause["json_extract(attribute, '$.refId')"] = joinClause

	builder := param.QueryBuilder{
		EventNames:      eNames,
		JoinWhereClause: joinWhereClause,
		SubQueryBuilder: &subQuery,
	}

	queryEvents, err := eventService.QueryEvents(builder)
	if err != nil {
		return
	}

	assert.Equal(t, len(expectedIds), len(queryEvents))

	for i, event := range queryEvents {
		assert.Equal(t, expectedIds[i], event.ID)
	}
}

func TestSimpleEvent(t *testing.T) {
	createSimpleEvent(t, "Generic EEEvent", `{"key": "generic"}`)
	event1 := createEvent(t, "deposit-received", `{"refId": "123456"}`, "123")
	_ = createEvent(t, "deposit-received", `{"refId": "123457"}`, "323")
	event3 := createEvent(t, "deposit-received", `{"refId": "123458"}`, "123")
	_ = createEvent(t, "deposit-received", `{"refId": "123459"}`, "123")
	event5 := createEvent(t, "deposit-received", `{"refId": "123460"}`, "323")
	_ = createEvent(t, "deposit-received-email-sent", `{"refId": "123459"}`, "123")
	_ = createEvent(t, "deposit-received-email-sent", `{"refId": "123457"}`, "123")

	queryDepositReceivedEvents(t, []uint{event1.ID, event3.ID, event5.ID})

	event6 := createTimedEvent(t, "timed-deposit-received", `{"refId": "123456"}`, "123", time.Now().Add(time.Second*2))

	queryTimedEvent(t, []uint{event6.ID})
	log.Printf("checking timed events. please wait for 3 seconds")
	time.Sleep(time.Second * 3)
	queryTimedEvent(t, []uint{})

	event6 = createTimedEvent(t, "timed-deposit-received", `{"refId": "323456"}`, "123", time.Now().Add(time.Minute*10))
	event7 := createTimedEvent(t, "timed-deposit-received", `{"refId": "343456"}`, "123", time.Now().Add(time.Minute*10))
	queryTimedEvent(t, []uint{event6.ID, event7.ID})
	_ = createEvent(t, "timed-deposit-received-email-sent", `{"refId": "323456"}`, "123")

	queryTimedEvent(t, []uint{event7.ID})
}

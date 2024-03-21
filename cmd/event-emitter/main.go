package main

import (
	service3 "github.com/PayRam/event-emitter/service"
	"github.com/PayRam/event-emitter/service/param"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(sqlite.Open("/Users/sameer/payram/db/payram.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	service := service3.NewEventServiceWithDB(db)
	//service := service3.NewEventService("/Users/sameer/payram/db/payram.db")

	// Example usage
	err = service.CreateEvent(param.EEEvent{
		EventName: "Sample EEEvent",
		ProfileID: "123",
		Attribute: `{"key": "value"}`,
	})
	if err != nil {
		log.Printf("failed to create event: %v", err)
	}

	// Query example
	events, err := service.QueryEvents(param.QuerySpec{
		EventName: new(string),
	})
	//*events[0].EventName = "Sample EEEvent" // assuming you want to query by EventName
	if err != nil {
		log.Printf("failed to query events: %v", err)
	}

	log.Printf("Queried events: %+v", events)
}

package main

import (
	db2 "github.com/PayRam/event-emitter/internal/db"
	"github.com/PayRam/event-emitter/internal/models"
	service3 "github.com/PayRam/event-emitter/service"
	"github.com/PayRam/event-emitter/service/param"
	"log"
)

func main() {
	db := db2.InitDB("/Users/sameer/payram/db/payram.db") // Initialize the database connection

	// Example usage
	service := service3.NewEventServiceWithDB(db)
	err := service.CreateEvent(models.EEEvent{
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

package main

import (
	db2 "github.com/PayRam/event-emitter/internal/db"
	service2 "github.com/PayRam/event-emitter/internal/service"
	event_emitter "github.com/PayRam/event-emitter/pkg/event-emitter"
	"log"
	"time"
)

func main() {
	db := db2.InitDB() // Initialize the database connection

	// Example usage
	service := service2.NewEventService(db)
	err := service.CreateEvent(db2.Event{
		EventName: "Sample Event",
		ProfileID: "123",
		CreatedAt: time.Now(),
		Attribute: `{"key": "value"}`,
	})
	if err != nil {
		log.Printf("failed to create event: %v", err)
	}

	// Query example
	events, err := service.QueryEvents(event_emitter.QuerySpec{
		EventName: new(string),
	})
	//*events[0].EventName = "Sample Event" // assuming you want to query by EventName
	if err != nil {
		log.Printf("failed to query events: %v", err)
	}

	log.Printf("Queried events: %+v", events)
}

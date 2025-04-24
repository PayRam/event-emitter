package main

import (
	service3 "github.com/PayRam/event-emitter/service"
	"github.com/PayRam/event-emitter/service/param"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(sqlite.Open("your.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	service := service3.NewEventServiceWithDB(db)
	//service := service3.NewEventService("/Users/sameer/payram/db/payram.db")

	//example usage
	_, err = service.CreateSimpleEvent("Generic EEEvent", `{"key": "generic"}`)
	//err = service.CreateEvent("deposit-received", "123", `{"refId": "123456"}`)
	//err = service.CreateEvent("deposit-received", "323", `{"refId": "123457"}`)
	//err = service.CreateEvent("deposit-received", "123", `{"refId": "123458"}`)
	//err = service.CreateEvent("deposit-received", "123", `{"refId": "123459"}`)
	//err = service.CreateEvent("deposit-received", "323", `{"refId": "123460"}`)
	//err = service.CreateEvent("deposit-received-email-sent", "123", `{"refId": "123459"}`)

	subQuery := param.QueryBuilder{
		EventNames: []string{"deposit-received-email-sent"},
	}

	eNames := []string{"deposit-received"}

	joinWhereClause := make(map[string]param.JoinClause)
	joinClause := param.JoinClause{
		Clause:  `attribute::jsonb ->>'refId'`, // ✅ PostgreSQL jsonb syntax
		Exclude: true,
	}
	joinWhereClause["attribute::jsonb ->>'refId'"] = joinClause

	builder := param.QueryBuilder{
		EventNames:      eNames,
		JoinWhereClause: joinWhereClause,
		SubQueryBuilder: &subQuery,
	}

	queryEvents, err := service.QueryEvents(builder)
	if err != nil {
		return
	}

	log.Printf("-----------------------------------------")
	for _, event := range queryEvents {
		log.Printf("Event: %+v", event)
		log.Printf("-----------------------------------------")
	}
}

# event-emitter

## Usage

Import in your project
```
require github.com/PayRam/event-emitter v0.1.0
```
Run Jobs
```
//example usage
err = service.CreateGenericEvent("Generic EEEvent", `{"key": "generic"}`)
err = service.CreateEvent("deposit-received", "123", `{"refId": "123456"}`)
err = service.CreateEvent("deposit-received", "323", `{"refId": "123457"}`)
err = service.CreateEvent("deposit-received", "123", `{"refId": "123458"}`)
err = service.CreateEvent("deposit-received", "123", `{"refId": "123459"}`)
err = service.CreateEvent("deposit-received", "323", `{"refId": "123460"}`)
err = service.CreateEvent("deposit-received-email-sent", "123", `{"refId": "123459"}`)

subQuery := param.QueryBuilder{
    EventName: []string{"deposit-received-email-sent"},
}

eNames := []string{"deposit-received"}

joinWhereClause := make(map[string]param.JoinClause)
joinClause := param.JoinClause{
    Clause:  "json_extract(attribute, '$.refId')",
    Exclude: true,
}
joinWhereClause["json_extract(attribute, '$.refId')"] = joinClause

builder := param.QueryBuilder{
    EventName:         eNames,
    JoinWhereClause:   joinWhereClause,
    QueryBuilderParam: &subQuery,
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
```

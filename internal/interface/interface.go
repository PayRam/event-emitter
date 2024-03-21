package _interface

import (
	"github.com/PayRam/event-emitter/internal/models"
	"github.com/PayRam/event-emitter/service/param"
)

type EventService interface {
	CreateEvent(event models.EEEvent) error
	QueryEvents(query param.QuerySpec) ([]models.EEEvent, error)
}

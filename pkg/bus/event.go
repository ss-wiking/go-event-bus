package bus

import (
	"github.com/google/uuid"
	"time"
)

// Dispatchable contract to prepare Event for JSON serialization
type Dispatchable interface {
	SetId(uuid.UUID)
	SetCreatedAt(time.Time)
	SetEventName(name string)
}

type Event struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	EventName string    `json:"event_name"`
}

func (e *Event) SetId(id uuid.UUID) {
	e.Id = id
}

func (e *Event) SetCreatedAt(t time.Time) {
	e.CreatedAt = t
}

func (e *Event) SetEventName(name string) {
	e.EventName = name
}

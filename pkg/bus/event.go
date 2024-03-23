package bus

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Payload   []byte    `json:"payload"`
}

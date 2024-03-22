package bus

import "encoding/json"

// EventHandler take an event and do things
type EventHandler interface {
	Handle(message string) error
}

func DecodeMessage(message *string, event interface{}) error {
	err := json.Unmarshal([]byte(*message), event)

	return err
}

package bus

// EventHandler take an event and do things
type EventHandler interface {
	Handle(message string) error
}

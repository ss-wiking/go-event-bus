package bus

// EventListener listens to queue, deserializes events and executes handlers
type EventListener interface {
	Listen(queue string)
}

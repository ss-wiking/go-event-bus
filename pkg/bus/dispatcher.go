package bus

// EventDispatcher pushes Event to the queue
type EventDispatcher interface {
	Dispatch(queue, name string, event interface{}) error
}

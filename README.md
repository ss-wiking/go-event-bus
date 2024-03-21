# GO Event Bus

---

Simple golang implementation of event bus

## Description

### `Dispatchable` interface and `Event` struct

are base (abstract) entities. All custom events must be composed with `Event` struct.<br>
**Attention**: `EventName` is a mandatory field of `Event` struct to fill.<br>
Example:
```go
package main

import "github.com/ss-wiking/go-event-bus/pkg/bus"

type UserCreatedEvent struct {
	bus.Event
	UserId int `json:"user_id"`
}

func main() {
	event := UserCreatedEvent{
		Event:  bus.Event{EventName: UserCreatedEventKey},
		UserId: 5678,
	}
}
```

### `EventDispatcher`
is the service that serializes and pushes event to the queue via `Dispatch` method.
Can be built via `bus.NewEventDispatcher` function.

### `EventListener`
is the service that listens to the queue, retrieves messages, and calls a list of event subscribed handlers.
Can be built via `bus.NewEventListener` function.

### `EventHandler` interface implementations
receives raw message payload as a string and performs any actions with this message
package redis

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/ss-wiking/go-event-bus/pkg/bus"
	"time"
)

type EventListener struct {
	client   *redis.Client
	handlers map[string][]bus.EventHandler
}

func NewEventListener(client *redis.Client, handlers map[string][]bus.EventHandler) *EventListener {
	return &EventListener{
		client:   client,
		handlers: handlers,
	}
}

func (l *EventListener) Listen(queue string) {
	ctx := context.Background()

	for {
		message, err := l.client.LPop(ctx, queue).Bytes()
		if err != nil {
			time.Sleep(3 * time.Second)
		}

		event, err := l.deserialize(message)
		if err != nil {
			// broken message
			continue
		}

		l.processEvent(&event)
	}
}

func (l *EventListener) processEvent(event *bus.Event) {
	// get handlers by event name
	eventHandlers, exists := l.handlers[event.Name]
	if !exists {
		return
	}

	for _, handler := range eventHandlers {
		go func(handler bus.EventHandler) {
			handler.Handle(*event)
		}(handler)
	}
}

func (l *EventListener) deserialize(message []byte) (bus.Event, error) {
	var event bus.Event
	err := json.Unmarshal(message, &event)

	return event, err
}

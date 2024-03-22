package bus

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/tidwall/gjson"
)

// EventListener listens to queue, deserializes events and executes handlers
type EventListener interface {
	Listen(queue string) error
}

type RedisEventListener struct {
	client   *redis.Client
	handlers map[string][]EventHandler
}

func NewEventListener(client *redis.Client, handlers map[string][]EventHandler) EventListener {
	return &RedisEventListener{
		client:   client,
		handlers: handlers,
	}
}

func (l *RedisEventListener) Listen(queue string) error {
	ctx := context.TODO()

	subscription := l.client.Subscribe(ctx, queue)
	channel := subscription.Channel()

	defer func(subscription *redis.PubSub) {
		_ = subscription.Close()
	}(subscription)

	for {
		message, ok := <-channel

		if !ok {
			return errors.New("listening connection refused")
		}

		go l.processMessage(message)
	}
}

func (l *RedisEventListener) processMessage(message *redis.Message) {
	// resolve handlers list by event name
	eventName := gjson.Get(message.Payload, "event_name").String()
	eventHandlers, exists := l.handlers[eventName]
	if !exists {
		return
	}

	for _, handler := range eventHandlers {
		go func(handler EventHandler) {
			_ = handler.Handle(message.Payload)
		}(handler)
	}
}

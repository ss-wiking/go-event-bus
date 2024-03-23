package redis

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/ss-wiking/go-event-bus/pkg/bus"
	"time"
)

type EventDispatcher struct {
	client *redis.Client
}

func NewEventDispatcher(client *redis.Client) *EventDispatcher {
	return &EventDispatcher{client: client}
}

func (d *EventDispatcher) Dispatch(queue, name string, event interface{}) error {
	payload, err := d.serialize(event)
	if err != nil {
		return err
	}

	message, _ := d.serialize(&bus.Event{
		Id:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		Name:      name,
		Payload:   payload,
	})
	if err != nil {
		return err
	}

	res := d.client.RPush(context.Background(), queue, message)

	return res.Err()
}

func (d *EventDispatcher) serialize(event interface{}) ([]byte, error) {
	return json.Marshal(event)
}

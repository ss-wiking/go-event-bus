package bus

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"time"
)

// EventDispatcher pushes Event to the queue
type EventDispatcher interface {
	Dispatch(dispatchable Dispatchable) error
}

type RedisEventDispatcher struct {
	client *redis.Client
	queue  string
}

func NewEventDispatcher(client *redis.Client, queue string) EventDispatcher {
	return &RedisEventDispatcher{
		client: client,
		queue:  queue,
	}
}

func (d *RedisEventDispatcher) Dispatch(dispatchable Dispatchable) error {
	prepared, err := d.prepareMessage(dispatchable)
	if err != nil {
		return err
	}

	ctx := context.TODO()

	err = d.client.Publish(ctx, d.queue, string(prepared)).Err()

	return err
}

func (d *RedisEventDispatcher) prepareMessage(dispatchable Dispatchable) ([]byte, error) {
	// Set some mandatory fields (maybe it will best practice to store this in ctx)
	dispatchable.SetId(uuid.New())
	dispatchable.SetCreatedAt(time.Now().UTC())
	dispatchable.SetEventName(RecognizeEventName(dispatchable))

	// JSON encode
	prepared, err := json.Marshal(dispatchable)

	if err != nil {
		return nil, err
	}

	return prepared, err
}

package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/ss-wiking/go-event-bus/pkg/bus"
	"time"
)

const (
	QueueName           = "event-bus"
	UserCreatedEventKey = "user.created"
)

func main() {
	handlers := map[string][]bus.EventHandler{
		UserCreatedEventKey: {
			new(UserCreatedHandler),
		},
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	listener := bus.NewEventListener(client, handlers)

	// listen channel async
	go func() {
		err := listener.Listen(QueueName)
		if err != nil {
			fmt.Println(err)
		}
	}()

	dispatcher := bus.NewEventDispatcher(client, QueueName)

	events := []bus.Dispatchable{
		&UserCreatedEvent{
			Event:  bus.Event{EventName: UserCreatedEventKey},
			UserId: 1234,
		},
		&UserCreatedEvent{
			Event:  bus.Event{EventName: UserCreatedEventKey},
			UserId: 5678,
		},
	}

	for _, event := range events {
		event := event

		go func() {
			err := dispatcher.Dispatch(event)
			if err != nil {
				fmt.Println("Error occurred")
				fmt.Println(err)
			} else {
				fmt.Println("Event pushed")
			}
		}()
	}

	time.Sleep(10 * time.Second)
}

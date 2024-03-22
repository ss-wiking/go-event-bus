package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/ss-wiking/go-event-bus/pkg/bus"
	"time"
)

const (
	QueueName = "event-bus"
)

var events = []bus.Dispatchable{
	&UserCreatedEvent{
		UserId: 1234,
	},
	&UserCreatedEvent{
		UserId: 5678,
	},
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	listener := bus.NewEventListener(client, GetHandlersList())

	// listen channel async
	go func() {
		err := listener.Listen(QueueName)
		if err != nil {
			fmt.Println(err)
		}
	}()

	dispatcher := bus.NewEventDispatcher(client, QueueName)

	for _, event := range events {
		go func(event bus.Dispatchable) {
			err := dispatcher.Dispatch(event)
			if err != nil {
				fmt.Println("Error occurred")
				fmt.Println(err)
			} else {
				fmt.Println("Event pushed")
			}
		}(event)
	}

	time.Sleep(10 * time.Second)
}

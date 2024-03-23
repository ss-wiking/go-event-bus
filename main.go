package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/ss-wiking/go-event-bus/internal/app"
	redisbus "github.com/ss-wiking/go-event-bus/pkg/bus/redis"
	"os"
	"sync"
	"time"
)

func main() {
	queueName := os.Getenv("QUEUE_NAME")
	client := app.InitRedisClient()

	listener := redisbus.NewEventListener(client, app.GetHandlersList())

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		listener.Listen(queueName)
	}()

	dispatcher := redisbus.NewEventDispatcher(client)

	app.SpamEvents(
		dispatcher,
		queueName,
		10,
		500*time.Millisecond,
	)
}

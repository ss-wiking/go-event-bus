package app

import (
	"github.com/ss-wiking/go-event-bus/internal/app/events"
	"github.com/ss-wiking/go-event-bus/pkg/bus"
	"math/rand"
	"time"
)

func SpamEvents(dispatcher bus.EventDispatcher, queue string, count int, pause time.Duration) {
	for i := 0; i < count; i++ {
		id := rand.Int()

		var event interface{}
		var name string
		if id%2 == 0 {
			event = &events.BotFailedEvent{BotId: id}
			name = "bot.failed"
		} else {
			event = &events.FlowCompletedEvent{FlowId: id}
			name = "flow.completed"
		}

		err := dispatcher.Dispatch(queue, name, event)
		if err != nil {
			panic(err)
		}

		time.Sleep(pause)
	}
}

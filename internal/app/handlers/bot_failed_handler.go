package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/ss-wiking/go-event-bus/internal/app/events"
	"github.com/ss-wiking/go-event-bus/pkg/bus"
)

type BotFailedHandler struct {
	// Dependency Injection and misc here
}

func (h *BotFailedHandler) Handle(message bus.Event) {
	event := new(events.BotFailedEvent)
	err := json.Unmarshal(message.Payload, &event)
	if err != nil {
		return
	}

	fmt.Println(fmt.Sprintf("[BotFailedHandler] Bot ID: %d", event.BotId))
}

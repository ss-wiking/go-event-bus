package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/ss-wiking/go-event-bus/internal/app/events"
	"github.com/ss-wiking/go-event-bus/pkg/bus"
)

type FlowCompletedHandler struct {
	// Dependency Injection and misc here
}

func (h *FlowCompletedHandler) Handle(message bus.Event) {
	event := new(events.FlowCompletedEvent)
	err := json.Unmarshal(message.Payload, &event)
	if err != nil {
		return
	}

	fmt.Println(fmt.Sprintf("[FlowCompletedHandler] Bot ID: %d", event.FlowId))
}

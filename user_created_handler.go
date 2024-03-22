package main

import (
	"fmt"
	"github.com/ss-wiking/go-event-bus/pkg/bus"
)

type UserCreatedHandler struct {
	// Dependency Injection and misc here
}

func (h *UserCreatedHandler) Handle(message string) error {
	event := new(UserCreatedEvent)
	err := bus.DecodeMessage(&message, event)

	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Processing UserCreatedEvent message: %s", event.Id))

	fmt.Println(fmt.Sprintf("Created at: %s", event.CreatedAt))
	fmt.Println(fmt.Sprintf("User ID: %d", event.UserId))

	fmt.Println(fmt.Sprintf("Message UserCreatedEvent processed: %s", event.Id))

	return nil
}

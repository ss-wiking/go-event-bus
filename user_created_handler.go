package main

import (
	"encoding/json"
	"fmt"
)

type UserCreatedHandler struct {
	// Dependency Injection and misc here
}

func (h *UserCreatedHandler) Handle(message string) error {
	// decode message to the struct
	event := new(UserCreatedEvent)

	err := json.Unmarshal([]byte(message), &event)

	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Processing UserCreatedEvent message: %s", event.Id))

	fmt.Println(fmt.Sprintf("Created at: %s", event.CreatedAt))
	fmt.Println(fmt.Sprintf("User ID: %d", event.UserId))

	fmt.Println(fmt.Sprintf("Message UserCreatedEvent processed: %s", event.Id))

	return nil
}

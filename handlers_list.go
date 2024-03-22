package main

import "github.com/ss-wiking/go-event-bus/pkg/bus"

func GetHandlersList() map[string][]bus.EventHandler {
	return map[string][]bus.EventHandler{
		bus.RecognizeEventName(new(UserCreatedEvent)): {
			new(UserCreatedHandler),
		},
	}
}

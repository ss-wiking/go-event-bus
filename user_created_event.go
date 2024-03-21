package main

import "github.com/ss-wiking/go-event-bus/pkg/bus"

type UserCreatedEvent struct {
	bus.Event
	UserId int `json:"user_id"`
}

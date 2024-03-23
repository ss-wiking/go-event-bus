package app

import (
	"github.com/ss-wiking/go-event-bus/internal/app/handlers"
	"github.com/ss-wiking/go-event-bus/pkg/bus"
)

func GetHandlersList() map[string][]bus.EventHandler {
	return map[string][]bus.EventHandler{
		"bot.failed": {
			new(handlers.BotFailedHandler),
		},
		"flow.completed": {
			new(handlers.FlowCompletedHandler),
		},
	}
}

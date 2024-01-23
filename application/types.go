package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-twitch/services/twitch"
)

type Application interface {
	Run() error
	Shutdown()
}

type Impl struct {
	twitchService twitch.Service
	broker        amqp.MessageBroker
}

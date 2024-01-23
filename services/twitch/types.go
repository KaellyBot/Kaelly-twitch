package twitch

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-twitch/repositories/streamers"
)

const (
	routingkey = "news.twitch"
)

type Service interface {
	Consume() error
}

type Impl struct {
	streamerRepo streamers.Repository
	broker       amqp.MessageBroker
}

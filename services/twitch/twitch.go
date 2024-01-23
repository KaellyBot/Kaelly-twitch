package twitch

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-twitch/repositories/streamers"
)

func New(streamerRepo streamers.Repository, broker amqp.MessageBroker) (*Impl, error) {
	return &Impl{
		broker:       broker,
		streamerRepo: streamerRepo,
	}, nil
}

func (service *Impl) Consume() error {
	// TODO
	return service.dispatchTwitchEvent()
}

func (service *Impl) dispatchTwitchEvent() error {
	// TODO
	return nil
}

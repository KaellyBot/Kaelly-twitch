package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-twitch/models/constants"
	"github.com/kaellybot/kaelly-twitch/repositories/streamers"
	"github.com/kaellybot/kaelly-twitch/services/twitch"
	"github.com/kaellybot/kaelly-twitch/utils/databases"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() (*Impl, error) {
	// misc
	db, err := databases.New()
	if err != nil {
		return nil, err
	}

	broker, err := amqp.New(constants.RabbitMQClientID, viper.GetString(constants.RabbitMQAddress), nil)
	if err != nil {
		return nil, err
	}

	// repositories
	streamerRepo := streamers.New(db)

	// services
	twitchService, err := twitch.New(streamerRepo, broker)
	if err != nil {
		return nil, err
	}

	return &Impl{
		twitchService: twitchService,
		broker:        broker,
	}, nil
}

func (app *Impl) Run() error {
	return app.twitchService.Consume()
}

func (app *Impl) Shutdown() {
	app.broker.Shutdown()
	log.Info().Msgf("Application is no longer running")
}

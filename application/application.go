package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-twitch/models/constants"
	"github.com/kaellybot/kaelly-twitch/repositories/streamers"
	"github.com/kaellybot/kaelly-twitch/services/twitch"
	"github.com/kaellybot/kaelly-twitch/utils/databases"
	"github.com/kaellybot/kaelly-twitch/utils/insights"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() (*Impl, error) {
	// misc
	broker := amqp.New(constants.RabbitMQClientID, viper.GetString(constants.RabbitMQAddress))
	db := databases.New()
	probes := insights.NewProbes(broker.IsConnected, db.IsConnected)
	prom := insights.NewPrometheusMetrics()

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
		db:            db,
		probes:        probes,
		prom:          prom,
	}, nil
}

func (app *Impl) Run() error {
	app.probes.ListenAndServe()
	app.prom.ListenAndServe()

	if err := app.db.Run(); err != nil {
		return err
	}

	if err := app.broker.Run(); err != nil {
		return err
	}

	return app.twitchService.Consume()
}

func (app *Impl) Shutdown() {
	app.broker.Shutdown()
	app.db.Shutdown()
	app.prom.Shutdown()
	app.probes.Shutdown()
	log.Info().Msgf("Application is no longer running")
}

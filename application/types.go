package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-twitch/services/twitch"
	"github.com/kaellybot/kaelly-twitch/utils/databases"
	"github.com/kaellybot/kaelly-twitch/utils/insights"
)

type Application interface {
	Run() error
	Shutdown()
}

type Impl struct {
	twitchService twitch.Service
	broker        amqp.MessageBroker
	db            databases.MySQLConnection
	probes        insights.Probes
	prom          insights.PrometheusMetrics
}

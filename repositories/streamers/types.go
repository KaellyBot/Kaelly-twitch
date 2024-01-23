package streamers

import (
	"github.com/kaellybot/kaelly-twitch/models/entities"
	"github.com/kaellybot/kaelly-twitch/utils/databases"
)

type Repository interface {
	GetStreamers() ([]entities.Streamer, error)
}

type Impl struct {
	db databases.MySQLConnection
}

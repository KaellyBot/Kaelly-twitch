package streamers

import (
	"github.com/kaellybot/kaelly-twitch/models/entities"
	"github.com/kaellybot/kaelly-twitch/utils/databases"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) GetStreamers() ([]entities.Streamer, error) {
	var streamers []entities.Streamer
	response := repo.db.GetDB().Model(&entities.Streamer{}).Find(&streamers)
	return streamers, response.Error
}

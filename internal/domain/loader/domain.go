package loader

import (
	"nas-torrent-bot/internal/dig/config"
)

type Loader struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Loader {
	return &Loader{
		cfg: cfg,
	}
}

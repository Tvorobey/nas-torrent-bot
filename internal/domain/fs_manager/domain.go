package fs_manager

import "nas-torrent-bot/internal/dig/config"

type FSManager struct {
	cfg *config.Config
}

func New(cfg *config.Config) *FSManager {
	return &FSManager{cfg: cfg}
}

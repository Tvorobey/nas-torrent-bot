package fs_watcher

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"nas-torrent-bot/internal/dig/config"
)

type Watcher struct {
	w      *fsnotify.Watcher
	Sender SendMessageUseCase

	cfg *config.Config
}

func New(
	sender SendMessageUseCase,
	cfg *config.Config,
) (*Watcher, error) {
	fsNotify, err := fsnotify.NewWatcher()

	if err != nil {
		return nil, errors.New("fsnotify.NewWatcher")
	}

	return &Watcher{
			w:      fsNotify,
			Sender: sender,
			cfg:    cfg,
		},
		nil
}

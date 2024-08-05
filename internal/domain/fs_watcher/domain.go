package fs_watcher

import (
	"errors"
	"nas-torrent-bot/internal/dig/config"
	"nas-torrent-bot/internal/usecase/send_message"

	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	w      *fsnotify.Watcher
	Sender *send_message.SendMessageUseCase

	cfg *config.Config
}

func New(
	sender *send_message.SendMessageUseCase,
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

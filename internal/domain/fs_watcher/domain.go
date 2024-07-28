package fs_watcher

import (
	"errors"
	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	w        *fsnotify.Watcher
	WatchDir string
	Sender   Sender
}

func New(
	watchDir string,
	sender Sender,
) (*Watcher, error) {
	fsNotify, err := fsnotify.NewWatcher()

	if err != nil {
		return nil, errors.New("fsnotify.NewWatcher")
	}

	return &Watcher{
			w:        fsNotify,
			WatchDir: watchDir,
			Sender:   sender,
		},
		nil
}

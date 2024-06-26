package fs_watcher

import (
	"errors"
	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	w *fsnotify.Watcher

	events chan string
}

func New() (*Watcher, error) {
	fsNotify, err := fsnotify.NewWatcher()

	if err != nil {
		return nil, errors.New("fsnotify.NewWatcher")
	}

	return &Watcher{
			w:      fsNotify,
			events: make(chan string, 1),
		},
		nil
}

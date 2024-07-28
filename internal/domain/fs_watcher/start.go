package fs_watcher

import (
	"context"
	"fmt"
	"path"

	"github.com/fsnotify/fsnotify"

	"nas-torrent-bot/internal/domain/fs_watcher/entity"
)

func (w *Watcher) Start(ctx context.Context, watchDir string) error {
	if err := w.w.Add(watchDir); err != nil {
		return fmt.Errorf("dw.w.Add: %s", err.Error())
	}

	go func() {
		for {
			select {
			case event, ok := <-w.w.Events:
				if !ok {
					return
				}
				if event.Op == fsnotify.Create {
					fullPath := event.Name
					_, file := path.Split(fullPath)
					ext := path.Ext(file)

					if _, ok := entity.ExtBlackList[ext]; !ok {
						w.Sender.SendMessageToAll(fmt.Sprintf("Файл %s успешно скачан", file))
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

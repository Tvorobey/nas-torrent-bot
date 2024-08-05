package fs_watcher

import (
	"context"
	"fmt"
	"path"

	"github.com/fsnotify/fsnotify"

	"nas-torrent-bot/internal/domain/fs_watcher/entity"
)

const (
	fileDownloadedMessage = "Файл %s успешно скачан.\nВведи команду /move %s <to>, где to папка, в которую надо переместить файл"
)

func (w *Watcher) Start(ctx context.Context) error {
	if err := w.w.Add(w.cfg.WatchDir); err != nil {
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
						w.Sender.SendMessageToAll(fmt.Sprintf(fileDownloadedMessage, file, file))
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

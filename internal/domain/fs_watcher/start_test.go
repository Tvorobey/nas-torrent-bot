package fs_watcher

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/stretchr/testify/assert"
	"nas-torrent-bot/internal/dig/config"
	"nas-torrent-bot/internal/domain/fs_watcher/mocks"
	"testing"
	"time"
)

func TestWatcher_Start(t *testing.T) {
	mockUseCase := &mocks.SendMessageUseCaseMock{
		SendMessageToAllFunc: func(message string) {},
	}

	cfg := &config.Config{
		WatchDir: ".",
	}

	w, err := New(
		mockUseCase,
		cfg,
	)
	assert.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err = w.Start(ctx)
		assert.NoError(t, err)
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		w.w.Events <- fsnotify.Event{
			Name: "testfile.txt",
			Op:   fsnotify.Create,
		}
	}()

	time.Sleep(200 * time.Millisecond)

	sendCalls := mockUseCase.SendMessageToAllCalls()
	assert.Len(t, sendCalls, 1)
	message := "Файл testfile.txt успешно скачан\n" + fmt.Sprintf(fileDownloadedMessage, "testfile.txt")
	assert.Equal(
		t,
		message,
		sendCalls[0].Message,
	)
	cancel()
	time.Sleep(100 * time.Millisecond)
}

package fs_watcher

import (
	"context"
	"nas-torrent-bot/internal/dig/config"
	"nas-torrent-bot/internal/domain/fs_watcher/mocks"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/stretchr/testify/assert"
)

func TestWatcher_Start(t *testing.T) {
	type args struct {
		cfg   *config.Config
		event fsnotify.Event
	}
	type fields struct {
		sender *mocks.SenderMock
	}
	tests := []struct {
		name        string
		args        args
		fields      fields
		wantErr     assert.ErrorAssertionFunc
		want        string
		assertCalls func(t *testing.T, f fields)
	}{
		{
			name: "error_:_watcher",
			args: args{
				cfg: &config.Config{
					WatchDir: "blabla",
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.ErrorContains(t, err, "dw.w.Add:")

				return true
			},
			assertCalls: func(t *testing.T, f fields) {
				assert.Nil(t, f.sender)
			},
		},
		{
			name: "black_list_ext",
			args: args{
				cfg: &config.Config{
					WatchDir: "./",
				},
				event: fsnotify.Event{
					Name: "file.download",
					Op:   fsnotify.Create,
				},
			},
			wantErr: assert.NoError,
			want:    "file.download",
			assertCalls: func(t *testing.T, f fields) {
				assert.Nil(t, f.sender)
			},
		},
		{
			name: "success",
			args: args{
				cfg: &config.Config{
					WatchDir: "./",
				},
				event: fsnotify.Event{
					Name: "movie.mov",
					Op:   fsnotify.Create,
				},
			},
			fields: fields{
				sender: &mocks.SenderMock{
					SendMessageToAllFunc: func(message string) {},
				},
			},
			wantErr: assert.NoError,
			want:    "movie.mov",
			assertCalls: func(t *testing.T, f fields) {
				sendCalls := f.sender.SendMessageToAllCalls()
				assert.Len(t, sendCalls, 1)
				assert.Equal(
					t,
					"Файл movie.mov успешно скачан",
					sendCalls[0].Message,
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, err := New(
				tt.fields.sender,
				tt.args.cfg,
			)

			assert.NoError(t, err)

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			err = w.Start(ctx)
			tt.wantErr(t, err)

			if err != nil {
				return
			}

			w.w.Events <- tt.args.event
		})
	}
}

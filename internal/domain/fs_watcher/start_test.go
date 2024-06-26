package fs_watcher

import (
	"context"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/stretchr/testify/assert"
)

func TestWatcher_Start(t *testing.T) {
	w, err := fsnotify.NewWatcher()
	assert.NoError(t, err)

	type args struct {
		watchDir string
		event    fsnotify.Event
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
		want    string
	}{
		{
			name: "error_:_watcher",
			args: args{
				watchDir: "blabla",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.ErrorContains(t, err, "dw.w.Add:")

				return true
			},
		},
		{
			name: "black_list_ext",
			args: args{
				watchDir: "./",
				event: fsnotify.Event{
					Name: "file.download",
					Op:   fsnotify.Create,
				},
			},
			wantErr: assert.NoError,
			want:    "file.download",
		},
		{
			name: "success",
			args: args{
				watchDir: "./",
				event: fsnotify.Event{
					Name: "movie.mov",
					Op:   fsnotify.Create,
				},
			},
			wantErr: assert.NoError,
			want:    "movie.mov",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Watcher{
				w:      w,
				events: make(chan string, 1),
			}

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			err := w.Start(ctx, tt.args.watchDir)
			tt.wantErr(t, err)

			if err != nil {
				return
			}

			events := w.GetEventsChan()

			w.w.Events <- tt.args.event

			for {
				select {
				case event, ok := <-events:
					if !ok {
						return
					}

					assert.Equal(
						t,
						tt.want,
						event,
					)
				}
			}
		})
	}
}

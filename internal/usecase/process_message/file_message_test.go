package process_message

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	loaderEntity "nas-torrent-bot/internal/domain/loader/entity"
	"nas-torrent-bot/internal/usecase/process_message/entity"
	"nas-torrent-bot/internal/usecase/process_message/mocks"
	"testing"
)

func TestProcessFileMessage(t *testing.T) {
	type fields struct {
		storage *mocks.StorageMock
		loader  *mocks.LoaderMock
	}
	type args struct {
		in entity.FileMessageIn
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        string
		assertCalls func(t *testing.T, f fields)
	}{
		{
			name: "unknown_user",
			fields: fields{
				storage: &mocks.StorageMock{
					ExistsFunc: func(userID int64) bool {
						return false
					},
				},
			},
			args: args{
				in: entity.FileMessageIn{
					UserID: 11,
				},
			},
			want: entity.StartRulesAnswer,
			assertCalls: func(t *testing.T, f fields) {
				existsCalls := f.storage.ExistsCalls()
				assert.Len(t, existsCalls, 1)
				assert.Equal(
					t,
					11,
					existsCalls[0].UserID,
				)
			},
		},
		{
			name: "download_error",
			fields: fields{
				storage: &mocks.StorageMock{
					ExistsFunc: func(userID int64) bool {
						return true
					},
				},
				loader: &mocks.LoaderMock{
					DownloadFunc: func(in loaderEntity.In) error {
						return assert.AnError
					},
				},
			},
			args: args{
				in: entity.FileMessageIn{
					UserID:   11,
					FileName: "file.txt",
					Url:      "some_url",
				},
			},
			want: fmt.Sprintf(entity.FailedDownloadFile, assert.AnError),
			assertCalls: func(t *testing.T, f fields) {
				existsCalls := f.storage.ExistsCalls()
				assert.Len(t, existsCalls, 1)
				assert.Equal(
					t,
					11,
					existsCalls[0].UserID,
				)

				downloadCalls := f.loader.DownloadCalls()
				assert.Len(t, downloadCalls, 1)
				assert.Equal(
					t,
					loaderEntity.In{
						FileName: "file.txt",
						Url:      "some_url",
					},
					downloadCalls[0].In,
				)
			},
		},
		{
			name: "success",
			fields: fields{
				storage: &mocks.StorageMock{
					ExistsFunc: func(userID int64) bool {
						return true
					},
				},
				loader: &mocks.LoaderMock{
					DownloadFunc: func(in loaderEntity.In) error {
						return nil
					},
				},
			},
			args: args{
				in: entity.FileMessageIn{
					UserID:   11,
					FileName: "file.txt",
					Url:      "some_url",
				},
			},
			want: entity.FileSuccessfullyDownloaded,
			assertCalls: func(t *testing.T, f fields) {
				existsCalls := f.storage.ExistsCalls()
				assert.Len(t, existsCalls, 1)
				assert.Equal(
					t,
					11,
					existsCalls[0].UserID,
				)

				downloadCalls := f.loader.DownloadCalls()
				assert.Len(t, downloadCalls, 1)
				assert.Equal(
					t,
					loaderEntity.In{
						FileName: "file.txt",
						Url:      "some_url",
					},
					downloadCalls[0].In,
				)
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			uc := New(
				tt.fields.storage,
				tt.fields.loader,
				nil,
				nil,
			)
			assert.Equalf(t, tt.want, uc.ProcessFileMessage(tt.args.in), "ProcessFileMessage(%v)", tt.args.in)
		})
	}
}

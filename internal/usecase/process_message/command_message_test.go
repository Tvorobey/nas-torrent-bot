package process_message

import (
	"fmt"
	"nas-torrent-bot/internal/dig/config"
	"nas-torrent-bot/internal/usecase/process_message/entity"
	"nas-torrent-bot/internal/usecase/process_message/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessCommandMessage(t *testing.T) {
	type fields struct {
		storage   *mocks.StorageMock
		loader    *mocks.LoaderMock
		fsManager *mocks.FSManagerMock
		cfg       *config.Config
	}
	type args struct {
		in entity.CommandMessageIn
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        string
		assertCalls func(t *testing.T, f fields)
	}{
		{
			name: "start_:_invalid_secret",
			fields: fields{
				storage: &mocks.StorageMock{
					ExistsFunc: func(userID int64) bool {
						return true
					},
				},
				cfg: &config.Config{
					SecretPhrase: "olo",
				},
			},
			args: args{
				in: entity.CommandMessageIn{
					UserID:  11,
					Command: "start",
					Args:    "lala",
				},
			},
			want: entity.InvalidSecret,
			assertCalls: func(t *testing.T, f fields) {
				assert.Nil(t, f.loader)
				assert.Nil(t, f.fsManager)

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
			name: "start_:_ok",
			fields: fields{
				storage: &mocks.StorageMock{
					ExistsFunc: func(userID int64) bool {
						return true
					},
					AddFunc: func(userID int64, chatID int64) {

					},
				},
				cfg: &config.Config{
					SecretPhrase: "olo",
				},
			},
			args: args{
				in: entity.CommandMessageIn{
					UserID:  11,
					ChatID:  22,
					Command: "start",
					Args:    "olo",
				},
			},
			want: fmt.Sprintf("Теперь я с тобой дружу!\n %s", entity.RulesAnswer),
			assertCalls: func(t *testing.T, f fields) {
				assert.Nil(t, f.loader)
				assert.Nil(t, f.fsManager)

				existsCalls := f.storage.ExistsCalls()
				assert.Len(t, existsCalls, 1)
				assert.Equal(
					t,
					11,
					existsCalls[0].UserID,
				)

				addCalls := f.storage.AddCalls()
				assert.Len(t, addCalls, 1)
				assert.Equal(
					t,
					11,
					addCalls[0].UserID,
				)
				assert.Equal(
					t,
					22,
					addCalls[0].ChatID,
				)
			},
		},
		{
			name: "move_:_error",
			fields: fields{
				storage: &mocks.StorageMock{
					ExistsFunc: func(userID int64) bool {
						return true
					},
				},
				fsManager: &mocks.FSManagerMock{
					MoveFunc: func(from string, to string) error {
						return assert.AnError
					},
				},
			},
			args: args{
				in: entity.CommandMessageIn{
					UserID:  11,
					ChatID:  22,
					Command: "move",
					Args:    "from to",
				},
			},
			want: fmt.Sprintf(entity.FailedMoveFile, assert.AnError),
			assertCalls: func(t *testing.T, f fields) {
				assert.Nil(t, f.loader)
				assert.Nil(t, f.cfg)

				existsCalls := f.storage.ExistsCalls()
				assert.Len(t, existsCalls, 1)
				assert.Equal(
					t,
					11,
					existsCalls[0].UserID,
				)

				moveCalls := f.fsManager.MoveCalls()
				assert.Len(t, moveCalls, 1)
				assert.Equal(
					t,
					"from",
					moveCalls[0].From,
				)
				assert.Equal(
					t,
					"to",
					moveCalls[0].To,
				)
			},
		},
		{
			name: "move_:_success",
			fields: fields{
				storage: &mocks.StorageMock{
					ExistsFunc: func(userID int64) bool {
						return true
					},
				},
				fsManager: &mocks.FSManagerMock{
					MoveFunc: func(from string, to string) error {
						return nil
					},
				},
			},
			args: args{
				in: entity.CommandMessageIn{
					UserID:  11,
					ChatID:  22,
					Command: "move",
					Args:    "some file.txt to my folder",
				},
			},
			want: fmt.Sprintf(entity.SuccessMovedFile, "some file.txt"),
			assertCalls: func(t *testing.T, f fields) {
				assert.Nil(t, f.loader)
				assert.Nil(t, f.cfg)

				existsCalls := f.storage.ExistsCalls()
				assert.Len(t, existsCalls, 1)
				assert.Equal(
					t,
					11,
					existsCalls[0].UserID,
				)

				moveCalls := f.fsManager.MoveCalls()
				assert.Len(t, moveCalls, 1)
				assert.Equal(
					t,
					"some file.txt",
					moveCalls[0].From,
				)
				assert.Equal(
					t,
					"my folder",
					moveCalls[0].To,
				)
			},
		},
		{
			name: "unknown_command",
			fields: fields{
				storage: &mocks.StorageMock{
					ExistsFunc: func(userID int64) bool {
						return true
					},
				},
			},
			args: args{
				in: entity.CommandMessageIn{
					UserID:  11,
					Command: "unknown",
				},
			},
			want: fmt.Sprintf(entity.InvalidCommand, "unknown"),
			assertCalls: func(t *testing.T, f fields) {
				assert.Nil(t, f.loader)
				assert.Nil(t, f.cfg)
				assert.Nil(t, f.fsManager)

				existsCalls := f.storage.ExistsCalls()
				assert.Len(t, existsCalls, 1)
				assert.Equal(
					t,
					11,
					existsCalls[0].UserID,
				)
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			uc := &ProcessMessageUseCase{
				storage:   tt.fields.storage,
				loader:    tt.fields.loader,
				fsManager: tt.fields.fsManager,
				cfg:       tt.fields.cfg,
			}
			assert.Equalf(t, tt.want, uc.ProcessCommandMessage(tt.args.in), "ProcessCommandMessage(%v)", tt.args.in)
		})
	}
}

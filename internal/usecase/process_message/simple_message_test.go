package process_message

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"nas-torrent-bot/internal/usecase/process_message/entity"
	"nas-torrent-bot/internal/usecase/process_message/mocks"
)

func TestProcessSimpleMessage(t *testing.T) {
	type fields struct {
		storage *mocks.StorageMock
	}
	type args struct {
		in entity.SimpleMessageIn
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        string
		assertCalls func(t *testing.T, f fields)
	}{
		{
			name: "ok_:_first_time_visit",
			fields: fields{
				storage: &mocks.StorageMock{
					ExistsFunc: func(userID int64) bool {
						return false
					},
				},
			},
			args: args{
				in: entity.SimpleMessageIn{
					UserID: 11,
					ChatID: 22,
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
			name: "ok_:_exists",
			fields: fields{
				storage: &mocks.StorageMock{
					ExistsFunc: func(userID int64) bool {
						return true
					},
				},
			},
			args: args{
				in: entity.SimpleMessageIn{
					UserID: 11,
					ChatID: 22,
				},
			},
			want: entity.RulesAnswer,
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
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			uc := New(
				tt.fields.storage,
				nil,
				nil,
				nil,
			)

			if got := uc.ProcessSimpleMessage(tt.args.in); got != tt.want {
				t.Errorf("ProcessSimpleMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

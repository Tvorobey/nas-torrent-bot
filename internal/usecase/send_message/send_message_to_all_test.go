package send_message

import (
	"nas-torrent-bot/internal/usecase/send_message/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendMessageToAll(t *testing.T) {
	type fields struct {
		Bot     *mocks.BotMock
		Storage *mocks.StorageMock
	}
	type args struct {
		message string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		assertCalls func(t *testing.T, f fields)
	}{
		{
			name: "ok_:_empty_chats",
			fields: fields{
				Storage: &mocks.StorageMock{
					GetAllChatsFunc: func() []int64 {
						return nil
					},
				},
			},
			args: args{
				message: "Hellow",
			},
			assertCalls: func(t *testing.T, f fields) {
				assert.Nil(t, f.Bot)
			},
		},
		{
			name: "ok_:_has_chats",
			fields: fields{
				Storage: &mocks.StorageMock{
					GetAllChatsFunc: func() []int64 {
						return []int64{11, 22}
					},
				},
				Bot: &mocks.BotMock{
					SendMessageFunc: func(chatIDs []int64, message string) {

					},
				},
			},
			args: args{
				message: "Hellow",
			},
			assertCalls: func(t *testing.T, f fields) {
				sendCalls := f.Bot.SendMessageCalls()
				assert.Len(t, sendCalls, 1)
				assert.Equal(
					t,
					[]int64{11, 22},
					sendCalls[0].ChatIDs,
				)
				assert.Equal(
					t,
					"Hellow",
					sendCalls[0].Message,
				)
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			uc := New(
				tt.fields.Bot,
				tt.fields.Storage,
			)
			uc.SendMessageToAll(tt.args.message)
		})
	}
}

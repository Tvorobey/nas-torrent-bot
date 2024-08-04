package bot

import "nas-torrent-bot/internal/usecase/process_message/entity"

//go:generate moq -skip-ensure -pkg mocks -out ./mocks/message.go . MessageUseCase
type MessageUseCase interface {
	ProcessCommandMessage(in entity.CommandMessageIn) string
	ProcessFileMessage(in entity.FileMessageIn) string
	ProcessSimpleMessage(in entity.SimpleMessageIn) string
}

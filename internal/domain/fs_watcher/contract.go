package fs_watcher

//go:generate moq -skip-ensure -pkg mocks -out ./mocks/send_message_use_case.go . SendMessageUseCase
type SendMessageUseCase interface {
	SendMessageToAll(message string)
}

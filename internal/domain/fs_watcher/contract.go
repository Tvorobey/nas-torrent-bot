package fs_watcher

//go:generate moq -skip-ensure -pkg mocks -out ./mocks/sender.go . Sender
type Sender interface {
	SendMessageToAll(message string)
}

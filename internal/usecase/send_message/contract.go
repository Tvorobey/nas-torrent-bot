package send_message

//go:generate moq -skip-ensure -pkg mocks -out ./mocks/bot.go . Bot
type Bot interface {
	SendMessage(chatIDs []int64, message string)
}

//go:generate moq -skip-ensure -pkg mocks -out ./mocks/storage.go . Storage
type Storage interface {
	GetAllChats() []int64
}

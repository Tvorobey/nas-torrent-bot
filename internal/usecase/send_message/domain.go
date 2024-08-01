package send_message

type SendMessageUseCase struct {
	Bot     Bot
	Storage Storage
}

func New(
	bot Bot,
	storage Storage,
) *SendMessageUseCase {
	return &SendMessageUseCase{
		Bot:     bot,
		Storage: storage,
	}
}

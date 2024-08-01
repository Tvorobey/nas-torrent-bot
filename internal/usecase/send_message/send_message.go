package send_message

func (uc *SendMessageUseCase) SendMessageToAll(message string) {
	chatIDs := uc.Storage.GetAllChats()

	if len(chatIDs) == 0 {
		return
	}

	uc.Bot.SendMessage(chatIDs, message)
}

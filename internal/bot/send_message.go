package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// nolint:errcheck
func (b *Bot) SendMessage(chatIDs []int64, message string) {
	for _, id := range chatIDs {
		msg := tgbotapi.NewMessage(id, message)
		b.tgBot.Send(msg)
	}
}

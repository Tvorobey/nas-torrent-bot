package bot

import (
	"context"
	"nas-torrent-bot/internal/usecase/process_message/entity"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

func (b *Bot) Start(ctx context.Context) error {
	botApi, err := tgbotapi.NewBotAPI(b.cfg.BotToken)
	if err != nil {
		return errors.Wrap(err, "tgbotapi.NewBotAPI")
	}

	b.tgBot = botApi

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.tgBot.GetUpdatesChan(u)

	b.log.Info("Bot started")

	go func() {
		updates := updates
		for {
			select {
			case update := <-updates:
				if update.Message != nil {
					chatID := update.Message.Chat.ID
					userID := update.Message.From.ID

					if doc := update.Message.Document; doc != nil {
						file, err := b.tgBot.GetFile(tgbotapi.FileConfig{FileID: doc.FileID})
						if err != nil {
							b.SendMessage([]int64{chatID}, "Не удалось получить данные о файле. Попробуй еще раз")
						}

						ans := b.uc.ProcessFileMessage(entity.FileMessageIn{
							UserID:   userID,
							FileName: doc.FileName,
							Url:      file.Link(b.cfg.BotToken),
						})

						b.SendMessage([]int64{chatID}, ans)
					} else if update.Message.IsCommand() {
						ans := b.uc.ProcessCommandMessage(entity.CommandMessageIn{
							UserID:  userID,
							ChatID:  chatID,
							Command: entity.Command(update.Message.Command()),
							Args:    update.Message.CommandArguments(),
						})

						b.SendMessage([]int64{chatID}, ans)
					} else {
						ans := b.uc.ProcessSimpleMessage(entity.SimpleMessageIn{
							UserID: userID,
							ChatID: chatID,
						})

						b.SendMessage([]int64{chatID}, ans)
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

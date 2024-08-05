package bot

import (
	"nas-torrent-bot/internal/dig/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type Bot struct {
	cfg   *config.Config
	tgBot *tgbotapi.BotAPI
	log   *zap.Logger
	uc    MessageUseCase
}

func New(
	cfg *config.Config,
	log *zap.Logger,
	uc MessageUseCase,
) *Bot {
	return &Bot{
		cfg: cfg,
		log: log,
		uc:  uc,
	}
}

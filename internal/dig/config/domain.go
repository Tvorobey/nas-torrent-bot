package config

import (
	"nas-torrent-bot/internal/dig/config/entity"
	"os"
)

type Config struct {
	BotToken     string
	WatchDir     string
	DownloadDir  string
	LogLevel     string
	SecretPhrase string
}

func NewConfig() *Config {
	return &Config{
		BotToken:     os.Getenv(entity.BotTokenEnv),
		WatchDir:     os.Getenv(entity.WatchDirEnv),
		DownloadDir:  os.Getenv(entity.DownloadDirEnv),
		LogLevel:     os.Getenv(entity.LogLevelEnv),
		SecretPhrase: os.Getenv(entity.SecretEnv),
	}
}

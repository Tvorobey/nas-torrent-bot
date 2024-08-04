package main

import (
	"context"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"nas-torrent-bot/internal/bot"
	"nas-torrent-bot/internal/dig/config"
	"nas-torrent-bot/internal/domain/fs_manager"
	"nas-torrent-bot/internal/domain/fs_watcher"
	"nas-torrent-bot/internal/domain/loader"
	"nas-torrent-bot/internal/domain/storage"
	"nas-torrent-bot/internal/usecase/process_message"
	"nas-torrent-bot/internal/usecase/send_message"
	"os"
	"os/signal"
	"syscall"
)

func NewLogger(config *config.Config) (*zap.Logger, error) {
	var logLevel zapcore.Level
	err := logLevel.UnmarshalText([]byte(config.LogLevel))
	if err != nil {
		return nil, err
	}

	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(logLevel),
		Development: true,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	return cfg.Build()
}

func initInternal(ctx context.Context, container *dig.Container) error {
	err := container.Provide(bot.New)
	return err
}

func initDomains(container *dig.Container) error {
	err := container.Provide(config.NewConfig)
	err = container.Provide(NewLogger)
	err = container.Provide(fs_watcher.New)
	err = container.Provide(loader.New)
	err = container.Provide(fs_manager.New)
	err = container.Provide(storage.New)

	return err
}

func initInterfaces(container *dig.Container) error {
	err := container.Provide(func(loader *loader.Loader) process_message.Loader {
		return loader
	})

	err = container.Provide(func(s *storage.Storage) process_message.Storage {
		return s
	})

	err = container.Provide(func(fs *fs_manager.FSManager) process_message.FSManager {
		return fs
	})

	err = container.Provide(func(
		storage process_message.Storage,
		loader process_message.Loader,
		fsManager process_message.FSManager,
		cfg *config.Config) bot.MessageUseCase {
		return process_message.New(
			storage,
			loader,
			fsManager,
			cfg,
		)
	})

	err = container.Provide(func(bot *bot.Bot) send_message.Bot {
		return bot
	})

	err = container.Provide(func(s *storage.Storage) send_message.Storage {
		return s
	})

	err = container.Provide(func(bot send_message.Bot, storage send_message.Storage) *send_message.SendMessageUseCase {
		return send_message.New(bot, storage)
	})

	return err
}

func digInvoke(ctx context.Context, container *dig.Container) error {
	err := container.Invoke(func(cfg *config.Config, logger *zap.Logger) {
		logger.Info("Application starting", zap.String("watchDir", cfg.WatchDir), zap.String("downloadDir", cfg.DownloadDir))
	})

	err = container.Invoke(func(bot *bot.Bot) error {
		return bot.Start(ctx)
	})

	err = container.Invoke(func(watcher *fs_watcher.Watcher, cfg *config.Config, logger *zap.Logger) {
		logger.Info("Starting watcher for", zap.String("dir", cfg.WatchDir))
		err := watcher.Start(ctx)
		if err != nil {
			logger.Fatal("Failed to start watcher", zap.Error(err))
		}
	})

	return err
}

func main() {
	ctx := context.Background()

	container := dig.New()
	err := initDomains(container)
	if err != nil {
		log.Fatalf("failed init domains: %v", err)
	}

	err = initInterfaces(container)
	if err != nil {
		log.Fatalf("failed init interfaces: %v", err)
	}

	err = initInternal(ctx, container)
	if err != nil {
		log.Fatalf("failed init internal: %v", err)
	}

	err = digInvoke(ctx, container)
	if err != nil {
		log.Fatalf("failed invole dig: %v", err)
	}

	// Grace shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)
	sig := <-sigs
	log.Fatal("Signal: ", sig)
}

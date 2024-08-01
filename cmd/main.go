package main

import (
	"context"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"nas-torrent-bot/internal/dig/config"
	"nas-torrent-bot/internal/domain/loader"
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

func initDomains(container *dig.Container) error {
	err := container.Provide(config.NewConfig)
	err = container.Provide(NewLogger)
	//err = container.Provide(fs_watcher.New)
	err = container.Provide(loader.New)

	return err
}

func digInvoke(ctx context.Context, container *dig.Container) error {
	err := container.Invoke(func(cfg *config.Config, logger *zap.Logger) {
		logger.Info("Application starting")
	})

	//err = container.Invoke(func(watcher *fs_watcher.Watcher, cfg *config.Config, logger *zap.Logger) {
	//	logger.Info("Starting watcher for", zap.String("dir", cfg.WatchDir))
	//	err := watcher.Start(ctx)
	//	if err != nil {
	//		logger.Fatal("Failed to start watcher", zap.Error(err))
	//	}
	//})

	return err
}

func main() {
	ctx := context.Background()

	container := dig.New()
	err := initDomains(container)
	if err != nil {
		log.Fatalf("failed init domains: %v", err)
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

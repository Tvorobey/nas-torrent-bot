package process_message

import "nas-torrent-bot/internal/dig/config"

type ProcessMessageUseCase struct {
	storage   Storage
	loader    Loader
	fsManager FSManager

	cfg *config.Config
}

func New(
	storage Storage,
	loader Loader,
	fsManager FSManager,
	cfg *config.Config,
) *ProcessMessageUseCase {
	return &ProcessMessageUseCase{
		storage:   storage,
		loader:    loader,
		fsManager: fsManager,
		cfg:       cfg,
	}
}

package process_message

import "nas-torrent-bot/internal/domain/loader/entity"

//go:generate moq -skip-ensure -pkg mocks -out ./mocks/storage.go . Storage
type Storage interface {
	Exists(userID int64) bool
	Add(userID, chatID int64)
}

//go:generate moq -skip-ensure -pkg mocks -out ./mocks/loader.go . Loader
type Loader interface {
	Download(in entity.In) error
}

//go:generate moq -skip-ensure -pkg mocks -out ./mocks/fs_manager.go . FSManager
type FSManager interface {
	Move(from, to string) error
}

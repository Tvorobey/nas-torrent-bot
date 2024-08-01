package storage

import "nas-torrent-bot/internal/domain/storage/entity"

type Storage struct {
	storage entity.UserIDChatIDMap
}

func New() *Storage {
	return &Storage{
		storage: map[int64]int64{},
	}
}

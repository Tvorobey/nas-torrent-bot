package storage

func (s *Storage) Add(userID, chatID int64) {
	s.storage[userID] = chatID
}

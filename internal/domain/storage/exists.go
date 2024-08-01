package storage

func (s *Storage) Exists(userID int64) bool {
	_, ok := s.storage[userID]
	return ok
}

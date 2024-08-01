package storage

func (s *Storage) GetAllChats() []int64 {
	chatIDs := make([]int64, len(s.storage))

	for key, _ := range s.storage {
		chatIDs = append(chatIDs, s.storage[key])
	}

	return chatIDs
}

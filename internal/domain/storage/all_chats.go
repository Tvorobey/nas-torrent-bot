package storage

func (s *Storage) GetAllChats() []int64 {
	chatIDs := make([]int64, len(s.storage))

	for _, chatID := range s.storage {
		chatIDs = append(chatIDs, chatID)
	}

	return chatIDs
}

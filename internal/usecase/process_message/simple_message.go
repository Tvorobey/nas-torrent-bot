package process_message

import (
	"nas-torrent-bot/internal/usecase/process_message/entity"
)

func (uc *ProcessMessageUseCase) ProcessSimpleMessage(in entity.SimpleMessageIn) string {
	if uc.storage.Exists(in.UserID) {
		return entity.RulesAnswer
	}

	return entity.StartRulesAnswer
}

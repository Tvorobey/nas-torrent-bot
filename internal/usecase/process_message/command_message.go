package process_message

import (
	"fmt"
	"nas-torrent-bot/internal/usecase/process_message/entity"
	"strings"
)

func (uc *ProcessMessageUseCase) ProcessCommandMessage(in entity.CommandMessageIn) string {
	switch in.Command {
	case entity.CommandStart:
		return uc.startCommand(in)
	case entity.CommandMove:
		return uc.moveCommand(in)
	default:
		return fmt.Sprintf(entity.InvalidCommand, in.Command)
	}
}

func (uc *ProcessMessageUseCase) startCommand(in entity.CommandMessageIn) string {
	if in.Args != uc.cfg.SecretPhrase {
		return entity.InvalidSecret
	}

	uc.storage.Add(in.UserID, in.ChatID)

	return fmt.Sprintf("Теперь я с тобой дружу!\n %s", entity.RulesAnswer)
}

func (uc *ProcessMessageUseCase) moveCommand(in entity.CommandMessageIn) string {
	if !uc.storage.Exists(in.UserID) {
		return entity.StartRulesAnswer
	}

	splitArgs := strings.Split(in.Args, " ")
	if err := uc.fsManager.Move(splitArgs[0], splitArgs[1]); err != nil {
		return fmt.Sprintf(entity.FailedMoveFile, err)
	}

	return fmt.Sprintf(entity.SuccessMovedFile, splitArgs[0])
}

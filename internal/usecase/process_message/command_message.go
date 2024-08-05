package process_message

import (
	"fmt"
	"nas-torrent-bot/internal/domain/fs_watcher"
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

	separatorPos := 0

	for i, arg := range splitArgs {
		if arg == fs_watcher.Separator {
			separatorPos = i

			break
		}
	}

	fileName := strings.Join(splitArgs[:separatorPos], " ")
	dest := strings.Join(splitArgs[separatorPos+1:], " ")

	if err := uc.fsManager.Move(fileName, dest); err != nil {
		return fmt.Sprintf(entity.FailedMoveFile, err)
	}

	return fmt.Sprintf(entity.SuccessMovedFile, fileName)
}

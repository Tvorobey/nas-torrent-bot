package process_message

import (
	"fmt"
	loaderEntity "nas-torrent-bot/internal/domain/loader/entity"
	"nas-torrent-bot/internal/usecase/process_message/entity"
)

func (uc *ProcessMessageUseCase) ProcessFileMessage(in entity.FileMessageIn) string {
	if !uc.storage.Exists(in.UserID) {
		return entity.StartRulesAnswer
	}

	if err := uc.loader.Download(loaderEntity.In{
		FileName: in.FileName,
		Url:      in.Url,
	},
	); err != nil {
		return fmt.Sprintf(entity.FailedDownloadFile, err)
	}

	return entity.FileSuccessfullyDownloaded
}

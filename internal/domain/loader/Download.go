package loader

import (
	"fmt"
	"io"
	"nas-torrent-bot/internal/domain/loader/entity"
	"net/http"
	"os"
	"path/filepath"
)

func (l *Loader) Download(in entity.In) error {
	filePath := filepath.Join(l.cfg.DownloadDir, in.FileName)

	// Создаем файл в указанной директории
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", filePath, err)
	}
	defer out.Close()

	// Отправляем запрос на скачивание файла
	resp, err := http.Get(in.Url)
	if err != nil {
		return fmt.Errorf("failed to download file: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Копируем содержимое ответа в файл
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}

	return nil
}

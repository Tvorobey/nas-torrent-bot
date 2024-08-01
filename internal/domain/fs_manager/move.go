package fs_manager

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

func (fs *FSManager) Move(fileName, to string) error {
	dir := fs.cfg.DownloadDir
	currentPath := path.Join(dir, fileName)
	destPath := path.Join(dir, to, fileName)

	fileStat, err := os.Stat(currentPath)
	if err != nil {
		return err
	}

	if fileStat.IsDir() {
		err = fs.moveDir(currentPath, destPath)
		os.RemoveAll(path.Join(fs.cfg.DownloadDir, fileName))
		return err
	}

	return fs.moveFile(currentPath, destPath)
}

func (fs *FSManager) moveDir(srcDir, dstDir string) error {
	srcDir = filepath.Clean(srcDir)
	dstDir = filepath.Clean(dstDir)

	err := os.MkdirAll(dstDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	return filepath.Walk(srcDir, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(srcDir, srcPath)
		dstPath := filepath.Join(dstDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		srcFile, err := os.Open(srcPath)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		return err
	})
}

func (fs *FSManager) moveFile(srcPath, dstPath string) error {
	// Переименовываем (перемещаем) файл
	err := os.Rename(srcPath, dstPath)
	if err != nil {
		return err
	}
	return nil
}

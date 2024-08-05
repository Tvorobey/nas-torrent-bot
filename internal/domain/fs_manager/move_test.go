package fs_manager

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// createTestDir создает временную директорию с тестовыми файлами
func createTestDir(t *testing.T, dir string) {
	err := os.MkdirAll(filepath.Join(dir, "subdir"), os.ModePerm)
	if err != nil {
		t.Fatalf("failed to create subdir: %v", err)
	}

	files := []string{
		"file1.txt",
		"file2.txt",
		filepath.Join("subdir", "file3.txt"),
	}

	for _, file := range files {
		filePath := filepath.Join(dir, file)
		err := ioutil.WriteFile(filePath, []byte("test content"), os.ModePerm)
		if err != nil {
			t.Fatalf("failed to create file %s: %v", filePath, err)
		}
	}
}

// removeTestDir удаляет временную директорию и ее содержимое
func removeTestDir(dir string) error {
	return os.RemoveAll(dir)
}

// TestCopyDirectory проверяет корректность работы функции CopyDirectory
func TestCopyDirectory(t *testing.T) {
	srcDir := t.TempDir() // Создаем временную директорию для исходных данных
	dstDir := t.TempDir() // Создаем временную директорию для назначения

	// Создаем тестовые данные в исходной директории
	createTestDir(t, srcDir)

	fs := New(nil)

	// Копируем директорию
	if err := fs.moveDir(srcDir, dstDir); err != nil {
		t.Fatalf("failed to copy directory: %v", err)
	}

	// Проверяем содержимое целевой директории
	expectedFiles := []string{
		"file1.txt",
		"file2.txt",
		filepath.Join("subdir", "file3.txt"),
	}

	for _, file := range expectedFiles {
		dstPath := filepath.Join(dstDir, file)
		if _, err := os.Stat(dstPath); os.IsNotExist(err) {
			t.Errorf("file %s does not exist in destination directory", dstPath)
		}
	}

	// Проверяем содержимое файлов
	for _, file := range expectedFiles {
		dstPath := filepath.Join(dstDir, file)
		content, err := ioutil.ReadFile(dstPath)
		if err != nil {
			t.Errorf("failed to read file %s: %v", dstPath, err)
		}
		if string(content) != "test content" {
			t.Errorf("file %s has incorrect content", dstPath)
		}
	}
}

// TestMoveFile проверяет корректность работы функции MoveFile.
func TestMoveFile(t *testing.T) {
	// Создаем временные директории для исходной и целевой
	srcDir := t.TempDir()
	dstDir := t.TempDir()

	// Определяем имена файлов
	srcFile := "example file.txt"
	srcPath := filepath.Join(srcDir, srcFile)
	dstPath := filepath.Join(dstDir, srcFile)

	// Создаем файл в исходной директории
	if err := ioutil.WriteFile(srcPath, []byte("test content"), 0644); err != nil {
		t.Fatalf("failed to create source file: %v", err)
	}

	fs := New(nil)

	// Выполняем перемещение файла
	if err := fs.moveFile(srcPath, dstPath); err != nil {
		t.Fatalf("failed to move file: %v", err)
	}

	// Проверяем, что файл существует в целевой директории
	if _, err := os.Stat(dstPath); os.IsNotExist(err) {
		t.Fatalf("file does not exist in destination directory: %s", dstPath)
	}

	// Проверяем, что файл отсутствует в исходной директории
	if _, err := os.Stat(srcPath); !os.IsNotExist(err) {
		t.Fatalf("file still exists in source directory: %s", srcPath)
	}

	// Проверяем содержимое файла в целевой директории
	content, err := ioutil.ReadFile(dstPath)
	if err != nil {
		t.Fatalf("failed to read file in destination directory: %v", err)
	}
	if string(content) != "test content" {
		t.Errorf("file content mismatch: got %q, want %q", content, "test content")
	}
}

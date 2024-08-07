package loader

// nolint:staticcheck
import (
	"io/ioutil"
	"nas-torrent-bot/internal/dig/config"
	"nas-torrent-bot/internal/domain/loader/entity"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestLoader_Download(t *testing.T) {
	mux := http.NewServeMux()
	testFileContent := []byte("test file content")
	mux.HandleFunc("/testfile.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Write(testFileContent)
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	tempDir := t.TempDir()
	tempFilePath := filepath.Join(tempDir, "testfile.txt")

	l := New(
		&config.Config{
			DownloadDir: tempDir,
		},
	)

	if err := l.Download(
		entity.In{
			FileName: "testfile.txt",
			Url:      server.URL + "/testfile.txt",
		},
	); err != nil {
		t.Fatalf("failed to download file: %v", err)
	}

	if _, err := os.Stat(tempFilePath); os.IsNotExist(err) {
		t.Fatalf("file does not exist: %v", err)
	}

	downloadedContent, err := ioutil.ReadFile(tempFilePath)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	if string(downloadedContent) != string(testFileContent) {
		t.Errorf("file content mismatch: got %q, want %q", string(downloadedContent), string(testFileContent))
	}
}

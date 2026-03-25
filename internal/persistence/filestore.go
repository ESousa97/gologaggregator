package persistence

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"gologaggregator/internal/store"
)

const maxFileSize = 10 * 1024 * 1024 // 10MB

// FileStore manages disk persistence with rotation
type FileStore struct {
	mu       sync.Mutex
	filePath string
}

// NewFileStore creates a new persistence manager for the given file path
func NewFileStore(path string) *FileStore {
	return &FileStore{
		filePath: path,
	}
}

// WriteBatch appends a batch of logs to the file and handles rotation
func (fs *FileStore) WriteBatch(batch []store.LogEntry) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(fs.filePath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Check if rotation is needed
	if info, err := os.Stat(fs.filePath); err == nil {
		if info.Size() >= maxFileSize {
			if err := fs.rotate(); err != nil {
				return fmt.Errorf("rotation failed: %w", err)
			}
		}
	}

	// Open file in append mode
	f, err := os.OpenFile(fs.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	for _, entry := range batch {
		line := fmt.Sprintf("[%s] [%s] %s\n", 
			entry.Timestamp.Format(time.RFC3339), 
			entry.Level, 
			entry.Message)
		
		if _, err := f.WriteString(line); err != nil {
			return fmt.Errorf("failed to write log: %w", err)
		}
	}

	return nil
}

// rotate renames the current log file with a timestamp
func (fs *FileStore) rotate() error {
	timestamp := time.Now().Format("20060102-150405")
	newName := fmt.Sprintf("%s.%s", fs.filePath, timestamp)
	return os.Rename(fs.filePath, newName)
}

// Package store provides thread-safe in-memory storage capabilities
// for managing log entries with efficient indexing and search.
package store

import (
	"strings"
	"sync"
	"time"
)

// LogEntry represents a parsed log entry stored in memory.
// It includes metadata like level and timestamp for efficient querying.
type LogEntry struct {
	// Level indicates the severity of the log (e.g., INFO, ERROR, DEBUG).
	Level string `json:"level"`
	// Message contains the descriptive content of the log.
	Message string `json:"message"`
	// Timestamp records exactly when the log was parsed or received.
	Timestamp time.Time `json:"timestamp"`
}

// PersistenceManager defines the behavior for writing log batches to permanent storage.
// This interface allows for different persistence backends (e.g., file, database).
type PersistenceManager interface {
	// WriteBatch writes a collection of log entries to the storage medium.
	WriteBatch(batch []LogEntry) error
}

// MemoryStore implements a thread-safe, in-memory storage for logs.
// It uses a circular buffer approach to maintain a fixed capacity,
// overwriting the oldest logs once the limit is reached.
type MemoryStore struct {
	mu          sync.RWMutex
	logs        []LogEntry
	capacity    int
	cursor      int // points to the next available slot
	persistence PersistenceManager
}

// NewMemoryStore creates a new store with the specified capacity and persistence manager
func NewMemoryStore(capacity int, pm PersistenceManager) *MemoryStore {
	return &MemoryStore{
		logs:        make([]LogEntry, 0, capacity),
		capacity:    capacity,
		persistence: pm,
	}
}

// Persist delegates batch writing to the persistence manager
func (s *MemoryStore) Persist(batch []LogEntry) error {
	if s.persistence == nil {
		return nil
	}
	return s.persistence.WriteBatch(batch)
}

// Append adds a new log entry to the store, maintaining the capacity limit
func (s *MemoryStore) Append(entry LogEntry) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.logs) < s.capacity {
		s.logs = append(s.logs, entry)
	} else {
		// Overwrite old logs if capacity is reached
		s.logs[s.cursor] = entry
		s.cursor = (s.cursor + 1) % s.capacity
	}
}

// GetAll returns a copy of all logs currently in store
func (s *MemoryStore) GetAll() []LogEntry {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Return a copy to avoid race conditions on the slice elements
	result := make([]LogEntry, len(s.logs))
	copy(result, s.logs)
	return result
}

// SearchFilters defines the criteria for searching logs
type SearchFilters struct {
	Level   string
	From    time.Time
	To      time.Time
	Keyword string
}

// Search filters logs based on the provided criteria
func (s *MemoryStore) Search(filters SearchFilters) []LogEntry {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var results []LogEntry
	for _, entry := range s.logs {
		// Filter by Level
		if filters.Level != "" && entry.Level != filters.Level {
			continue
		}

		// Filter by Time Range (From)
		if !filters.From.IsZero() && entry.Timestamp.Before(filters.From) {
			continue
		}

		// Filter by Time Range (To)
		if !filters.To.IsZero() && entry.Timestamp.After(filters.To) {
			continue
		}

		// Filter by Keyword
		if filters.Keyword != "" {
			// Case-insensitive search for better UX
			content := entry.Level + ": " + entry.Message
			if !containsIgnoreCase(content, filters.Keyword) {
				continue
			}
		}

		results = append(results, entry)
	}

	return results
}

func containsIgnoreCase(s, substr string) bool {
	s, substr = strings.ToLower(s), strings.ToLower(substr)
	return strings.Contains(s, substr)
}

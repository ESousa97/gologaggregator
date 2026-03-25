package store

import (
	"strings"
	"sync"
	"time"
)

// LogEntry represents a parsed log entry stored in memory
type LogEntry struct {
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// MemoryStore implements a thread-safe in-memory storage for logs
// with a fixed capacity (circular buffer style)
type MemoryStore struct {
	mu       sync.RWMutex
	logs     []LogEntry
	capacity int
	cursor   int // points to the next available slot
}

// NewMemoryStore creates a new store with the specified capacity
func NewMemoryStore(capacity int) *MemoryStore {
	return &MemoryStore{
		logs:     make([]LogEntry, 0, capacity),
		capacity: capacity,
	}
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

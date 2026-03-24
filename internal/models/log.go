package models

import "time"

// LogEntry represents a single log message within the system
type LogEntry struct {
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

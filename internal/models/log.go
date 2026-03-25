// Package models defines the domain entities and data transfer objects used
// across the log aggregator system.
package models

import "time"

// LogEntry represents a single log message within the system.
// It is used as the standard structure for both ingestion and querying.
type LogEntry struct {
	// Content is the raw or structured body of the log message.
	Content string `json:"content"`
	// Timestamp records the exact moment the log was created or received.
	Timestamp time.Time `json:"timestamp"`
}

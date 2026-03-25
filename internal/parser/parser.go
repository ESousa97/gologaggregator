// Package parser provides utilities for transforming raw log strings into
// structured data entries.
package parser

import (
	"strings"
	"time"

	"github.com/ESousa97/gologaggregator/internal/store"
)

// ParseRawMessage extracts log level and message content from a raw string.
// It expects the format "LEVEL: MESSAGE". If the format is invalid,
// it defaults to the "UNKNOWN" level.
func ParseRawMessage(raw string) store.LogEntry {
	parts := strings.SplitN(raw, ":", 2)

	entry := store.LogEntry{
		Timestamp: time.Now(),
	}

	if len(parts) == 2 {
		// Found "LEVEL: MESSAGE" pattern
		entry.Level = strings.ToUpper(strings.TrimSpace(parts[0]))
		entry.Message = strings.TrimSpace(parts[1])
	} else {
		// Fallback to "UNKNOWN" level if the delimiter is not found
		entry.Level = "UNKNOWN"
		entry.Message = strings.TrimSpace(raw)
	}

	return entry
}

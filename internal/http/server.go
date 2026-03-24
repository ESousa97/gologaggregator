package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"gologaggregator/internal/models"
)

// Server implements the HTTP listener for log messages
type Server struct {
	Address string
}

// Start initializes the HTTP server with its routes
func (s *Server) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/logs", s.handleLog)

	log.Printf("[HTTP] Server listening on %s", s.Address)
	
	server := &http.Server{
		Addr:    s.Address,
		Handler: mux,
	}

	return server.ListenAndServe()
}

// handleLog processes incoming HTTP POST requests with log entries
func (s *Server) handleLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var entry models.LogEntry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Validate or set default timestamp if not provided
	if entry.Timestamp.IsZero() {
		entry.Timestamp = time.Now()
	}

	fmt.Printf("[HTTP-LOG] [%s] %s\n", entry.Timestamp.Format(time.RFC3339), entry.Content)
	w.WriteHeader(http.StatusAccepted)
}

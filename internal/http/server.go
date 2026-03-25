package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"gologaggregator/internal/store"
)

// Server implements the HTTP listener for log messages
type Server struct {
	Address       string
	IngestionChan chan<- string
	Store         *store.MemoryStore
}

// Start initializes the HTTP server with its routes
func (s *Server) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/logs", s.handleLog)
	mux.HandleFunc("/search", s.handleSearch)

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

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// P3: Async - send to buffered channel for later processing
	s.IngestionChan <- string(body)

	w.WriteHeader(http.StatusAccepted)
}

// handleSearch allows querying logs based on filters
func (s *Server) handleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	filters := store.SearchFilters{
		Level:   query.Get("level"),
		Keyword: query.Get("keyword"),
	}

	// Parse timestamps if provided
	if from := query.Get("from"); from != "" {
		if t, err := time.Parse(time.RFC3339, from); err == nil {
			filters.From = t
		}
	}
	if to := query.Get("to"); to != "" {
		if t, err := time.Parse(time.RFC3339, to); err == nil {
			filters.To = t
		}
	}

	results := s.Store.Search(filters)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

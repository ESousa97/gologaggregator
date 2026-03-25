package http

import (
	"io"
	"log"
	"net/http"
)

// Server implements the HTTP listener for log messages
type Server struct {
	Address       string
	IngestionChan chan<- string
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

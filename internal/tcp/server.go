package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"gologaggregator/internal/models"
)

// Server implements the TCP listener for log messages
type Server struct {
	Address string
}

// Start initializes the TCP listener and accepts connections
func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		return fmt.Errorf("failed to start TCP listener on %s: %w", s.Address, err)
	}
	defer listener.Close()

	log.Printf("[TCP] Server listening on %s", s.Address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("[TCP] Error accepting connection: %v", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

// handleConnection processes incoming log entries over TCP
func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		rawMessage := scanner.Text()
		entry := s.parseMessage(rawMessage)
		
		fmt.Printf("[TCP-LOG] [%s] %s\n", entry.Timestamp.Format(time.RFC3339), entry.Content)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("[TCP] Error reading from connection: %v", err)
	}
}

// parseMessage extracts content and timestamp from the raw TCP payload
// Expected format: "TIMESTAMP|CONTENT" or just "CONTENT" (defaults to current time)
func (s *Server) parseMessage(raw string) models.LogEntry {
	parts := strings.SplitN(raw, "|", 2)
	
	if len(parts) == 2 {
		t, err := time.Parse(time.RFC3339, parts[0])
		if err == nil {
			return models.LogEntry{
				Timestamp: t,
				Content:   parts[1],
			}
		}
	}

	return models.LogEntry{
		Timestamp: time.Now(),
		Content:   raw,
	}
}

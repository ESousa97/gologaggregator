// Package tcp implements the high-throughput raw TCP ingestion layer.
package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// Server implements the TCP listener for log messages.
// It is designed for high-performance, line-delimited text ingestion.
type Server struct {
	// Address is the host:port string the server will bind to.
	Address string
	// IngestionChan is the write-only channel for sending raw logs to the pipeline.
	IngestionChan chan<- string
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

		// P3: Async - send to buffered channel for later processing
		s.IngestionChan <- rawMessage
	}

	if err := scanner.Err(); err != nil {
		log.Printf("[TCP] Error reading from connection: %v", err)
	}
}

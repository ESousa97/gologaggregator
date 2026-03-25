package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gologaggregator/internal/config"
	"gologaggregator/internal/http"
	"gologaggregator/internal/pipeline"
	"gologaggregator/internal/store"
	"gologaggregator/internal/tcp"
)

func main() {
	// P3: Config validation on boot
	cfg := config.Load()

	// Initialize thread-safe in-memory store
	// Capacity: 10,000 logs
	logStore := store.NewMemoryStore(10000)

	// Initialize log processing pipeline
	// Batch size: 100, Timeout: 5s, Workers: 4, Buffer: 1000
	proc := pipeline.NewProcessor(pipeline.BatchConfig{
		MaxSize:     100,
		MaxWaitTime: 5 * time.Second,
		WorkerCount: 4,
		BufferSize:  1000,
	}, logStore)
	proc.Start()

	// Initialize servers
	tcpServer := &tcp.Server{
		Address:       cfg.TCPAddress,
		IngestionChan: proc.IngestionChan,
	}
	httpServer := &http.Server{
		Address:       cfg.HTTPAddress,
		IngestionChan: proc.IngestionChan,
	}

	errChan := make(chan error, 2)

	// Start TCP server in a goroutine
	go func() {
		if err := tcpServer.Start(); err != nil {
			errChan <- err
		}
	}()

	// Start HTTP server in a goroutine
	go func() {
		if err := httpServer.Start(); err != nil {
			errChan <- err
		}
	}()

	// Graceful shutdown handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		log.Fatalf("Critical error during server startup: %v", err)
	case sig := <-sigChan:
		log.Printf("Shutting down... Received signal: %v", sig)
	}

	// Wait for processing to drain
	proc.Stop()
}

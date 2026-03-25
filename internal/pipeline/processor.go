// Package pipeline manages the concurrent flow of log ingestion,
// buffering, batching, and dispatching to storage engines.
package pipeline

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ESousa97/gologaggregator/internal/parser"
	"github.com/ESousa97/gologaggregator/internal/store"
)

// LogStore defines the behavior for storing and persisting log entries.
// It decouples the pipeline from the specific storage implementation.
type LogStore interface {
	// Append adds a single log entry to the volatile storage.
	Append(entry store.LogEntry)
	// Persist handles disk persistence for a batch of log entries.
	Persist(batch []store.LogEntry) error
}

// BatchConfig defines the parameters for log batching
type BatchConfig struct {
	MaxSize     int
	MaxWaitTime time.Duration
	WorkerCount int
	BufferSize  int
}

// Processor manages the ingestion channel and orchestrates the worker pool.
// It handles gravity-based log batching and ensures ordered processing.
type Processor struct {
	// IngestionChan is the entry point for raw log strings.
	IngestionChan chan string
	config        BatchConfig
	store         LogStore
	wg            sync.WaitGroup
}

// NewProcessor creates a new pipeline processor with defined configuration
func NewProcessor(cfg BatchConfig, logStore LogStore) *Processor {
	return &Processor{
		IngestionChan: make(chan string, cfg.BufferSize),
		config:        cfg,
		store:         logStore,
	}
}

// Start launches the worker pool to consume from the ingestion channel
func (p *Processor) Start() {
	log.Printf("[PIPELINE] Starting %d workers with batch size %d and timeout %s",
		p.config.WorkerCount, p.config.MaxSize, p.config.MaxWaitTime)

	for i := 1; i <= p.config.WorkerCount; i++ {
		p.wg.Add(1)
		go p.worker(i)
	}
}

// worker consumes from the ingestion channel and batches messages
func (p *Processor) worker(id int) {
	defer p.wg.Done()

	batch := make([]store.LogEntry, 0, p.config.MaxSize)
	ticker := time.NewTicker(p.config.MaxWaitTime)
	defer ticker.Stop()

	for {
		select {
		case msg, ok := <-p.IngestionChan:
			if !ok {
				// Channel closed, process remaining batch and exit
				if len(batch) > 0 {
					p.processBatch(id, batch)
				}
				return
			}

			// P5: Parse and validate on the border
			entry := parser.ParseRawMessage(msg)

			// Index into memory store immediately for fast retrieval
			p.store.Append(entry)

			batch = append(batch, entry)
			if len(batch) >= p.config.MaxSize {
				p.processBatch(id, batch)
				batch = make([]store.LogEntry, 0, p.config.MaxSize)
				ticker.Reset(p.config.MaxWaitTime)
			}

		case <-ticker.C:
			if len(batch) > 0 {
				p.processBatch(id, batch)
				batch = make([]store.LogEntry, 0, p.config.MaxSize)
			}
		}
	}
}

// processBatch simulates sending the batch to the next stage (e.g., storage, API)
func (p *Processor) processBatch(workerID int, batch []store.LogEntry) {
	fmt.Printf("[PIPELINE-WORKER-%d] Processing batch of %d parsed messages\n", workerID, len(batch))

	// P3: Async - Persist batch to disk
	if err := p.store.Persist(batch); err != nil {
		log.Printf("[PIPELINE-WORKER-%d] Error persisting batch: %v", workerID, err)
	}
}

// Stop closes the ingestion channel and waits for workers to finish
func (p *Processor) Stop() {
	close(p.IngestionChan)
	p.wg.Wait()
	log.Println("[PIPELINE] All workers stopped gracefully")
}

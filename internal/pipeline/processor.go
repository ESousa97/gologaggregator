package pipeline

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// BatchConfig defines the parameters for log batching
type BatchConfig struct {
	MaxSize     int
	MaxWaitTime time.Duration
	WorkerCount int
	BufferSize  int
}

// Processor manages the ingestion channel and worker pool
type Processor struct {
	IngestionChan chan string
	config        BatchConfig
	wg            sync.WaitGroup
}

// NewProcessor creates a new pipeline processor with defined configuration
func NewProcessor(cfg BatchConfig) *Processor {
	return &Processor{
		IngestionChan: make(chan string, cfg.BufferSize),
		config:        cfg,
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
	
	batch := make([]string, 0, p.config.MaxSize)
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
			
			batch = append(batch, msg)
			if len(batch) >= p.config.MaxSize {
				p.processBatch(id, batch)
				batch = make([]string, 0, p.config.MaxSize)
				ticker.Reset(p.config.MaxWaitTime)
			}

		case <-ticker.C:
			if len(batch) > 0 {
				p.processBatch(id, batch)
				batch = make([]string, 0, p.config.MaxSize)
			}
		}
	}
}

// processBatch simulates sending the batch to the next stage (e.g., storage, API)
func (p *Processor) processBatch(workerID int, batch []string) {
	fmt.Printf("[PIPELINE-WORKER-%d] Processing batch of %d messages\n", workerID, len(batch))
	// In a real scenario, this would send to a database or external service
	// for _, msg := range batch {
	//     // ... logic
	// }
}

// Stop closes the ingestion channel and waits for workers to finish
func (p *Processor) Stop() {
	close(p.IngestionChan)
	p.wg.Wait()
	log.Println("[PIPELINE] All workers stopped gracefully")
}

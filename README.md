<div align="center">
  <h1>gologaggregator</h1>
  <p>High-performance, multi-protocol log aggregator built in Go with concurrent ingestion and processing.</p>

  <img src="assets/github-go.png" alt="gologaggregator Banner" width="600px">

  <br>

[![Go Report Card](https://goreportcard.com/badge/github.com/ESousa97/gologaggregator)](https://goreportcard.com/report/github.com/ESousa97/gologaggregator)
[![Go Reference](https://pkg.go.dev/badge/github.com/ESousa97/gologaggregator.svg)](https://pkg.go.dev/github.com/ESousa97/gologaggregator)
[![License: MIT](https://img.shields.io/github/license/ESousa97/gologaggregator)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ESousa97/gologaggregator)](https://github.com/ESousa97/gologaggregator)
[![Last Commit](https://img.shields.io/github/last-commit/ESousa97/gologaggregator)](https://github.com/ESousa97/gologaggregator/commits/main)

</div>

---

`gologaggregator` is a high-performance log aggregation service designed to handle multiple ingestion protocols simultaneously. Built with scalability in mind, it utilizes Go's powerful concurrency primitives to ingest, process, and index logs with minimal latency and high resilience.

## Features

- **Multi-Protocol Ingestion**: Native support for TCP (raw text) and HTTP (JSON) protocols.
- **Concurrent Pipeline**: Efficient log processing using Worker Pools and buffered channels.
- **Real-time Processing**: High-speed log batching and dispatching logic.
- **Structured Indexing**: In-memory optimized storage with support for log level parsing.
- **Query Engine**: Powerful API for filtering and searching through aggregated logs.
- **Resilient Storage**: Persistence layer with Write-Ahead Logging (WAL) and log rotation.

## Tech Stack

| Technology | Role                                        |
| ---------- | ------------------------------------------- |
| Go 1.25+   | Core language and concurrent execution      |
| Net/HTTP   | Standard library for protocol implementation|
| Docker     | Containerization and deployment             |

## Architecture

The project follows strict modularization and clean architecture principles:

- **`internal/tcp`**: Raw TCP server for high-throughput text ingestion.
- **`internal/http`**: RESTful API for structured JSON ingestion and log searching.
- **`internal/pipeline`**: Concurrent log processing with Worker Pools and batching logic.
- **`internal/store`**: Thread-safe in-memory circular buffer with search capabilities.
- **`internal/parser`**: Log parsing logic to extract levels and messages.
- **`internal/persistence`**: Disk storage manager with automatic log rotation.
- **`internal/models`**: Domain models and log definitions.
- **`internal/config`**: Typed configuration and environment management.

## Installation and Usage

### From Source

```bash
git clone https://github.com/ESousa97/gologaggregator.git
cd gologaggregator
go run ./cmd/aggregator/main.go
```

## Roadmap

### Phase 1: Multi-Protocol Ingestion (Ingress) ✅
**Goal:** Demonstrate mastery over varied networks and protocols (TCP and HTTP).
- [x] TCP server on port 5000 (raw text).
- [x] HTTP server on port 8080 (JSON).
- [x] Extraction of content and timestamp using concurrent goroutines.

### Phase 2: Buffer and Processing Pipeline (The Pipeline) ✅
**Goal:** Prevent system lockup under heavy load (backpressure) using channels and buffers.
- [x] Implementation of an internal buffered `chan string`.
- [x] Worker Pool for log consumption.
- [x] Batching of 100 messages or every 5 seconds.

### Phase 3: Parsing and In-Memory Indexing (Indexing) ✅
**Goal:** Transform raw text into structured and searchable data.
- [x] Parser for log level identification (`INFO`, `ERROR`, `DEBUG`).
- [x] In-memory data structure protected by `sync.RWMutex`.
- [x] Optimized indexing for the last 10,000 logs.

### Phase 4: Search and Filtering API (The Query Engine) ✅
**Goal:** Allow users to query logs efficiently.
- [x] `GET /search` endpoint with level, time range, and keyword filters.
- [x] Structured JSON response format.
- [x] Efficient search logic over the in-memory index.

### Phase 5: Persistence and Resilience (The Storage Layer) ✅
**Goal:** Ensure logs are not lost if the service restarts (WAL).
- [x] Disk persistence using `.log` files.
- [x] Simple log rotation logic (10MB limit).
- [x] Dockerization and `docker-compose.yml` for load simulation.

## Contributing

Contributions are welcome! Please follow the engineering standards defined in the project.

## License

Distributed under the MIT License. See `LICENSE` for more details.

<div align="center">

## Author

**Enoque Sousa**

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=flat&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/enoque-sousa-bb89aa168/)
[![GitHub](https://img.shields.io/badge/GitHub-100000?style=flat&logo=github&logoColor=white)](https://github.com/ESousa97)
[![Portfolio](https://img.shields.io/badge/Portfolio-FF5722?style=flat&logo=target&logoColor=white)](https://enoquesousa.vercel.app)

**[⬆ Back to top](#gologaggregator)**

Made with ❤️ by [Enoque Sousa](https://github.com/ESousa97)

**Project Status:** Active — Study Project

</div>

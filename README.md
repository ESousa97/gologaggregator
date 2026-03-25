<div align="center">
  <h1>gologaggregator</h1>
  <p>High-performance, multi-protocol log aggregator built in Go with concurrent ingestion and processing.</p>

  <img src="assets/github-go.png" alt="gologaggregator Banner" width="600px">

  <br>

[![CI](https://img.shields.io/github/actions/workflow/status/ESousa97/gologaggregator/ci.yml?branch=main&label=CI&logo=github)](https://github.com/ESousa97/gologaggregator/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ESousa97/gologaggregator)](https://goreportcard.com/report/github.com/ESousa97/gologaggregator)
[![Go Reference](https://pkg.go.dev/badge/github.com/ESousa97/gologaggregator.svg)](https://pkg.go.dev/github.com/ESousa97/gologaggregator)
[![License: MIT](https://img.shields.io/github/license/ESousa97/gologaggregator?color=blue)](https://github.com/ESousa97/gologaggregator/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ESousa97/gologaggregator)](https://github.com/ESousa97/gologaggregator)
[![Last Commit](https://img.shields.io/github/last-commit/ESousa97/gologaggregator)](https://github.com/ESousa97/gologaggregator/commits/main)

</div>

---

`gologaggregator` is a high-performance log aggregation service that supports multiple ingestion protocols simultaneously. Designed for scalability, it leverages Go's concurrency primitives to ingest, process, and index logs with minimal latency and high resilience through disk persistence (WAL).

## Demonstration

### Ingestion via HTTP (JSON)
```bash
curl -X POST http://localhost:8080/logs \
  -H "Content-Type: application/json" \
  -d '{"content": "INFO: System initialized successfully"}'
```

### Ingestion via TCP (Plain Text)
```bash
echo "ERROR: Database connection failure" | nc localhost 5000
```

### Searching Logs
```bash
curl "http://localhost:8080/search?level=ERROR"
```

## Tech Stack

| Technology | Role |
|---|---|
| Go 1.25 | Core language and concurrent execution |
| Net/HTTP | Protocol implementation and REST API |
| Docker | Containerization and load simulation |
| Sync/Mutex | Thread-safety guarantee for in-memory index |

## Prerequisites

- Go >= 1.25 (defined in `go.mod`)
- Docker & Docker Compose (optional)
- Network tools (curl, nc/telnet)

## Installation and Usage

### As a Binary

```bash
go install github.com/ESousa97/gologaggregator/cmd/aggregator@latest
```

### From Source

```bash
git clone https://github.com/ESousa97/gologaggregator.git
cd gologaggregator
# Optional: cp .env.example .env
make build
make run
```

## Makefile Targets

| Target | Description |
|---|---|
| `make build` | Compiles the project binary to `bin/` |
| `make run` | Runs the aggregator directly via `go run` |
| `make test` | Runs all unit tests in the project |
| `make lint` | Runs the linter (golangci-lint or go vet) |
| `make docker-up` | Starts the full environment via docker-compose |
| `make docker-down` | Stops the docker-compose services |
| `make clean` | Removes build artifacts and temporary files |
| `make spell` | Runs spell checker (requires cspell) |

## Architecture

The project adopts a modular architecture focused on separation of concerns:

- **`internal/tcp` & `internal/http`**: Decoupled ingestion drivers.
- **`internal/pipeline`**: Core of the system, manages Worker Pools and log buffering.
- **`internal/parser`**: Logic for level extraction and message structuring.
- **`internal/store`**: Thread-safe in-memory index (Circular Buffer).
- **`internal/persistence`**: Disk writing layer with rotation support.

## API Reference

Detailed documentation of the API and internal packages can be found at [pkg.go.dev](https://pkg.go.dev/github.com/ESousa97/gologaggregator).

## Configuration

The application uses environment variables or defaults.

| Variable | Description | Type | Default |
|---|---|---|---|
| `SERVER_HOST` | Host for TCP/HTTP servers | string | `0.0.0.0` |
| `TCP_PORT` | Port for raw TCP ingestion | int | `5000` |
| `HTTP_PORT` | Port for REST API and HTTP ingestion | int | `8080` |

## Roadmap

- [x] Multi-Protocol Ingestion (TCP/HTTP)
- [x] Concurrent processing pipeline with Worker Pool
- [x] In-memory indexing with search and filtering
- [x] Structured Query API
- [x] Disk persistence and log rotation
- [ ] Web Interface for real-time visualization (Dashboard)
- [ ] Support for exports to S3/ElasticSearch

## Contributing

Contributions are welcome! See the full guide at [CONTRIBUTING.md](CONTRIBUTING.md).

## License

Distributed under the MIT license. See [LICENSE](LICENSE) for more details.

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

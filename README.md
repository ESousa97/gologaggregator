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
| Zod/ZAP    | (Planned) Logging and validation            |

## Architecture

The project follows strict modularization and clean architecture principles:

- **`internal/tcp`**: Raw TCP server for high-throughput text ingestion.
- **`internal/http`**: RESTful API for structured JSON log ingestion.
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

### Phase 1: Ingestão Multi-Protocolo (Ingress) ✅
**Objetivo:** Mostrar domínio sobre redes e protocolos variados (TCP e HTTP).
- [x] Servidor TCP na porta 5000 (raw text).
- [x] Servidor HTTP na porta 8080 (JSON).
- [x] Extração de conteúdo e timestamp simultâneos via goroutines.

### Phase 2: Buffer e Pipeline de Processamento (The Pipeline) ⏳
**Objetivo:** Evitar que o sistema trave sob carga pesada (backpressure) usando canais e buffers.
- [ ] Implementação de canal interno (`chan string`) com buffer.
- [ ] Worker Pool para consumo de logs.
- [ ] Agrupamento em lotes (batching) de 100 mensagens ou a cada 5 segundos.

### Phase 3: Parsing e Indexação em Memória (Indexing) ⏳
**Objetivo:** Transformar texto bruto em dados estruturados e pesquisáveis.
- [ ] Parser para identificação de níveis (`INFO`, `ERROR`, `DEBUG`).
- [ ] Estrutura de dados em memória protegida por `sync.RWMutex`.
- [ ] Indexação otimizada para os últimos 10.000 logs.

### Phase 4: API de Busca e Filtragem (The Query Engine) ⏳
**Objetivo:** Permitir que o usuário consulte os logs de forma eficiente.
- [ ] Endpoint `GET /search` com filtros de nível, tempo e keyword.
- [ ] Retorno em formato JSON estruturado.
- [ ] Lógica de busca eficiente sobre o índice em memória.

### Phase 5: Persistência e Resiliência (The Storage Layer) ⏳
**Objetivo:** Garantir que os logs não sumam se o serviço reiniciar (WAL).
- [ ] Persistência em disco com arquivos `.log`.
- [ ] Lógica de rotação de logs (limite de 10MB).
- [ ] Dockerização e `docker-compose.yml` para simulação de carga.

## Contributing

Contribuições são bem-vindas! Siga os padrões de engenharia definidos no projeto.

## License

Distribuído sob a licença MIT. Veja `LICENSE` para mais detalhes.

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

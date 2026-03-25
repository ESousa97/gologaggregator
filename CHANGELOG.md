# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-03-24

### Added
- Initial implementation of the log aggregator.
- Ingestion via TCP (port 5000) and HTTP (port 8080).
- Concurrent processing pipeline with Worker Pool.
- Thread-safe in-memory index using a Circular Buffer.
- Search API (`GET /search`) with filters by level and term.
- Disk persistence with support for log rotation.
- Dockerization and Docker Compose for orchestration.
- Makefile for task automation.
- Complete documentation (`README.md`, `CONTRIBUTING.md`, etc.).
- Doc comments (Godoc) for all exported code.

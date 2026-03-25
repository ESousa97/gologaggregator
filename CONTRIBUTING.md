# Contribution Guide

Thank you for your interest in contributing to `gologaggregator`! This project follows rigorous engineering standards to ensure performance and maintainability.

## Environment Setup

1. **Go 1.25**: Ensure you have at least this version installed.
2. **Dependencies**: The project uses only the Standard Library for the core, minimizing the dependency burden.
3. **Build Tools**: Use the included `Makefile` for common operations:
   ```bash
   make test
   make build
   ```

## Coding Style

- Follow the [Effective Go](https://golang.org/doc/effective_go.html) guidelines.
- Run `go fmt ./...` before committing.
- All exported (public) items **must** have documentation comments (Godoc).
- Keep functions short and with a single responsibility.

## Pull Request Process

1. Create a branch from `main` with a descriptive name (e.g., `feat/http-streaming` or `fix/tcp-deadlock`).
2. Ensure that `make test` passes locally.
3. Commits should follow the [Conventional Commits](https://www.conventionalcommits.org/) standard.
4. Open the PR clearly describing the changes and the reason behind them.

## Reporting Bugs

Use GitHub Issues to report problems. Be as detailed as possible, including:
- Steps to reproduce.
- Expected behavior vs. Actual behavior.
- Logs and environment info (OS, Go version).

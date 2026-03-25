# Project: gologaggregator
# Description: High-performance log aggregator built in Go.

BINARY_NAME=aggregator
MAIN_PATH=./cmd/aggregator/main.go

.PHONY: all build run test lint clean docker-up docker-down help spell

all: build

## spell: Run spell checker (requires cspell)
spell:
	@echo "Running spell checker..."
	@if command -v cspell > /dev/null; then \
		cspell "**/*.{go,md,txt,yml,yaml}"; \
	else \
		echo "Error: cspell is not installed. Run 'npm install -g cspell' to install it."; \
		exit 1; \
	fi

## build: Build the binary
build:
	@echo "Building binary..."
	go build -o bin/$(BINARY_NAME) $(MAIN_PATH)

## run: Run the application
run:
	@echo "Running application..."
	go run $(MAIN_PATH)

## test: Run tests
test:
	@echo "Running tests..."
	go test -v ./...

## lint: Run golangci-lint (if installed)
lint:
	@echo "Running linter..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		go vet ./...; \
	fi

## clean: Remove build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf bin/

## docker-up: Start services using docker-compose
docker-up:
	@echo "Starting Docker services..."
	docker-compose up -d

## docker-down: Stop services using docker-compose
docker-down:
	@echo "Stopping Docker services..."
	docker-compose down

## help: Show this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

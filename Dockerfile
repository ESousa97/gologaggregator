# Stage 1: Build
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o aggregator ./cmd/aggregator/main.go

# Stage 2: Final Image
FROM alpine:latest

WORKDIR /app

# Create logs directory
RUN mkdir logs

# Copy the binary from builder
COPY --from=builder /app/aggregator .

# Expose ports for TCP and HTTP
EXPOSE 8080 9090

# Run the application
CMD ["./aggregator"]

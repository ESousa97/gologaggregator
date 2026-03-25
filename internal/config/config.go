// Package config handles the application's configuration by loading and
// validating environment variables.
package config

import (
	"log"
	"os"
)

// AppConfig defines the configuration structure for the log aggregator.
// It contains the network addresses for both TCP and HTTP ingestion engines.
type AppConfig struct {
	// TCPAddress is the full host:port string for the TCP server.
	TCPAddress string
	// HTTPAddress is the full host:port string for the HTTP API server.
	HTTPAddress string
}

// Load fetches configurations from environment variables or defaults
func Load() *AppConfig {
	tcpPort := getEnv("TCP_PORT", "5000")
	httpPort := getEnv("HTTP_PORT", "8080")
	host := getEnv("SERVER_HOST", "0.0.0.0")

	cfg := &AppConfig{
		TCPAddress:  host + ":" + tcpPort,
		HTTPAddress: host + ":" + httpPort,
	}

	validate(cfg)
	return cfg
}

// getEnv retrieves environment variable or returns default
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// validate ensures all configuration fields are present and valid
func validate(cfg *AppConfig) {
	if cfg.TCPAddress == "" || cfg.HTTPAddress == "" {
		log.Fatal("Invalid configuration: Addresses cannot be empty")
	}
}

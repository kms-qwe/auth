package config

import (
	"time"

	"github.com/joho/godotenv"
)

// Load loads environment variables from a specified file.
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

// GRPCConfig defines the configuration for gRPC.
type GRPCConfig interface {
	Address() string
}

// PGConfig defines the configuration for PostgreSQL.
type PGConfig interface {
	DSN() string
}

type RedisConfig interface {
	Address() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
	TTL() time.Duration
}

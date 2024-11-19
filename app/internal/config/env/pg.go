package env

import (
	"errors"
	"os"

	"github.com/kms-qwe/auth/internal/config"
)

const (
	dsnEnvName = "PG_DSN"
)

type pgConfig struct {
	dsn string
}

// DSN provides postgres DSN
func (p *pgConfig) DSN() string {
	return p.dsn
}

// NewPGConfig creates a new PostgreSQL configuration based on environment variables.
func NewPGConfig() (config.PGConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("pg dsn not found")
	}
	return &pgConfig{
		dsn: dsn,
	}, nil
}

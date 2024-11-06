package pgv1

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	pgClient "github.com/kms-qwe/auth/internal/client/postgres"
)

type client struct {
	masterDBC pgClient.DB
}

func NewPgClient(ctx context.Context, dsn string) (pgClient.Client, error) {
	dbc, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	return &client{
		masterDBC: NewDB(dbc),
	}, nil
}

func (c *client) DB() pgClient.DB {
	return c.masterDBC
}
func (c *client) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}

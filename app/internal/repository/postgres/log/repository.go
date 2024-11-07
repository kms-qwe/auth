package user

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	pgClient "github.com/kms-qwe/auth/internal/client/postgres"
	"github.com/kms-qwe/auth/internal/repository"
)

const (
	tableName = "log"

	idColumn      = "id"
	messageColumn = "message"
	logTimeColumn = "log_time"
)

type repo struct {
	db pgClient.Client
}

func NewLogRepository(db pgClient.Client) repository.LogRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Log(ctx context.Context, log string) error {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(messageColumn).
		Values(log)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to create query: %w", err)
	}

	q := pgClient.Query{
		Name:     "log_repository.Log",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to create log: %w", err)
	}

	return nil
}

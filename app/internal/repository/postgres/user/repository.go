package user

import (
	"context"
	"fmt"
	"log"

	"github.com/kms-qwe/auth/internal/model"
	"github.com/kms-qwe/auth/internal/repository"
	"github.com/kms-qwe/auth/internal/repository/postgres/user/converter"
	modelRepo "github.com/kms-qwe/auth/internal/repository/postgres/user/model"
	pgClient "github.com/kms-qwe/platform_common/pkg/client/postgres"
)

type repo struct {
	db pgClient.Client
}

// Create adds user to db
func (r *repo) Create(ctx context.Context, info *model.UserInfo) (int64, error) {

	query, args, err := queryCreateUser(ctx, info)
	if err != nil {
		return 0, fmt.Errorf("failed to build query: %w", err)
	}

	q := pgClient.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to scan id: %w", err)
	}

	return id, nil
}

// Get select user from db
func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {

	query, args, err := queryGetUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	q := pgClient.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	repoUser := &modelRepo.User{}

	err = r.db.DB().ScanOneContext(ctx, repoUser, q, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return converter.ToUserFromRepo(repoUser), nil
}

// Update updates user info in db
func (r *repo) Update(ctx context.Context, userInfoUpdate *model.UserInfoUpdate) error {

	query, args, err := queryUpdateUser(ctx, userInfoUpdate)
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	q := pgClient.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	log.Printf("updated %d rows\n", res.RowsAffected())

	return nil
}

// Delete deletes user from db
func (r *repo) Delete(ctx context.Context, id int64) error {

	query, args, err := queryDeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	q := pgClient.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	log.Printf("delete %d rowd\n", res.RowsAffected())
	return nil
}

// NewUserRepository initializes a new PostgreSQL storage instance using the provided DSN.
func NewUserRepository(pgClient pgClient.Client) repository.UserRepository {

	return &repo{
		db: pgClient,
	}
}

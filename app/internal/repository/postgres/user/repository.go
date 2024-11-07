package user

import (
	"context"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	pgClient "github.com/kms-qwe/auth/internal/client/postgres"
	"github.com/kms-qwe/auth/internal/model"
	"github.com/kms-qwe/auth/internal/repository"
	"github.com/kms-qwe/auth/internal/repository/postgres/user/converter"
	modelRepo "github.com/kms-qwe/auth/internal/repository/postgres/user/model"
)

const (
	tableName = "users"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	passwordColumn  = "password"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db pgClient.Client
}

// Create adds user to db
func (r *repo) Create(ctx context.Context, info *model.UserInfo) (int64, error) {
	repoInfo := converter.ToRepoFromUserInfo(info)

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn).
		Values(repoInfo.Email, repoInfo.Email, repoInfo.Password, repoInfo.Role).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
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
	builder := sq.Select(idColumn, nameColumn, emailColumn, passwordColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
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
	repoUserInfoUpdate := converter.ToRepoFromUserInfoUpdate(userInfoUpdate)

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(nameColumn, repoUserInfoUpdate.Name).
		Set(emailColumn, repoUserInfoUpdate.Email).
		Set(roleColumn, repoUserInfoUpdate.Role).
		Set(updatedAtColumn, sq.Expr("NOW()")).
		Where(sq.Eq{"id": repoUserInfoUpdate.ID})

	query, args, err := builder.ToSql()
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
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
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

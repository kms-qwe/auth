package postgres

import (
	"context"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kms-qwe/microservices_course_auth/internal/model"
	"github.com/kms-qwe/microservices_course_auth/internal/storage"
)

type pgStorage struct {
	pool *pgxpool.Pool
}

func (pg *pgStorage) AddNewUser(ctx context.Context, info *model.UserInfo) (int64, error) {
	builderInsert := sq.Insert("userV1.user").
		PlaceholderFormat(sq.Dollar).
		Columns("name", "email", "password", "role").
		Values(info.Name, info.Email, info.Password, info.Role).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	var userID int64
	err = pg.pool.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		return 0, err
	}

	log.Printf("inserted user with id: %d\n", userID)

	return userID, nil
}
func (pg *pgStorage) GetUser(ctx context.Context, id int64) (*model.User, error) {
	builderSelectOne := sq.Select("name", "email", "password", "role", "created_at", "updated_at").
		From("userV1.user").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		return nil, err
	}
	var name, email, password string
	var role int32
	var createdAt time.Time
	var updatedAt *time.Time

	err = pg.pool.QueryRow(ctx, query, args...).Scan(&name, &email, &password, &role, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, name: %s, email: %s, password: %s, role: %d, created_at: %s, updated_at: %v\n", id, name, email, password, role, createdAt, updatedAt)

	return &model.User{
		ID: id,
		Info: &model.UserInfo{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     role,
		},
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
func (pg *pgStorage) UpdateUserInfo(ctx context.Context, id int64, name, email string, role int32) error {
	builderUpdate := sq.Update("userV1.user").
		PlaceholderFormat(sq.Dollar).
		Set("name", name).
		Set("email", email).
		Set("role", role).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": id})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	res, err := pg.pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	log.Printf("updated %d rows\n", res.RowsAffected())

	return nil
}
func (pg *pgStorage) DeleteUser(ctx context.Context, id int64) error {
	builderUpdate := sq.Delete("userV1.user").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}
	res, err := pg.pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	log.Printf("deleted %d rows\n", res.RowsAffected())

	return nil
}

// NewPgStorage initializes a new PostgreSQL storage instance using the provided DSN.
func NewPgStorage(ctx context.Context, DSN string) (storage.Storage, error) {
	pool, err := pgxpool.New(ctx, DSN)
	if err != nil {
		return nil, err
	}
	return &pgStorage{
		pool: pool,
	}, nil
}

package storage

import (
	"context"

	"github.com/kms-qwe/auth/internal/model"
)

// Storage interface defines methods for user data storage operations.
type Storage interface {
	AddNewUser(ctx context.Context, info *model.UserInfo) (int64, error)
	GetUser(ctx context.Context, id int64) (*model.User, error)
	UpdateUserInfo(ctx context.Context, id int64, name, email string, role int32) error
	DeleteUser(ctx context.Context, id int64) error
}

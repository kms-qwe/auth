package storage

import (
	"context"

	"github.com/kms-qwe/microservices_course_auth/internal/model"
)

type Storage interface {
	AddNewUser(ctx context.Context, info *model.UserInfo) (int64, error)
	GetUser(ctx context.Context, id int64) (*model.User, error)
	UpdateUserInfo(ctx context.Context, id int64, name, email string, role int32) error
	DeleteUser(ctx context.Context, id int64) error
}

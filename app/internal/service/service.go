package service

import (
	"context"

	"github.com/kms-qwe/auth/internal/model"
)

type UserService interface {
	Create(ctx context.Context, info *model.UserInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, info *model.UserInfoUpdate) error
	Delete(ctx context.Context, id int64) error
}

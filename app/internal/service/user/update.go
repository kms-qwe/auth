package user

import (
	"context"

	"github.com/kms-qwe/auth/internal/model"
)

func (s *serv) Update(ctx context.Context, info *model.UserInfoUpdate) error {
	return s.userRepository.Update(ctx, info)
}

package user

import (
	"context"
	"fmt"

	"github.com/kms-qwe/auth/internal/model"
)

func (s *serv) Update(ctx context.Context, info *model.UserInfoUpdate) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		errTx = s.userRepository.Update(ctx, info)
		if errTx != nil {
			return errTx
		}

		errTx = s.logRepository.Log(ctx, fmt.Sprintf("user updated:%#v", info))
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

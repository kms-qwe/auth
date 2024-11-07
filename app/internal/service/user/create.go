package user

import (
	"context"
	"fmt"

	"github.com/kms-qwe/auth/internal/model"
)

func (s *serv) Create(ctx context.Context, info *model.UserInfo) (int64, error) {

	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.userRepository.Create(ctx, info)
		if errTx != nil {
			return errTx
		}

		errTx = s.logRepository.Log(ctx, fmt.Sprintf("create user: %#v", info))
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, nil
	}

	return id, nil
}

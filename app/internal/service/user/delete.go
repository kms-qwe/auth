package user

import (
	"context"
	"errors"
	"fmt"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.userRepository.Delete(ctx, id)
		if errTx != nil {
			return errTx
		}

		errTx = errors.New("fail on delete migration ")
		if errTx != nil {
			return errTx
		}

		errTx = s.logRepository.Log(ctx, fmt.Sprintf("user deleted: %d", id))
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

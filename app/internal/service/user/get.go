package user

import (
	"context"
	"fmt"

	"github.com/kms-qwe/auth/internal/model"
)

// Get gets a new user using the provided id
func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	var user = &model.User{}
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		user, errTx = s.userRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		fmt.Printf("user after GET: %#v", user)

		errTx = s.logRepository.Log(ctx, fmt.Sprintf("get user: %#v", user))
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	fmt.Printf("ITOG USER %#v\n", user)
	return user, nil
}

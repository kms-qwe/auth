package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/kms-qwe/auth/internal/cache"
	"github.com/kms-qwe/auth/internal/model"
	"github.com/kms-qwe/auth/internal/repository"
	"github.com/kms-qwe/auth/internal/service"

	pgClient "github.com/kms-qwe/platform_common/pkg/client/postgres"
)

type serv struct {
	userRepository repository.UserRepository
	logRepository  repository.LogRepository
	txManager      pgClient.TxManager

	userCache cache.UserCache
}

// NewUserService creates new a UserService with provided  UserRepository LogRepository TxManager
func NewUserService(
	userRepository repository.UserRepository,
	logRepository repository.LogRepository,
	txManager pgClient.TxManager,
	userCache cache.UserCache,
) service.UserService {
	return &serv{
		userRepository: userRepository,
		logRepository:  logRepository,
		txManager:      txManager,
		userCache:      userCache,
	}
}

// Create creates a new user using the provided user model
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
		return 0, err
	}

	return id, nil
}

// Delete delete a new user using the provided id
func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.userRepository.Delete(ctx, id)
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

	err = s.userCache.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// Get gets a new user using the provided id
func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {

	user, err := s.userCache.Get(ctx, id)
	if err == nil {
		return user, nil
	}
	if !errors.Is(err, model.ErrorUserNotFound) {
		return nil, err
	}
	user = &model.User{}
	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
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

	err = s.userCache.Set(ctx, user)
	if err != nil {
		return nil, err
	}

	fmt.Printf("ITOG USER %#v\n", user)

	return user, nil
}

// Update updates a new user using the provided update info
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

package user

import (
	pgClient "github.com/kms-qwe/auth/internal/client/postgres"
	"github.com/kms-qwe/auth/internal/repository"
	"github.com/kms-qwe/auth/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
	txManager      pgClient.TxManager
}

func NewUserService(
	userRepository repository.UserRepository,
	txManager pgClient.TxManager,
) service.UserService {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
	}
}

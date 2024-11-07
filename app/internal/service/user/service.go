package user

import (
	pgClient "github.com/kms-qwe/auth/internal/client/postgres"
	"github.com/kms-qwe/auth/internal/repository"
	"github.com/kms-qwe/auth/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
	logRepository  repository.LogRepository
	txManager      pgClient.TxManager
}

func NewUserService(
	userRepository repository.UserRepository,
	logRepository repository.LogRepository,
	txManager pgClient.TxManager,
) service.UserService {
	return &serv{
		userRepository: userRepository,
		logRepository:  logRepository,
		txManager:      txManager,
	}
}

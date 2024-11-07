package app

import (
	"context"
	"log"

	"github.com/kms-qwe/auth/internal/api/grpc/user"
	"github.com/kms-qwe/auth/internal/client/postgres"
	pgv1 "github.com/kms-qwe/auth/internal/client/postgres/pg_v1"
	"github.com/kms-qwe/auth/internal/client/postgres/transaction"
	"github.com/kms-qwe/auth/internal/closer"
	"github.com/kms-qwe/auth/internal/config"
	"github.com/kms-qwe/auth/internal/config/env"
	"github.com/kms-qwe/auth/internal/repository"
	logpg "github.com/kms-qwe/auth/internal/repository/postgres/log"
	usepg "github.com/kms-qwe/auth/internal/repository/postgres/user"
	"github.com/kms-qwe/auth/internal/service"
	useserv "github.com/kms-qwe/auth/internal/service/user"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	pgClient       postgres.Client
	txManager      postgres.TxManager
	userRepository repository.UserRepository
	logRepository  repository.LogRepository

	userService service.UserService

	userImpl *user.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}
func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get postgres config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) PGClient(ctx context.Context) postgres.Client {
	if s.pgClient == nil {
		pgClient, err := pgv1.NewPgClient(ctx, s.PGConfig().DSN())

		if err != nil {
			log.Fatalf("failed to create pg client: %v", err)
		}

		err = pgClient.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		s.pgClient = pgClient

		closer.Add(s.pgClient.Close)
	}

	return s.pgClient
}

func (s *serviceProvider) TxManager(ctx context.Context) postgres.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.PGClient(ctx).DB())

	}

	return s.txManager
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = usepg.NewUserRepository(s.PGClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) LogRepository(ctx context.Context) repository.LogRepository {
	if s.logRepository == nil {
		s.logRepository = logpg.NewLogRepository(s.PGClient(ctx))
	}

	return s.logRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = useserv.NewUserService(s.UserRepository(ctx), s.LogRepository(ctx), s.TxManager(ctx))
	}

	return s.userService
}

func (s *serviceProvider) UserImpl(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}

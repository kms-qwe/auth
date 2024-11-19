package app

import (
	"context"
	"log"

	"github.com/kms-qwe/auth/internal/api/grpc/user"
	"github.com/kms-qwe/auth/internal/cache"
	"github.com/kms-qwe/platform_common/pkg/client/cache/redis"
	"github.com/kms-qwe/platform_common/pkg/client/postgres"
	pg "github.com/kms-qwe/platform_common/pkg/client/postgres/pg"

	redigo "github.com/gomodule/redigo/redis"
	userCache "github.com/kms-qwe/auth/internal/cache/redis/user"
	"github.com/kms-qwe/auth/internal/config"
	"github.com/kms-qwe/auth/internal/config/env"
	"github.com/kms-qwe/auth/internal/repository"
	logpg "github.com/kms-qwe/auth/internal/repository/postgres/log"
	userPg "github.com/kms-qwe/auth/internal/repository/postgres/user"
	"github.com/kms-qwe/auth/internal/service"
	userServ "github.com/kms-qwe/auth/internal/service/user"
	cacheClient "github.com/kms-qwe/platform_common/pkg/client/cache"
	"github.com/kms-qwe/platform_common/pkg/client/postgres/transaction"
	"github.com/kms-qwe/platform_common/pkg/closer"
)

type serviceProvider struct {
	pgConfig    config.PGConfig
	grpcConfig  config.GRPCConfig
	redisConfig config.RedisConfig

	pgClient       postgres.Client
	txManager      postgres.TxManager
	userRepository repository.UserRepository
	logRepository  repository.LogRepository

	redisPool   *redigo.Pool
	redisClient cacheClient.RedisCache
	userCache   cache.UserCache

	userService service.UserService

	userGrpcHandlers *user.GrpcHandlers
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// PGConfig provides pgconfig
func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Panicf("failed to get postgres config: %s", err.Error())
		}

		s.pgConfig = cfg
	}
	return s.pgConfig
}

// GRPCConfig provides grpc config
func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Panicf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

// RedisConfig provides redis config
func (s *serviceProvider) RedisConfig() config.RedisConfig {
	if s.redisConfig == nil {
		cfg, err := env.NewRedisConfig()
		if err != nil {
			log.Panicf("failed to get redis config: %s", err.Error())
		}

		s.redisConfig = cfg
	}

	return s.redisConfig
}

// PGClient provides pg client
func (s *serviceProvider) PGClient(ctx context.Context) postgres.Client {
	if s.pgClient == nil {
		pgClient, err := pg.NewPgClient(ctx, s.PGConfig().DSN())

		if err != nil {
			log.Panicf("failed to create pg client: %v", err)
		}

		err = pgClient.DB().Ping(ctx)
		if err != nil {
			log.Panicf("ping error: %s", err.Error())
		}
		s.pgClient = pgClient

		closer.Add(s.pgClient.Close)
	}

	return s.pgClient
}

// TxManager provides tx manager
func (s *serviceProvider) TxManager(ctx context.Context) postgres.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.PGClient(ctx).DB())

	}

	return s.txManager
}

// UserRepository provides User Repository
func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userPg.NewUserRepository(s.PGClient(ctx))
	}

	return s.userRepository
}

// LogRepository provides  LogRepository
func (s *serviceProvider) LogRepository(ctx context.Context) repository.LogRepository {
	if s.logRepository == nil {
		s.logRepository = logpg.NewLogRepository(s.PGClient(ctx))
	}

	return s.logRepository
}

// RedisPool provides RedisPool
func (s *serviceProvider) RedisPool() *redigo.Pool {
	if s.redisPool == nil {
		s.redisPool = &redigo.Pool{
			MaxIdle:     s.RedisConfig().MaxIdle(),
			IdleTimeout: s.RedisConfig().IdleTimeout(),
			DialContext: func(ctx context.Context) (redigo.Conn, error) {
				return redigo.DialContext(ctx, "tcp", s.RedisConfig().Address())
			},
		}
	}

	return s.redisPool
}

// RedisClient provides RedisClient
func (s *serviceProvider) RedisClient() cacheClient.RedisCache {
	if s.redisClient == nil {
		s.redisClient = redis.NewClient(s.RedisPool(), s.RedisConfig().ConnectionTimeout())
	}

	return s.redisClient
}

// UserCache provides UserCache
func (s *serviceProvider) UserCache() cache.UserCache {
	if s.userCache == nil {
		s.userCache = userCache.NewUserCache(s.RedisClient(), s.RedisConfig().TTL())
	}

	return s.userCache
}

// UserService provides UserService
func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userServ.NewUserService(s.UserRepository(ctx), s.LogRepository(ctx), s.TxManager(ctx), s.UserCache())
	}

	return s.userService
}

// UserGrpcHandlers provides UserGrpcHandlers
func (s *serviceProvider) UserGrpcHandlers(ctx context.Context) *user.GrpcHandlers {
	if s.userGrpcHandlers == nil {
		s.userGrpcHandlers = user.NewUserGrpcHandlers(s.UserService(ctx))
	}

	return s.userGrpcHandlers
}

package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/kms-qwe/auth/internal/cache"
	cacheMock "github.com/kms-qwe/auth/internal/cache/mocks"
	"github.com/kms-qwe/auth/internal/repository"
	repositoryMock "github.com/kms-qwe/auth/internal/repository/mocks"
	"github.com/kms-qwe/auth/internal/service/user"
	pgClient "github.com/kms-qwe/platform_common/pkg/client/postgres"
	pgClientMock "github.com/kms-qwe/platform_common/pkg/client/postgres/mocks"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type logRepositoryMockFunc func(mc *minimock.Controller) repository.LogRepository
	type txManagerMockFunc func(mc *minimock.Controller) pgClient.TxManager
	type userCacheMockFunc func(mc *minimock.Controller) cache.UserCache

	type args struct {
		ctx context.Context
		req int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()
		// name     = gofakeit.Name()
		// email    = gofakeit.Email()
		// password = gofakeit.Password(true, true, true, true, false, 12)
		// role     = constant.Role(gofakeit.Number(0, 3))
		// createdAt = gofakeit.Date()
		// updatedAt = gofakeit.Date()

		userRepoErr  = fmt.Errorf("user repo error")
		logRepoErr   = fmt.Errorf("log repo error")
		txManagerErr = fmt.Errorf("tx manager error")
		cacheErr     = fmt.Errorf("cache error")

		reqCorrect = id

		reqEmpty = int64(0)

		logCorrect = fmt.Sprintf("user deleted: %d", reqCorrect)
		logEmpty   = fmt.Sprintf("user deleted: %d", reqEmpty)
	)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		userRepositoryMock userRepositoryMockFunc
		logRepositoryMock  logRepositoryMockFunc
		txManagerMock      txManagerMockFunc
		cacheMock          userCacheMockFunc
	}{
		{
			name: "t1: succes case",
			args: args{
				ctx: ctx,
				req: reqCorrect,
			},
			want: id,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMock.NewUserRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, reqCorrect).Return(nil)
				return mock
			},
			logRepositoryMock: func(mc *minimock.Controller) repository.LogRepository {
				mock := repositoryMock.NewLogRepositoryMock(mc)
				mock.LogMock.Expect(ctx, logCorrect).Return(nil)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) pgClient.TxManager {
				mock := pgClientMock.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f pgClient.Handler) error {
					return f(ctx)
				})
				return mock
			},
			cacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMock.NewUserCacheMock(mc)
				mock.DeleteMock.Expect(ctx, reqCorrect).Return(nil)
				return mock
			},
		},

		{
			name: "t2: empty case",
			args: args{
				ctx: ctx,
				req: reqEmpty,
			},
			want: id,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMock.NewUserRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, reqEmpty).Return(nil)
				return mock
			},
			logRepositoryMock: func(mc *minimock.Controller) repository.LogRepository {
				mock := repositoryMock.NewLogRepositoryMock(mc)
				mock.LogMock.Expect(ctx, logEmpty).Return(nil)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) pgClient.TxManager {
				mock := pgClientMock.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f pgClient.Handler) error {
					return f(ctx)
				})
				return mock
			},
			cacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMock.NewUserCacheMock(mc)
				mock.DeleteMock.Expect(ctx, reqEmpty).Return(nil)
				return mock
			},
		},

		{
			name: "t3: user repo error case",
			args: args{
				ctx: ctx,
				req: reqCorrect,
			},
			want: 0,
			err:  userRepoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMock.NewUserRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, reqCorrect).Return(userRepoErr)
				return mock
			},
			logRepositoryMock: func(mc *minimock.Controller) repository.LogRepository {
				mock := repositoryMock.NewLogRepositoryMock(mc)
				// mock.LogMock.Expect(ctx, logCorrect).Return(nil)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) pgClient.TxManager {
				mock := pgClientMock.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f pgClient.Handler) error {
					return f(ctx)
				})
				return mock
			},
			cacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMock.NewUserCacheMock(mc)
				return mock
			},
		},

		{
			name: "t4: log repo error case",
			args: args{
				ctx: ctx,
				req: reqCorrect,
			},
			want: 0,
			err:  logRepoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMock.NewUserRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, reqCorrect).Return(nil)
				return mock
			},
			logRepositoryMock: func(mc *minimock.Controller) repository.LogRepository {
				mock := repositoryMock.NewLogRepositoryMock(mc)
				mock.LogMock.Expect(ctx, logCorrect).Return(logRepoErr)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) pgClient.TxManager {
				mock := pgClientMock.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f pgClient.Handler) error {
					return f(ctx)
				})
				return mock
			},
			cacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMock.NewUserCacheMock(mc)
				return mock
			},
		},

		{
			name: "t5: tx manager error case",
			args: args{
				ctx: ctx,
				req: reqCorrect,
			},
			want: 0,
			err:  txManagerErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMock.NewUserRepositoryMock(mc)
				return mock
			},
			logRepositoryMock: func(mc *minimock.Controller) repository.LogRepository {
				mock := repositoryMock.NewLogRepositoryMock(mc)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) pgClient.TxManager {
				mock := pgClientMock.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f pgClient.Handler) error {
					return txManagerErr
				})
				return mock
			},
			cacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMock.NewUserCacheMock(mc)
				return mock
			},
		},
		{
			name: "t6: cache error case",
			args: args{
				ctx: ctx,
				req: reqCorrect,
			},
			want: 0,
			err:  cacheErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMock.NewUserRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, reqCorrect).Return(nil)
				return mock
			},
			logRepositoryMock: func(mc *minimock.Controller) repository.LogRepository {
				mock := repositoryMock.NewLogRepositoryMock(mc)
				mock.LogMock.Expect(ctx, logCorrect).Return(nil)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) pgClient.TxManager {
				mock := pgClientMock.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f pgClient.Handler) error {
					return f(ctx)
				})
				return mock
			},
			cacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMock.NewUserCacheMock(mc)
				mock.DeleteMock.Expect(ctx, reqCorrect).Return(cacheErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userService := user.NewUserService(tt.userRepositoryMock(mc), tt.logRepositoryMock(mc), tt.txManagerMock(mc), tt.cacheMock(mc))

			err := userService.Delete(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}

}

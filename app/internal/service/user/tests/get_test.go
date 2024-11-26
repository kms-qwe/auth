package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/kms-qwe/auth/internal/cache"
	cacheMock "github.com/kms-qwe/auth/internal/cache/mocks"
	"github.com/kms-qwe/auth/internal/constant"
	"github.com/kms-qwe/auth/internal/model"
	"github.com/kms-qwe/auth/internal/repository"
	repositoryMock "github.com/kms-qwe/auth/internal/repository/mocks"
	"github.com/kms-qwe/auth/internal/service/user"
	pgClient "github.com/kms-qwe/platform_common/pkg/client/postgres"
	pgClientMock "github.com/kms-qwe/platform_common/pkg/client/postgres/mocks"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
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

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		password  = gofakeit.Password(true, true, true, true, false, 12)
		role      = constant.Role(gofakeit.Int32() % 3)
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		userRepoErr    = fmt.Errorf("user repo error")
		logRepoErr     = fmt.Errorf("log repo error")
		txManagerErr   = fmt.Errorf("tx manager error")
		cacheGetErr    = fmt.Errorf("cache get error")
		cacheSetErr    = fmt.Errorf("cache set error")
		cacheExpireErr = fmt.Errorf("cache expire error")

		reqCorrect = id

		reqEmpty = int64(0)

		resCorrect = &model.User{
			ID: id,
			Info: &model.UserInfo{
				Name:            name,
				Email:           email,
				Password:        password,
				PasswordConfirm: password,
				Role:            role,
			},
			CreatedAt: createdAt,
			UpdatedAt: &updatedAt,
		}

		resEmpty = &model.User{}

		logCorrect = fmt.Sprintf("get user: %#v", resCorrect)
	)

	tests := []struct {
		name               string
		args               args
		want               *model.User
		err                error
		userRepositoryMock userRepositoryMockFunc
		logRepositoryMock  logRepositoryMockFunc
		txManagerMock      txManagerMockFunc
		cacheMock          userCacheMockFunc
	}{
		{
			name: "t1: succes case from cache",
			args: args{
				ctx: ctx,
				req: reqCorrect,
			},
			want: resCorrect,
			err:  nil,
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
				return mock
			},
			cacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMock.NewUserCacheMock(mc)
				mock.GetMock.Expect(ctx, reqCorrect).Return(resCorrect, nil)
				return mock
			},
		},
		{
			name: "t1_1: succes case from repo",
			args: args{
				ctx: ctx,
				req: reqCorrect,
			},
			want: resCorrect,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMock.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, reqCorrect).Return(resCorrect, nil)
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
				mock.GetMock.Expect(ctx, reqCorrect).Return(nil, model.ErrorUserNotFound)
				mock.SetMock.Expect(ctx, resCorrect).Return(nil)
				mock.ExpireMock.Expect(ctx, resCorrect.ID, 0).Return(nil)

				return mock
			},
		},
		{
			name: "t2: empty case",
			args: args{
				ctx: ctx,
				req: reqEmpty,
			},
			want: resEmpty,
			err:  nil,
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
				return mock
			},
			cacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMock.NewUserCacheMock(mc)
				mock.GetMock.Expect(ctx, reqEmpty).Return(resEmpty, nil)
				return mock
			},
		},

		{
			name: "t3: repo error case",
			args: args{
				ctx: ctx,
				req: reqCorrect,
			},
			want: nil,
			err:  userRepoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMock.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, reqCorrect).Return(nil, userRepoErr)
				return mock
			},
			logRepositoryMock: func(mc *minimock.Controller) repository.LogRepository {
				mock := repositoryMock.NewLogRepositoryMock(mc)
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
				mock.GetMock.Expect(ctx, reqCorrect).Return(nil, model.ErrorUserNotFound)
				return mock
			},
		},
		{
			name: "t4: log error case",
			args: args{
				ctx: ctx,
				req: reqCorrect,
			},
			want: nil,
			err:  logRepoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMock.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, reqCorrect).Return(resCorrect, nil)
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
				mock.GetMock.Expect(ctx, reqCorrect).Return(nil, model.ErrorUserNotFound)
				return mock
			},
		},
		{
			name: "4: tx manager error case",
			args: args{
				ctx: ctx,
				req: reqCorrect,
			},
			want: nil,
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
				mock.GetMock.Expect(ctx, reqCorrect).Return(nil, model.ErrorUserNotFound)
				return mock
			},
		},
		{
			name: "t6: cache get error case",
			args: args{
				ctx: ctx,
				req: reqCorrect,
			},
			want: nil,
			err:  cacheGetErr,
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
				return mock
			},
			cacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMock.NewUserCacheMock(mc)
				mock.GetMock.Expect(ctx, reqCorrect).Return(nil, cacheGetErr)
				return mock
			},
		},
		{
			name: "t6: cache set error case",
			args: args{
				ctx: ctx,
				req: reqCorrect,
			},
			want: nil,
			err:  cacheSetErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMock.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, reqCorrect).Return(resCorrect, nil)
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
				mock.GetMock.Expect(ctx, reqCorrect).Return(nil, model.ErrorUserNotFound)
				mock.SetMock.Expect(ctx, resCorrect).Return(cacheSetErr)
				return mock
			},
		},
		{
			name: "t7: cache expire error case",
			args: args{
				ctx: ctx,
				req: reqCorrect,
			},
			want: nil,
			err:  cacheExpireErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMock.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, reqCorrect).Return(resCorrect, nil)
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
				mock.GetMock.Expect(ctx, reqCorrect).Return(nil, model.ErrorUserNotFound)
				mock.SetMock.Expect(ctx, resCorrect).Return(nil)
				mock.ExpireMock.Expect(ctx, resCorrect.ID, 0).Return(cacheExpireErr)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userService := user.NewUserService(tt.userRepositoryMock(mc), tt.logRepositoryMock(mc), tt.txManagerMock(mc), tt.cacheMock(mc))

			res, err := userService.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}

}

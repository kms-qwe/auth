package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/kms-qwe/auth/internal/api/grpc/user"
	"github.com/kms-qwe/auth/internal/constant"
	"github.com/kms-qwe/auth/internal/model"
	"github.com/kms-qwe/auth/internal/service"
	serviceMocks "github.com/kms-qwe/auth/internal/service/mocks"
	desc "github.com/kms-qwe/auth/pkg/user_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestUpdate(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		name  = gofakeit.Name()
		email = gofakeit.Email()
		role  = desc.Role(gofakeit.Number(0, 10))

		serviceErr = fmt.Errorf("service error")

		req1 = &desc.UpdateRequest{
			UserUpdate: &desc.UserInfoUpdate{
				Id:    id,
				Name:  wrapperspb.String(name),
				Email: wrapperspb.String(email),
				Role:  role,
			},
		}

		serviceUserUpdateInfo1 = &model.UserInfoUpdate{
			ID:    id,
			Name:  &name,
			Email: &email,
			Role:  constant.Role(role),
		}

		req2 = &desc.UpdateRequest{
			UserUpdate: &desc.UserInfoUpdate{
				Id:    id,
				Name:  nil,
				Email: nil,
				Role:  role,
			},
		}

		serviceUserUpdateInfo2 = &model.UserInfoUpdate{
			ID:    id,
			Name:  nil,
			Email: nil,
			Role:  constant.Role(role),
		}

		res = &emptypb.Empty{}
	)

	// t.Cleanup(mc.Finish) не нужен с новой версией minimock

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "succes case",
			args: args{
				ctx: ctx,
				req: req1,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, serviceUserUpdateInfo1).Return(nil)
				return mock
			},
		},
		{
			name: "nil email and email",
			args: args{
				ctx: ctx,
				req: req2,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, serviceUserUpdateInfo2).Return(nil)
				return mock
			},
		},
		{
			// репрезентует случай, когда сервис выдает ошибку (пока просто заглушка)
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req1,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, serviceUserUpdateInfo1).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := user.NewUserGrpcHandlers(userServiceMock)

			res, err := api.Update(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}

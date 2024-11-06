package user

import (
	"context"
	"log"

	"github.com/kms-qwe/auth/internal/converter"
	desc "github.com/kms-qwe/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := i.userService.Update(ctx, converter.ToUserInfoUpdateFromDesc(req.UserUpdate))
	if err != nil {
		return nil, err
	}

	log.Printf("update user: id: %d, name: %s, email: %s, role: %d\n", req.UserUpdate.Id, req.UserUpdate.Name, req.UserUpdate.Email, req.UserUpdate.Role)

	return &emptypb.Empty{}, nil
}

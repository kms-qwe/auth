package user

import (
	"context"
	"log"

	desc "github.com/kms-qwe/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.userService.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	log.Printf("delete user with id: %d", req.Id)

	return &emptypb.Empty{}, nil
}

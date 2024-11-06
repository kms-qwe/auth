package user

import (
	"github.com/kms-qwe/auth/internal/service"
	desc "github.com/kms-qwe/auth/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}

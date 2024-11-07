package user

import (
	"github.com/kms-qwe/auth/internal/service"
	desc "github.com/kms-qwe/auth/pkg/user_v1"
)

// Implementation represents the gRPC handlers that implement the UserV1Server interface
// and use the UserService for business logic operations.
type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

// NewImplementation creates a new instance of GRPCHandlers with the provided UserService.
func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}

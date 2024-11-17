package user

import (
	"context"
	"log"

	"github.com/kms-qwe/auth/internal/converter"
	"github.com/kms-qwe/auth/internal/service"
	desc "github.com/kms-qwe/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GrpcHandlers represents the gRPC handlers that implement the UserV1Server interface
// and use the UserService for business logic operations.
type GrpcHandlers struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

// NewUserGrpcHandlers creates a new instance of GRPCHandlers with the provided UserService.
func NewUserGrpcHandlers(userService service.UserService) *GrpcHandlers {
	return &GrpcHandlers{
		userService: userService,
	}
}

// Create handles the request for creating a new user.
func (g *GrpcHandlers) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := g.userService.Create(ctx, converter.ToUserInfoFromAPI(req.Info))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted user with id: %d", id)

	return &desc.CreateResponse{Id: id}, nil
}

// Delete handles the request for deleting a user.
func (g *GrpcHandlers) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := g.userService.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	log.Printf("delete user with id: %d", req.Id)

	return &emptypb.Empty{}, nil
}

// Get handles the request for retrieving user data.
func (g *GrpcHandlers) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := g.userService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	log.Printf("get user: %#v\n", user)
	return &desc.GetResponse{
		User: converter.ToAPIFromUser(user),
	}, nil
}

// Update handles the request for updating user data.
func (g *GrpcHandlers) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := g.userService.Update(ctx, converter.ToUserInfoUpdateFromAPI(req.UserUpdate))
	if err != nil {
		return nil, err
	}

	log.Printf("update user: id: %d, name: %s, email: %s, role: %d\n", req.UserUpdate.Id, req.UserUpdate.Name, req.UserUpdate.Email, req.UserUpdate.Role)

	return &emptypb.Empty{}, nil
}

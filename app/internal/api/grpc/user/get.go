package user

import (
	"context"
	"log"

	"github.com/kms-qwe/auth/internal/converter"
	desc "github.com/kms-qwe/auth/pkg/user_v1"
)

// Get handles the request for retrieving user data.
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := i.userService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	log.Printf("get user: id: %d, name: %s, email: %s, created_at: %v, updated_at: %v\n", user.ID, user.Info.Name, user.Info.Email, user.CreatedAt, user.UpdatedAt)
	return &desc.GetResponse{
		User: converter.ToDescFromUser(user),
	}, nil
}

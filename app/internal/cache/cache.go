package cache

import (
	"context"
	"time"

	"github.com/kms-qwe/auth/internal/model"
)

// UserCache interface defines methods for user cache operations.
type UserCache interface {
	Expire(ctx context.Context, id int64, ttl time.Duration) error
	// Set adds user to cache
	Set(ctx context.Context, user *model.User) error
	// Get gets user in cache
	Get(ctx context.Context, id int64) (*model.User, error)
	// Delete deletes user fromm cache
	Delete(ctx context.Context, id int64) error
}

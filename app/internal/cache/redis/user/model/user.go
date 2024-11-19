package model

import (
	"github.com/kms-qwe/auth/internal/constant"
)

// User model for redis
type User struct {
	ID        int64         `redis:"id"`
	Name      string        `redis:"name"`
	Email     string        `redis:"email"`
	Password  string        `redis:"password"`
	Role      constant.Role `redis:"role"`
	CreatedAt int64         `redis:"created_at"`
	UpdatedAt *int64        `redis:"updated_at"`
}

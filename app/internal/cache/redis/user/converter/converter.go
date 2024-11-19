package converter

import (
	"time"

	modelCache "github.com/kms-qwe/auth/internal/cache/redis/user/model"
	"github.com/kms-qwe/auth/internal/model"
)

// ToCacheFromUser convert serivce model to cache model
func ToCacheFromUser(user *model.User) *modelCache.User {
	return &modelCache.User{
		ID:        user.ID,
		Name:      user.Info.Name,
		Email:     user.Info.Email,
		Password:  user.Info.Password,
		Role:      user.Info.Role,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: TimePtrToInt64Ptr(user.UpdatedAt),
	}
}

// TimePtrToInt64Ptr converts time ptr to int64 ptr
func TimePtrToInt64Ptr(t *time.Time) *int64 {
	var num int64
	if t != nil {
		num = t.Unix()
	}

	return &num
}

// Int64PtrToTimePtr converts int64 ptr to time ptr
func Int64PtrToTimePtr(n *int64) *time.Time {
	var t time.Time
	if n != nil {
		t = time.Unix(0, *n)
	}

	return &t
}

// ToUserFromCache convert cache model to service model
func ToUserFromCache(user *modelCache.User) *model.User {
	return &model.User{
		ID: user.ID,
		Info: &model.UserInfo{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
			Role:     user.Role,
		},
		CreatedAt: time.Unix(0, user.CreatedAt),
		UpdatedAt: Int64PtrToTimePtr(user.UpdatedAt),
	}
}

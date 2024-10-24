package model

import (
	"time"
)

type User struct {
	Id        int64
	Info      *UserInfo
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserInfo struct {
	Name     string
	Email    string
	Password string
	Role     int32
}

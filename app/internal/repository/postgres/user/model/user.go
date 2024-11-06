package model

import (
	"database/sql"
	"time"
)

// User модель для работы с postgres
type User struct {
	ID        int64        `db:"id"`         // Unique identifier for the user.
	Info      *UserInfo    `db:""`           // Additional information about the user.
	CreatedAt time.Time    `db:"created_at"` // Timestamp of when the user was created.
	UpdatedAt sql.NullTime `db:"updated_at"` // Timestamp of the last update to the user.
}

// UserInfo модель для работы с postres
type UserInfo struct {
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     int32  `db:"role"`
}

type UserInfoUpdate struct {
	ID    int64          `db:"id"`
	Name  sql.NullString `db:"name"`
	Email sql.NullString `db:"email"`
	Role  int32          `db:"role"`
}

package model

import (
	"database/sql"
	"time"
)

// User represents a user in the system with related information.
// It includes an ID, user info, and timestamps for creation and updates.
type User struct {
	ID        int64        // Unique identifier for the user.
	Info      *UserInfo    // Additional information about the user.
	CreatedAt time.Time    // Timestamp of when the user was created.
	UpdatedAt sql.NullTime // Timestamp of the last update to the user.
}

// UserInfo holds information about a user.
type UserInfo struct {
	Name     string
	Email    string
	Password string
	Role     int32
}

// UserInfoUpdate holds information about a user update.
type UserInfoUpdate struct {
	ID    int64
	Name  sql.NullString
	Email sql.NullString
	Role  int32
}

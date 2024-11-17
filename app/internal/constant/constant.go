package constant

// Role represents user role
type Role int32

const (
	RoleUnknown = iota // RoleUnknown represents a role for unknown
	RoleUser           // RoleUser represents a role for user
	RoleAdmin          // RoleAdmin represents a role for admin
)

package constant

// Role represents user role
type Role int32

const (
	RoleUnknown = iota
	RoleUser
	RoleAdmin
)

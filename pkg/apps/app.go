package apps

import (
	"gobble/pkg/common"
	"gobble/pkg/users"
)

// App interface defines the public methods each media service should
// implement
type App interface {
	GetUsers() ([]*users.User, error)
	GetSystemInfo() (*common.SystemInfo, error)
}

package services

import "gobble/pkg/users"

// Service interface defines the public methods each media service should
// implement
type Service interface {
	GetUsers() []users.User
}

package user

import (
	user "api-user-go/internal/domain/user"
)

type RepositoryUser interface {
	SaveUser(user *user.User) (*user.User, error)
	// findUserById(id int) (*user.User, error)
	// findUserByEmail(email string) (*user.User, error)
	// findAllUsers() ([]user.User, error)
	// updateUser(user *user.User) error
	// deleteUser(id int) error
}

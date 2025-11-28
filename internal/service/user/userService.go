package user

import domainUser "api-user-go/internal/domain/user"

type UserService interface {
	CreateUser(user *domainUser.User) error
}

package impl

import (
	domainUser "api-user-go/internal/domain/user"
	"api-user-go/internal/service/user"

	"go.uber.org/zap"
)

type UserServiceImpl struct {
	log *zap.Logger
}

func NewUserServiceImpl() user.UserService {
	return &UserServiceImpl{}
}

func (u *UserServiceImpl) CreateUser(user *domainUser.User) error {
	u.log.Info("iniciando creacion de usuario", zap.String("user", user.FirstName), zap.String("email", user.Email))

}

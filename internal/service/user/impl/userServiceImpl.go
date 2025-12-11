package impl

import (
	"api-user-go/internal/mapper"
	repositoryInterface "api-user-go/internal/repository/user"
	"api-user-go/internal/service/user"
	"api-user-go/internal/service/user/dto"
	"api-user-go/pkg/util/bycry"
	"context"
	"fmt"

	"go.uber.org/zap"
)

type UserService struct {
	repositoryUser repositoryInterface.RepositoryUser
	log            *zap.Logger
}

func NewUserService(repo repositoryInterface.RepositoryUser, log *zap.Logger) user.UserService {
	return &UserService{
		repositoryUser: repo,
		log:            log,
	}
}

func (u *UserService) CreateUser(userDTO *dto.CreateUserRequest, ctx context.Context) (*dto.UserCreateResponse, error) {
	u.log.Info("iniciando creacion de usuario", zap.String("user", userDTO.FirstName), zap.String("email", userDTO.Email))

	passHash, err := bycry.HashPassword(userDTO.Password, bycry.DefaultCost)
	if err != nil {
		u.log.Error("error al hashear password", zap.Error(err))
		return nil, fmt.Errorf("UserService.CreateUser: %w", err)
	}

	user := mapper.ToDomainUser(*userDTO, passHash)

	userCreated, err := u.repositoryUser.SaveUser(user)
	if err != nil {
		u.log.Error("error al crear usuario", zap.Error(err))
		return nil, fmt.Errorf("UserService.CreateUser: %w", err)
	}

	userResponse := mapper.ToUserResponseDTO(userCreated)
	u.log.Info("usuario creado exitosamente", zap.String("user", userCreated.FirstName), zap.String("email", userCreated.Email))
	return &userResponse, nil
}

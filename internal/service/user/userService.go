package user

import (
	"api-user-go/internal/service/user/dto"
	"context"
)

type UserService interface {
	CreateUser(userDTO *dto.CreateUserRequest, ctx context.Context) (*dto.UserCreateResponse, error)
}

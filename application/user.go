package application

import (
	"context"
	"github.com/trewanek/layered-architecture/domain/model"
	"github.com/trewanek/layered-architecture/domain/service"
)

type UserUseCase struct {
	userService service.UserService
}

func NewUserUseCase(userSer service.UserService) *UserUseCase {
	return &UserUseCase{
		userService: userSer,
	}
}

func (u *UserUseCase) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	return u.userService.GetUserByID(ctx, userID)
}

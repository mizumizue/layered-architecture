package service

import (
	"context"
	"github.com/trewanek/layered-architecture/domain/model"
	"github.com/trewanek/layered-architecture/domain/repository"
)

type UserService interface {
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (u *userServiceImpl) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	return u.repo.GetUserByID(ctx, userID)
}

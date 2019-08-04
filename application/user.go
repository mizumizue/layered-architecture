package application

import (
	"context"
	"github.com/trewanek/layered-architecture/domain/exception"
	"github.com/trewanek/layered-architecture/domain/model"
	"github.com/trewanek/layered-architecture/domain/repository"
	"golang.org/x/xerrors"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repository repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		repository: repository,
	}
}

func (u *UserUseCase) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	user, err := u.repository.GetUserByID(ctx, userID)
	if err != nil {
		switch err.(type) {
		case *exception.ResourceNotFoundError:
			return nil, xerrors.Errorf("not found: %w", err)
		case *exception.InternalError:
			return nil, xerrors.Errorf("internal error: %w", err)
		}
		return nil, xerrors.Errorf("unknown error: %w", err)
	}
	return user, nil
}

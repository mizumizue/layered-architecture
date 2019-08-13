package repository

import (
	"context"
	"github.com/trewanek/layered-architecture/domain/model"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
}

package repository

import (
	"context"
	"github.com/trewanek/layered-architecture/domain/model"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
}

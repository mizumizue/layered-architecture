package application

import (
	"context"
	"github.com/trewanek/LayeredArchitectureWithGolang/application/errors"
	"github.com/trewanek/LayeredArchitectureWithGolang/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetUserByID(ctx context.Context, userID string) (*domain.User, error) {
	user, err := domain.GetUserByID(ctx, userID)
	if err != nil {
		unwrapped := errors.UnWrap(err, 2)
		s := status.Convert(unwrapped)
		if s.Code() == codes.NotFound {
			return nil, errors.Errorf(err, errors.ResourceNotFound)
		}
		return nil, errors.Errorf(err, errors.Unknown)
	}
	return user, nil
}

package application

import (
	"context"
	"github.com/trewanek/LayeredArchitectureWithGolang/domain"
	"github.com/trewanek/LayeredArchitectureWithGolang/infrastructure"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetUserByID(ctx context.Context, userID string) (*domain.User, error) {
	user, err := domain.GetUserByID(ctx, userID)
	if err != nil {
		switch UnWrap(err).(type) {
		case *infrastructure.NotFoundError:
			return nil, Errorf(err, ResourceNotFound)
		case *infrastructure.CovertDocumentRefToStructError:
			return nil, Errorf(err, Internal)
		}
		return nil, Errorf(err, Unknown)
	}
	return user, nil
}

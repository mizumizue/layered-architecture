package domain

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/trewanek/LayeredArchitectureWithGolang/infrastructure"
	"golang.org/x/xerrors"
	"os"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Address   string
}

func NewUserFromUserDto(dto *infrastructure.UserDto) *User {
	return &User{
		ID:        dto.ID,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Address:   dto.Address,
	}
}

const (
	projectID     = "ProjectID"
	nameSeparator = " "
)

func (u *User) GetFullName() string {
	return u.FirstName + nameSeparator + u.LastName
}

func GetUserByID(ctx context.Context, userID string) (*User, error) {
	client, err := firestore.NewClient(ctx, os.Getenv(projectID))
	defer client.Close()
	if err != nil {
		return nil, xerrors.Errorf(
			"create database client failed. userID: %d, error in package domain#GetUserByID: %w", userID, err)
	}

	dto, err := infrastructure.GetUserByID(ctx, client, userID)
	if err != nil {
		return nil, xerrors.Errorf(
			"infrastructure.GetUserByID failed. userID: %d, error in package domain#GetUserByID: %w", userID, err)
	}

	return &User{
		ID:        dto.ID,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Address:   dto.Address,
	}, nil
}

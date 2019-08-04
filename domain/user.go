package domain

import (
	"context"
	"github.com/trewanek/LayeredArchitectureWithGolang/infrastructure"
	"golang.org/x/xerrors"
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
	projectID     = "PROJECT_ID"
	nameSeparator = " "
)

func (u *User) GetFullName() string {
	return u.FirstName + nameSeparator + u.LastName
}

func GetUserByID(ctx context.Context, userID string) (*User, error) {
	dbconn, err := infrastructure.NewDBConn(ctx)
	defer dbconn.Close()
	dto, err := infrastructure.GetUserByID(ctx, dbconn, userID)
	if err != nil {
		return nil, xerrors.Errorf("get user by id failed: %w", err)
	}

	return &User{
		ID:        dto.ID,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Address:   dto.Address,
	}, nil
}

package infrastructure

import (
	"context"
)

func getUserByIDMock(ctx context.Context, client *MockClient, userID string) (*UserDto, error) {
	// TODO implements
	return &UserDto{
		ID:        "dummy-user",
		FirstName: "Dummy",
		LastName:  "Taro",
		Address:   "Tokyo Chiyoda",
	}, nil
}

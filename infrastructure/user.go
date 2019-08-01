package infrastructure

import (
	"cloud.google.com/go/firestore"
	"context"
	"golang.org/x/xerrors"
)

const (
	path = "users/"
)

type UserSnapshot struct {
	FirstName string `firestore:"FirstName"`
	LastName  string `firestore:"LastName"`
	Address   string `firestore:"Address"`
}

type UserDto struct {
	ID        string
	FirstName string
	LastName  string
	Address   string
}

func NewUserDto(userID string, snapshot *UserSnapshot) *UserDto {
	return &UserDto{
		ID:        userID,
		FirstName: snapshot.FirstName,
		LastName:  snapshot.LastName,
		Address:   snapshot.Address,
	}
}

func GetUserByID(ctx context.Context, client *firestore.Client, userID string) (*UserDto, error) {
	docRef := client.Doc(path + userID)
	snapshot, err := docRef.Get(ctx)
	if err != nil {
		return nil, xerrors.Errorf(
			"to get snapshot failed. userID: %d, error in package infrastructure#GetUserByID: %w", userID, err)
	}

	var userSnapshot UserSnapshot
	err = snapshot.DataTo(&userSnapshot)
	if err != nil {
		return nil, xerrors.Errorf(
			"unmarshal snapshot to struct failed. userID: %d, error in package infrastructure#GetUserByID: %w", userID, err)
	}

	return NewUserDto(userID, &userSnapshot), nil
}

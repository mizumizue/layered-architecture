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
		return nil, xerrors.Errorf("get user failed from firestore document ref: %v", err)
	}

	var userSnapshot UserSnapshot
	err = snapshot.DataTo(&userSnapshot)
	if err != nil {
		return nil, xerrors.Errorf("convert snapshot to struct failed: %v", err)
	}

	return NewUserDto(userID, &userSnapshot), nil
}

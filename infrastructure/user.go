package infrastructure

import (
	"cloud.google.com/go/firestore"
	"context"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func GetUserByID(ctx context.Context, dbconn *DBConn, userID string) (*UserDto, error) {
	c, ok := dbconn.client.(*firestore.Client)
	if ok {
		return getUserByID(ctx, c, userID)
	}
	cc, ok := dbconn.client.(*MockClient)
	if ok {
		return getUserByIDMock(ctx, cc, userID)
	}
	panic(xerrors.New("db client unknown, convert failed"))
}

func getUserByID(ctx context.Context, client *firestore.Client, userID string) (*UserDto, error) {
	docRef := client.Doc(path + userID)

	snapshot, err := docRef.Get(ctx)
	if err != nil {
		s := status.Convert(err)
		if s.Code() == codes.NotFound {
			return nil, Errorf(err, NotFound)
		}
		return nil, Errorf(err, Unknown)
	}

	var userSnapshot UserSnapshot
	err = snapshot.DataTo(&userSnapshot)
	if err != nil {
		return nil, Errorf(err, CovertDocumentRefToStruct)
	}

	return NewUserDto(userID, &userSnapshot), nil
}

package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/trewanek/layered-architecture/domain/exception"
	"github.com/trewanek/layered-architecture/domain/model"
	"github.com/trewanek/layered-architecture/domain/repository"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
)

const (
	projectID = "PROJECT_ID"
	path      = "users/"
)

type UserSnapshot struct {
	FirstName string `firestore:"FirstName"`
	LastName  string `firestore:"LastName"`
	Address   string `firestore:"Address"`
}

type UserRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	client, err := firestore.NewClient(ctx, os.Getenv(projectID))
	if err != nil {
		return nil, xerrors.Errorf("create firestore client failed: %w", err)
	}
	defer client.Close()

	docRef := client.Doc(path + userID)
	snapshot, err := docRef.Get(ctx)
	if err != nil {
		s := status.Convert(err)
		if s.Code() == codes.NotFound {
			return nil, exception.Errorf(
				xerrors.Errorf("resource not found: %w", err), exception.ResourceNotFound)
		}
		return nil, exception.Errorf(xerrors.Errorf("unknown error: %w", err), exception.Internal)
	}

	var userSnapshot UserSnapshot
	err = snapshot.DataTo(&userSnapshot)
	if err != nil {
		return nil, exception.Errorf(xerrors.Errorf("convert data to struct error: %w", err), exception.Internal)
	}
	return newUserFromSnapshot(userID, &userSnapshot), nil
}

func newUserFromSnapshot(userID string, snapshot *UserSnapshot) *model.User {
	return &model.User{
		ID:        userID,
		FirstName: snapshot.FirstName,
		LastName:  snapshot.LastName,
		Address:   snapshot.Address,
	}
}

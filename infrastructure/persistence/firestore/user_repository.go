package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/trewanek/layered-architecture/domain"
	"github.com/trewanek/layered-architecture/domain/model"
	"github.com/trewanek/layered-architecture/domain/repository"
	"google.golang.org/api/iterator"
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
	Age       int    `firestore:"Age"`
}

type UserRepository struct {
}

func NewUserRepository() repository.UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]*model.User, error) {
	var client *firestore.Client
	var err error
	if client, err = firestore.NewClient(ctx, os.Getenv(projectID)); err != nil {
		return nil, err
	}
	defer client.Close()

	var users []*model.User
	itr := client.Collection("users").Documents(ctx)
	for {
		var err error
		snap, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, domain.Errorf(iteratorNextFailed(err), domain.Internal)
		}
		var us UserSnapshot
		if err = snap.DataTo(&us); err != nil {
			return nil, domain.Errorf(convertToDataFailed(err), domain.Internal)
		}
		users = append(users, newUserFromSnapshot(snap.Ref.ID, &us))
	}
	return users, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	var client *firestore.Client
	var err error
	if client, err = firestore.NewClient(ctx, os.Getenv(projectID)); err != nil {
		return nil, err
	}
	defer client.Close()

	docRef := client.Doc(path + userID)
	snapshot, err := docRef.Get(ctx)
	if err != nil {
		s := status.Convert(err)
		if s.Code() == codes.NotFound {
			return nil, domain.Errorf(documentNotFound(err), domain.ResourceNotFound)
		}
		return nil, domain.Errorf(firestoreUnknownErr(err), domain.Internal)
	}

	var userSnapshot UserSnapshot
	err = snapshot.DataTo(&userSnapshot)
	if err != nil {
		return nil, domain.Errorf(convertToDataFailed(err), domain.Internal)
	}
	return newUserFromSnapshot(snapshot.Ref.ID, &userSnapshot), nil
}

func newUserFromSnapshot(userID string, snapshot *UserSnapshot) *model.User {
	return &model.User{
		ID:        userID,
		FirstName: snapshot.FirstName,
		LastName:  snapshot.LastName,
		Address:   snapshot.Address,
		Age:       snapshot.Age,
	}
}

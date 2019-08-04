package infrastructure

import (
	"cloud.google.com/go/firestore"
	"context"
	"golang.org/x/xerrors"
	"log"
	"os"
	"reflect"
)

const (
	projectID = "PROJECT_ID"
)

type DBConn struct {
	client interface{}
}

type MockClient struct {
}

func NewDBConn(ctx context.Context) (*DBConn, error) {
	if os.Getenv(projectID) == "" {
		return &DBConn{client: &MockClient{}}, nil
	}
	client, err := firestore.NewClient(ctx, os.Getenv(projectID))
	if err != nil {
		return nil, xerrors.Errorf("create firestore client failed: %w", err)
	}
	return &DBConn{client: client}, nil
}

func (dbconn *DBConn) isFirestoreClient(i interface{}) bool {
	return reflect.TypeOf(i) == reflect.TypeOf(firestore.Client{})
}

func (dbconn *DBConn) Close() {
	switch c := dbconn.client.(type) {
	case *firestore.Client:
		err := c.Close()
		if err != nil {
			log.Printf("firestore client close err: %+v", err)
		}
		log.Println("firestore client close completed")
		return
	default:
		return
	}
}

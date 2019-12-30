package firebase

import (
	"cloud.google.com/go/firestore"
	"golang.org/x/net/context"
)

func NewFirestore(ctx context.Context) (*firestore.Client, error) {
	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return firestoreClient, nil
}

var defaultFirestoreClient *firestore.Client

func InitDefaultFirestoreClient(ctx context.Context) {

	client, err := NewFirestore(ctx)
	if err != nil {
		panic("failed to new firestore client")
	}
	defaultFirestoreClient = client
}

func CloseDefaultFirestoreClient() {
	if err := defaultFirestoreClient.Close(); err != nil {
		panic(err)
	}
}

func DefaultFirestoreClient() *firestore.Client {
	return defaultFirestoreClient
}

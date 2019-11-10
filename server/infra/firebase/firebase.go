package firebase

import (
	"fmt"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

var (
	app *firebase.App
)

func Init(ctx context.Context) error {
	opt := option.WithCredentialsFile("./secret/firebase-account-key.json")
	firebaseApp, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}

	app = firebaseApp
	return nil
}

package firebase

import (
	"log"

	"todo/auth"

	firebaseAuth "firebase.google.com/go/auth"
	"golang.org/x/net/context"
)

type FirebaseAuth struct {
	client *firebaseAuth.Client
}

func NewAuth(ctx context.Context) (*FirebaseAuth, error) {
	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	instance := &FirebaseAuth{
		client: authClient,
	}
	return instance, nil
}

// VerifyIDToken firebaseトークンからアカウント取得
func (a *FirebaseAuth) VerifyIDToken(ctx context.Context, idToken string) (auth.Account, error) {
	// auth.TokenはFirebaseIDTokenをデコードしたもの。UIDがUserID
	token, err := a.client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
		return auth.Account{}, err
	}
	account := auth.Account{
		UserID: token.UID,
	}

	return account, nil
}

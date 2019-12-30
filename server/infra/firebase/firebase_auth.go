package firebase

import (
	"todo/auth"

	firebaseAuth "firebase.google.com/go/auth"
	"github.com/pkg/errors"
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
		return auth.Account{}, errors.WithStack(err)
	}
	account := auth.Account{
		UserID: token.UID,
	}

	return account, nil
}

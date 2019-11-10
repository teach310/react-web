package auth

import "context"

type contextKey int

const AccountContextKey = contextKey(1)

// Account ctxに流す
type Account struct {
	UserID string
}

func SetAccountOnContect(ctx context.Context, account Account) context.Context {
	return context.WithValue(ctx, AccountContextKey, account)
}

func GetAccountFromContext(ctx context.Context) (Account, bool) {
	if account, ok := ctx.Value(AccountContextKey).(Account); ok {
		return account, true
	}
	return Account{}, false
}
